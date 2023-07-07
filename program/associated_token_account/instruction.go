package associated_token_account

import (
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/near/borsh-go"
)

type Instruction borsh.Enum

const (
	InstructionCreate Instruction = iota
	InstructionCreateIdempotent
	InstructionRecoverNested
)

type CreateAssociatedTokenAccountParam struct {
	Funder                 common.PublicKey
	Owner                  common.PublicKey
	Mint                   common.PublicKey
	AssociatedTokenAccount common.PublicKey
}

// CreateAssociatedTokenAccount creates an associated token account for the given wallet address and token mint. Return an error if the account exists.
// Deprecated: please use `Create`
func CreateAssociatedTokenAccount(param CreateAssociatedTokenAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionCreate,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SPLAssociatedTokenAccountProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Funder, IsSigner: true, IsWritable: true},
			{PubKey: param.AssociatedTokenAccount, IsSigner: false, IsWritable: true},
			{PubKey: param.Owner, IsSigner: false, IsWritable: false},
			{PubKey: param.Mint, IsSigner: false, IsWritable: false},
			{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}

type CreateParam struct {
	Funder                 common.PublicKey
	Owner                  common.PublicKey
	Mint                   common.PublicKey
	AssociatedTokenAccount common.PublicKey
}

// Create creates an associated token account for the given wallet address and token mint. Return an error if the account exists.
func Create(param CreateParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionCreate,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SPLAssociatedTokenAccountProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Funder, IsSigner: true, IsWritable: true},
			{PubKey: param.AssociatedTokenAccount, IsSigner: false, IsWritable: true},
			{PubKey: param.Owner, IsSigner: false, IsWritable: false},
			{PubKey: param.Mint, IsSigner: false, IsWritable: false},
			{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}

type CreateIdempotentParam struct {
	Funder                 common.PublicKey
	Owner                  common.PublicKey
	Mint                   common.PublicKey
	AssociatedTokenAccount common.PublicKey
}

// CreateIdempotent creates an associated token account for the given wallet address and token mint,
// if it doesn't already exist. Returns an error if the account exists, but with a different owner.
func CreateIdempotent(param CreateIdempotentParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionCreateIdempotent,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SPLAssociatedTokenAccountProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Funder, IsSigner: true, IsWritable: true},
			{PubKey: param.AssociatedTokenAccount, IsSigner: false, IsWritable: true},
			{PubKey: param.Owner, IsSigner: false, IsWritable: false},
			{PubKey: param.Mint, IsSigner: false, IsWritable: false},
			{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}

type RecoverNestedParam struct {
	Owner                             common.PublicKey
	OwnerMint                         common.PublicKey
	OwnerAssociatedTokenAccount       common.PublicKey
	NestedMint                        common.PublicKey
	NestedMintAssociatedTokenAccount  common.PublicKey
	DestinationAssociatedTokenAccount common.PublicKey
}

// RecoverNested transfers from and closes a nested associated token account: an associated token account owned by an associated token account.
func RecoverNested(param RecoverNestedParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionRecoverNested,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SPLAssociatedTokenAccountProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.NestedMintAssociatedTokenAccount, IsSigner: false, IsWritable: true},
			{PubKey: param.NestedMint, IsSigner: false, IsWritable: false},
			{PubKey: param.DestinationAssociatedTokenAccount, IsSigner: false, IsWritable: true},
			{PubKey: param.OwnerAssociatedTokenAccount, IsSigner: false, IsWritable: true},
			{PubKey: param.OwnerMint, IsSigner: false, IsWritable: false},
			{PubKey: param.Owner, IsSigner: true, IsWritable: true},
			{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}
