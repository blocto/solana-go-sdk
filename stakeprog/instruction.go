package stakeprog

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
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

type Authorized struct {
	Staker     common.PublicKey
	Withdrawer common.PublicKey
}

func Initialize(initAccount common.PublicKey, auth Authorized, lockup Lockup) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction Instruction
		Auth        Authorized
		Lockup      Lockup
	}{
		Instruction: InstructionInitialize,
		Auth:        auth,
		Lockup:      lockup,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: initAccount, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}

func Authorize(stakePubkey, authPubkey, newAuthPubkey common.PublicKey, authType StakeAuthorizationType, custodianPubkey common.PublicKey) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction            Instruction
		NewAuthorized          common.PublicKey
		StakeAuthorizationType StakeAuthorizationType
	}{
		Instruction:            InstructionAuthorize,
		NewAuthorized:          newAuthPubkey,
		StakeAuthorizationType: authType,
	})
	if err != nil {
		panic(err)
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
	}
}

func DelegateStake(stakePubkey, authPubkey, votePubkey common.PublicKey) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionDelegateStake,
	})
	if err != nil {
		panic(err)
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
	}
}

func Split(stakePubkey, authPubkey, splitStakePubkey common.PublicKey, lamports uint64) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionSplit,
		Lamports:    lamports,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: stakePubkey, IsSigner: false, IsWritable: true},
			{PubKey: splitStakePubkey, IsSigner: false, IsWritable: true},
			{PubKey: authPubkey, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

func Withdraw(stakePubkey, authPubkey, toPubkey common.PublicKey, lamports uint64, custodianPubkey common.PublicKey) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionWithdraw,
		Lamports:    lamports,
	})
	if err != nil {
		panic(err)
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
	}
}

func Deactivate(stakePubkey, authPubkey common.PublicKey) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionDeactivate,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: stakePubkey, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
			{PubKey: authPubkey, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

func SetLockup() types.Instruction {
	panic("not implement yet")
}

func Merge(dest, src, auth common.PublicKey) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionMerge,
	})
	if err != nil {
		panic(err)
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
	}
}

func AuthorizeWithSeed() types.Instruction {
	panic("not implement yet")
}
