package tokenmeta

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/bincode"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction uint8

const (
	InstructionCreateMetadataAccount Instruction = iota
)

type TokenData struct {
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	OptionsCreator       bool
	Creators             []uint8
}

func CreateMetadataAccount(metadata, mint, mintAuthority, payer, updateAuthority common.PublicKey, updateAuthorityIsSigner, isMutable bool, mintData Data) types.Instruction {
	var creators []uint8

	if mintData.Creators != nil {
		for _, creator := range *mintData.Creators {
			data, err := bincode.SerializeData(creator)
			if err != nil {
				panic(err)
			}

			creators = append(creators, data...)
		}
	}

	tokenData := TokenData{
		Name:                 mintData.Name,
		Symbol:               mintData.Symbol,
		Uri:                  mintData.Uri,
		SellerFeeBasisPoints: mintData.SellerFeeBasisPoints,
		OptionsCreator:       len(creators) > 0,
		Creators:             creators,
	}

	data, err := bincode.SerializeData(struct {
		Instruction Instruction
		Data        TokenData
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccount,
		Data:        tokenData,
		IsMutable:   isMutable,
	})

	if err != nil {
		panic(err)
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
				IsWritable: false,
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
	}
}
