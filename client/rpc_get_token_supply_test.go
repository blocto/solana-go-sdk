package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/rpc"
)

func TestClient_GetTokenSupply(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenSupply", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187574619},"value":{"amount":"0","decimals":9,"uiAmount":0.0,"uiAmountString":"0"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenSupply(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					)
				},
				ExpectedValue: TokenAmount{
					Amount:         0,
					Decimals:       9,
					UIAmountString: "0",
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTokenSupplyWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenSupply", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187574619},"value":{"amount":"0","decimals":9,"uiAmount":0.0,"uiAmountString":"0"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenSupplyWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetTokenSupplyConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: TokenAmount{
					Amount:         0,
					Decimals:       9,
					UIAmountString: "0",
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTokenSupplyAndContext(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenSupply", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187574619},"value":{"amount":"0","decimals":9,"uiAmount":0.0,"uiAmountString":"0"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenSupplyAndContext(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					)
				},
				ExpectedValue: rpc.ValueWithContext[TokenAmount]{
					Context: rpc.Context{
						Slot:       187574619,
						ApiVersion: "1.14.10",
					},
					Value: TokenAmount{
						Amount:         0,
						Decimals:       9,
						UIAmountString: "0",
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTokenSupplyAndContextWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenSupply", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187574619},"value":{"amount":"0","decimals":9,"uiAmount":0.0,"uiAmountString":"0"}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenSupplyAndContextWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetTokenSupplyConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[TokenAmount]{
					Context: rpc.Context{
						Slot:       187574619,
						ApiVersion: "1.14.10",
					},
					Value: TokenAmount{
						Amount:         0,
						Decimals:       9,
						UIAmountString: "0",
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
