package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetTokenSupply(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenSupply", "params":["4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":85609218},"value":{"amount":"10000000000","decimals":9,"uiAmount":10.0,"uiAmountString":"10"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTokenSupply(
						context.TODO(),
						"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetTokenSupplyResultValue]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetTokenSupplyResultValue]{
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
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTokenSupplyWithConfig(
						context.TODO(),
						"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
						GetTokenSupplyConfig{
							Commitment: CommitmentProcessed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetTokenSupplyResultValue]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetTokenSupplyResultValue]{
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
		},
	)
}
