package types

import (
	"crypto/ed25519"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/mr-tron/base58"
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
	parseUvarint := func(tx *[]byte) (uint64, error) {
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

	var numRequireSignatures, numReadonlySignedAccounts, numReadonlyUnsignedAccounts uint8
	var t uint64
	list := []*uint8{&numRequireSignatures, &numReadonlySignedAccounts, &numReadonlyUnsignedAccounts}
	for i := 0; i < len(list); i++ {
		t, err = parseUvarint(&tx)
		if t > 255 {
			return Transaction{}, fmt.Errorf("message header #%d parse error: %v", i+1, err)
		}
		*list[i] = uint8(t)
	}
	if uint64(numRequireSignatures) != signatureCount {
		return Transaction{}, errors.New("numRequireSignatures is not equal to signatureCount")
	}

	accountCount, err := parseUvarint(&tx)
	if len(tx) < int(accountCount)*32 {
		return Transaction{}, errors.New("parse account error")
	}
	accounts := make([]common.PublicKey, 0, accountCount)
	for i := 0; i < int(accountCount); i++ {
		accounts = append(accounts, common.PublicKeyFromHex(hex.EncodeToString(tx[:32])))
		tx = tx[32:]
	}

	if len(tx) < 32 {
		return Transaction{}, errors.New("parse blockhash error")
	}
	blockHash := base58.Encode(tx[:32])
	tx = tx[32:]

	instructionCount, err := parseUvarint(&tx)
	if err != nil {
		return Transaction{}, fmt.Errorf("parse instruction count error: %v", err)
	}

	instructions := make([]CompiledInstruction, 0, instructionCount)
	for i := 0; i < int(instructionCount); i++ {
		programID, err := parseUvarint(&tx)
		if err != nil {
			return Transaction{}, fmt.Errorf("parse instruction #%d programID error: %v", i+1, err)
		}
		accountCount, err := parseUvarint(&tx)
		if err != nil {
			return Transaction{}, fmt.Errorf("parse instruction #%d account count error: %v", i+1, err)
		}
		accounts := make([]int, 0, accountCount)
		for j := 0; j < int(accountCount); j++ {
			accountIdx, err := parseUvarint(&tx)
			if err != nil {
				return Transaction{}, fmt.Errorf("parse instruction #%d account #%d idx error: %v", i+1, j+1, err)
			}
			accounts = append(accounts, int(accountIdx))
		}
		dataLen, err := parseUvarint(&tx)
		if err != nil {
			return Transaction{}, fmt.Errorf("parse instruction #%d data length error: %v", i+1, err)
		}
		var data []byte
		data, tx = tx[:dataLen], tx[dataLen:]

		instructions = append(instructions, CompiledInstruction{
			ProgramIDIndex: int(programID),
			Accounts:       accounts,
			Data:           data,
		})
	}

	return Transaction{
		Signatures: signatures,
		Message: Message{
			Header: MessageHeader{
				NumRequireSignatures:        numRequireSignatures,
				NumReadonlySignedAccounts:   numReadonlySignedAccounts,
				NumReadonlyUnsignedAccounts: numReadonlyUnsignedAccounts,
			},
			Accounts:        accounts,
			RecentBlockHash: blockHash,
			Instructions:    instructions,
		},
	}, nil
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
