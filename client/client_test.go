package client

import (
	"errors"
	"testing"

	"github.com/portto/solana-go-sdk/rpc"
	"github.com/stretchr/testify/assert"
)

func Test_checkJsonRpcResponse(t *testing.T) {
	type args struct {
		res rpc.JsonRpcResponse[rpc.GetBlock]
		err error
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			args: args{
				res: rpc.JsonRpcResponse[rpc.GetBlock]{
					JsonRpc: "2.0",
					Id:      1,
					Result:  rpc.GetBlock{},
				},
				err: nil,
			},
			expectedErr: nil,
		},
		{
			args: args{
				res: rpc.JsonRpcResponse[rpc.GetBlock]{
					JsonRpc: "2.0",
					Id:      1,
					Result:  rpc.GetBlock{},
				},
				err: errors.New("rpc error"),
			},
			expectedErr: errors.New("rpc error"),
		},
		{
			args: args{
				res: rpc.JsonRpcResponse[rpc.GetBlock]{
					JsonRpc: "2.0",
					Id:      1,
					Result:  rpc.GetBlock{},
					Error: &rpc.JsonRpcError{
						Code:    -1,
						Message: "error",
						Data:    nil,
					},
				},
				err: nil,
			},
			expectedErr: &rpc.JsonRpcError{
				Code:    -1,
				Message: "error",
				Data:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkJsonRpcResponse(tt.args.res, tt.args.err)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
