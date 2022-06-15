package memoprog

import (
	"reflect"
	"testing"

	"github.com/OldSmokeGun/solana-go-sdk/common"
	"github.com/OldSmokeGun/solana-go-sdk/types"
)

func TestBuildMemo(t *testing.T) {
	type args struct {
		param BuildMemoParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: BuildMemoParam{
					SignerPubkeys: []common.PublicKey{
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						common.PublicKeyFromString("S1gner1111111111111111111111111111111111112"),
					},
					Memo: []byte("👻"),
				},
			},
			want: types.Instruction{
				ProgramID: common.MemoProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
						IsSigner:   true,
						IsWritable: false,
					},
					{
						PubKey:     common.PublicKeyFromString("S1gner1111111111111111111111111111111111112"),
						IsSigner:   true,
						IsWritable: false,
					},
				},
				Data: []byte("👻"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildMemo(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
