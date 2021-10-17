package tokenmeta

import (
	"testing"

	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
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
