package tokenprog

import (
	"github.com/portto/solana-go-sdk/common"
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
)

// InitializeMint init a mint, if you don't need to freeze, pass the empty pubKey common.PublicKey{}
func InitializeMint(decimals uint8, mint, mintAuthority common.PublicKey, freezeAuthority common.PublicKey) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction     Instruction
		Decimals        uint8
		MintAuthority   common.PublicKey
		Option          bool
		FreezeAuthority common.PublicKey
	}{
		Instruction:     InstructionInitializeMint,
		Decimals:        decimals,
		MintAuthority:   mintAuthority,
		Option:          freezeAuthority == common.PublicKey{},
		FreezeAuthority: freezeAuthority,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: mint, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}

func InitializeAccount(accountPublicKey, mintPublicKey, ownerPublickey common.PublicKey) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction Instruction
	}{
		Instruction: InstructionInitializeAccount,
	})
	if err != nil {
		panic(err)
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
	}
}

func InitializeMultisig() types.Instruction {
	panic("not implement yet")
}

func Transfer(srcPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionTransfer,
		Amount:      amount,
	})
	if err != nil {
		panic(err)
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
	}
}

func Approve() types.Instruction {
	panic("not implement yet")
}

func Revoke() types.Instruction {
	panic("not implement yet")
}

func SetAuthority() types.Instruction {
	panic("not implement yet")
}

func MintTo(mintPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) types.Instruction {
	data, err := common.SerializeData(struct {
		Instruction Instruction
		Amount      uint64
	}{
		Instruction: InstructionMintTo,
		Amount:      amount,
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

func Burn(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) types.Instruction {
	data, err := common.SerializeData(struct {
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

func CloseAccount() types.Instruction {
	panic("not implement yet")
}

func FreezeAccount() types.Instruction {
	panic("not implement yet")
}

func ThawAccount() types.Instruction {
	panic("not implement yet")
}

func TransferChecked(srcPubkey, destPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := common.SerializeData(struct {
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

func ApproveChecked() types.Instruction {
	panic("not implement yet")
}

func MintToChecked(mintPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := common.SerializeData(struct {
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
	data, err := common.SerializeData(struct {
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

func InitializeAccount2() types.Instruction {
	panic("not implement yet")
}
