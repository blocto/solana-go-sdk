package rpc

import (
	"context"
	"testing"
)

func TestGetTransactionCount(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransactionCount"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":2168509541,"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetTransactionCount(
					context.TODO(),
				)
			},
			ExpectedResponse: GetTransactionCountResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: 2168509541,
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransactionCount", "params":[{"commitment": "processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":2168514398,"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetTransactionCountWithConfig(
					context.TODO(),
					GetTransactionCountConfig{
						Commitment: CommitmentProcessed,
					},
				)
			},
			ExpectedResponse: GetTransactionCountResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: 2168514398,
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
