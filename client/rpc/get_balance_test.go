package rpc

import (
	"context"
	"testing"
)

func TestClient_GetBalance(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":73914708},"value":6999995000},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetBalance(
					context.TODO(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
				)
			},
			ExpectedResponse: GetBalanceResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetBalanceResult{
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
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetBalanceWithCfg(
					context.Background(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
					GetBalanceConfig{
						Commitment: CommitmentFinalized,
					},
				)
			},
			ExpectedResponse: GetBalanceResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetBalanceResult{
					Context: Context{
						Slot: 73914708,
					},
					Value: 6999995000,
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
