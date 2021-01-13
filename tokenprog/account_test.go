package tokenprog

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
)

func TestAccountFromData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Account
		wantErr bool
	}{
		{
			args: args{
				data: []byte{105, 145, 9, 101, 129, 184, 46, 130, 176, 132, 102, 98, 17, 241, 215, 189, 90, 219, 106, 196, 196, 121, 174, 243, 65, 40, 132, 7, 252, 112, 238, 112, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 0, 186, 69, 61, 244, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: &Account{
				Mint:            common.PublicKeyFromString("8765cK2Vucsic6NA5nm4cfkrCzusaFVqBf6Pk31tGkXH"),
				Owner:           common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				Amount:          1049000000000,
				Delegate:        nil,
				State:           AccountStateInitialized,
				IsNative:        nil,
				DelegatedAmount: 0,
				CloseAuthority:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AccountFromData(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountFromData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountFromData() = %v, want %v", got, tt.want)
			}
		})
	}
}
