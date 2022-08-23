package types

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/bincode"

	"github.com/mr-tron/base58"
)

type MessageHeader struct {
	NumRequireSignatures        uint8
	NumReadonlySignedAccounts   uint8
	NumReadonlyUnsignedAccounts uint8
}

type MessageVersion string

const (
	MessageVersionLegacy = "legacy"
	MessageVersionV0     = "v0"
)

type Message struct {
	Version            MessageVersion
	Header             MessageHeader
	Accounts           []common.PublicKey
	RecentBlockHash    string
	Instructions       []CompiledInstruction
	AddressLookupTable *CompiledAddressLookupTable
}

type CompiledAddressLookupTable struct {
	AccountKey      common.PublicKey
	WritableIndexes []uint8
	ReadonlyIndexes []uint8
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

	if len(m.Version) > 0 && m.Version != MessageVersionLegacy {
		versionNum, err := strconv.Atoi(string(m.Version[1:]))
		if err != nil || versionNum > 255 {
			return nil, fmt.Errorf("failed to parse message version")
		}
		if versionNum > 128 {
			return nil, fmt.Errorf("unexpected message version")
		}
		b = append([]byte{byte(versionNum + 128)}, b...)

		if m.AddressLookupTable != nil &&
			(len(m.AddressLookupTable.WritableIndexes) != 0 || len(m.AddressLookupTable.ReadonlyIndexes) != 0) {
			b = append(b, 1)
			b = append(b, m.AddressLookupTable.AccountKey.Bytes()...)
			b = append(b, bincode.UintToVarLenBytes(uint64(len(m.AddressLookupTable.WritableIndexes)))...)
			b = append(b, m.AddressLookupTable.WritableIndexes...)
			b = append(b, bincode.UintToVarLenBytes(uint64(len(m.AddressLookupTable.ReadonlyIndexes)))...)
			b = append(b, m.AddressLookupTable.ReadonlyIndexes...)
		} else {
			b = append(b, 0)
		}
	}

	return b, nil
}

// DecompileInstructions hasn't support v0 message decode
func (m *Message) DecompileInstructions() []Instruction {
	switch m.Version {
	case MessageVersionLegacy:
		return m.decompileLegacyMessageInstructions()
	case MessageVersionV0:
		panic("hasn't supported")
	default:
		return m.decompileLegacyMessageInstructions()
	}
}

func (m Message) decompileLegacyMessageInstructions() []Instruction {
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
	var version MessageVersion
	if v := uint8(messageData[0]); v > 127 {
		version = MessageVersion(fmt.Sprintf("v%v", v-128))
		messageData = messageData[1:]
	} else {
		version = MessageVersionLegacy
	}

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

	var compiledAddressLookupTable *CompiledAddressLookupTable = nil
	if version != MessageVersionLegacy {
		hasCompiledAddressLookupTable := messageData[0]
		messageData = messageData[1:]
		if hasCompiledAddressLookupTable == 1 {
			addressLookupTablePubkey := common.PublicKeyFromBytes(messageData[:32])
			messageData = messageData[32:]

			writableAccountIdxCount, err := parseUvarint(&messageData)
			if err != nil {
				return Message{}, fmt.Errorf("failed to parse address lookup table writable account idx count, err: %v", err)
			}
			var writableAccountIdxList []uint8
			writableAccountIdxList, messageData = messageData[:writableAccountIdxCount], messageData[writableAccountIdxCount:]

			readOnlyAccountIdxCount, err := parseUvarint(&messageData)
			if err != nil {
				return Message{}, fmt.Errorf("failed to parse address lookup table readOnly account idx count, err: %v", err)
			}
			var readOnlyAccountIdxList []uint8
			readOnlyAccountIdxList, messageData = messageData[:readOnlyAccountIdxCount], messageData[readOnlyAccountIdxCount:]

			compiledAddressLookupTable = &CompiledAddressLookupTable{
				AccountKey:      addressLookupTablePubkey,
				WritableIndexes: writableAccountIdxList,
				ReadonlyIndexes: readOnlyAccountIdxList,
			}
		}
	}

	return Message{
		Version: version,
		Header: MessageHeader{
			NumRequireSignatures:        numRequireSignatures,
			NumReadonlySignedAccounts:   numReadonlySignedAccounts,
			NumReadonlyUnsignedAccounts: numReadonlyUnsignedAccounts,
		},
		Accounts:           accounts,
		RecentBlockHash:    blockHash,
		Instructions:       instructions,
		AddressLookupTable: compiledAddressLookupTable,
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
	// v0 transaction
	AddressLookupTable          common.PublicKey
	AddressLookupTableAddresses []common.PublicKey
}

type accountMeta struct {
	IsSigner   bool
	IsWritable bool
	IsProgram  bool
}

func compileKey(instructions []Instruction) map[common.PublicKey]*accountMeta {
	m := map[common.PublicKey]*accountMeta{}

	for _, instruction := range instructions {
		// compile program
		v, exist := m[instruction.ProgramID]
		if !exist {
			m[instruction.ProgramID] = &accountMeta{
				IsSigner:   false,
				IsWritable: false,
				IsProgram:  true,
			}
		} else {
			v.IsProgram = true
		}

		// compile accounts
		for i := 0; i < len(instruction.Accounts); i++ {
			account := instruction.Accounts[i]
			v, exist := m[account.PubKey]
			if !exist {
				m[account.PubKey] = &accountMeta{
					IsSigner:   account.IsSigner,
					IsWritable: account.IsWritable,
				}
			} else {
				v.IsSigner = v.IsSigner || account.IsSigner
				v.IsWritable = v.IsWritable || account.IsWritable
			}
		}
	}

	return m
}

func NewMessage(param NewMessageParam) Message {
	compiledKeys := compileKey(param.Instructions)

	writableSignedAccount := []common.PublicKey{}
	readOnlySignedAccount := []common.PublicKey{}
	writableUnsignedAccount := []common.PublicKey{}
	readOnlyUnsignedAccount := []common.PublicKey{}
	addressLookupTableWritable := []common.PublicKey{}
	addressLookupTableReadonly := []common.PublicKey{}

	addressLookupTableWritableIdx := []uint8{}
	addressLookupTableReadonlyIdx := []uint8{}

	addressLookupTableMap := map[common.PublicKey]uint8{}
	for i, k := range param.AddressLookupTableAddresses {
		addressLookupTableMap[k] = uint8(i)
	}

	classify := func(key common.PublicKey, meta *accountMeta) {
		if meta.IsSigner {
			if meta.IsWritable {
				writableSignedAccount = append(writableSignedAccount, key)
			} else {
				readOnlySignedAccount = append(readOnlySignedAccount, key)
			}
		} else {
			if meta.IsWritable {
				idx, exist := addressLookupTableMap[key]
				if exist && !meta.IsProgram {
					addressLookupTableWritable = append(addressLookupTableWritable, key)
					addressLookupTableWritableIdx = append(addressLookupTableWritableIdx, idx)
				} else {
					writableUnsignedAccount = append(writableUnsignedAccount, key)
				}
			} else {
				idx, exist := addressLookupTableMap[key]
				if exist && !meta.IsProgram {
					addressLookupTableReadonly = append(addressLookupTableReadonly, key)
					addressLookupTableReadonlyIdx = append(addressLookupTableReadonlyIdx, idx)
				} else {
					readOnlyUnsignedAccount = append(readOnlyUnsignedAccount, key)
				}
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
		for key, meta := range compiledKeys {
			if param.FeePayer == key {
				continue
			}
			classify(key, meta)
		}
		sortAllAccount()
		writableSignedAccount = append([]common.PublicKey{param.FeePayer}, writableSignedAccount...)
	} else {
		for key, meta := range compiledKeys {
			classify(key, meta)
		}
		sortAllAccount()
	}

	l := 0 +
		len(writableSignedAccount) +
		len(readOnlySignedAccount) +
		len(writableUnsignedAccount) +
		len(readOnlyUnsignedAccount) +
		len(addressLookupTableWritable) +
		len(addressLookupTableReadonly)

	publicKeys := make([]common.PublicKey, 0, l)
	publicKeys = append(publicKeys, writableSignedAccount...)
	publicKeys = append(publicKeys, readOnlySignedAccount...)
	publicKeys = append(publicKeys, writableUnsignedAccount...)
	publicKeys = append(publicKeys, readOnlyUnsignedAccount...)
	publicKeys = append(publicKeys, addressLookupTableWritable...)
	publicKeys = append(publicKeys, addressLookupTableReadonly...)
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

	var version MessageVersion = MessageVersionLegacy
	if param.AddressLookupTable != (common.PublicKey{}) {
		version = MessageVersionV0
	}

	var compiledAddressLookupTable *CompiledAddressLookupTable = nil
	if len(addressLookupTableWritable) != 0 || len(addressLookupTableReadonly) != 0 {
		compiledAddressLookupTable = &CompiledAddressLookupTable{
			AccountKey:      param.AddressLookupTable,
			WritableIndexes: addressLookupTableWritableIdx,
			ReadonlyIndexes: addressLookupTableReadonlyIdx,
		}
	}
	return Message{
		Version: version,
		Header: MessageHeader{
			NumRequireSignatures:        uint8(len(writableSignedAccount) + len(readOnlySignedAccount)),
			NumReadonlySignedAccounts:   uint8(len(readOnlySignedAccount)),
			NumReadonlyUnsignedAccounts: uint8(len(readOnlyUnsignedAccount)),
		},
		Accounts:           publicKeys[:len(publicKeys)-len(addressLookupTableWritable)-len(addressLookupTableReadonly)],
		RecentBlockHash:    param.RecentBlockhash,
		Instructions:       compiledInstructions,
		AddressLookupTable: compiledAddressLookupTable,
	}
}
