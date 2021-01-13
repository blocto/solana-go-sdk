package tokenprog

import (
	"encoding/binary"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

func Transfer(srcPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) types.Instruction {
	instruction := []byte{0x03}
	number := make([]byte, 8)
	binary.LittleEndian.PutUint64(number, amount)
	data := make([]byte, 0, 9)
	data = append(data, instruction...)
	data = append(data, number...)

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: srcPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) > 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}
	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}
