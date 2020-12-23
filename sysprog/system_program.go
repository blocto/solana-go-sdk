package sysprog

import (
	"encoding/binary"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

var SystemProgramID = common.PublicKeyFromString("11111111111111111111111111111111")

type CreateAccountParam struct {
	FromPublicKey    common.PublicKey
	NewAccountPubkey common.PublicKey
	Lamports         uint64
	Space            uint64
	ProgramId        common.PublicKey
}

func CreateAccount(param CreateAccountParam) types.Instruction {
	instruction := make([]byte, 4)
	binary.LittleEndian.PutUint32(instruction, 0)
	lamports := make([]byte, 8)
	binary.LittleEndian.PutUint64(lamports, param.Lamports)
	space := make([]byte, 8)
	binary.LittleEndian.PutUint64(space, param.Space)

	data := make([]byte, 0, 52)
	data = append(data, instruction...)
	data = append(data, lamports...)
	data = append(data, space...)
	data = append(data, param.ProgramId[:]...)

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.FromPublicKey, IsSigner: true, IsWritable: true},
			{PubKey: param.NewAccountPubkey, IsSigner: true, IsWritable: true},
		},
		ProgramID: SystemProgramID,
		Data:      data,
	}
}

func Transfer(from, to common.PublicKey, number uint64) types.Instruction {
	instruction := make([]byte, 4)
	binary.LittleEndian.PutUint32(instruction, 2)
	lamports := make([]byte, 8)
	binary.LittleEndian.PutUint64(lamports, number)

	data := make([]byte, 0, 12)
	data = append(data, instruction...)
	data = append(data, lamports...)

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: from, IsSigner: true, IsWritable: true},
			{PubKey: to, IsSigner: false, IsWritable: true},
		},
		ProgramID: SystemProgramID,
		Data:      data,
	}
}
