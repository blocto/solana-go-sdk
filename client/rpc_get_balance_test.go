package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/rpc"
)

func TestClient_GetBalance(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187552526},"value":2039280},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetBalance(
						context.TODO(),
						"CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK",
					)
				},
				ExpectedValue: uint64(2039280),
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetBalanceWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187552526},"value":2039280},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetBalanceWithConfig(
						context.TODO(),
						"CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK",
						GetBalanceConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: uint64(2039280),
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetBalanceAndContext(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187552526},"value":2039280},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetBalanceAndContext(
						context.TODO(),
						"CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK",
					)
				},
				ExpectedValue: rpc.ValueWithContext[uint64]{
					Context: rpc.Context{
						Slot:       187552526,
						ApiVersion: "1.14.10",
					},
					Value: uint64(2039280),
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetBalanceAndContextWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187552526},"value":2039280},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetBalanceAndContextWithConfig(
						context.TODO(),
						"CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK",
						GetBalanceConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[uint64]{
					Context: rpc.Context{
						Slot:       187552526,
						ApiVersion: "1.14.10",
					},
					Value: uint64(2039280),
				},
				ExpectedError: nil,
			},
		},
	)
}
