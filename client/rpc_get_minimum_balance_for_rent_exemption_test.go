package client

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/rpc"
)

func TestClient_GetMinimumBalanceForRentExemption(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMinimumBalanceForRentExemption", "params":[100]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":1586880,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetMinimumBalanceForRentExemption(
						context.TODO(),
						100,
					)
				},
				ExpectedValue: uint64(1586880),
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetMinimumBalanceForRentExemptionWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMinimumBalanceForRentExemption", "params":[100, {"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":1586880,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetMinimumBalanceForRentExemptionWithConfig(
						context.TODO(),
						100,
						GetMinimumBalanceForRentExemptionConfig{
							Commitment: rpc.CommitmentProcessed,
						},
					)
				},
				ExpectedValue: uint64(1586880),
				ExpectedError: nil,
			},
		},
	)
}
