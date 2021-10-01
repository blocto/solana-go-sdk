package types

import (
	"crypto/ed25519"
	"reflect"
	"testing"

	"github.com/olegfomenko/solana-go-sdk/common"
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

func TestAccount_Sign(t *testing.T) {
	type fields struct {
		PublicKey  common.PublicKey
		PrivateKey ed25519.PrivateKey
	}
	type args struct {
		message []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			fields: fields{
				PublicKey:  common.PublicKeyFromString("2SeBK1pUxnVbY82vN4TEJiWh4GwaGDkffxPegQP3DFPk"),
				PrivateKey: []byte{214, 49, 53, 208, 232, 140, 85, 41, 45, 128, 173, 3, 79, 105, 136, 236, 132, 164, 35, 93, 59, 196, 51, 59, 127, 139, 1, 155, 245, 83, 230, 184, 21, 109, 62, 131, 66, 207, 210, 237, 39, 93, 125, 50, 137, 69, 236, 28, 138, 68, 1, 30, 175, 228, 109, 140, 77, 52, 105, 79, 223, 111, 131, 31},
			},
			args: args{
				message: []byte("hello"),
			},
			want: []byte{0x64, 0x63, 0xcf, 0x8f, 0x9c, 0x5a, 0x83, 0x35, 0x3c, 0x1c, 0x3c, 0x2c, 0x9d, 0xcf, 0x70, 0x8d, 0xf9, 0xd2, 0xc2, 0x4b, 0x2f, 0xc5, 0x24, 0x66, 0x4a, 0x75, 0xfb, 0xc8, 0x99, 0xed, 0xd1, 0x57, 0xfe, 0xab, 0x31, 0x90, 0x5b, 0x3c, 0x30, 0x5e, 0x48, 0x98, 0xf, 0x35, 0xea, 0xa3, 0x13, 0xa3, 0x50, 0xa, 0xce, 0x43, 0x48, 0x3d, 0x81, 0x6e, 0xc5, 0x15, 0x68, 0xfd, 0x8, 0x1f, 0x24, 0x5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Account{
				PublicKey:  tt.fields.PublicKey,
				PrivateKey: tt.fields.PrivateKey,
			}
			if got := a.Sign(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
