package sysprog

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

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
				from:   common.PublicKeyFromHex("ced387e6c36f57fe93ef8f516e9f318c6d89e0c51831df3d7b084e6d6e88e4f0"),
				to:     common.PublicKeyFromHex("86acd1d5e3893d6c74abcd7c36443d6e501ff0756c8961de26f2449c1b411d8e"),
				number: 1,
			},
			want: types.Instruction{
				ProgramID: SystemProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromHex("ced387e6c36f57fe93ef8f516e9f318c6d89e0c51831df3d7b084e6d6e88e4f0"),
						IsSigner:   true,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromHex("86acd1d5e3893d6c74abcd7c36443d6e501ff0756c8961de26f2449c1b411d8e"),
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
			if got := Transfer(tt.args.from, tt.args.to, tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
