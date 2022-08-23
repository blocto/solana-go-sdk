package rpc

import (
	"context"
	"testing"
)

func TestGetBalance(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":73914708},"value":6999995000},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBalance(
					context.TODO(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBalance]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBalance{
					Context: Context{
						Slot: 73914708,
					},
					Value: 6999995000,
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", {"commitment": "finalized"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":73914708},"value":6999995000},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBalanceWithConfig(
					context.Background(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
					GetBalanceConfig{
						Commitment: CommitmentFinalized,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBalance]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBalance{
					Context: Context{
						Slot: 73914708,
					},
					Value: 6999995000,
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.10.34","slot":155451486},"value":114638463277},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBalance(
					context.TODO(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBalance]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBalance{
					Context: Context{
						Slot:       155451486,
						ApiVersion: "1.10.34",
					},
					Value: 114638463277,
				},
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
