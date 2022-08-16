package rpc

import (
	"context"
	"testing"
)

func TestGetBlocks(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlocks", "params":[86686567, 86686578]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":[86686567,86686572,86686573,86686574,86686575,86686576,86686577,86686578],"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlocks(
					context.TODO(),
					86686567,
					86686578,
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlocks]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result:  []uint64{86686567, 86686572, 86686573, 86686574, 86686575, 86686576, 86686577, 86686578},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlocks", "params":[86686567, 86686578, {"commitment": "confirmed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":[86686567,86686572,86686573,86686574,86686575,86686576,86686577,86686578],"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlocksWithConfig(
					context.TODO(),
					86686567,
					86686578,
					GetBlocksConfig{
						Commitment: CommitmentConfirmed,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlocks]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result:  []uint64{86686567, 86686572, 86686573, 86686574, 86686575, 86686576, 86686577, 86686578},
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			testRpcCall(t, tt)
		})
	}
}
