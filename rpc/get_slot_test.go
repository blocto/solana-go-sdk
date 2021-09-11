package rpc

import (
	"context"
	"testing"
)

func TestGetSlot(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlot"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":78413497,"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetSlot(
					context.TODO(),
				)
			},
			ExpectedResponse: GetSlotResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: 78413497,
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlot", "params":[{"commitment": "processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":78478796,"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetSlotWithCfg(
					context.TODO(),
					GetSlotConfig{
						Commitment: CommitmentProcessed,
					},
				)
			},
			ExpectedResponse: GetSlotResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: 78478796,
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
