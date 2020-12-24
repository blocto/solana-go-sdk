package types

import (
	"crypto/ed25519"
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
		Message:    compileMessage(param.FeePayer, param.Instructions, param.RecentBlockHash),
	}

	signTx, err := tx.sign(param.Signers)
	if err != nil {
		return nil, err
	}

	return signTx.Serialize()
}
