package types

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/sasaxie/go-client-api/common/base58"
)

type MessageHeader struct {
	NumRequireSignatures        uint8
	NumReadonlySignedAccounts   uint8
	NumReadonlyUnsignedAccounts uint8
}

type Message struct {
	Header          MessageHeader
	Accounts        []common.PublicKey
	RecentBlockHash string
	Instructions    []CompiledInstruction
}

func (m *Message) Serialize() ([]byte, error) {
	b := []byte{}
	b = append(b, m.Header.NumRequireSignatures)
	b = append(b, m.Header.NumReadonlySignedAccounts)
	b = append(b, m.Header.NumReadonlyUnsignedAccounts)

	b = append(b, common.UintToVarLenBytes(uint64(len(m.Accounts)))...)
	for _, key := range m.Accounts {
		b = append(b, key[:]...)
	}

	blockHash, err := base58.Decode(m.RecentBlockHash)
	if err != nil {
		return nil, err
	}
	b = append(b, blockHash...)

	b = append(b, common.UintToVarLenBytes(uint64(len(m.Instructions)))...)
	for _, instruction := range m.Instructions {
		b = append(b, byte(instruction.ProgramIDIndex))
		b = append(b, common.UintToVarLenBytes(uint64(len(instruction.Accounts)))...)
		for _, accountIdx := range instruction.Accounts {
			b = append(b, byte(accountIdx))
		}

		b = append(b, byte(len(instruction.Data)))
		b = append(b, instruction.Data...)
	}
	return b, nil
}

func NewMessage(feePayer common.PublicKey, instructions []Instruction, recentBlockHash string) Message {
	accountMap := map[common.PublicKey]*AccountMeta{}
	for _, instruction := range instructions {
		// program is a readonly unsigned account
		_, exist := accountMap[instruction.ProgramID]
		if !exist {
			accountMap[instruction.ProgramID] = &AccountMeta{
				PubKey:     instruction.ProgramID,
				IsSigner:   false,
				IsWritable: false,
			}
		}
		for i := 0; i < len(instruction.Accounts); i++ {
			account := instruction.Accounts[i]
			a, exist := accountMap[account.PubKey]
			if !exist {
				accountMap[account.PubKey] = &account
			} else {
				a.IsSigner = a.IsSigner || account.IsSigner
				a.IsWritable = a.IsWritable || account.IsWritable
			}
		}
	}

	writableSignedAccount := []common.PublicKey{}
	readOnlySignedAccount := []common.PublicKey{}
	writableUnsignedAccount := []common.PublicKey{}
	readOnlyUnsignedAccount := []common.PublicKey{}
	classify := func(account *AccountMeta) {
		if account.IsSigner {
			if account.IsWritable {
				writableSignedAccount = append(writableSignedAccount, account.PubKey)
			} else {
				readOnlySignedAccount = append(readOnlySignedAccount, account.PubKey)
			}
		} else {
			if account.IsWritable {
				writableUnsignedAccount = append(writableUnsignedAccount, account.PubKey)
			} else {
				readOnlyUnsignedAccount = append(readOnlyUnsignedAccount, account.PubKey)
			}
		}
	}
	if feePayer != common.ZeroPublicKey {
		for _, account := range accountMap {
			if feePayer == account.PubKey {
				continue
			}
			classify(account)
		}
		writableSignedAccount = append([]common.PublicKey{feePayer}, writableSignedAccount...)
	} else {
		for _, account := range accountMap {
			classify(account)
		}
	}

	publicKeys := []common.PublicKey{}
	publicKeys = append(publicKeys, writableSignedAccount...)
	publicKeys = append(publicKeys, readOnlySignedAccount...)
	publicKeys = append(publicKeys, writableUnsignedAccount...)
	publicKeys = append(publicKeys, readOnlyUnsignedAccount...)
	publicKeyToIdx := map[common.PublicKey]int{}
	for idx, publicKey := range publicKeys {
		publicKeyToIdx[publicKey] = idx
	}

	compiledInstructions := []CompiledInstruction{}
	for _, instruction := range instructions {
		accountIdx := []int{}
		for _, account := range instruction.Accounts {
			accountIdx = append(accountIdx, publicKeyToIdx[account.PubKey])
		}
		compiledInstructions = append(compiledInstructions, CompiledInstruction{
			ProgramIDIndex: publicKeyToIdx[instruction.ProgramID],
			Accounts:       accountIdx,
			Data:           instruction.Data,
		})
	}

	return Message{
		Header: MessageHeader{
			NumRequireSignatures:        uint8(len(writableSignedAccount) + len(readOnlySignedAccount)),
			NumReadonlySignedAccounts:   uint8(len(readOnlySignedAccount)),
			NumReadonlyUnsignedAccounts: uint8(len(readOnlyUnsignedAccount)),
		},
		Accounts:        publicKeys,
		RecentBlockHash: recentBlockHash,
		Instructions:    compiledInstructions,
	}
}
