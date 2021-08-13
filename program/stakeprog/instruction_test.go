package stakeprog

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

func TestSplit(t *testing.T) {
	type args struct {
		stakePubkey      common.PublicKey
		authPubkey       common.PublicKey
		splitStakePubkey common.PublicKey
		lamports         uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				stakePubkey:      common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				authPubkey:       common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				splitStakePubkey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				lamports:         1,
			},
			want: types.Instruction{
				ProgramID: common.StakeProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{3, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.stakePubkey, tt.args.authPubkey, tt.args.splitStakePubkey, tt.args.lamports); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		dest common.PublicKey
		src  common.PublicKey
		auth common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				dest: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				src:  common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				auth: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
			},
			want: types.Instruction{
				ProgramID: common.StakeProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarStakeHistoryPubkey, IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{7, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.dest, tt.args.src, tt.args.auth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizeWithSeed(t *testing.T) {
	type args struct {
		stakePubkey     common.PublicKey
		authBasePubkey  common.PublicKey
		authSeed        string
		authOwnerPubkey common.PublicKey
		newAuthPubkey   common.PublicKey
		authType        StakeAuthorizationType
		custodianPubkey common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				stakePubkey:     common.PublicKeyFromString("6JNQUmE1MdB4E1Caj2A443Za15ju2XFCyjumnddjeNrP"),
				authBasePubkey:  common.PublicKeyFromString("Gx6FKjrt1EbBKsA8DFSgkj6egv8R5AoATBk1j2J3GHxU"),
				authSeed:        "any seed here",
				authOwnerPubkey: common.PublicKeyFromString("Stake11111111111111111111111111111111111111"),
				newAuthPubkey:   common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
				authType:        0,
				custodianPubkey: common.PublicKey{},
			},
			want: types.Instruction{
				ProgramID: common.StakeProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("6JNQUmE1MdB4E1Caj2A443Za15ju2XFCyjumnddjeNrP"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("Gx6FKjrt1EbBKsA8DFSgkj6egv8R5AoATBk1j2J3GHxU"), IsSigner: true, IsWritable: false},
					{PubKey: common.SysVarClockPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{8, 0, 0, 0, 6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222, 0, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 0, 97, 110, 121, 32, 115, 101, 101, 100, 32, 104, 101, 114, 101, 6, 161, 216, 23, 145, 55, 84, 42, 152, 52, 55, 189, 254, 42, 122, 178, 85, 127, 83, 92, 138, 120, 114, 43, 104, 164, 157, 192, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthorizeWithSeed(tt.args.stakePubkey, tt.args.authBasePubkey, tt.args.authSeed, tt.args.authOwnerPubkey, tt.args.newAuthPubkey, tt.args.authType, tt.args.custodianPubkey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeWithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
