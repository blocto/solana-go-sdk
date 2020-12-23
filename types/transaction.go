package types

import (
	"crypto/ed25519"
	"errors"

	"github.com/portto/solana-go-sdk/common"
)

type Signature []byte

type Transaction struct {
	Signatures []Signature
	Message    Message
}

func (tx *Transaction) sign(privateKeys []ed25519.PrivateKey) (*Transaction, error) {
	messageData, err := tx.Message.Serialize()
	if err != nil {
		return nil, err
	}

	tx.Signatures = make([]Signature, 0, len(privateKeys))
	for _, privateKey := range privateKeys {
		signature := ed25519.Sign(privateKey, messageData)
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

type CreateRawTransactionParam struct {
	Instructions    []Instruction
	PrivateKeys     []ed25519.PrivateKey
	FeePayer        common.PublicKey
	RecentBlockHash string
}

func NewTransaction(param CreateRawTransactionParam) ([]byte, error) {
	if param.RecentBlockHash == "" {
		return nil, errors.New("recent block hash is required")
	}
	if len(param.Instructions) < 1 {
		return nil, errors.New("no instructions provided")
	}

	message := compileMessage(param.FeePayer, param.Instructions, param.RecentBlockHash)
	if int(message.Header.NumRequireSignatures) != len(param.PrivateKeys) {
		return nil, errors.New("")
	}

	tx := Transaction{
		Signatures: []Signature{},
		Message:    message,
	}
	signTx, err := tx.sign(param.PrivateKeys)
	if err != nil {
		return nil, err
	}

	return signTx.Serialize()
}
