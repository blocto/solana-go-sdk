package rpc

import (
	"context"
	"testing"
)

func TestGetTokenSupply(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenSupply", "params":["4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":85609218},"value":{"amount":"10000000000","decimals":9,"uiAmount":10.0,"uiAmountString":"10"}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetTokenSupply(
					context.TODO(),
					"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
				)
			},
			ExpectedResponse: JsonRpcResponse[GetTokenSupply]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetTokenSupply{
					Context: Context{
						Slot: 85609218,
					},
					Value: GetTokenSupplyResultValue{
						Amount:         "10000000000",
						Decimals:       9,
						UIAmountString: "10",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenSupply", "params":["4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3", {"commitment":"processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":85609258},"value":{"amount":"10000000000","decimals":9,"uiAmount":10.0,"uiAmountString":"10"}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetTokenSupplyWithConfig(
					context.TODO(),
					"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
					GetTokenSupplyConfig{
						Commitment: CommitmentProcessed,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetTokenSupply]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetTokenSupply{
					Context: Context{
						Slot: 85609258,
					},
					Value: GetTokenSupplyResultValue{
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
