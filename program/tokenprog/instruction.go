package tokenprog

import (
	"github.com/near/borsh-go"
	"github.com/olegfomenko/solana-go-sdk/common"
	"github.com/olegfomenko/solana-go-sdk/types"
	"github.com/pkg/errors"
)

type Instruction uint8

const (
	InstructionInitializeMint Instruction = iota
	InstructionInitializeAccount
	InstructionInitializeMultisig
	InstructionTransfer
	InstructionApprove
	InstructionRevoke
	InstructionSetAuthority
	InstructionMintTo
	InstructionBurn
	InstructionCloseAccount
	InstructionFreezeAccount
	InstructionThawAccount
	InstructionTransferChecked
	InstructionApproveChecked
	InstructionMintToChecked
	InstructionBurnChecked
	InstructionInitializeAccount2
	InstructionSyncNative
	InitializeAccount3
	InitializeMultisig2
	InitializeMint2
)

// InitializeMint init a mint, if you don't need to freeze, pass the empty pubKey common.PublicKey{}
func InitializeMint(decimals uint8, mint, mintAuthority common.PublicKey, freezeAuthority common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction     Instruction
		Decimals        uint8
		MintAuthority   common.PublicKey
		Option          bool
		FreezeAuthority common.PublicKey
	}{
		Instruction:     InstructionInitializeMint,
		Decimals:        decimals,
		MintAuthority:   mintAuthority,
		Option:          freezeAuthority != common.PublicKey{},
		FreezeAuthority: freezeAuthority,
	})

	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: mint, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}, nil
}

// InitializeAccount init a token account which can receive token
func InitializeAccount(accountPublicKey, mintPublicKey, ownerPublickey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionInitializeAccount,
	})

	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := []types.AccountMeta{
		{PubKey: accountPublicKey, IsSigner: false, IsWritable: true},
		{PubKey: mintPublicKey, IsSigner: false, IsWritable: false},
		{PubKey: ownerPublickey, IsSigner: false, IsWritable: false},
		{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
	}
	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func InitializeMultisig(authPubkey common.PublicKey, signerPubkeys []common.PublicKey, miniRequired uint8) (types.Instruction, error) {
	if len(signerPubkeys) < 1 {
		panic("minimum of signer is 1")
	}
	if len(signerPubkeys) > 11 {
		panic("maximum of signer is 11")
	}
	if miniRequired > uint8(len(signerPubkeys)) {
		panic("required number too big")
	}

	data, err := borsh.Serialize(struct {
		Instruction     Instruction
		MinimumRequired uint8
	}{
		Instruction:     InstructionInitializeMultisig,
		MinimumRequired: miniRequired,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 2+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: authPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func Transfer(srcPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionTransfer,
		Amount:      amount,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: srcPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}
	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func Approve(sourcePubkey, delegatePubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionApprove,
		Amount:      amount,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: sourcePubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: delegatePubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func Revoke(srcPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionRevoke,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 2+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: srcPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

type AuthorityType uint8

const (
	AuthorityTypeMintTokens AuthorityType = iota
	AuthorityTypeFreezeAccount
	AuthorityTypeAccountOwner
	AuthorityTypeCloseAccount
)

func SetAuthority(accountPubkey, newAuthPubkey common.PublicKey, authType AuthorityType, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction   Instruction
		AuthorityType AuthorityType
		Option        bool
		NewAuthPubkey common.PublicKey
	}{
		Instruction:   InstructionSetAuthority,
		AuthorityType: authType,
		Option:        newAuthPubkey != common.PublicKey{},
		NewAuthPubkey: newAuthPubkey,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 2+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func MintTo(mintPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionMintTo,
		Amount:      amount,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func Burn(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionBurn,
		Amount:      amount,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

//CloseAccount will close an account and transfer its all SOL to dest, only account's token balance is zero can be closed.
func CloseAccount(accountPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionCloseAccount,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func FreezeAccount(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionFreezeAccount,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func ThawAccount(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionThawAccount,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func TransferChecked(srcPubkey, destPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Amount      uint64
		Decimals    uint8
	}{
		Instruction: InstructionTransferChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 4+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: srcPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func ApproveChecked(sourcePubkey, mintPubkey, delegatePubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Amount      uint64
		Decimals    uint8
	}{
		Instruction: InstructionApproveChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 4+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: sourcePubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: delegatePubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func MintToChecked(mintPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Amount      uint64
		Decimals    uint8
	}{
		Instruction: InstructionMintToChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func BurnChecked(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Amount      uint64
		Decimals    uint8
	}{
		Instruction: InstructionBurnChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func InitializeAccount2(accountPubkey, mintPubkey, ownerPubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Owner       common.PublicKey
	}{
		Instruction: InstructionInitializeAccount2,
		Owner:       ownerPubkey,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
			{PubKey: mintPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}, nil
}

// SyncNative will update your wrapped SOL balance
func SyncNative(accountPubkey common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionSyncNative,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
		},
		Data: data,
	}, nil
}
