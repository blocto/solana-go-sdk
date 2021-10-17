package stakeprog

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/types"
)

func TestSplit(t *testing.T) {
	type args struct {
		param SplitParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: SplitParam{
					Stake:      common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Auth:       common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					SplitStake: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Lamports:   1,
				},
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
			if got := Split(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		param MergeParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: MergeParam{
					From: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					To:   common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Auth: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
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
			if got := Merge(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizeWithSeed(t *testing.T) {
	type args struct {
		param AuthorizeWithSeedParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: AuthorizeWithSeedParam{
					Stake:     common.PublicKeyFromString("6JNQUmE1MdB4E1Caj2A443Za15ju2XFCyjumnddjeNrP"),
					AuthBase:  common.PublicKeyFromString("Gx6FKjrt1EbBKsA8DFSgkj6egv8R5AoATBk1j2J3GHxU"),
					AuthSeed:  "any seed here",
					AuthOwner: common.PublicKeyFromString("Stake11111111111111111111111111111111111111"),
					NewAuth:   common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
					AuthType:  StakeAuthorizationTypeStaker,
					Custodian: nil,
				},
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
			if got := AuthorizeWithSeed(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeWithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetLockup(t *testing.T) {
	type args struct {
		param SetLockupParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: SetLockupParam{
					Stake:  common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Auth:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Lockup: LockupParam{},
				},
			},
			want: types.Instruction{
				ProgramID: common.StakeProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
						IsSigner:   true,
						IsWritable: false,
					},
				},
				Data: []byte{6, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: SetLockupParam{
					Stake: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Auth:  common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Lockup: LockupParam{
						UnixTimestamp: pointer.Int64(1),
					},
				},
			},
			want: types.Instruction{
				ProgramID: common.StakeProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
						IsSigner:   true,
						IsWritable: false,
					},
				},
				Data: []byte{6, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: SetLockupParam{
					Stake: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Auth:  common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Lockup: LockupParam{
						Epoch: pointer.Uint64(1),
					},
				},
			},
			want: types.Instruction{
				ProgramID: common.StakeProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
						IsSigner:   true,
						IsWritable: false,
					},
				},
				Data: []byte{6, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: SetLockupParam{
					Stake: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Auth:  common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Lockup: LockupParam{
						Cusodian: pointer.Pubkey(common.PublicKeyFromString("DTA7FmUNYuQs2mScj2Lx8gQV63SEL1zGtzCSvPxtijbi")),
					},
				},
			},
			want: types.Instruction{
				ProgramID: common.StakeProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
						IsSigner:   true,
						IsWritable: false,
					},
				},
				Data: []byte{6, 0, 0, 0, 0, 0, 1, 184, 255, 164, 117, 118, 144, 60, 115, 73, 97, 27, 182, 246, 34, 53, 198, 220, 186, 65, 56, 58, 57, 174, 141, 147, 215, 61, 242, 136, 190, 88, 189},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetLockup(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLockup() = %v, want %v", got, tt.want)
			}
		})
	}
}
