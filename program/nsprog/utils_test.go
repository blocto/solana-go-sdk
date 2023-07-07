package nsprog

import (
	"reflect"
	"testing"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/stretchr/testify/assert"
)

func TestGetTwitterRegistryKey(t *testing.T) {
	type args struct {
		twitterHandle string
	}
	tests := []struct {
		name string
		args args
		want common.PublicKey
	}{
		{
			args: args{
				twitterHandle: "gghost07114721",
			},
			want: common.PublicKeyFromString("5r2pKbCFibGZp18u51tcvzQpsNsA98TyCF1UbDmbSUk5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetTwitterRegistryKey(tt.args.twitterHandle))
		})
	}
}

func TestGetHashName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				name: "blocto",
			},
			want: []byte{0x62, 0x94, 0x3a, 0xc2, 0x9e, 0x7b, 0x9d, 0x4e, 0x38, 0x53, 0xb2, 0x84, 0xdd, 0x7f, 0x1, 0x66, 0xeb, 0x5f, 0x0, 0xe3, 0x1f, 0x25, 0x53, 0x51, 0x83, 0x61, 0x38, 0x33, 0xcd, 0xc5, 0xf9, 0x3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHashName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHashName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNameAccountKey(t *testing.T) {
	type args struct {
		hashName   []byte
		nameClass  common.PublicKey
		nameParent common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want common.PublicKey
	}{
		{
			name: "domain: blocto.sol",
			args: args{
				hashName:   GetHashName("blocto"),
				nameClass:  common.PublicKey{},
				nameParent: SolTldAuthority,
			},
			want: common.PublicKeyFromString("6yAP2rFW7wQiqVmySE4DTfQSWmp6fR1geGyWx6SQMAhS"),
		},
		{
			name: "domain yihau.blocto.sol",
			args: args{
				hashName:   GetHashName("\x00yihau"),
				nameClass:  common.PublicKey{},
				nameParent: GetNameAccountKey(GetHashName("blocto"), common.PublicKey{}, SolTldAuthority),
			},
			want: common.PublicKeyFromString("5Cjg2Xah4Cc24yM7zsfbyBuXKZ6Wm9ZJqHa5n47vnvNz"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNameAccountKey(tt.args.hashName, tt.args.nameClass, tt.args.nameParent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNameAccountKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
