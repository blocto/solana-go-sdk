package assotokenprog

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type CreateAssociatedTokenAccountParam struct {
	Funder                 common.PublicKey
	Owner                  common.PublicKey
	Mint                   common.PublicKey
	AssociatedTokenAccount common.PublicKey
}

// CreateAssociatedTokenAccount is the only instruction in associated token program
func CreateAssociatedTokenAccount(param CreateAssociatedTokenAccountParam) types.Instruction {
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
		Data: []byte{},
	}
}
