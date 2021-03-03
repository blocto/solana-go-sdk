package tokenprog

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

// InitMint create a new mint
func InitMint(decimals uint8, mint, mintAuthority common.PublicKey, freezeAuthority *common.PublicKey) types.Instruction {
	var option uint8 = 1
	if freezeAuthority == nil {
		option = 0
		randomPublicKey := types.NewAccount().PublicKey
		freezeAuthority = &randomPublicKey
	}

	data, err := common.SerializeData(struct {
		Instruction     TokenInstruction
		Decimals        uint8
		MintAuthority   common.PublicKey
		Option          uint8
		FreezeAuthority common.PublicKey
	}{
		Instruction:     TokenInstructionInitializeMint,
		Decimals:        decimals,
		MintAuthority:   mintAuthority,
		Option:          option,
		FreezeAuthority: *freezeAuthority,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: mint, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}
