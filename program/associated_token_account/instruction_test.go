package associated_token_account

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateAssociatedTokenAccount(t *testing.T) {
	type args struct {
		param CreateAssociatedTokenAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: CreateAssociatedTokenAccountParam{
					Funder:                 common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Owner:                  common.PublicKeyFromString("5JksDo879mvhxnBPLKPQLvgemxi4et75ipWC9BaLTHBK"),
					Mint:                   common.PublicKeyFromString("G1dYC47buM23b4kdWsa7utfEGM95t2LL3fZn535W5pYC"),
					AssociatedTokenAccount: common.PublicKeyFromString("8qJdAUsYNCRDDfs7ANyCoLPUj9CfnTM1aJU6Sndbviro"),
				},
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
				Data: []byte{0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAssociatedTokenAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAssociatedTokenAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIdempotent(t *testing.T) {
	type args struct {
		param CreateIdempotentParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: CreateIdempotentParam{
					Funder:                 common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Owner:                  common.PublicKeyFromString("5JksDo879mvhxnBPLKPQLvgemxi4et75ipWC9BaLTHBK"),
					Mint:                   common.PublicKeyFromString("G1dYC47buM23b4kdWsa7utfEGM95t2LL3fZn535W5pYC"),
					AssociatedTokenAccount: common.PublicKeyFromString("8qJdAUsYNCRDDfs7ANyCoLPUj9CfnTM1aJU6Sndbviro"),
				},
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
				Data: []byte{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CreateIdempotent(tt.args.param))
		})
	}
}

func TestRecoverNested(t *testing.T) {
	type args struct {
		param RecoverNestedParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: RecoverNestedParam{
					Owner:                             common.PublicKeyFromString("GmNDCuWcaWKzrt7hMo7m7FC7zjUAaZ22hVb5j5LKQtsJ"),
					OwnerMint:                         common.PublicKeyFromString("BE8XnSd5rXK2WS1C6eyX1zsz7Q1cnvM6W8fP9JMyf513"),
					OwnerAssociatedTokenAccount:       common.PublicKeyFromString("FMX5cjKZCT3kxEBmhxVJAEtq9JcZApf4cktcBtxRGeF4"),
					NestedMint:                        common.PublicKeyFromString("C3BoE7oNqkS2ufmYFis5PNnGPLoPrBJWYiF4vP43p7qC"),
					NestedMintAssociatedTokenAccount:  common.PublicKeyFromString("73Ze74x2JRnJgYHp8PATBnDGUXGREmz3YRSh9KUxAVUE"),
					DestinationAssociatedTokenAccount: common.PublicKeyFromString("9DyGS9BVS1BwfCTFLVTYu7doKoX3jDduhvAkAPktfQAs"),
				},
			},
			want: types.Instruction{
				ProgramID: common.SPLAssociatedTokenAccountProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("73Ze74x2JRnJgYHp8PATBnDGUXGREmz3YRSh9KUxAVUE"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("C3BoE7oNqkS2ufmYFis5PNnGPLoPrBJWYiF4vP43p7qC"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("9DyGS9BVS1BwfCTFLVTYu7doKoX3jDduhvAkAPktfQAs"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("FMX5cjKZCT3kxEBmhxVJAEtq9JcZApf4cktcBtxRGeF4"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BE8XnSd5rXK2WS1C6eyX1zsz7Q1cnvM6W8fP9JMyf513"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("GmNDCuWcaWKzrt7hMo7m7FC7zjUAaZ22hVb5j5LKQtsJ"), IsSigner: true, IsWritable: true},
					{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
				},
				Data: []byte{2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RecoverNested(tt.args.param))
		})
	}
}
