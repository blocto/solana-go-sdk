package tokenmeta

import (
	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction uint8

const (
	InstructionCreateMetadataAccount Instruction = iota
)

type CreateMetadataAccountParam struct {
	Metadata                common.PublicKey
	Mint                    common.PublicKey
	MintAuthority           common.PublicKey
	Payer                   common.PublicKey
	UpdateAuthority         common.PublicKey
	UpdateAuthorityIsSigner bool
	IsMutable               bool
	MintData                Data
}

func CreateMetadataAccount(param CreateMetadataAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Data        Data
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccount,
		Data:        param.MintData,
		IsMutable:   param.IsMutable,
	})

	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   param.UpdateAuthorityIsSigner,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}
