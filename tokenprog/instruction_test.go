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
				amount:        99999,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{7, 159, 134, 1, 0, 0, 0, 0, 0},
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
				amount:        99999,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{14, 159, 134, 1, 0, 0, 0, 0, 0, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MintToChecked(tt.args.mintPubkey, tt.args.destPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MintToChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	type args struct {
		srcPubkey     common.PublicKey
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
				srcPubkey:     common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				destPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				amount:        99999,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{3, 159, 134, 1, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transfer(tt.args.srcPubkey, tt.args.destPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransferChecked(t *testing.T) {
	type args struct {
		srcPubkey     common.PublicKey
		destPubkey    common.PublicKey
		mintPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		amount        uint64
		decimals      uint8
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				srcPubkey:     common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				destPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				mintPubkey:    common.PublicKeyFromString("HFCNHUwPxRqqW6gaLd3uUjJcEUfjnRptJzh4xvnNmavv"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				amount:        99999,
				decimals:      4,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("HFCNHUwPxRqqW6gaLd3uUjJcEUfjnRptJzh4xvnNmavv"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{12, 159, 134, 1, 0, 0, 0, 0, 0, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransferChecked(tt.args.srcPubkey, tt.args.destPubkey, tt.args.mintPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransferChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}
