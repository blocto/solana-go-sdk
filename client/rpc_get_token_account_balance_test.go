package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/rpc"
)

func TestClient_GetTokenAccountBalance(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountBalance", "params":["AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187578908},"value":{"amount":"9000000000","decimals":9,"uiAmount":9.0,"uiAmountString":"9"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenAccountBalance(
						context.Background(),
						"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
					)
				},
				ExpectedValue: TokenAmount{
					Amount:         9000000000,
					Decimals:       9,
					UIAmountString: "9",
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTokenAccountBalanceWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountBalance", "params":["AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187578908},"value":{"amount":"9000000000","decimals":9,"uiAmount":9.0,"uiAmountString":"9"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenAccountBalanceWithConfig(
						context.Background(),
						"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
						GetTokenAccountBalanceConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: TokenAmount{
					Amount:         9000000000,
					Decimals:       9,
					UIAmountString: "9",
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTokenAccountBalanceAndContext(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountBalance", "params":["AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187578908},"value":{"amount":"9000000000","decimals":9,"uiAmount":9.0,"uiAmountString":"9"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenAccountBalanceAndContext(
						context.Background(),
						"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
					)
				},
				ExpectedValue: rpc.ValueWithContext[TokenAmount]{
					Context: rpc.Context{
						Slot:       187578908,
						ApiVersion: "1.14.10",
					},
					Value: TokenAmount{
						Amount:         9000000000,
						Decimals:       9,
						UIAmountString: "9",
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTokenAccountBalanceAndContextWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountBalance", "params":["AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187578908},"value":{"amount":"9000000000","decimals":9,"uiAmount":9.0,"uiAmountString":"9"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenAccountBalanceAndContextWithConfig(
						context.Background(),
						"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
						GetTokenAccountBalanceConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[TokenAmount]{
					Context: rpc.Context{
						Slot:       187578908,
						ApiVersion: "1.14.10",
					},
					Value: TokenAmount{
						Amount:         9000000000,
						Decimals:       9,
						UIAmountString: "9",
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
