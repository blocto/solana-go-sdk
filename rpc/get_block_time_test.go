package rpc

import (
	"context"
	"testing"
)

func TestGetBlockTime(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockTime", "params":[100000]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32009,"message":"Slot 100000 was skipped, or missing in long-term storage"},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetBlockTime(
					context.TODO(),
					100000,
				)
			},
			ExpectedResponse: GetBlockTimeResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error: &ErrorResponse{
						Code:    -32009,
						Message: "Slot 100000 was skipped, or missing in long-term storage",
					},
				},
				Result: 0,
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockTime", "params":[100048426]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32004,"message":"Block not available for slot 100048426"},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetBlockTime(
					context.TODO(),
					100048426,
				)
			},
			ExpectedResponse: GetBlockTimeResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error: &ErrorResponse{
						Code:    -32004,
						Message: "Block not available for slot 100048426",
					},
				},
				Result: 0,
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockTime", "params":[85588104]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":1633531934,"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetBlockTime(
					context.TODO(),
					85588104,
				)
			},
			ExpectedResponse: GetBlockTimeResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: 1633531934,
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
