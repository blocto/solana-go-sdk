package stakeprog

import (
	"github.com/near/borsh-go"
	"github.com/olegfomenko/solana-go-sdk/common"
	"github.com/olegfomenko/solana-go-sdk/types"
	"github.com/pkg/errors"
)

const AccountSize uint64 = 200

type Instruction uint32

const (
	InstructionInitialize Instruction = iota
	InstructionAuthorize
	InstructionDelegateStake
	InstructionSplit
	InstructionWithdraw
	InstructionDeactivate
	InstructionSetLockup
	InstructionMerge
	InstructionAuthorizeWithSeed
)

type StakeAuthorizationType uint32

const (
	StakeAuthorizationTypeStaker StakeAuthorizationType = iota
	StakeAuthorizationTypeWithdrawer
)

type Lockup struct {
	UnixTimestamp int64
	Epoch         uint64
	Cusodian      common.PublicKey
}

type LockupParam struct {
	UnixTimestamp *int64
	Epoch         *uint64
	Cusodian      *common.PublicKey
}

type Authorized struct {
	Staker     common.PublicKey
	Withdrawer common.PublicKey
}

func Initialize(initAccount common.PublicKey, auth Authorized, lockup Lockup) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Auth        Authorized
		Lockup      Lockup
	}{
		Instruction: InstructionInitialize,
		Auth:        auth,
		Lockup:      lockup,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: initAccount, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}, nil
}

func Authorize(stakePubkey, authPubkey, newAuthPubkey common.PublicKey, authType StakeAuthorizationType, custodianPubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction            Instruction
		NewAuthorized          common.PublicKey
		StakeAuthorizationType StakeAuthorizationType
	}{
		Instruction:            InstructionAuthorize,
		NewAuthorized:          newAuthPubkey,
		StakeAuthorizationType: authType,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 4)
	accounts = append(accounts,
		types.AccountMeta{PubKey: stakePubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
		types.AccountMeta{PubKey: authPubkey, IsSigner: true, IsWritable: false},
	)
	if custodianPubkey != (common.PublicKey{}) {
		accounts = append(accounts, types.AccountMeta{PubKey: custodianPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func DelegateStake(stakePubkey, authPubkey, votePubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionDelegateStake,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: stakePubkey, IsSigner: false, IsWritable: true},
			{PubKey: votePubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarStakeHistoryPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.StakeConfigPubkey, IsSigner: false, IsWritable: false},
			{PubKey: authPubkey, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}, nil
}

func Split(stakePubkey, authPubkey, splitStakePubkey common.PublicKey, lamports uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionSplit,
		Lamports:    lamports,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: stakePubkey, IsSigner: false, IsWritable: true},
			{PubKey: splitStakePubkey, IsSigner: false, IsWritable: true},
			{PubKey: authPubkey, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}, nil
}

func Withdraw(stakePubkey, authPubkey, toPubkey common.PublicKey, lamports uint64, custodianPubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionWithdraw,
		Lamports:    lamports,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 6)
	accounts = append(accounts,
		types.AccountMeta{PubKey: stakePubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: toPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
		types.AccountMeta{PubKey: common.SysVarStakeHistoryPubkey, IsSigner: false, IsWritable: false},
		types.AccountMeta{PubKey: authPubkey, IsSigner: true, IsWritable: false},
	)
	if custodianPubkey != (common.PublicKey{}) {
		accounts = append(accounts, types.AccountMeta{PubKey: custodianPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func Deactivate(stakePubkey, authPubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionDeactivate,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: stakePubkey, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
			{PubKey: authPubkey, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}, nil
}

func SetLockup(src, auth common.PublicKey, lockup LockupParam) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction   Instruction
		UnixTimestamp *int64
		Epoch         *uint64
		Cusodian      *common.PublicKey
	}{
		Instruction:   InstructionSetLockup,
		UnixTimestamp: lockup.UnixTimestamp,
		Epoch:         lockup.Epoch,
		Cusodian:      lockup.Cusodian,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: src, IsSigner: false, IsWritable: true},
			{PubKey: auth, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}, nil
}

func Merge(dest, src, auth common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionMerge,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: dest, IsSigner: false, IsWritable: true},
			{PubKey: src, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarStakeHistoryPubkey, IsSigner: false, IsWritable: false},
			{PubKey: auth, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}, nil
}

func AuthorizeWithSeed(
	stakePubkey common.PublicKey,
	authBasePubkey common.PublicKey,
	authSeed string,
	authOwnerPubkey common.PublicKey,
	newAuthPubkey common.PublicKey,
	authType StakeAuthorizationType,
	custodianPubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction            Instruction
		NewAuthorized          common.PublicKey
		StakeAuthorizationType StakeAuthorizationType
		AuthSeed               string
		AuthOwner              common.PublicKey
	}{
		Instruction:            InstructionAuthorizeWithSeed,
		NewAuthorized:          newAuthPubkey,
		StakeAuthorizationType: authType,
		AuthSeed:               authSeed,
		AuthOwner:              authOwnerPubkey,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 4)
	accounts = append(accounts,
		types.AccountMeta{PubKey: stakePubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authBasePubkey, IsSigner: true, IsWritable: false},
		types.AccountMeta{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
	)
	if custodianPubkey != (common.PublicKey{}) {
		accounts = append(accounts, types.AccountMeta{PubKey: custodianPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}
