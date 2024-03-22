package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/rpc"
)

func TestClient_GetRecentPrioritizationFees(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentPrioritizationFees", "params":[["CxELquR1gPP8wHe33gZ4QxqGB3sZ9RSwsJ2KshVewkFY"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"slot":348125,"prioritizationFee":0},{"slot":348126,"prioritizationFee":1000},{"slot":348127,"prioritizationFee":500},{"slot":348128,"prioritizationFee":0},{"slot":348129,"prioritizationFee":1234}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetRecentPrioritizationFees(context.Background(), []common.PublicKey{common.PublicKeyFromString("CxELquR1gPP8wHe33gZ4QxqGB3sZ9RSwsJ2KshVewkFY")})
				},
				ExpectedValue: rpc.PrioritizationFees{
					{
						Slot:              348125,
						PrioritizationFee: 0,
					},
					{
						Slot:              348126,
						PrioritizationFee: 1000,
					},
					{
						Slot:              348127,
						PrioritizationFee: 500,
					},
					{
						Slot:              348128,
						PrioritizationFee: 0,
					},
					{
						Slot:              348129,
						PrioritizationFee: 1234,
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentPrioritizationFees", "params":[[]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"slot":348125,"prioritizationFee":0},{"slot":348126,"prioritizationFee":1000},{"slot":348127,"prioritizationFee":500},{"slot":348128,"prioritizationFee":0},{"slot":348129,"prioritizationFee":1234}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetRecentPrioritizationFees(context.Background(), []common.PublicKey{})
				},
				ExpectedValue: rpc.PrioritizationFees{
					{
						Slot:              348125,
						PrioritizationFee: 0,
					},
					{
						Slot:              348126,
						PrioritizationFee: 1000,
					},
					{
						Slot:              348127,
						PrioritizationFee: 500,
					},
					{
						Slot:              348128,
						PrioritizationFee: 0,
					},
					{
						Slot:              348129,
						PrioritizationFee: 1234,
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetRecentPrioritizationFeesWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentPrioritizationFees", "params":[[],{}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"slot":348125,"prioritizationFee":0},{"slot":348126,"prioritizationFee":1000},{"slot":348127,"prioritizationFee":500},{"slot":348128,"prioritizationFee":0},{"slot":348129,"prioritizationFee":1234}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetRecentPrioritizationFeesWithConfig(
						context.Background(),
						[]common.PublicKey{},
						rpc.GetRecentPrioritizationFeesConfig{},
					)
				},
				ExpectedValue: rpc.PrioritizationFees{
					{
						Slot:              348125,
						PrioritizationFee: 0,
					},
					{
						Slot:              348126,
						PrioritizationFee: 1000,
					},
					{
						Slot:              348127,
						PrioritizationFee: 500,
					},
					{
						Slot:              348128,
						PrioritizationFee: 0,
					},
					{
						Slot:              348129,
						PrioritizationFee: 1234,
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentPrioritizationFees", "params":[["CxELquR1gPP8wHe33gZ4QxqGB3sZ9RSwsJ2KshVewkFY"], {"percentile": 5000}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"slot":348125,"prioritizationFee":0},{"slot":348126,"prioritizationFee":1000},{"slot":348127,"prioritizationFee":500},{"slot":348128,"prioritizationFee":0},{"slot":348129,"prioritizationFee":1234}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetRecentPrioritizationFeesWithConfig(
						context.Background(),
						[]common.PublicKey{common.PublicKeyFromString("CxELquR1gPP8wHe33gZ4QxqGB3sZ9RSwsJ2KshVewkFY")},
						rpc.GetRecentPrioritizationFeesConfig{Percentile: 5000},
					)
				},
				ExpectedValue: rpc.PrioritizationFees{
					{
						Slot:              348125,
						PrioritizationFee: 0,
					},
					{
						Slot:              348126,
						PrioritizationFee: 1000,
					},
					{
						Slot:              348127,
						PrioritizationFee: 500,
					},
					{
						Slot:              348128,
						PrioritizationFee: 0,
					},
					{
						Slot:              348129,
						PrioritizationFee: 1234,
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentPrioritizationFees", "params":[["CxELquR1gPP8wHe33gZ4QxqGB3sZ9RSwsJ2KshVewkFY"], {}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"slot":348125,"prioritizationFee":0},{"slot":348126,"prioritizationFee":1000},{"slot":348127,"prioritizationFee":500},{"slot":348128,"prioritizationFee":0},{"slot":348129,"prioritizationFee":1234}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetRecentPrioritizationFeesWithConfig(
						context.Background(),
						[]common.PublicKey{common.PublicKeyFromString("CxELquR1gPP8wHe33gZ4QxqGB3sZ9RSwsJ2KshVewkFY")},
						rpc.GetRecentPrioritizationFeesConfig{},
					)
				},
				ExpectedValue: rpc.PrioritizationFees{
					{
						Slot:              348125,
						PrioritizationFee: 0,
					},
					{
						Slot:              348126,
						PrioritizationFee: 1000,
					},
					{
						Slot:              348127,
						PrioritizationFee: 500,
					},
					{
						Slot:              348128,
						PrioritizationFee: 0,
					},
					{
						Slot:              348129,
						PrioritizationFee: 1234,
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentPrioritizationFees", "params":[[], {"percentile": 5000}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"slot":348125,"prioritizationFee":0},{"slot":348126,"prioritizationFee":1000},{"slot":348127,"prioritizationFee":500},{"slot":348128,"prioritizationFee":0},{"slot":348129,"prioritizationFee":1234}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetRecentPrioritizationFeesWithConfig(
						context.Background(),
						[]common.PublicKey{},
						rpc.GetRecentPrioritizationFeesConfig{Percentile: 5000},
					)
				},
				ExpectedValue: rpc.PrioritizationFees{
					{
						Slot:              348125,
						PrioritizationFee: 0,
					},
					{
						Slot:              348126,
						PrioritizationFee: 1000,
					},
					{
						Slot:              348127,
						PrioritizationFee: 500,
					},
					{
						Slot:              348128,
						PrioritizationFee: 0,
					},
					{
						Slot:              348129,
						PrioritizationFee: 1234,
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
