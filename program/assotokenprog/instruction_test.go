package assotokenprog

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

func TestCreateAssociatedTokenAccount(t *testing.T) {
	type args struct {
		funder    common.PublicKey
		wallet    common.PublicKey
		tokenMint common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				funder:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				wallet:    common.PublicKeyFromString("5JksDo879mvhxnBPLKPQLvgemxi4et75ipWC9BaLTHBK"),
				tokenMint: common.PublicKeyFromString("G1dYC47buM23b4kdWsa7utfEGM95t2LL3fZn535W5pYC"),
			},
			want: types.Instruction{
				ProgramID: common.SPLAssociatedTokenAccountProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: true},
					{PubKey: common.PublicKeyFromString("8qJdAUsYNCRDDfs7ANyCoLPUj9CfnTM1aJU6Sndbviro"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("5JksDo879mvhxnBPLKPQLvgemxi4et75ipWC9BaLTHBK"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("G1dYC47buM23b4kdWsa7utfEGM95t2LL3fZn535W5pYC"), IsSigner: false, IsWritable: false},
					{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
					{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAssociatedTokenAccount(tt.args.funder, tt.args.wallet, tt.args.tokenMint)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAssociatedTokenAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
