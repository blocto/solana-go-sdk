package sysprog

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"

	"github.com/olegfomenko/solana-go-sdk/common"
	"github.com/olegfomenko/solana-go-sdk/types"
)

func TestCreateAccountWithSeed(t *testing.T) {
	type args struct {
		fromPubkey       common.PublicKey
		newAccountPubkey common.PublicKey
		basePubkey       common.PublicKey
		programID        common.PublicKey
		seed             string
		lamports         uint64
		space            uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				fromPubkey:       common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				newAccountPubkey: common.PublicKeyFromString("DTA7FmUNYuQs2mScj2Lx8gQV63SEL1zGtzCSvPxtijbi"),
				basePubkey:       common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				programID:        common.SystemProgramID,
				seed:             "0",
				lamports:         0,
				space:            0,
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
				fromPubkey:       common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				newAccountPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				basePubkey:       common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				programID:        common.SystemProgramID,
				seed:             "0",
				lamports:         0,
				space:            0,
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
			got, err := CreateAccountWithSeed(tt.args.fromPubkey, tt.args.newAccountPubkey, tt.args.basePubkey, tt.args.programID, tt.args.seed, tt.args.lamports, tt.args.space)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccountWithSeed() = %v, want %v", got, tt.want)
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
			got, err := AllocateWithSeed(tt.args.accountPubkey, tt.args.basePubkey, tt.args.programID, tt.args.seed, tt.args.space)
			assert.NoError(t, err)
			if !reflect.DeepEqual(got, tt.want) {
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
			got, err := AssignWithSeed(tt.args.accountPubkey, tt.args.assignToProgramID, tt.args.basePubkey, tt.args.seed)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
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
			got, err := Allocate(tt.args.accountPubkey, tt.args.space)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Allocate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssign(t *testing.T) {
	type args struct {
		accountPubkey     common.PublicKey
		assignToProgramID common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				assignToProgramID: common.StakeProgramID,
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
			got, err := Assign(tt.args.accountPubkey, tt.args.assignToProgramID)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Assign() = %v, want %v", got, tt.want)
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
			got, err := TransferWithSeed(tt.args.from, tt.args.to, tt.args.base, tt.args.programID, tt.args.seed, tt.args.lamports)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransferWithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateAccount(t *testing.T) {
	type args struct {
		fromAccount  common.PublicKey
		newAccount   common.PublicKey
		owner        common.PublicKey
		initLamports uint64
		accountSpace uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				fromAccount:  common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				newAccount:   common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				owner:        common.StakeProgramID,
				initLamports: 1,
				accountSpace: 200,
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
			got, err := CreateAccount(tt.args.fromAccount, tt.args.newAccount, tt.args.owner, tt.args.initLamports, tt.args.accountSpace)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	type args struct {
		from   common.PublicKey
		to     common.PublicKey
		number uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				from:   common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				to:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				number: 1,
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
			got, err := Transfer(tt.args.from, tt.args.to, tt.args.number)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdvanceNonceAccount(t *testing.T) {
	type args struct {
		noncePubkey common.PublicKey
		authPubkey  common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				noncePubkey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:  common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
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
			got, err := AdvanceNonceAccount(tt.args.noncePubkey, tt.args.authPubkey)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdvanceNonceAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeNonceAccount(t *testing.T) {
	type args struct {
		noncePubkey common.PublicKey
		authPubkey  common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				noncePubkey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:  common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
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
			got, err := InitializeNonceAccount(tt.args.noncePubkey, tt.args.authPubkey)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeNonceAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithdrawNonceAccount(t *testing.T) {
	type args struct {
		noncePubkey common.PublicKey
		authPubkey  common.PublicKey
		toPubkey    common.PublicKey
		lamports    uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				noncePubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				authPubkey:  common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				toPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				lamports:    1,
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
			got, err := WithdrawNonceAccount(tt.args.noncePubkey, tt.args.authPubkey, tt.args.toPubkey, tt.args.lamports)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithdrawNonceAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizeNonceAccount(t *testing.T) {
	type args struct {
		noncePubkey   common.PublicKey
		oriAuthPubkey common.PublicKey
		newAuthPubkey common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				noncePubkey:   common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				oriAuthPubkey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				newAuthPubkey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
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
			got, err := AuthorizeNonceAccount(tt.args.noncePubkey, tt.args.oriAuthPubkey, tt.args.newAuthPubkey)
			assert.NoError(t, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeNonceAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
