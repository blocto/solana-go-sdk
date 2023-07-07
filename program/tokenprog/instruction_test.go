package tokenprog

import (
	"reflect"
	"testing"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
	"github.com/blocto/solana-go-sdk/types"
)

func TestInitializeMint(t *testing.T) {
	type args struct {
		param InitializeMintParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: InitializeMintParam{
					Decimals:   1,
					Mint:       common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					MintAuth:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					FreezeAuth: nil,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{0, 1, 159, 186, 247, 199, 172, 215, 195, 31, 127, 42, 207, 18, 192, 64, 156, 59, 98, 1, 180, 8, 69, 70, 199, 127, 220, 159, 6, 40, 64, 117, 246, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: InitializeMintParam{
					Decimals:   5,
					Mint:       common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					MintAuth:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					FreezeAuth: pointer.Get[common.PublicKey](common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7")),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{0, 5, 159, 186, 247, 199, 172, 215, 195, 31, 127, 42, 207, 18, 192, 64, 156, 59, 98, 1, 180, 8, 69, 70, 199, 127, 220, 159, 6, 40, 64, 117, 246, 19, 1, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeMint(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeMint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeAccount(t *testing.T) {
	type args struct {
		param InitializeAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: InitializeAccountParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Owner:   common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMintTo(t *testing.T) {
	type args struct {
		param MintToParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: MintToParam{
					Mint:    common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					To:      common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{},
					Amount:  99999,
				},
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
		{
			args: args{
				param: MintToParam{
					Mint: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					To:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
					Amount: 99999,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{7, 159, 134, 1, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MintTo(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MintTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMintToChecked(t *testing.T) {
	type args struct {
		param MintToCheckedParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: MintToCheckedParam{
					Mint:     common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					To:       common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers:  []common.PublicKey{},
					Amount:   99999,
					Decimals: 5,
				},
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
		{
			args: args{
				param: MintToCheckedParam{
					Mint: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					To:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
					Amount:   99999,
					Decimals: 5,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{14, 159, 134, 1, 0, 0, 0, 0, 0, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MintToChecked(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MintToChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	type args struct {
		param TransferParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: TransferParam{
					From:    common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					To:      common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{},
					Amount:  99999,
				},
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
			if got := Transfer(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransferChecked(t *testing.T) {
	type args struct {
		param TransferCheckedParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: TransferCheckedParam{
					From:     common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					To:       common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Mint:     common.PublicKeyFromString("HFCNHUwPxRqqW6gaLd3uUjJcEUfjnRptJzh4xvnNmavv"),
					Auth:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers:  []common.PublicKey{},
					Amount:   99999,
					Decimals: 4,
				},
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
		{
			args: args{
				param: TransferCheckedParam{
					From: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					To:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Mint: common.PublicKeyFromString("HFCNHUwPxRqqW6gaLd3uUjJcEUfjnRptJzh4xvnNmavv"),
					Auth: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
					Amount:   99999,
					Decimals: 4,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("HFCNHUwPxRqqW6gaLd3uUjJcEUfjnRptJzh4xvnNmavv"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{12, 159, 134, 1, 0, 0, 0, 0, 0, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransferChecked(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransferChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBurn(t *testing.T) {
	type args struct {
		param BurnParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: BurnParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{},
					Amount:  99999,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{8, 159, 134, 1, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: BurnParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
					Amount: 99999,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{8, 159, 134, 1, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Burn(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Burn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBurnChecked(t *testing.T) {
	type args struct {
		param BurnCheckedParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: BurnCheckedParam{
					Account:  common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Auth:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers:  []common.PublicKey{},
					Mint:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Amount:   99999,
					Decimals: 9,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{15, 159, 134, 1, 0, 0, 0, 0, 0, 9},
			},
		},
		{
			args: args{
				param: BurnCheckedParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
					Mint:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Amount:   99999,
					Decimals: 9,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{15, 159, 134, 1, 0, 0, 0, 0, 0, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BurnChecked(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BurnChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloseAccount(t *testing.T) {
	type args struct {
		param CloseAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: CloseAccountParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{},
					To:      common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{9},
			},
		},
		{
			args: args{
				param: CloseAccountParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
					To: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CloseAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloseAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeAccount2(t *testing.T) {
	type args struct {
		param InitializeAccount2Param
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: InitializeAccount2Param{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Owner:   common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{16, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeAccount2(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeAccount2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreezeAccount(t *testing.T) {
	type args struct {
		param FreezeAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: FreezeAccountParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{10},
			},
		},
		{
			args: args{
				param: FreezeAccountParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FreezeAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreezeAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThawAccount(t *testing.T) {
	type args struct {
		param ThawAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: ThawAccountParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{11},
			},
		},
		{
			args: args{
				param: ThawAccountParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThawAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThawAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApprove(t *testing.T) {
	type args struct {
		param ApproveParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: ApproveParam{
					From:    common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					To:      common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{},
					Amount:  99999,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{4, 159, 134, 1, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Approve(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Approve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevoke(t *testing.T) {
	type args struct {
		param RevokeParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: RevokeParam{
					From: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Auth: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Revoke(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Revoke() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApproveChecked(t *testing.T) {
	type args struct {
		param ApproveCheckedParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: ApproveCheckedParam{
					From:     common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					To:       common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"),
					Auth:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers:  []common.PublicKey{},
					Amount:   99999,
					Decimals: 9,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{13, 159, 134, 1, 0, 0, 0, 0, 0, 9},
			},
		},
		{
			args: args{
				param: ApproveCheckedParam{
					From: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					To:   common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"),
					Auth: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
					Amount:   99999,
					Decimals: 9,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{13, 159, 134, 1, 0, 0, 0, 0, 0, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApproveChecked(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApproveChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeMultisig(t *testing.T) {
	type args struct {
		param InitializeMultisigParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: InitializeMultisigParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
						common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"),
						common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					},
					MinRequired: 2,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{

					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeMultisig(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeMultisig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSyncNative(t *testing.T) {
	type args struct {
		param SyncNativeParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: SyncNativeParam{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
				},
				Data: []byte{17},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SyncNative(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SyncNative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetAuthority(t *testing.T) {
	type args struct {
		param SetAuthorityParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: SetAuthorityParam{
					Account:  common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					NewAuth:  pointer.Get[common.PublicKey](common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7")),
					AuthType: AuthorityTypeMintTokens,
					Auth:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Signers:  []common.PublicKey{},
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{6, 0, 1, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240},
			},
		},
		{
			args: args{
				param: SetAuthorityParam{
					Account:  common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					NewAuth:  pointer.Get[common.PublicKey](common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7")),
					AuthType: AuthorityTypeMintTokens,
					Auth:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					},
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{6, 0, 1, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240},
			},
		},
		{
			args: args{
				param: SetAuthorityParam{
					Account:  common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					NewAuth:  nil,
					AuthType: AuthorityTypeMintTokens,
					Auth:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Signers:  []common.PublicKey{},
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: SetAuthorityParam{
					Account:  common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					NewAuth:  nil,
					AuthType: AuthorityTypeFreezeAccount,
					Auth:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Signers:  []common.PublicKey{},
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{6, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: SetAuthorityParam{
					Account:  common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					NewAuth:  nil,
					AuthType: AuthorityTypeAccountOwner,
					Auth:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Signers:  []common.PublicKey{},
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{6, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: SetAuthorityParam{
					Account:  common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					NewAuth:  nil,
					AuthType: AuthorityTypeCloseAccount,
					Auth:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Signers:  []common.PublicKey{},
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{6, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetAuthority(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAuthority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeAccount3(t *testing.T) {
	type args struct {
		param InitializeAccount3Param
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: InitializeAccount3Param{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Mint:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Owner:   common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
				},
				Data: []byte{18, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeAccount3(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeAccount3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeMultisig2(t *testing.T) {
	type args struct {
		param InitializeMultisig2Param
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: InitializeMultisig2Param{
					Account: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Signers: []common.PublicKey{
						common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
						common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"),
						common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					},
					MinRequired: 2,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{

					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{19, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeMultisig2(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeMultisig2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeMint2(t *testing.T) {
	type args struct {
		param InitializeMint2Param
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: InitializeMint2Param{
					Decimals:   1,
					Mint:       common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					MintAuth:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					FreezeAuth: nil,
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
				},
				Data: []byte{20, 1, 159, 186, 247, 199, 172, 215, 195, 31, 127, 42, 207, 18, 192, 64, 156, 59, 98, 1, 180, 8, 69, 70, 199, 127, 220, 159, 6, 40, 64, 117, 246, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: InitializeMint2Param{
					Decimals:   5,
					Mint:       common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					MintAuth:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					FreezeAuth: pointer.Get[common.PublicKey](common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7")),
				},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
				},
				Data: []byte{20, 5, 159, 186, 247, 199, 172, 215, 195, 31, 127, 42, 207, 18, 192, 64, 156, 59, 98, 1, 180, 8, 69, 70, 199, 127, 220, 159, 6, 40, 64, 117, 246, 19, 1, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeMint2(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeMint2() = %v, want %v", got, tt.want)
			}
		})
	}
}
