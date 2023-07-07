package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/rpc"
)

func TestClient_IsBlockhashValid(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params": ["J24S4UPZ4drBXMbymT7uUUdRpo91jN4ooDguwhX2xVzq"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187612903},"value":true},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.IsBlockhashValid(
						context.Background(),
						"J24S4UPZ4drBXMbymT7uUUdRpo91jN4ooDguwhX2xVzq",
					)
				},
				ExpectedValue: true,
				ExpectedError: nil,
			},
		},
	)
}
func TestClient_IsBlockhashValidWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params": ["J24S4UPZ4drBXMbymT7uUUdRpo91jN4ooDguwhX2xVzq", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187612903},"value":true},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.IsBlockhashValidWithConfig(
						context.Background(),
						"J24S4UPZ4drBXMbymT7uUUdRpo91jN4ooDguwhX2xVzq",
						IsBlockhashValidConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: true,
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_IsBlockhashValidAndContext(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params": ["J24S4UPZ4drBXMbymT7uUUdRpo91jN4ooDguwhX2xVzq"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187612903},"value":true},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.IsBlockhashValidAndContext(
						context.Background(),
						"J24S4UPZ4drBXMbymT7uUUdRpo91jN4ooDguwhX2xVzq",
					)
				},
				ExpectedValue: rpc.ValueWithContext[bool]{
					Context: rpc.Context{
						Slot:       187612903,
						ApiVersion: "1.14.10",
					},
					Value: true,
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_IsBlockhashValidAndContextWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params": ["J24S4UPZ4drBXMbymT7uUUdRpo91jN4ooDguwhX2xVzq", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187612903},"value":true},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.IsBlockhashValidAndContextWithConfig(
						context.Background(),
						"J24S4UPZ4drBXMbymT7uUUdRpo91jN4ooDguwhX2xVzq",
						IsBlockhashValidConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[bool]{
					Context: rpc.Context{
						Slot:       187612903,
						ApiVersion: "1.14.10",
					},
					Value: true,
				},
				ExpectedError: nil,
			},
		},
	)
}
