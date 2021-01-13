package tokenprog

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

func InitializeAccount(accountPublicKey, mintPublicKey, ownerPublickey common.PublicKey) types.Instruction {
	data := []byte{0x01}
	accounts := []types.AccountMeta{
		{PubKey: accountPublicKey, IsSigner: false, IsWritable: true},
		{PubKey: mintPublicKey, IsSigner: false, IsWritable: false},
		{PubKey: ownerPublickey, IsSigner: false, IsWritable: false},
		{PubKey: common.PublicKeyFromString("SysvarRent111111111111111111111111111111111"), IsSigner: false, IsWritable: false},
	}
	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}
