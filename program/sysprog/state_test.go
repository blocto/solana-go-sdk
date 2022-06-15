package sysprog

import (
	"reflect"
	"testing"

	"github.com/OldSmokeGun/solana-go-sdk/common"
)

func TestNonceAccountDeserialize(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    NonceAccount
		wantErr bool
	}{
		{
			args: args{
				data: []byte{0, 0, 0, 0, 1, 0, 0, 0, 170, 118, 78, 20, 110, 21, 146, 201, 207, 34, 55, 190, 100, 27, 130, 117, 252, 159, 223, 230, 13, 166, 95, 130, 155, 86, 34, 134, 87, 106, 160, 233, 118, 21, 129, 71, 191, 98, 171, 247, 177, 47, 125, 104, 215, 37, 254, 44, 68, 82, 208, 182, 201, 123, 37, 207, 233, 116, 103, 34, 74, 217, 164, 8, 136, 19, 0, 0, 0, 0, 0, 0},
			},
			want: NonceAccount{
				Version:          0,
				State:            1,
				AuthorizedPubkey: common.PublicKeyFromString("CUQwQyNDPdGM2KfC7B4NJhrSwDwRjdqKetpwBHe9CvEk"),
				Nonce:            common.PublicKeyFromString("8wx8PoVMibdYTrfweG2wCFuYz7EhwkaZLm8hutyFgh8T"),
				FeeCalculator: FeeCalculator{
					LamportsPerSignature: 5000,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NonceAccountDeserialize(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NonceAccountDeserialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NonceAccountDeserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
