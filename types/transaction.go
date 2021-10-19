package types

import (
	"crypto/ed25519"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/bincode"
)

var (
	ErrTransactionAddNotNecessarySignatures = errors.New("add not necessary signatures")
)

type Signature []byte

type Transaction struct {
	Signatures []Signature
	Message    Message
}

type NewTransactionParam struct {
	Message Message
	Signers []Account
}

// NewTransaction create a new tx by message and signer. it will reserve signatures slot.
func NewTransaction(param NewTransactionParam) (Transaction, error) {
	signatures := make([]Signature, 0, param.Message.Header.NumRequireSignatures)
	for i := uint8(0); i < param.Message.Header.NumRequireSignatures; i++ {
		signatures = append(signatures, make([]byte, 64))
	}

	m := map[common.PublicKey]uint8{}
	for i := uint8(0); i < param.Message.Header.NumRequireSignatures; i++ {
		m[param.Message.Accounts[i]] = i
	}

	data, err := param.Message.Serialize()
	if err != nil {
		return Transaction{}, fmt.Errorf("failed to serialize message, err: %v", err)
	}
	for _, signer := range param.Signers {
		idx, ok := m[signer.PublicKey]
		if !ok {
			return Transaction{}, fmt.Errorf("%w, %v is not a signer", ErrTransactionAddNotNecessarySignatures, signer.PublicKey)
		}
		signatures[idx] = signer.Sign(data)
	}

	return Transaction{
		Signatures: signatures,
		Message:    param.Message,
	}, nil
}

// AddSignature will add or replace signature into the correct order signature's slot.
func (tx *Transaction) AddSignature(sig []byte) error {
	data, err := tx.Message.Serialize()
	if err != nil {
		return fmt.Errorf("failed to serialize message, err: %v", err)
	}
	for i := uint8(0); i < tx.Message.Header.NumRequireSignatures; i++ {
		a := tx.Message.Accounts[i]
		if ed25519.Verify(a.Bytes(), data, sig) {
			tx.Signatures[i] = sig
			return nil
		}
	}
	return fmt.Errorf("%w, no match signer", ErrTransactionAddNotNecessarySignatures)
}

// Serialize pack tx into byte array
func (tx *Transaction) Serialize() ([]byte, error) {
	if len(tx.Signatures) == 0 || len(tx.Signatures) != int(tx.Message.Header.NumRequireSignatures) {
		return nil, errors.New("Signature verification failed")
	}

	signatureCount := bincode.UintToVarLenBytes(uint64(len(tx.Signatures)))
	messageData, err := tx.Message.Serialize()
	if err != nil {
		return nil, err
	}

	output := make([]byte, 0, len(signatureCount)+len(signatureCount)*64+len(messageData))
	output = append(output, signatureCount...)
	for _, sig := range tx.Signatures {
		output = append(output, sig...)
	}
	output = append(output, messageData...)

	return output, nil
}

// TransactionDeserialize can deserialize a tx from byte array
func TransactionDeserialize(tx []byte) (Transaction, error) {
	signatureCount, err := parseUvarint(&tx)
	if err != nil {
		return Transaction{}, fmt.Errorf("parse signature count error: %v", err)
	}
	if signatureCount < 1 {
		return Transaction{}, errors.New("signature count must be greater than or equal to 1")
	}
	if len(tx) < int(signatureCount)*64 {
		return Transaction{}, errors.New("parse signature error")
	}
	signatures := make([]Signature, 0, signatureCount)
	for i := 0; i < int(signatureCount); i++ {
		signatures = append(signatures, tx[:64])
		tx = tx[64:]
	}

	message, err := MessageDeserialize(tx)
	if err != nil {
		return Transaction{}, fmt.Errorf("failed to parse message, err: %v", err)
	}

	if uint64(message.Header.NumRequireSignatures) != signatureCount {
		return Transaction{}, errors.New("numRequireSignatures is not equal to signatureCount")
	}

	return Transaction{
		Signatures: signatures,
		Message:    message,
	}, nil
}

// MustTransactionDeserialize can deserialize a tx from byte array, it panic if error ocour
func MustTransactionDeserialize(data []byte) Transaction {
	tx, err := TransactionDeserialize(data)
	if err != nil {
		panic(err)
	}
	return tx
}

func parseUvarint(tx *[]byte) (uint64, error) {
	if len(*tx) == 0 {
		return 0, errors.New("data is empty")
	}
	u, n := binary.Uvarint(*tx)
	if n <= 0 {
		return 0, errors.New("format error")
	}
	*tx = (*tx)[n:]
	return u, nil
}
