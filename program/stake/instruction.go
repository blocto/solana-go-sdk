package stake

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/bincode"
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

type LockupParam struct {
	UnixTimestamp *int64
	Epoch         *uint64
	Cusodian      *common.PublicKey
}

type Authorized struct {
	Staker     common.PublicKey
	Withdrawer common.PublicKey
}

type InitializeParam struct {
	Stake  common.PublicKey
	Auth   Authorized
	Lockup Lockup
}

func Initialize(param InitializeParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Auth        Authorized
		Lockup      Lockup
	}{
		Instruction: InstructionInitialize,
		Auth:        param.Auth,
		Lockup:      param.Lockup,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Stake, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}

type AuthorizeParam struct {
	Stake     common.PublicKey
	Auth      common.PublicKey
	NewAuth   common.PublicKey
	AuthType  StakeAuthorizationType
	Custodian *common.PublicKey
}

func Authorize(param AuthorizeParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction            Instruction
		NewAuthorized          common.PublicKey
		StakeAuthorizationType StakeAuthorizationType
	}{
		Instruction:            InstructionAuthorize,
		NewAuthorized:          param.NewAuth,
		StakeAuthorizationType: param.AuthType,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 4)
	accounts = append(accounts,
		types.AccountMeta{PubKey: param.Stake, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
		types.AccountMeta{PubKey: param.Auth, IsSigner: true, IsWritable: false},
	)
	if param.Custodian != nil {
		accounts = append(accounts, types.AccountMeta{PubKey: *param.Custodian, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type DelegateStakeParam struct {
	Stake common.PublicKey
	Auth  common.PublicKey
	Vote  common.PublicKey
}

func DelegateStake(param DelegateStakeParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
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
			{PubKey: param.Stake, IsSigner: false, IsWritable: true},
			{PubKey: param.Vote, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarStakeHistoryPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.StakeConfigPubkey, IsSigner: false, IsWritable: false},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

type SplitParam struct {
	Stake      common.PublicKey
	Auth       common.PublicKey
	SplitStake common.PublicKey
	Lamports   uint64
}

func Split(param SplitParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionSplit,
		Lamports:    param.Lamports,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Stake, IsSigner: false, IsWritable: true},
			{PubKey: param.SplitStake, IsSigner: false, IsWritable: true},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

type WithdrawParam struct {
	Stake     common.PublicKey
	Auth      common.PublicKey
	To        common.PublicKey
	Lamports  uint64
	Custodian *common.PublicKey
}

func Withdraw(param WithdrawParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Lamports    uint64
	}{
		Instruction: InstructionWithdraw,
		Lamports:    param.Lamports,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 6)
	accounts = append(accounts,
		types.AccountMeta{PubKey: param.Stake, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: param.To, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
		types.AccountMeta{PubKey: common.SysVarStakeHistoryPubkey, IsSigner: false, IsWritable: false},
		types.AccountMeta{PubKey: param.Auth, IsSigner: true, IsWritable: false},
	)
	if param.Custodian != nil {
		accounts = append(accounts, types.AccountMeta{PubKey: *param.Custodian, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type DeactivateParam struct {
	Stake common.PublicKey
	Auth  common.PublicKey
}

func Deactivate(param DeactivateParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
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
			{PubKey: param.Stake, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

type SetLockupParam struct {
	Stake  common.PublicKey
	Auth   common.PublicKey
	Lockup LockupParam
}

func SetLockup(param SetLockupParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction   Instruction
		UnixTimestamp *int64
		Epoch         *uint64
		Cusodian      *common.PublicKey
	}{
		Instruction:   InstructionSetLockup,
		UnixTimestamp: param.Lockup.UnixTimestamp,
		Epoch:         param.Lockup.Epoch,
		Cusodian:      param.Lockup.Cusodian,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Stake, IsSigner: false, IsWritable: true},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

type MergeParam struct {
	From common.PublicKey
	Auth common.PublicKey
	To   common.PublicKey
}

func Merge(param MergeParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
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
			{PubKey: param.To, IsSigner: false, IsWritable: true},
			{PubKey: param.From, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarStakeHistoryPubkey, IsSigner: false, IsWritable: false},
			{PubKey: param.Auth, IsSigner: true, IsWritable: false},
		},
		Data: data,
	}
}

type AuthorizeWithSeedParam struct {
	Stake     common.PublicKey
	AuthBase  common.PublicKey
	AuthSeed  string
	AuthOwner common.PublicKey
	NewAuth   common.PublicKey
	AuthType  StakeAuthorizationType
	Custodian *common.PublicKey
}

func AuthorizeWithSeed(param AuthorizeWithSeedParam) types.Instruction {

	data, err := bincode.SerializeData(struct {
		Instruction            Instruction
		NewAuthorized          common.PublicKey
		StakeAuthorizationType StakeAuthorizationType
		AuthSeed               string
		AuthOwner              common.PublicKey
	}{
		Instruction:            InstructionAuthorizeWithSeed,
		NewAuthorized:          param.NewAuth,
		StakeAuthorizationType: param.AuthType,
		AuthSeed:               param.AuthSeed,
		AuthOwner:              param.AuthOwner,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 4)
	accounts = append(accounts,
		types.AccountMeta{PubKey: param.Stake, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: param.AuthBase, IsSigner: true, IsWritable: false},
		types.AccountMeta{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
	)
	if param.Custodian != nil {
		accounts = append(accounts, types.AccountMeta{PubKey: *param.Custodian, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.StakeProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}
