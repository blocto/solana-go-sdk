package token_metadata

import (
	"testing"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/stretchr/testify/assert"
)

func TestGetEditionMark(t *testing.T) {
	type args struct {
		mint    common.PublicKey
		edition uint64
	}
	tests := []struct {
		name    string
		args    args
		want    common.PublicKey
		wantErr error
	}{
		{
			args: args{
				mint:    common.PublicKeyFromString("7WUw2LkJJ6kAjuJM4gf6XcJdLdpKPXEGZQf1E3qisXie"),
				edition: 1,
			},
			want: common.PublicKeyFromString("HXr9QjSNttPntDHPvLokf8MgXu3vUdytVMjRFzJSc2VR"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEditionMark(tt.args.mint, tt.args.edition)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetMasterEdition(t *testing.T) {
	type args struct {
		mint common.PublicKey
	}
	tests := []struct {
		name    string
		args    args
		want    common.PublicKey
		wantErr error
	}{
		{
			args: args{
				mint: common.PublicKeyFromString("7WUw2LkJJ6kAjuJM4gf6XcJdLdpKPXEGZQf1E3qisXie"),
			},
			want: common.PublicKeyFromString("2e446uJgJ3o2qBPAmCAubM3FXmbwxQuoWgqERo2Fcjka"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMasterEdition(tt.args.mint)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTokenMetaPubkey(t *testing.T) {
	type args struct {
		mint common.PublicKey
	}
	tests := []struct {
		name    string
		args    args
		want    common.PublicKey
		wantErr error
	}{
		{
			args: args{
				mint: common.PublicKeyFromString("7WUw2LkJJ6kAjuJM4gf6XcJdLdpKPXEGZQf1E3qisXie"),
			},
			want: common.PublicKeyFromString("3UzhPgdEwidPgvS51bymnCiAgFrYGygpBnBBwjnpKbnp"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTokenMetaPubkey(tt.args.mint)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
