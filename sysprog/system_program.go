package sysprog

import (
	"encoding/binary"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

func CreateAccount(fromAccount, newAccount, owner common.PublicKey, initLamports, accountSpace uint64) types.Instruction {
	instruction := make([]byte, 4)
	binary.LittleEndian.PutUint32(instruction, 0)
	lamports := make([]byte, 8)
	binary.LittleEndian.PutUint64(lamports, initLamports)
	space := make([]byte, 8)
	binary.LittleEndian.PutUint64(space, accountSpace)

	data := make([]byte, 0, 52)
	data = append(data, instruction...)
	data = append(data, lamports...)
	data = append(data, space...)
	data = append(data, owner[:]...)

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: fromAccount, IsSigner: true, IsWritable: true},
			{PubKey: newAccount, IsSigner: true, IsWritable: true},
		},
		ProgramID: common.SystemProgramID,
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
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}
