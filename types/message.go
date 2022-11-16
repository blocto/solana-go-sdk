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
	Version             MessageVersion
	Header              MessageHeader
	Accounts            []common.PublicKey
	RecentBlockHash     string
	Instructions        []CompiledInstruction
	AddressLookupTables []CompiledAddressLookupTable
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

		validAddressLookupCount := 0
		accountLookupTableSerializedData := []byte{}

		if len(m.AddressLookupTables) > 0 {
			for _, addressLookupTable := range m.AddressLookupTables {
				if len(addressLookupTable.WritableIndexes) != 0 || len(addressLookupTable.ReadonlyIndexes) != 0 {
					accountLookupTableSerializedData = append(accountLookupTableSerializedData, addressLookupTable.AccountKey.Bytes()...)
					accountLookupTableSerializedData = append(accountLookupTableSerializedData, bincode.UintToVarLenBytes(uint64(len(addressLookupTable.WritableIndexes)))...)
					accountLookupTableSerializedData = append(accountLookupTableSerializedData, addressLookupTable.WritableIndexes...)
					accountLookupTableSerializedData = append(accountLookupTableSerializedData, bincode.UintToVarLenBytes(uint64(len(addressLookupTable.ReadonlyIndexes)))...)
					accountLookupTableSerializedData = append(accountLookupTableSerializedData, addressLookupTable.ReadonlyIndexes...)
					validAddressLookupCount++
				}
			}

		}

		b = append(b, bincode.UintToVarLenBytes(uint64(validAddressLookupCount))...)
		b = append(b, accountLookupTableSerializedData...)
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
		if err != nil || t > 255 {
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

	compiledAddressLookupTables := []CompiledAddressLookupTable{}
	if version == MessageVersionV0 {
		addressLookupTableCount, err := parseUvarint(&messageData)
		if err != nil {
			return Message{}, fmt.Errorf("parse instruction count error: %v", err)
		}

		for i := uint64(0); i < addressLookupTableCount; i++ {
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

			compiledAddressLookupTables = append(
				compiledAddressLookupTables,
				CompiledAddressLookupTable{
					AccountKey:      addressLookupTablePubkey,
					WritableIndexes: writableAccountIdxList,
					ReadonlyIndexes: readOnlyAccountIdxList,
				},
			)
		}
	}

	return Message{
		Version: version,
		Header: MessageHeader{
			NumRequireSignatures:        numRequireSignatures,
			NumReadonlySignedAccounts:   numReadonlySignedAccounts,
			NumReadonlyUnsignedAccounts: numReadonlyUnsignedAccounts,
		},
		Accounts:            accounts,
		RecentBlockHash:     blockHash,
		Instructions:        instructions,
		AddressLookupTables: compiledAddressLookupTables,
	}, nil
}

func MustMessageDeserialize(messageData []byte) Message {
	message, err := MessageDeserialize(messageData)
	if err != nil {
		panic(err)
	}
	return message
}

type AddressLookupTableAccount struct {
	Key       common.PublicKey
	Addresses []common.PublicKey
}

type NewMessageParam struct {
	FeePayer        common.PublicKey
	Instructions    []Instruction
	RecentBlockhash string
	// v0 transaction
	AddressLookupTableAccounts []AddressLookupTableAccount
}

type CompiledKeys struct {
	Payer      *common.PublicKey
	KeyMetaMap map[common.PublicKey]CompiledKeyMeta
}

type CompiledKeyMeta struct {
	IsSigner   bool
	IsWritable bool
	IsInvoked  bool
}

func NewCompiledKeys(instructions []Instruction, payer *common.PublicKey) CompiledKeys {
	m := map[common.PublicKey]CompiledKeyMeta{}

	for _, instruction := range instructions {
		// compile program
		v := m[instruction.ProgramID]
		v.IsInvoked = true
		m[instruction.ProgramID] = v

		// compile accounts
		for i := 0; i < len(instruction.Accounts); i++ {
			account := instruction.Accounts[i]

			v := m[account.PubKey]
			v.IsSigner = v.IsSigner || account.IsSigner
			v.IsWritable = v.IsWritable || account.IsWritable
			m[account.PubKey] = v
		}
	}

	if payer != nil && *payer != (common.PublicKey{}) {
		v := m[*payer]
		v.IsSigner = true
		v.IsWritable = true
		m[*payer] = v
	}

	return CompiledKeys{
		Payer:      payer,
		KeyMetaMap: m,
	}
}

func NewMessage(param NewMessageParam) Message {
	writableSignedAccount := []common.PublicKey{}
	readOnlySignedAccount := []common.PublicKey{}
	writableUnsignedAccount := []common.PublicKey{}
	readOnlyUnsignedAccount := []common.PublicKey{}

	addressLookupTableAccountCount := len(param.AddressLookupTableAccounts)
	addressLookupTableWritable := make([][]common.PublicKey, addressLookupTableAccountCount)
	addressLookupTableWritableIdx := make([][]uint8, addressLookupTableAccountCount)
	addressLookupTableReadonly := make([][]common.PublicKey, addressLookupTableAccountCount)
	addressLookupTableReadonlyIdx := make([][]uint8, addressLookupTableAccountCount)

	addressLookupTableMaps := make([]map[common.PublicKey]uint8, 0, addressLookupTableAccountCount)
	for _, addressLookupTableAccount := range param.AddressLookupTableAccounts {
		m := map[common.PublicKey]uint8{}
		for i, address := range addressLookupTableAccount.Addresses {
			m[address] = uint8(i)
		}
		addressLookupTableMaps = append(addressLookupTableMaps, m)
	}

	compiledKeys := NewCompiledKeys(param.Instructions, &param.FeePayer)
	allKeys := make([]common.PublicKey, 0, len(compiledKeys.KeyMetaMap))
	for key := range compiledKeys.KeyMetaMap {
		allKeys = append(allKeys, key)
	}
	sort.Slice(allKeys, func(i, j int) bool {
		return bytes.Compare(allKeys[i].Bytes(), allKeys[j].Bytes()) < 0
	})

NEXT_ACCOUNT:
	for _, key := range allKeys {
		meta := compiledKeys.KeyMetaMap[key]
		if key == param.FeePayer {
			continue NEXT_ACCOUNT
		}
		if meta.IsSigner {
			if meta.IsWritable {
				writableSignedAccount = append(writableSignedAccount, key)
			} else {
				readOnlySignedAccount = append(readOnlySignedAccount, key)
			}
		} else {
			if meta.IsWritable {
				for n, addressLookupTableMap := range addressLookupTableMaps {
					idx, exist := addressLookupTableMap[key]
					if exist && !meta.IsInvoked {
						addressLookupTableWritable[n] = append(addressLookupTableWritable[n], key)
						addressLookupTableWritableIdx[n] = append(addressLookupTableWritableIdx[n], idx)
						continue NEXT_ACCOUNT
					}
				}
				// if not found in address lookup table
				writableUnsignedAccount = append(writableUnsignedAccount, key)
			} else {
				for n, addressLookupTableMap := range addressLookupTableMaps {
					idx, exist := addressLookupTableMap[key]
					if exist && !meta.IsInvoked {
						addressLookupTableReadonly[n] = append(addressLookupTableReadonly[n], key)
						addressLookupTableReadonlyIdx[n] = append(addressLookupTableReadonlyIdx[n], idx)
						continue NEXT_ACCOUNT
					}
				}
				// if not found in address lookup table
				readOnlyUnsignedAccount = append(readOnlyUnsignedAccount, key)
			}
		}
	}

	// add fee payer
	writableSignedAccount = append([]common.PublicKey{param.FeePayer}, writableSignedAccount...)

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

	compiledAddressLookupTables := []CompiledAddressLookupTable{}
	lookupAddressCount := 0
	for i := 0; i < addressLookupTableAccountCount; i++ {
		publicKeys = append(publicKeys, addressLookupTableWritable[i]...)
		lookupAddressCount += len(addressLookupTableWritable[i])
	}
	for i := 0; i < addressLookupTableAccountCount; i++ {
		publicKeys = append(publicKeys, addressLookupTableReadonly[i]...)
		lookupAddressCount += len(addressLookupTableReadonly[i])

		if len(addressLookupTableWritable[i]) > 0 || len(addressLookupTableReadonly[i]) > 0 {
			compiledAddressLookupTables = append(compiledAddressLookupTables, CompiledAddressLookupTable{
				AccountKey:      param.AddressLookupTableAccounts[i].Key,
				WritableIndexes: addressLookupTableWritableIdx[i],
				ReadonlyIndexes: addressLookupTableReadonlyIdx[i],
			})
		}
	}

	var version MessageVersion = MessageVersionLegacy
	if addressLookupTableAccountCount > 0 {
		version = MessageVersionV0
	}

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
		Version: version,
		Header: MessageHeader{
			NumRequireSignatures:        uint8(len(writableSignedAccount) + len(readOnlySignedAccount)),
			NumReadonlySignedAccounts:   uint8(len(readOnlySignedAccount)),
			NumReadonlyUnsignedAccounts: uint8(len(readOnlyUnsignedAccount)),
		},
		Accounts:            publicKeys[:len(publicKeys)-lookupAddressCount],
		RecentBlockHash:     param.RecentBlockhash,
		Instructions:        compiledInstructions,
		AddressLookupTables: compiledAddressLookupTables,
	}
}
