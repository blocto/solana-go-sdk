package client

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/rpc"
)

func TestClient_GetLatestBlockhash(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187545846},"value":{"blockhash":"DjQ4csyDJ9ZQvNNbK838ATs5UrqMq8s4Pd5i1ts22HAQ","lastValidBlockHeight":177067026}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetLatestBlockhash(context.Background())
				},
				ExpectedValue: rpc.GetLatestBlockhashValue{
					Blockhash:              "DjQ4csyDJ9ZQvNNbK838ATs5UrqMq8s4Pd5i1ts22HAQ",
					LatestValidBlockHeight: 177067026,
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetLatestBlockhashWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash", "params":[{"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187545846},"value":{"blockhash":"DjQ4csyDJ9ZQvNNbK838ATs5UrqMq8s4Pd5i1ts22HAQ","lastValidBlockHeight":177067026}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetLatestBlockhashWithConfig(
						context.Background(),
						GetLatestBlockhashConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: rpc.GetLatestBlockhashValue{
					Blockhash:              "DjQ4csyDJ9ZQvNNbK838ATs5UrqMq8s4Pd5i1ts22HAQ",
					LatestValidBlockHeight: 177067026,
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetLatestBlockhashAndContext(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187545846},"value":{"blockhash":"DjQ4csyDJ9ZQvNNbK838ATs5UrqMq8s4Pd5i1ts22HAQ","lastValidBlockHeight":177067026}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetLatestBlockhashAndContext(context.Background())
				},
				ExpectedValue: rpc.ValueWithContext[rpc.GetLatestBlockhashValue]{
					Context: rpc.Context{
						Slot:       187545846,
						ApiVersion: "1.14.10",
					},
					Value: rpc.GetLatestBlockhashValue{
						Blockhash:              "DjQ4csyDJ9ZQvNNbK838ATs5UrqMq8s4Pd5i1ts22HAQ",
						LatestValidBlockHeight: 177067026,
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetLatestBlockhashAndContextWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash", "params":[{"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187545846},"value":{"blockhash":"DjQ4csyDJ9ZQvNNbK838ATs5UrqMq8s4Pd5i1ts22HAQ","lastValidBlockHeight":177067026}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetLatestBlockhashAndContextWithConfig(
						context.Background(),
						GetLatestBlockhashConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[rpc.GetLatestBlockhashValue]{
					Context: rpc.Context{
						Slot:       187545846,
						ApiVersion: "1.14.10",
					},
					Value: rpc.GetLatestBlockhashValue{
						Blockhash:              "DjQ4csyDJ9ZQvNNbK838ATs5UrqMq8s4Pd5i1ts22HAQ",
						LatestValidBlockHeight: 177067026,
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
