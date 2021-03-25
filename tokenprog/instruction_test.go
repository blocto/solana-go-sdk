package tokenprog

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

func TestMintTo(t *testing.T) {
	type args struct {
		mintPubkey    common.PublicKey
		destPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		amount        uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				mintPubkey:    common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				destPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				amount:        1,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{7, 1, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MintTo(tt.args.mintPubkey, tt.args.destPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MintTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMintToChecked(t *testing.T) {
	type args struct {
		mintPubkey    common.PublicKey
		destPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		decimals      uint8
		amount        uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				mintPubkey:    common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				destPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				decimals:      5,
				amount:        1,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{14, 1, 0, 0, 0, 0, 0, 0, 0, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MintToChecked(tt.args.mintPubkey, tt.args.destPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.decimals, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MintToChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}
