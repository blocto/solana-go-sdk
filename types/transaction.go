package types

import (
	"crypto/ed25519"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/portto/solana-go-sdk/common"
)

type Signature []byte

type Transaction struct {
	Signatures []Signature
	Message    Message
}

func (tx *Transaction) sign(accounts []Account) (*Transaction, error) {
	if int(tx.Message.Header.NumRequireSignatures) != len(accounts) {
		return nil, errors.New("signer's num not match")
	}

	message, err := tx.Message.Serialize()
	if err != nil {
		return nil, err
	}

	accountMap := map[common.PublicKey]ed25519.PrivateKey{}
	for _, account := range accounts {
		accountMap[account.PublicKey] = account.PrivateKey
	}

	for i := 0; i < int(tx.Message.Header.NumRequireSignatures); i++ {
		privateKey, exist := accountMap[tx.Message.Accounts[i]]
		if !exist {
			return nil, fmt.Errorf("lack %s's private key", tx.Message.Accounts[i].ToBase58())
		}
		signature := ed25519.Sign(privateKey, message)
		tx.Signatures = append(tx.Signatures, signature)
	}
	return tx, nil
}

func (tx *Transaction) Serialize() ([]byte, error) {
	if len(tx.Signatures) == 0 || len(tx.Signatures) != int(tx.Message.Header.NumRequireSignatures) {
		return nil, errors.New("Signature verification failed")
	}

	signatureCount := common.UintToVarLenBytes(uint64(len(tx.Signatures)))
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

	if uint64(message.Header.NumRequireSignatures) != signatureCount {
		return Transaction{}, errors.New("numRequireSignatures is not equal to signatureCount")
	}

	return Transaction{
		Signatures: signatures,
		Message:    message,
	}, nil
}

func MustTransactionDeserialize(data []byte) Transaction {
	tx, err := TransactionDeserialize(data)
	if err != nil {
		panic(err)
	}
	return tx
}

type CreateRawTransactionParam struct {
	Instructions    []Instruction
	Signers         []Account
	FeePayer        common.PublicKey
	RecentBlockHash string
}

func CreateRawTransaction(param CreateRawTransactionParam) ([]byte, error) {
	if param.RecentBlockHash == "" {
		return nil, errors.New("recent block hash is required")
	}
	if len(param.Instructions) < 1 {
		return nil, errors.New("no instructions provided")
	}

	tx := Transaction{
		Signatures: []Signature{},
		Message:    NewMessage(param.FeePayer, param.Instructions, param.RecentBlockHash),
	}

	signTx, err := tx.sign(param.Signers)
	if err != nil {
		return nil, err
	}

	return signTx.Serialize()
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
