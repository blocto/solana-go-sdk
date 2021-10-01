package sysprog

import (
	"github.com/near/borsh-go"
	"github.com/olegfomenko/solana-go-sdk/common"
	"github.com/olegfomenko/solana-go-sdk/types"
	"github.com/pkg/errors"
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
)

func CreateAccount(fromAccount, newAccount, owner common.PublicKey, initLamports, accountSpace uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Lamports    uint64
		Space       uint64
		Owner       common.PublicKey
	}{
		Instruction: InstructionCreateAccount,
		Lamports:    initLamports,
		Space:       accountSpace,
		Owner:       owner,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: fromAccount, IsSigner: true, IsWritable: true},
			{PubKey: newAccount, IsSigner: true, IsWritable: true},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}, nil
}

func Assign(accountPubkey, assignToProgramID common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction       Instruction
		AssignToProgramID common.PublicKey
	}{
		Instruction:       InstructionAssign,
		AssignToProgramID: assignToProgramID,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: accountPubkey, IsSigner: true, IsWritable: true},
		},
		Data: data,
	}, nil
}

func Transfer(from, to common.PublicKey, lamports uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionTransfer,
		Lamports:    lamports,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: from, IsSigner: true, IsWritable: true},
			{PubKey: to, IsSigner: false, IsWritable: true},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}, nil
}

func CreateAccountWithSeed(fromPubkey, newAccountPubkey, basePubkey, programID common.PublicKey, seed string, lamports, space uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Base        common.PublicKey
		Seed        string
		Lamports    uint64
		Space       uint64
		ProgramID   common.PublicKey
	}{
		Instruction: InstructionCreateAccountWithSeed,
		Base:        basePubkey,
		Seed:        seed,
		Lamports:    lamports,
		Space:       space,
		ProgramID:   programID,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 3)
	accounts = append(accounts,
		types.AccountMeta{PubKey: fromPubkey, IsSigner: true, IsWritable: true},
		types.AccountMeta{PubKey: newAccountPubkey, IsSigner: false, IsWritable: true},
	)
	if basePubkey != fromPubkey {
		accounts = append(accounts, types.AccountMeta{PubKey: basePubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func AdvanceNonceAccount(noncePubkey, authPubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionAdvanceNonceAccount,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: noncePubkey, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: authPubkey, IsSigner: true, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}, nil
}

func WithdrawNonceAccount(noncePubkey, authPubkey, toPubkey common.PublicKey, lamports uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionWithdrawNonceAccount,
		Lamports:    lamports,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: noncePubkey, IsSigner: false, IsWritable: true},
			{PubKey: toPubkey, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
			{PubKey: authPubkey, IsSigner: true, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}, nil
}

func InitializeNonceAccount(noncePubkey, authPubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Auth        common.PublicKey
	}{
		Instruction: InstructionInitializeNonceAccount,
		Auth:        authPubkey,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: noncePubkey, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}, nil
}

func AuthorizeNonceAccount(noncePubkey, oriAuthPubkey, newAuthPubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Auth        common.PublicKey
	}{
		Instruction: InstructionAuthorizeNonceAccount,
		Auth:        newAuthPubkey,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: noncePubkey, IsSigner: false, IsWritable: true},
			{PubKey: oriAuthPubkey, IsSigner: true, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}, nil
}

func Allocate(accountPubkey common.PublicKey, space uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Space       uint64
	}{
		Instruction: InstructionAllocate,
		Space:       space,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: accountPubkey, IsSigner: true, IsWritable: true},
		},
		Data: data,
	}, nil
}

func AllocateWithSeed(accountPubkey, basePubkey, programID common.PublicKey, seed string, space uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Base        common.PublicKey
		Seed        string
		Space       uint64
		ProgramID   common.PublicKey
	}{
		Instruction: InstructionAllocateWithSeed,
		Base:        basePubkey,
		Seed:        seed,
		Space:       space,
		ProgramID:   programID,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
			{PubKey: basePubkey, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}, nil
}
func AssignWithSeed(accountPubkey, assignToProgramID, basePubkey common.PublicKey, seed string) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction       Instruction
		Base              common.PublicKey
		Seed              string
		AssignToProgramID common.PublicKey
	}{
		Instruction:       InstructionAssignWithSeed,
		Base:              basePubkey,
		Seed:              seed,
		AssignToProgramID: assignToProgramID,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
			{PubKey: basePubkey, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}, nil
}

func TransferWithSeed(from, to, base, programID common.PublicKey, seed string, lamports uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Lamports    uint64
		Seed        string
		ProgramID   common.PublicKey
	}{
		Instruction: InstructionTransferWithSeed,
		Lamports:    lamports,
		Seed:        seed,
		ProgramID:   programID,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.SystemProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: from, IsSigner: false, IsWritable: true},
			{PubKey: base, IsSigner: true, IsWritable: false},
			{PubKey: to, IsSigner: false, IsWritable: true},
		},
		Data: data,
	}, nil
}
