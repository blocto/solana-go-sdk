package types

import (
	"bytes"
	"errors"
	"fmt"
	"sort"

	"github.com/OldSmokeGun/solana-go-sdk/common"
	"github.com/OldSmokeGun/solana-go-sdk/pkg/bincode"

	"github.com/mr-tron/base58"
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

	b = append(b, bincode.UintToVarLenBytes(uint64(len(m.Accounts)))...)
	for _, key := range m.Accounts {
		b = append(b, key[:]...)
	}

	blockHash, err := base58.Decode(m.RecentBlockHash)
	if err != nil {
		return nil, err
	}
	b = append(b, blockHash...)

	b = append(b, bincode.UintToVarLenBytes(uint64(len(m.Instructions)))...)
	for _, instruction := range m.Instructions {
		b = append(b, byte(instruction.ProgramIDIndex))
		b = append(b, bincode.UintToVarLenBytes(uint64(len(instruction.Accounts)))...)
		for _, accountIdx := range instruction.Accounts {
			b = append(b, byte(accountIdx))
		}

		b = append(b, bincode.UintToVarLenBytes(uint64(len(instruction.Data)))...)
		b = append(b, instruction.Data...)
	}
	return b, nil
}

func (m *Message) DecompileInstructions() []Instruction {
	instructions := make([]Instruction, 0, len(m.Instructions))
	for _, cins := range m.Instructions {
		accounts := make([]AccountMeta, 0, len(cins.Accounts))
		for i := 0; i < len(cins.Accounts); i++ {
			accounts = append(accounts, AccountMeta{
				PubKey:   m.Accounts[cins.Accounts[i]],
				IsSigner: cins.Accounts[i] < int(m.Header.NumRequireSignatures),
				IsWritable: cins.Accounts[i] < int(m.Header.NumRequireSignatures-m.Header.NumReadonlySignedAccounts) ||
					(cins.Accounts[i] >= int(m.Header.NumRequireSignatures) &&
						cins.Accounts[i] < len(m.Accounts)-int(m.Header.NumReadonlyUnsignedAccounts)),
			})
		}
		instructions = append(instructions, Instruction{
			ProgramID: m.Accounts[cins.ProgramIDIndex],
			Accounts:  accounts,
			Data:      cins.Data,
		})
	}
	return instructions
}

func MessageDeserialize(messageData []byte) (Message, error) {
	var numRequireSignatures, numReadonlySignedAccounts, numReadonlyUnsignedAccounts uint8
	var t uint64
	var err error
	list := []*uint8{&numRequireSignatures, &numReadonlySignedAccounts, &numReadonlyUnsignedAccounts}
	for i := 0; i < len(list); i++ {
		t, err = parseUvarint(&messageData)
		if t > 255 {
			return Message{}, fmt.Errorf("message header #%d parse error: %v", i+1, err)
		}
		*list[i] = uint8(t)
	}

	accountCount, err := parseUvarint(&messageData)
	if err != nil {
		return Message{}, fmt.Errorf("falied to parse count of account, err: %v", err)
	}
	if len(messageData) < int(accountCount)*32 {
		return Message{}, errors.New("parse account error")
	}
	accounts := make([]common.PublicKey, 0, accountCount)
	for i := 0; i < int(accountCount); i++ {
		accounts = append(accounts, common.PublicKeyFromBytes(messageData[:32]))
		messageData = messageData[32:]
	}

	if len(messageData) < 32 {
		return Message{}, errors.New("parse blockhash error")
	}
	blockHash := base58.Encode(messageData[:32])
	messageData = messageData[32:]

	instructionCount, err := parseUvarint(&messageData)
	if err != nil {
		return Message{}, fmt.Errorf("parse instruction count error: %v", err)
	}

	instructions := make([]CompiledInstruction, 0, instructionCount)
	for i := 0; i < int(instructionCount); i++ {
		programID, err := parseUvarint(&messageData)
		if err != nil {
			return Message{}, fmt.Errorf("parse instruction #%d programID error: %v", i+1, err)
		}
		accountCount, err := parseUvarint(&messageData)
		if err != nil {
			return Message{}, fmt.Errorf("parse instruction #%d account count error: %v", i+1, err)
		}
		accounts := make([]int, 0, accountCount)
		for j := 0; j < int(accountCount); j++ {
			accountIdx, err := parseUvarint(&messageData)
			if err != nil {
				return Message{}, fmt.Errorf("parse instruction #%d account #%d idx error: %v", i+1, j+1, err)
			}
			accounts = append(accounts, int(accountIdx))
		}
		dataLen, err := parseUvarint(&messageData)
		if err != nil {
			return Message{}, fmt.Errorf("parse instruction #%d data length error: %v", i+1, err)
		}
		var data []byte
		data, messageData = messageData[:dataLen], messageData[dataLen:]

		instructions = append(instructions, CompiledInstruction{
			ProgramIDIndex: int(programID),
			Accounts:       accounts,
			Data:           data,
		})
	}

	return Message{
		Header: MessageHeader{
			NumRequireSignatures:        numRequireSignatures,
			NumReadonlySignedAccounts:   numReadonlySignedAccounts,
			NumReadonlyUnsignedAccounts: numReadonlyUnsignedAccounts,
		},
		Accounts:        accounts,
		RecentBlockHash: blockHash,
		Instructions:    instructions,
	}, nil
}

func MustMessageDeserialize(messageData []byte) Message {
	message, err := MessageDeserialize(messageData)
	if err != nil {
		panic(err)
	}
	return message
}

type NewMessageParam struct {
	FeePayer        common.PublicKey
	Instructions    []Instruction
	RecentBlockhash string
}

func NewMessage(param NewMessageParam) Message {
	accountMap := map[common.PublicKey]*AccountMeta{}
	for _, instruction := range param.Instructions {
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
	sortAllAccount := func() {
		sort.Slice(writableSignedAccount, func(i, j int) bool {
			return bytes.Compare(writableSignedAccount[i].Bytes(), writableSignedAccount[j].Bytes()) < 0
		})
		sort.Slice(readOnlySignedAccount, func(i, j int) bool {
			return bytes.Compare(readOnlySignedAccount[i].Bytes(), readOnlySignedAccount[j].Bytes()) < 0
		})
		sort.Slice(writableUnsignedAccount, func(i, j int) bool {
			return bytes.Compare(writableUnsignedAccount[i].Bytes(), writableUnsignedAccount[j].Bytes()) < 0
		})
		sort.Slice(readOnlyUnsignedAccount, func(i, j int) bool {
			return bytes.Compare(readOnlyUnsignedAccount[i].Bytes(), readOnlyUnsignedAccount[j].Bytes()) < 0
		})
	}
	if param.FeePayer != (common.PublicKey{}) {
		for _, account := range accountMap {
			if param.FeePayer == account.PubKey {
				continue
			}
			classify(account)
		}
		sortAllAccount()
		writableSignedAccount = append([]common.PublicKey{param.FeePayer}, writableSignedAccount...)
	} else {
		for _, account := range accountMap {
			classify(account)
		}
		sortAllAccount()
	}

	publicKeys := make([]common.PublicKey, 0, len(writableSignedAccount)+len(readOnlySignedAccount)+len(writableUnsignedAccount)+len(readOnlyUnsignedAccount))
	publicKeys = append(publicKeys, writableSignedAccount...)
	publicKeys = append(publicKeys, readOnlySignedAccount...)
	publicKeys = append(publicKeys, writableUnsignedAccount...)
	publicKeys = append(publicKeys, readOnlyUnsignedAccount...)
	publicKeyToIdx := map[common.PublicKey]int{}
	for idx, publicKey := range publicKeys {
		publicKeyToIdx[publicKey] = idx
	}

	compiledInstructions := []CompiledInstruction{}
	for _, instruction := range param.Instructions {
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
		RecentBlockHash: param.RecentBlockhash,
		Instructions:    compiledInstructions,
	}
}
