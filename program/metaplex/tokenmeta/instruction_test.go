package tokenmeta

import (
	"testing"

	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateMetadataAccount(t *testing.T) {
	type args struct {
		param CreateMetadataAccountParam
	}

	tests := []struct {
		name  string
		args  args
		check func(t *testing.T, instruction types.Instruction)
	}{
		{
			name: "Positive test creating metadata instruction",
			args: args{
				param: CreateMetadataAccountParam{
					Metadata:                common.PublicKeyFromString("DC2mkgwhy56w3viNtHDjJQmc7SGu2QX785bS4aexojwX"),
					Mint:                    common.PublicKeyFromString("GphF2vTuzhwhLWBWWvD8y5QLCPp1aQC5EnzrWsnbiWPx"),
					MintAuthority:           common.PublicKeyFromString("9BKWqDHfHZh9j39xakYVMdr6hXmCLHH5VfCpeq2idU9L"),
					Payer:                   common.PublicKeyFromString("9FYsKrNuEweb55Wa2jaj8wTKYDBvuCG3huhakEj96iN9"),
					UpdateAuthority:         common.PublicKeyFromString("HNGVuL5kqjDehw7KR63w9gxow32sX6xzRNgLb8GkbwCM"),
					UpdateAuthorityIsSigner: true,
					IsMutable:               true,
					MintData: Data{
						Name:                 "Test NFT",
						Symbol:               "TST",
						Uri:                  "https://test.com/metadata",
						SellerFeeBasisPoints: 10,
						Creators: &[]Creator{
							{
								Address:  common.PublicKeyFromString("7FzXBBPjzrNJbm9MrZKZcyvP3ojVeYPUG2XkBPVZvuBu"),
								Verified: true,
								Share:    100,
							},
						},
					},
				},
			},
			check: func(t *testing.T, instruction types.Instruction) {
				data, err := borsh.Serialize(struct {
					Instruction Instruction
					Data        Data
					IsMutable   bool
				}{
					Instruction: InstructionCreateMetadataAccount,
					Data: Data{
						Name:                 "Test NFT",
						Symbol:               "TST",
						Uri:                  "https://test.com/metadata",
						SellerFeeBasisPoints: 10,
						Creators: &[]Creator{
							{
								Address:  common.PublicKeyFromString("7FzXBBPjzrNJbm9MrZKZcyvP3ojVeYPUG2XkBPVZvuBu"),
								Verified: true,
								Share:    100,
							},
						},
					},
					IsMutable: true,
				})

				assert.NoError(t, err)

				want := types.Instruction{
					ProgramID: common.MetaplexTokenMetaProgramID,
					Accounts: []types.AccountMeta{
						{PubKey: common.PublicKeyFromString("DC2mkgwhy56w3viNtHDjJQmc7SGu2QX785bS4aexojwX"), IsSigner: false, IsWritable: true},
						{PubKey: common.PublicKeyFromString("GphF2vTuzhwhLWBWWvD8y5QLCPp1aQC5EnzrWsnbiWPx"), IsSigner: false, IsWritable: false},
						{PubKey: common.PublicKeyFromString("9BKWqDHfHZh9j39xakYVMdr6hXmCLHH5VfCpeq2idU9L"), IsSigner: true, IsWritable: false},
						{PubKey: common.PublicKeyFromString("9FYsKrNuEweb55Wa2jaj8wTKYDBvuCG3huhakEj96iN9"), IsSigner: true, IsWritable: true},
						{PubKey: common.PublicKeyFromString("HNGVuL5kqjDehw7KR63w9gxow32sX6xzRNgLb8GkbwCM"), IsSigner: true, IsWritable: false},
						{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
						{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
					},
					Data: data,
				}

				assert.Equal(t, want, instruction)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateMetadataAccount(tt.args.param)
			tt.check(t, got)
		})
	}
}

func TestCreateMasterEdition(t *testing.T) {
	type args struct {
		param CreateMasterEditionParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: CreateMasterEditionParam{
					Edition:         common.PublicKeyFromString("edition111111111111111111111111111111111111"),
					Mint:            common.PublicKeyFromString("mint111111111111111111111111111111111111111"),
					UpdateAuthority: common.PublicKeyFromString("updateAuthority1111111111111111111111111111"),
					MintAuthority:   common.PublicKeyFromString("mintAuthority111111111111111111111111111111"),
					Metadata:        common.PublicKeyFromString("metadata11111111111111111111111111111111111"),
					Payer:           common.PublicKeyFromString("payer11111111111111111111111111111111111111"),
					MaxSupply:       nil,
				},
			},
			want: types.Instruction{
				ProgramID: common.MetaplexTokenMetaProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("edition111111111111111111111111111111111111"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("mint111111111111111111111111111111111111111"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("updateAuthority1111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("mintAuthority111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("payer11111111111111111111111111111111111111"), IsSigner: true, IsWritable: true},
					{PubKey: common.PublicKeyFromString("metadata11111111111111111111111111111111111"), IsSigner: false, IsWritable: false},
					{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
					{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{10, 0},
			},
		},
		{
			args: args{
				param: CreateMasterEditionParam{
					Edition:         common.PublicKeyFromString("edition111111111111111111111111111111111111"),
					Mint:            common.PublicKeyFromString("mint111111111111111111111111111111111111111"),
					UpdateAuthority: common.PublicKeyFromString("updateAuthority1111111111111111111111111111"),
					MintAuthority:   common.PublicKeyFromString("mintAuthority111111111111111111111111111111"),
					Metadata:        common.PublicKeyFromString("metadata11111111111111111111111111111111111"),
					Payer:           common.PublicKeyFromString("payer11111111111111111111111111111111111111"),
					MaxSupply:       pointer.Uint64(2),
				},
			},
			want: types.Instruction{
				ProgramID: common.MetaplexTokenMetaProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("edition111111111111111111111111111111111111"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("mint111111111111111111111111111111111111111"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("updateAuthority1111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("mintAuthority111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("payer11111111111111111111111111111111111111"), IsSigner: true, IsWritable: true},
					{PubKey: common.PublicKeyFromString("metadata11111111111111111111111111111111111"), IsSigner: false, IsWritable: false},
					{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
					{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{10, 1, 2, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateMasterEdition(tt.args.param)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMintNewEditionFromMasterEditionViaToken(t *testing.T) {
	type args struct {
		param MintNewEditionFromMasterEditionViaTokeParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: MintNewEditionFromMasterEditionViaTokeParam{
					NewMetaData:                common.PublicKeyFromString("newMetaData11111111111111111111111111111111"),
					NewEdition:                 common.PublicKeyFromString("newEdition111111111111111111111111111111111"),
					MasterEdition:              common.PublicKeyFromString("masterEdition111111111111111111111111111111"),
					NewMint:                    common.PublicKeyFromString("newMint111111111111111111111111111111111111"),
					EditionMark:                common.PublicKeyFromString("editionMark11111111111111111111111111111111"),
					NewMintAuthority:           common.PublicKeyFromString("newMintAuthority111111111111111111111111111"),
					Payer:                      common.PublicKeyFromString("payer11111111111111111111111111111111111111"),
					TokenAccountOwner:          common.PublicKeyFromString("tokenAccountOwner11111111111111111111111111"),
					TokenAccount:               common.PublicKeyFromString("tokenAccount1111111111111111111111111111111"),
					NewMetadataUpdateAuthority: common.PublicKeyFromString("newMetadataUpdateAuthority11111111111111111"),
					MasterMetadata:             common.PublicKeyFromString("masterMetadata11111111111111111111111111111"),
					Edition:                    1,
				},
			},
			want: types.Instruction{
				ProgramID: common.MetaplexTokenMetaProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("newMetaData11111111111111111111111111111111"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("newEdition111111111111111111111111111111111"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("masterEdition111111111111111111111111111111"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("newMint111111111111111111111111111111111111"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("editionMark11111111111111111111111111111111"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("newMintAuthority111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("payer11111111111111111111111111111111111111"), IsSigner: true, IsWritable: true},
					{PubKey: common.PublicKeyFromString("tokenAccountOwner11111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("tokenAccount1111111111111111111111111111111"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("newMetadataUpdateAuthority11111111111111111"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("masterMetadata11111111111111111111111111111"), IsSigner: false, IsWritable: false},
					{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
					{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{11, 1, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MintNewEditionFromMasterEditionViaToken(tt.args.param)
			assert.Equal(t, tt.want, got)
		})
	}
}
