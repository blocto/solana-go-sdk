package address_lookup_table

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/types"
)

func TestCreateLookupTable(t *testing.T) {
	type args struct {
		params CreateLookupTableParams
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				params: CreateLookupTableParams{
					LookupTable: common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
					Authority:   common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
					Payer:       common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					RecentSlot:  1,
					BumpSeed:    255,
				},
			},
			want: types.Instruction{
				ProgramID: common.AddressLookupTableProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
						IsSigner:   true,
						IsWritable: false,
					},
					{
						PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						IsSigner:   true,
						IsWritable: true,
					},
					{
						PubKey:     common.SystemProgramID,
						IsSigner:   false,
						IsWritable: false,
					},
				},
				Data: []byte{
					0, 0, 0, 0,
					1, 0, 0, 0, 0, 0, 0, 0,
					255,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLookupTable(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLookupTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreezeLookupTable(t *testing.T) {
	type args struct {
		params FreezeLookupTableParams
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				params: FreezeLookupTableParams{
					LookupTable: common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
					Authority:   common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
				},
			},
			want: types.Instruction{
				ProgramID: common.AddressLookupTableProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
						IsSigner:   true,
						IsWritable: false,
					},
				},
				Data: []byte{
					1, 0, 0, 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FreezeLookupTable(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreezeLookupTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtendLookupTable(t *testing.T) {
	type args struct {
		params ExtendLookupTableParams
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				params: ExtendLookupTableParams{
					LookupTable: common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
					Authority:   common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
					Payer:       nil,
					Addresses: []common.PublicKey{
						common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
						common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					},
				},
			},
			want: types.Instruction{
				ProgramID: common.AddressLookupTableProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
						IsSigner:   true,
						IsWritable: false,
					},
				},
				Data: []byte{
					2, 0, 0, 0,
					2, 0, 0, 0, 0, 0, 0, 0,
					127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185,
					29, 11, 113, 4, 101, 239, 39, 167, 201, 112, 156, 239, 236, 36, 251, 140, 76, 199, 150, 228, 218, 214, 20, 123, 180, 181, 103, 160, 71, 251, 237, 123,
				},
			},
		},
		{
			args: args{
				params: ExtendLookupTableParams{
					LookupTable: common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
					Authority:   common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
					Payer:       pointer.Get[common.PublicKey](common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7")),
					Addresses: []common.PublicKey{
						common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
						common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					},
				},
			},
			want: types.Instruction{
				ProgramID: common.AddressLookupTableProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
						IsSigner:   true,
						IsWritable: false,
					},
					{
						PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						IsSigner:   true,
						IsWritable: true,
					},
					{
						PubKey:     common.SystemProgramID,
						IsSigner:   false,
						IsWritable: false,
					},
				},
				Data: []byte{
					2, 0, 0, 0,
					2, 0, 0, 0, 0, 0, 0, 0,
					127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185,
					29, 11, 113, 4, 101, 239, 39, 167, 201, 112, 156, 239, 236, 36, 251, 140, 76, 199, 150, 228, 218, 214, 20, 123, 180, 181, 103, 160, 71, 251, 237, 123,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtendLookupTable(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtendLookupTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeactivateLookupTable(t *testing.T) {
	type args struct {
		params DeactivateLookupTableParams
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				params: DeactivateLookupTableParams{
					LookupTable: common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
					Authority:   common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
				},
			},
			want: types.Instruction{
				ProgramID: common.AddressLookupTableProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
						IsSigner:   true,
						IsWritable: false,
					},
				},
				Data: []byte{
					3, 0, 0, 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeactivateLookupTable(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeactivateLookupTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloseLookupTable(t *testing.T) {
	type args struct {
		params CloseLookupTableParams
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				params: CloseLookupTableParams{
					LookupTable: common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
					Authority:   common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
					Recipient:   common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
			},
			want: types.Instruction{
				ProgramID: common.AddressLookupTableProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
						IsSigner:   false,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
						IsSigner:   true,
						IsWritable: false,
					},
					{
						PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						IsSigner:   false,
						IsWritable: true,
					},
				},
				Data: []byte{
					4, 0, 0, 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CloseLookupTable(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloseLookupTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
