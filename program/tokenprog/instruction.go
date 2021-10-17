package tokenprog

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/bincode"
	"github.com/portto/solana-go-sdk/types"
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

type InitializeMintParam struct {
	Decimals   uint8
	Mint       common.PublicKey
	MintAuth   common.PublicKey
	FreezeAuth *common.PublicKey
}

// InitializeMint init a mint, if you don't need to freeze, pass the empty pubKey common.PublicKey{}
func InitializeMint(param InitializeMintParam) types.Instruction {
	var freezeAuth common.PublicKey
	if param.FreezeAuth != nil {
		freezeAuth = *param.FreezeAuth
	}
	data, err := bincode.SerializeData(struct {
		Instruction     Instruction
		Decimals        uint8
		MintAuthority   common.PublicKey
		Option          bool
		FreezeAuthority common.PublicKey
	}{
		Instruction:     InstructionInitializeMint,
		Decimals:        param.Decimals,
		MintAuthority:   param.MintAuth,
		Option:          param.FreezeAuth != nil,
		FreezeAuthority: freezeAuth,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: param.Mint, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}

type InitializeAccountParam struct {
	Account common.PublicKey
	Mint    common.PublicKey
	Owner   common.PublicKey
}

// InitializeAccount init a token account which can receive token
func InitializeAccount(param InitializeAccountParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionInitializeAccount,
	})
	if err != nil {
		panic(err)
	}

	accounts := []types.AccountMeta{
		{PubKey: param.Account, IsSigner: false, IsWritable: true},
		{PubKey: param.Mint, IsSigner: false, IsWritable: false},
		{PubKey: param.Owner, IsSigner: false, IsWritable: false},
		{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
	}
	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type InitializeMultisigParam struct {
	Account     common.PublicKey
	Signers     []common.PublicKey
	MinRequired uint8
}

func InitializeMultisig(param InitializeMultisigParam) types.Instruction {
	if len(param.Signers) < 1 {
		panic("minimum of signer is 1")
	}
	if len(param.Signers) > 11 {
		panic("maximum of signer is 11")
	}
	if param.MinRequired > uint8(len(param.Signers)) {
		panic("required number too big")
	}

	data, err := bincode.SerializeData(struct {
		Instruction     Instruction
		MinimumRequired uint8
	}{
		Instruction:     InstructionInitializeMultisig,
		MinimumRequired: param.MinRequired,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 2+len(param.Signers))
	accounts = append(accounts,
		types.AccountMeta{PubKey: param.Account, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
	)
	for _, signerPubkey := range param.Signers {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type TransferParam struct {
	From    common.PublicKey
	To      common.PublicKey
	Auth    common.PublicKey
	Signers []common.PublicKey
	Amount  uint64
}

func Transfer(param TransferParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionTransfer,
		Amount:      param.Amount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(param.Signers))
	accounts = append(accounts, types.AccountMeta{PubKey: param.From, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: param.To, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: param.Auth, IsSigner: len(param.Signers) == 0, IsWritable: false})
	for _, signerPubkey := range param.Signers {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}
	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type ApproveParam struct {
	From    common.PublicKey
	To      common.PublicKey
	Auth    common.PublicKey
	Signers []common.PublicKey
	Amount  uint64
}

func Approve(param ApproveParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionApprove,
		Amount:      param.Amount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(param.Signers))
	accounts = append(accounts, types.AccountMeta{PubKey: param.From, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: param.To, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: param.Auth, IsSigner: len(param.Signers) == 0, IsWritable: false})
	for _, signerPubkey := range param.Signers {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type RevokeParam struct {
	From    common.PublicKey
	Auth    common.PublicKey
	Signers []common.PublicKey
}

func Revoke(param RevokeParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionRevoke,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 2+len(param.Signers))
	accounts = append(accounts,
		types.AccountMeta{PubKey: param.From, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: param.Auth, IsSigner: len(param.Signers) == 0, IsWritable: false},
	)
	for _, signerPubkey := range param.Signers {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type AuthorityType uint8

const (
	AuthorityTypeMintTokens AuthorityType = iota
	AuthorityTypeFreezeAccount
	AuthorityTypeAccountOwner
	AuthorityTypeCloseAccount
)

type SetAuthorityParam struct {
	Account  common.PublicKey
	NewAuth  *common.PublicKey
	AuthType AuthorityType
	Auth     common.PublicKey
	Signers  []common.PublicKey
}

func SetAuthority(param SetAuthorityParam) types.Instruction {
	var newAuth common.PublicKey
	if param.NewAuth != nil {
		newAuth = *param.NewAuth
	}
	data, err := bincode.SerializeData(struct {
		Instruction   Instruction
		AuthorityType AuthorityType
		Option        bool
		NewAuthPubkey common.PublicKey
	}{
		Instruction:   InstructionSetAuthority,
		AuthorityType: param.AuthType,
		Option:        param.NewAuth != nil,
		NewAuthPubkey: newAuth,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 2+len(param.Signers))
	accounts = append(accounts,
		types.AccountMeta{PubKey: param.Account, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: param.Auth, IsSigner: len(param.Signers) == 0, IsWritable: false},
	)
	for _, signerPubkey := range param.Signers {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

type MintToParam struct {
	Mint    common.PublicKey
	To      common.PublicKey
	Auth    common.PublicKey
	Signers []common.PublicKey
	Amount  uint64
}

func MintTo(param MintToParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionMintTo,
		Amount:      param.Amount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(param.Signers))
	accounts = append(accounts,
		types.AccountMeta{PubKey: param.Mint, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: param.To, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: param.Auth, IsSigner: len(param.Signers) == 0, IsWritable: false},
	)
	for _, signerPubkey := range param.Signers {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func Burn(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionBurn,
		Amount:      amount,
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
	}
}

// Close an account and transfer its all SOL to dest, only account's token balance is zero can be closed.
func CloseAccount(accountPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionCloseAccount,
	})
	if err != nil {
		panic(err)
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
	}
}

func FreezeAccount(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionFreezeAccount,
	})
	if err != nil {
		panic(err)
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
	}
}

func ThawAccount(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionThawAccount,
	})
	if err != nil {
		panic(err)
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
	}
}

func TransferChecked(srcPubkey, destPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Amount      uint64
		Decimals    uint8
	}{
		Instruction: InstructionTransferChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		panic(err)
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
	}
}

func ApproveChecked(sourcePubkey, mintPubkey, delegatePubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Amount      uint64
		Decimals    uint8
	}{
		Instruction: InstructionApproveChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		panic(err)
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
	}
}

func MintToChecked(mintPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Amount      uint64
		Decimals    uint8
	}{
		Instruction: InstructionMintToChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		panic(err)
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
	}
}

func BurnChecked(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := bincode.SerializeData(struct {
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
	}
}

func InitializeAccount2(accountPubkey, mintPubkey, ownerPubkey common.PublicKey) types.Instruction {
	data, err := bincode.SerializeData(struct {
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
	}
}

// SyncNative will update your wrapped SOL balance
func SyncNative(accountPubkey common.PublicKey) types.Instruction {
	data, err := bincode.SerializeData(struct {
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
	}
}
