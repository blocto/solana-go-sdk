package client

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/rpc"
)

func TestClient_GetTransactionCount(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransactionCount"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":2168509541,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTransactionCount(
						context.TODO(),
					)
				},
				ExpectedValue: uint64(2168509541),
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTransactionCountWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransactionCount", "params":[{"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":2168514398,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTransactionCountWithConfig(
						context.TODO(),
						GetTransactionCountConfig{
							Commitment: rpc.CommitmentProcessed,
						},
					)
				},
				ExpectedValue: uint64(2168514398),
				ExpectedError: nil,
			},
		},
	)
}
