package types

import (
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/stretchr/testify/assert"
)

func TestAccountFromBase58(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want Account
		err  error
	}{
		{
			args: args{
				key: "5HNxRJoirY4oRTcRwiEYFALSSLn9nMAyLQKDuSuiCJ966816BjwGuamRdTLTsR2FBHiB7CQkGaw6B4ehBMogPRvW",
			},
			want: Account{
				PublicKey:  common.PublicKeyFromString("2SeBK1pUxnVbY82vN4TEJiWh4GwaGDkffxPegQP3DFPk"),
				PrivateKey: []byte{214, 49, 53, 208, 232, 140, 85, 41, 45, 128, 173, 3, 79, 105, 136, 236, 132, 164, 35, 93, 59, 196, 51, 59, 127, 139, 1, 155, 245, 83, 230, 184, 21, 109, 62, 131, 66, 207, 210, 237, 39, 93, 125, 50, 137, 69, 236, 28, 138, 68, 1, 30, 175, 228, 109, 140, 77, 52, 105, 79, 223, 111, 131, 31},
			},
		},
		{
			args: args{
				key: "d63135d0e88c55292d80ad034f6988ec84a4235d3bc4333b7f8b019bf553e6b8156d3e8342cfd2ed275d7d328945ec1c8a44011eafe46d8c4d34694fdf6f831f",
			},
			want: Account{},
			err:  ErrAccountFailedToBase58Decode,
		},
		{
			args: args{
				key: "2SeBK1pUxnVbY82vN4TEJiWh4GwaGDkffxPegQP3DFPk",
			},
			want: Account{},
			err:  ErrAccountPrivateKeyLengthMismatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AccountFromBase58(tt.args.key)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAccountFromHex(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want Account
		err  error
	}{
		{
			args: args{
				key: "d63135d0e88c55292d80ad034f6988ec84a4235d3bc4333b7f8b019bf553e6b8156d3e8342cfd2ed275d7d328945ec1c8a44011eafe46d8c4d34694fdf6f831f",
			},
			want: Account{
				PublicKey:  common.PublicKeyFromString("2SeBK1pUxnVbY82vN4TEJiWh4GwaGDkffxPegQP3DFPk"),
				PrivateKey: []byte{214, 49, 53, 208, 232, 140, 85, 41, 45, 128, 173, 3, 79, 105, 136, 236, 132, 164, 35, 93, 59, 196, 51, 59, 127, 139, 1, 155, 245, 83, 230, 184, 21, 109, 62, 131, 66, 207, 210, 237, 39, 93, 125, 50, 137, 69, 236, 28, 138, 68, 1, 30, 175, 228, 109, 140, 77, 52, 105, 79, 223, 111, 131, 31},
			},
		},
		{
			args: args{
				key: "5HNxRJoirY4oRTcRwiEYFALSSLn9nMAyLQKDuSuiCJ966816BjwGuamRdTLTsR2FBHiB7CQkGaw6B4ehBMogPRvW",
			},
			want: Account{},
			err:  ErrAccountFailedToHexDecode,
		},
		{
			args: args{
				key: "d63135d0e88c55292d80ad034f6988ec84a4235d3bc4333b7f8b019bf553e6b8",
			},
			want: Account{},
			err:  ErrAccountPrivateKeyLengthMismatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AccountFromHex(tt.args.key)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
