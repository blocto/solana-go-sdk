package address_lookup_table

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/stretchr/testify/assert"
)

func TestDeriveLookupTableAddress(t *testing.T) {
	type args struct {
		authorityAddr   common.PublicKey
		recentBlockSlot uint64
	}
	tests := []struct {
		name  string
		args  args
		want  common.PublicKey
		want1 uint8
	}{
		{
			args: args{
				authorityAddr:   common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
				recentBlockSlot: 1,
			},
			want:  common.PublicKeyFromString("HJ6JRbBAPFfeUtiiD2VKAoTH9w7ZCyCGZSaevFFCZtsJ"),
			want1: 255,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DeriveLookupTableAddress(tt.args.authorityAddr, tt.args.recentBlockSlot)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeriveLookupTableAddress() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DeriveLookupTableAddress() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeserializeLookupTable(t *testing.T) {
	type args struct {
		data         []byte
		accountOwner common.PublicKey
	}
	tests := []struct {
		name    string
		args    args
		want    AddressLookupTable
		wantErr error
	}{
		{
			args: args{
				data:         []byte{},
				accountOwner: common.TokenProgramID,
			},
			want:    AddressLookupTable{},
			wantErr: ErrInvalidAccountOwner,
		},
		{
			args: args{
				data:         []byte{1, 0, 0, 0, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 215, 20, 147, 30, 186, 106, 25, 168, 244, 220, 108, 1, 154, 255, 38, 79, 95, 191, 104, 197, 162, 142, 224, 179, 185, 135, 85, 206, 57, 214, 73, 211, 0, 0},
				accountOwner: common.AddressLookupTableProgramID,
			},
			want: AddressLookupTable{
				ProgramState:               ProgramStateLookupTable,
				DeactivationSlot:           ^uint64(0),
				LastExtendedSlot:           0,
				LastExtendedSlotStartIndex: 0,
				Authority:                  pointer.Pubkey(common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz")),
				padding:                    0,
				Addresses:                  []common.PublicKey{},
			},
			wantErr: nil,
		},
		{
			args: args{
				data:         []byte{1, 0, 0, 0, 255, 255, 255, 255, 255, 255, 255, 255, 230, 107, 61, 9, 0, 0, 0, 0, 0, 1, 215, 20, 147, 30, 186, 106, 25, 168, 244, 220, 108, 1, 154, 255, 38, 79, 95, 191, 104, 197, 162, 142, 224, 179, 185, 135, 85, 206, 57, 214, 73, 211, 0, 0, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185, 29, 11, 113, 4, 101, 239, 39, 167, 201, 112, 156, 239, 236, 36, 251, 140, 76, 199, 150, 228, 218, 214, 20, 123, 180, 181, 103, 160, 71, 251, 237, 123},
				accountOwner: common.AddressLookupTableProgramID,
			},
			want: AddressLookupTable{
				ProgramState:               ProgramStateLookupTable,
				DeactivationSlot:           ^uint64(0),
				LastExtendedSlot:           155020262,
				LastExtendedSlotStartIndex: 0,
				Authority:                  pointer.Pubkey(common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz")),
				padding:                    0,
				Addresses: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeserializeLookupTable(tt.args.data, tt.args.accountOwner)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
