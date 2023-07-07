package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/rpc"
)

func TestClient_GetSlot(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlot"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":187548524,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSlot(
						context.Background(),
					)
				},
				ExpectedValue: uint64(187548524),
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetSlotWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlot", "params":[{"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":187548524,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSlotWithConfig(
						context.Background(),
						GetSlotConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: uint64(187548524),
				ExpectedError: nil,
			},
		},
	)
}
