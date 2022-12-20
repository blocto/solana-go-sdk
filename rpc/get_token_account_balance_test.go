package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetTokenBalance(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountBalance", "params":["AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":80218700},"value":{"amount":"10000000000","decimals":9,"uiAmount":10.0,"uiAmountString":"10"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTokenAccountBalance(
						context.TODO(),
						"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
					)
				},
				ExpectedValue: JsonRpcResponse[GetTokenAccountBalance]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetTokenAccountBalance{
						Context: Context{
							Slot: 80218700,
						},
						Value: TokenAccountBalance{
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
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTokenAccountBalanceWithConfig(
						context.TODO(),
						"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
						GetTokenAccountBalanceConfig{
							Commitment: CommitmentProcessed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetTokenAccountBalance]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetTokenAccountBalance{
						Context: Context{
							Slot: 80219466,
						},
						Value: TokenAccountBalance{
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
