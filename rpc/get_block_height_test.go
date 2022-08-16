package rpc

import (
	"context"
	"testing"
)

func TestGetBlockHeight(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockHeight"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":83518197,"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockHeight(context.TODO())
			},
			ExpectedResponse: GetBlockHeightResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: 83518197,
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockHeight", "params":[{"commitment": "confirmed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":83518231,"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockHeightWithConfig(
					context.Background(),
					GetBlockHeightConfig{
						Commitment: CommitmentConfirmed,
					},
				)
			},
			ExpectedResponse: GetBlockHeightResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: 83518231,
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
