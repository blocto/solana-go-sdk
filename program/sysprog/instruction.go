package sysprog

import (
	"github.com/OldSmokeGun/solana-go-sdk/common"
	"github.com/OldSmokeGun/solana-go-sdk/pkg/bincode"
	"github.com/OldSmokeGun/solana-go-sdk/types"
)

type Instruction uint32

const (
	InstructionCreateAccount Instruction = iota
	InstructionAssign
	InstructionTransfer
	InstructionCreateAccountWithSeed
	InstructionAdvanceNonceAccount
	InstructionWithdrawNonceAccount
	InstructionInitializeNonceAccount
	InstructionAuthorizeNonceAccount
	InstructionAllocate
	InstructionAllocateWithSeed
	InstructionAssignWithSeed
	InstructionTransferWithSeed
	InstructionUpgradeNonceAccount
)

type CreateAccountParam struct {
	From     common.PublicKey
	New      common.PublicKey
	Owner    common.PublicKey
	Lamports uint64
	Space    uint64
}

func CreateAccount(param CreateAccountParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Lamports    uint64
		Space       uint64
		Owner       common.PublicKey
	}{
		Instruction: InstructionCreateAccount,
		Lamports:    param.Lamports,
		Space:       param.Space,
		Owner:       param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.From, IsSigner: true, IsWritable: true},
			{PubKey: param.New, IsSigner: true, IsWritable: true},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type AssignParam struct {
	From  common.PublicKey
	Owner common.PublicKey
}

func Assign(param AssignParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction       Instruction
		AssignToProgramID common.PublicKey
	}{
		Instruction:       InstructionAssign,
		AssignToProgramID: param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.From, IsSigner: true, IsWritable: true},
		},
		Data: data,
	}
}

type TransferParam struct {
	From   common.PublicKey
	To     common.PublicKey
	Amount uint64
}

func Transfer(param TransferParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionTransfer,
		Lamports:    param.Amount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.From, IsSigner: true, IsWritable: true},
			{PubKey: param.To, IsSigner: false, IsWritable: true},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type CreateAccountWithSeedParam struct {
	From     common.PublicKey
	New      common.PublicKey
	Base     common.PublicKey
	Owner    common.PublicKey
	Seed     string
	Lamports uint64
	Space    uint64
}

func CreateAccountWithSeed(param CreateAccountWithSeedParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Base        common.PublicKey
		Seed        string
		Lamports    uint64
		Space       uint64
		ProgramID   common.PublicKey
	}{
		Instruction: InstructionCreateAccountWithSeed,
		Base:        param.Base,
		Seed:        param.Seed,
		Lamports:    param.Lamports,
		Space:       param.Space,
		ProgramID:   param.Owner,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3)
	accounts = append(accounts,
		types.AccountMeta{PubKey: param.From, IsSigner: true, IsWritable: true},
		types.AccountMeta{PubKey: param.New, IsSigner: false, IsWritable: true},
	)
	if param.Base != param.From {
		accounts = append(accounts, types.AccountMeta{PubKey: param.Base, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type AdvanceNonceAccountParam struct {
	Nonce common.PublicKey
	Auth  common.PublicKey
}

func AdvanceNonceAccount(param AdvanceNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionAdvanceNonceAccount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.Nonce, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type WithdrawNonceAccountParam struct {
	Nonce  common.PublicKey
	Auth   common.PublicKey
	To     common.PublicKey
	Amount uint64
}

func WithdrawNonceAccount(param WithdrawNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionWithdrawNonceAccount,
		Lamports:    param.Amount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.Nonce, IsSigner: false, IsWritable: true},
			{PubKey: param.To, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type InitializeNonceAccountParam struct {
	Nonce common.PublicKey
	Auth  common.PublicKey
}

func InitializeNonceAccount(param InitializeNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Auth        common.PublicKey
	}{
		Instruction: InstructionInitializeNonceAccount,
		Auth:        param.Auth,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.Nonce, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type AuthorizeNonceAccountParam struct {
	Nonce   common.PublicKey
	Auth    common.PublicKey
	NewAuth common.PublicKey
}

func AuthorizeNonceAccount(param AuthorizeNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Auth        common.PublicKey
	}{
		Instruction: InstructionAuthorizeNonceAccount,
		Auth:        param.NewAuth,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.Nonce, IsSigner: false, IsWritable: true},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}

type AllocateParam struct {
	Account common.PublicKey
	Space   uint64
}

func Allocate(param AllocateParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Space       uint64
	}{
		Instruction: InstructionAllocate,
		Space:       param.Space,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Account, IsSigner: true, IsWritable: true},
		},
		Data: data,
	}
}

type AllocateWithSeedParam struct {
	Account common.PublicKey
	Base    common.PublicKey
	Owner   common.PublicKey
	Seed    string
	Space   uint64
}

func AllocateWithSeed(param AllocateWithSeedParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Base        common.PublicKey
		Seed        string
		Space       uint64
		ProgramID   common.PublicKey
	}{
		Instruction: InstructionAllocateWithSeed,
		Base:        param.Base,
		Seed:        param.Seed,
		Space:       param.Space,
		ProgramID:   param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Account, IsSigner: false, IsWritable: true},
			{PubKey: param.Base, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

type AssignWithSeedParam struct {
	Account common.PublicKey
	Owner   common.PublicKey
	Base    common.PublicKey
	Seed    string
}

func AssignWithSeed(param AssignWithSeedParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction       Instruction
		Base              common.PublicKey
		Seed              string
		AssignToProgramID common.PublicKey
	}{
		Instruction:       InstructionAssignWithSeed,
		Base:              param.Base,
		Seed:              param.Seed,
		AssignToProgramID: param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Account, IsSigner: false, IsWritable: true},
			{PubKey: param.Base, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

type TransferWithSeedParam struct {
	From   common.PublicKey
	To     common.PublicKey
	Base   common.PublicKey
	Owner  common.PublicKey
	Seed   string
	Amount uint64
}

func TransferWithSeed(param TransferWithSeedParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Lamports    uint64
		Seed        string
		ProgramID   common.PublicKey
	}{
		Instruction: InstructionTransferWithSeed,
		Lamports:    param.Amount,
		Seed:        param.Seed,
		ProgramID:   param.Owner,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.From, IsSigner: false, IsWritable: true},
			{PubKey: param.Base, IsSigner: true, IsWritable: false},
			{PubKey: param.To, IsSigner: false, IsWritable: true},
		},
		Data: data,
	}
}

type UpgradeNonceAccountParam struct {
	NonceAccountPubkey common.PublicKey
}

func UpgradeNonceAccount(param UpgradeNonceAccountParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionUpgradeNonceAccount,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.NonceAccountPubkey, IsSigner: false, IsWritable: true},
		},
		Data: data,
	}
}
