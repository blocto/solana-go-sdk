package assotokenprog

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

// CreateAssociatedTokenAccount is the only instruction in associated token program
func CreateAssociatedTokenAccount(funder, wallet, tokenMint common.PublicKey) (types.Instruction, error) {
	assosiatedAccount, _, _ := common.FindAssociatedTokenAddress(wallet, tokenMint)
	return types.Instruction{
		ProgramID: common.SPLAssociatedTokenAccountProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: funder, IsSigner: true, IsWritable: true},
			{PubKey: assosiatedAccount, IsSigner: false, IsWritable: true},
			{PubKey: wallet, IsSigner: false, IsWritable: false},
			{PubKey: tokenMint, IsSigner: false, IsWritable: false},
			{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: []byte{},
	}, nil
}
