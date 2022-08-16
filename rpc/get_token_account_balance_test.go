package rpc

import (
	"context"
	"testing"
)

func TestGetTokenBalance(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountBalance", "params":["AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":80218700},"value":{"amount":"10000000000","decimals":9,"uiAmount":10.0,"uiAmountString":"10"}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetTokenAccountBalance(
					context.TODO(),
					"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
				)
			},
			ExpectedResponse: GetTokenAccountBalanceResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetTokenAccountBalanceResult{
					Context: Context{
						Slot: 80218700,
					},
					Value: GetTokenAccountBalanceResultValue{
						Amount:         "10000000000",
						Decimals:       9,
						UIAmountString: "10",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountBalance", "params":["AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ", {"commitment":"processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":80219466},"value":{"amount":"10000000000","decimals":9,"uiAmount":10.0,"uiAmountString":"10"}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetTokenAccountBalanceWithConfig(
					context.TODO(),
					"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
					GetTokenAccountBalanceConfig{
						Commitment: CommitmentProcessed,
					},
				)
			},
			ExpectedResponse: GetTokenAccountBalanceResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetTokenAccountBalanceResult{
					Context: Context{
						Slot: 80219466,
					},
					Value: GetTokenAccountBalanceResultValue{
						Amount:         "10000000000",
						Decimals:       9,
						UIAmountString: "10",
					},
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
