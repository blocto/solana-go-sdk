package tokenmeta

import (
	"github.com/near/borsh-go"
	"github.com/olegfomenko/solana-go-sdk/common"
	"github.com/olegfomenko/solana-go-sdk/types"
	"github.com/pkg/errors"
)

type Instruction uint8

const (
	InstructionCreateMetadataAccount Instruction = iota
	InstructionUpdatePrimarySaleHappenedViaToken
)

func CreateMetadataAccount(metadata, mint, mintAuthority, payer, updateAuthority common.PublicKey, updateAuthorityIsSigner, isMutable bool, mintData Data) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Data        Data
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccount,
		Data:        mintData,
		IsMutable:   isMutable,
	})

	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     mintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     updateAuthority,
				IsSigner:   updateAuthorityIsSigner,
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
	}, nil
}

func UpdatePrimarySaleHappenedViaToken(token, owner common.PublicKey) (types.Instruction, error) {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionCreateMetadataAccount,
	})

	if err != nil {
		return types.Instruction{}, errors.Wrap(err, "failed serialize")
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     token,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     owner,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: data,
	}, nil
}
