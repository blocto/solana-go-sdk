package sysprog

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

func TestCreateAccount(t *testing.T) {
	type args struct {
		param CreateAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: CreateAccountParam{
					From:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					New:      common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Owner:    common.StakeProgramID,
					Lamports: 1,
					Space:    200,
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: true},
				},
				Data: []byte{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 200, 0, 0, 0, 0, 0, 0, 0, 6, 161, 216, 23, 145, 55, 84, 42, 152, 52, 55, 189, 254, 42, 122, 178, 85, 127, 83, 92, 138, 120, 114, 43, 104, 164, 157, 192, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssign(t *testing.T) {
	type args struct {
		param AssignParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: AssignParam{
					From:  common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Owner: common.StakeProgramID,
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: true},
				},
				Data: []byte{1, 0, 0, 0, 6, 161, 216, 23, 145, 55, 84, 42, 152, 52, 55, 189, 254, 42, 122, 178, 85, 127, 83, 92, 138, 120, 114, 43, 104, 164, 157, 192, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Assign(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Assign() = %v, want %v", got, tt.want)
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
					From:   common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					To:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Amount: 1,
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						IsSigner:   true,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
						IsSigner:   false,
						IsWritable: true,
					},
				},
				Data: []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
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

func TestCreateAccountWithSeed(t *testing.T) {
	type args struct {
		param CreateAccountWithSeedParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: CreateAccountWithSeedParam{
					From:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					New:      common.PublicKeyFromString("DTA7FmUNYuQs2mScj2Lx8gQV63SEL1zGtzCSvPxtijbi"),
					Base:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Owner:    common.SystemProgramID,
					Seed:     "0",
					Lamports: 0,
					Space:    0,
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: true},
					{PubKey: common.PublicKeyFromString("DTA7FmUNYuQs2mScj2Lx8gQV63SEL1zGtzCSvPxtijbi"), IsSigner: false, IsWritable: true},
				},
				Data: []byte{3, 0, 0, 0, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 1, 0, 0, 0, 0, 0, 0, 0, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			args: args{
				param: CreateAccountWithSeedParam{
					From:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					New:      common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Base:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Owner:    common.SystemProgramID,
					Seed:     "0",
					Lamports: 0,
					Space:    0,
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: true},
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{3, 0, 0, 0, 159, 186, 247, 199, 172, 215, 195, 31, 127, 42, 207, 18, 192, 64, 156, 59, 98, 1, 180, 8, 69, 70, 199, 127, 220, 159, 6, 40, 64, 117, 246, 19, 1, 0, 0, 0, 0, 0, 0, 0, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAccountWithSeed(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccountWithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdvanceNonceAccount(t *testing.T) {
	type args struct {
		param AdvanceNonceAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: AdvanceNonceAccountParam{
					Nonce: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:  common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{4, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdvanceNonceAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdvanceNonceAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithdrawNonceAccount(t *testing.T) {
	type args struct {
		param WithdrawNonceAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: WithdrawNonceAccountParam{
					Nonce:  common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Auth:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					To:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Amount: 1,
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: true},
					{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{5, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithdrawNonceAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithdrawNonceAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeNonceAccount(t *testing.T) {
	type args struct {
		param InitializeNonceAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: InitializeNonceAccountParam{
					Nonce: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Auth:  common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{6, 0, 0, 0, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeNonceAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeNonceAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizeNonceAccount(t *testing.T) {
	type args struct {
		param AuthorizeNonceAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: AuthorizeNonceAccountParam{
					Nonce:   common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
					Auth:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					NewAuth: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{7, 0, 0, 0, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthorizeNonceAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeNonceAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocateWithSeed(t *testing.T) {
	type args struct {
		accountPubkey common.PublicKey
		basePubkey    common.PublicKey
		programID     common.PublicKey
		seed          string
		space         uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				basePubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				programID:     common.SystemProgramID,
				seed:          "0",
				space:         256,
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{9, 0, 0, 0, 159, 186, 247, 199, 172, 215, 195, 31, 127, 42, 207, 18, 192, 64, 156, 59, 98, 1, 180, 8, 69, 70, 199, 127, 220, 159, 6, 40, 64, 117, 246, 19, 1, 0, 0, 0, 0, 0, 0, 0, 48, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllocateWithSeed(tt.args.accountPubkey, tt.args.basePubkey, tt.args.programID, tt.args.seed, tt.args.space); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllocateWithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssignWithSeed(t *testing.T) {
	type args struct {
		accountPubkey     common.PublicKey
		assignToProgramID common.PublicKey
		basePubkey        common.PublicKey
		seed              string
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey:     common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				basePubkey:        common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				assignToProgramID: common.StakeProgramID,
				seed:              "0",
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{10, 0, 0, 0, 159, 186, 247, 199, 172, 215, 195, 31, 127, 42, 207, 18, 192, 64, 156, 59, 98, 1, 180, 8, 69, 70, 199, 127, 220, 159, 6, 40, 64, 117, 246, 19, 1, 0, 0, 0, 0, 0, 0, 0, 48, 6, 161, 216, 23, 145, 55, 84, 42, 152, 52, 55, 189, 254, 42, 122, 178, 85, 127, 83, 92, 138, 120, 114, 43, 104, 164, 157, 192, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssignWithSeed(tt.args.accountPubkey, tt.args.assignToProgramID, tt.args.basePubkey, tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssignWithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocate(t *testing.T) {
	type args struct {
		accountPubkey common.PublicKey
		space         uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				space:         500,
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: true},
				},
				Data: []byte{8, 0, 0, 0, 244, 1, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Allocate(tt.args.accountPubkey, tt.args.space); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Allocate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransferWithSeed(t *testing.T) {
	type args struct {
		from      common.PublicKey
		to        common.PublicKey
		base      common.PublicKey
		programID common.PublicKey
		seed      string
		lamports  uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				from:      common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				to:        common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				base:      common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				programID: common.SystemProgramID,
				seed:      "0",
				lamports:  99999,
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: false, IsWritable: true},
				},
				Data: []byte{11, 0, 0, 0, 159, 134, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransferWithSeed(tt.args.from, tt.args.to, tt.args.base, tt.args.programID, tt.args.seed, tt.args.lamports); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransferWithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
