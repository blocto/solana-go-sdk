package address_lookup_table

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/bincode"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction uint32

const (
	InstructionCreateLookupTable Instruction = iota
	InstructionFreezeLookupTable
	InstructionExtendLookupTable
	InstructionDeactivateLookupTable
	InstructionCloseLookupTable
)

type CreateLookupTableParams struct {
	LookupTable common.PublicKey
	Authority   common.PublicKey
	Payer       common.PublicKey
	RecentSlot  uint64
	BumpSeed    uint8
}

func CreateLookupTable(params CreateLookupTableParams) types.Instruction {
	return types.Instruction{
		ProgramID: common.AddressLookupTableProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     params.LookupTable,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     params.Authority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     params.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: bincode.MustSerializeData(struct {
			Instruction Instruction
			RecentSlot  uint64
			BumpSeed    uint8
		}{
			Instruction: InstructionCreateLookupTable,
			RecentSlot:  params.RecentSlot,
			BumpSeed:    params.BumpSeed,
		}),
	}
}

type FreezeLookupTableParams struct {
	LookupTable common.PublicKey
	Authority   common.PublicKey
}

func FreezeLookupTable(params FreezeLookupTableParams) types.Instruction {
	return types.Instruction{
		ProgramID: common.AddressLookupTableProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     params.LookupTable,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     params.Authority,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: bincode.MustSerializeData(struct {
			Instruction Instruction
		}{
			Instruction: InstructionFreezeLookupTable,
		}),
	}
}

type ExtendLookupTableParams struct {
	LookupTable common.PublicKey
	Authority   common.PublicKey
	Payer       *common.PublicKey
	Addresses   []common.PublicKey
}

func ExtendLookupTable(params ExtendLookupTableParams) types.Instruction {
	accounts := make([]types.AccountMeta, 0, 4)
	accounts = append(accounts,
		types.AccountMeta{
			PubKey:     params.LookupTable,
			IsSigner:   false,
			IsWritable: true,
		},
		types.AccountMeta{
			PubKey:     params.Authority,
			IsSigner:   true,
			IsWritable: false,
		},
	)
	if params.Payer != nil {
		accounts = append(accounts,
			types.AccountMeta{
				PubKey:     *params.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			types.AccountMeta{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
		)
	}

	return types.Instruction{
		ProgramID: common.AddressLookupTableProgramID,
		Accounts:  accounts,
		Data: bincode.MustSerializeData(struct {
			Instruction  Instruction
			NewAddresses []common.PublicKey
		}{
			Instruction:  InstructionExtendLookupTable,
			NewAddresses: params.Addresses,
		}),
	}
}

type DeactivateLookupTableParams struct {
	LookupTable common.PublicKey
	Authority   common.PublicKey
}

func DeactivateLookupTable(params DeactivateLookupTableParams) types.Instruction {
	return types.Instruction{
		ProgramID: common.AddressLookupTableProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     params.LookupTable,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     params.Authority,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: bincode.MustSerializeData(struct {
			Instruction Instruction
		}{
			Instruction: InstructionDeactivateLookupTable,
		}),
	}
}

type CloseLookupTableParams struct {
	LookupTable common.PublicKey
	Authority   common.PublicKey
	Recipient   common.PublicKey
}

func CloseLookupTable(params CloseLookupTableParams) types.Instruction {
	return types.Instruction{
		ProgramID: common.AddressLookupTableProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     params.LookupTable,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     params.Authority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     params.Recipient,
				IsSigner:   false,
				IsWritable: true,
			},
		},
		Data: bincode.MustSerializeData(struct {
			Instruction Instruction
		}{
			Instruction: InstructionCloseLookupTable,
		}),
	}
}
