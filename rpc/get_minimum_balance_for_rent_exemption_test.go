package rpc

import (
	"context"
	"testing"
)

func TestGetMinimumBalanceForRentExemption(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMinimumBalanceForRentExemption", "params":[100]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":1586880,"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetMinimumBalanceForRentExemption(
					context.TODO(),
					100,
				)
			},
			ExpectedResponse: JsonRpcResponse[uint64]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result:  1586880,
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMinimumBalanceForRentExemption", "params":[100, {"commitment": "processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":1586880,"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetMinimumBalanceForRentExemptionWithConfig(
					context.TODO(),
					100,
					GetMinimumBalanceForRentExemptionConfig{
						Commitment: CommitmentProcessed,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[uint64]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result:  1586880,
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
