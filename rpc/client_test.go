package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_preparePayload(t *testing.T) {
	type args struct {
		params []any
	}
	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			args: args{
				params: []any{"getBalance", "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", nil},
			},
			want: `{"id":1,"jsonrpc":"2.0","method":"getBalance","params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",null]}`,
			err:  nil,
		},
		{
			args: args{
				params: []any{"getBalance", "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", GetBalanceConfig{}},
			},
			want: `{"id":1,"jsonrpc":"2.0","method":"getBalance","params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",{}]}`,
			err:  nil,
		},
		{
			args: args{
				params: []any{"getBalance", "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", GetBalanceConfig{Commitment: CommitmentFinalized}},
			},
			want: `{"id":1,"jsonrpc":"2.0","method":"getBalance","params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",{"commitment":"finalized"}]}`,
			err:  nil,
		},
		{
			args: args{
				params: []any{"getBalance"},
			},
			want: `{"id":1,"jsonrpc":"2.0","method":"getBalance"}`,
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := preparePayload(tt.args.params)
			assert.Equal(t, err, tt.err)
			assert.JSONEq(t, string(got), string(tt.want))
		})
	}
}
