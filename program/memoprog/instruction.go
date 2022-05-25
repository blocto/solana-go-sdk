package memoprog

import (
	"github.com/OldSmokeGun/solana-go-sdk/common"
	"github.com/OldSmokeGun/solana-go-sdk/types"
)

type BuildMemoParam struct {
	SignerPubkeys []common.PublicKey
	Memo          []byte
}

func BuildMemo(param BuildMemoParam) types.Instruction {
	accounts := make([]types.AccountMeta, 0, len(param.SignerPubkeys))
	for _, signerPubkey := range param.SignerPubkeys {
		accounts = append(accounts, types.AccountMeta{
			PubKey:     signerPubkey,
			IsSigner:   true,
			IsWritable: false,
		})
	}
	return types.Instruction{
		ProgramID: common.MemoProgramID,
		Accounts:  accounts,
		Data:      param.Memo,
	}
}
