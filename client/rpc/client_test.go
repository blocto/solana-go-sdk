package rpc

import (
	"reflect"
	"testing"
)

func Test_preparePayload(t *testing.T) {
	type args struct {
		params []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{
				params: []interface{}{"getBalance", "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", nil},
			},
			want:    []byte(`{"id":1,"jsonrpc":"2.0","method":"getBalance","params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",null]}`),
			wantErr: false,
		},
		{
			args: args{
				params: []interface{}{"getBalance", "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", GetBalanceConfig{}},
			},
			want:    []byte(`{"id":1,"jsonrpc":"2.0","method":"getBalance","params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",{}]}`),
			wantErr: false,
		},
		{
			args: args{
				params: []interface{}{"getBalance", "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", GetBalanceConfig{Commitment: CommitmentFinalized}},
			},
			want:    []byte(`{"id":1,"jsonrpc":"2.0","method":"getBalance","params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",{"commitment":"finalized"}]}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := preparePayload(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("preparePayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preparePayload() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
