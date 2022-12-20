package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetMinimumBalanceForRentExemption(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMinimumBalanceForRentExemption", "params":[100]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":1586880,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMinimumBalanceForRentExemption(
						context.TODO(),
						100,
					)
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  1586880,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMinimumBalanceForRentExemption", "params":[100, {"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":1586880,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMinimumBalanceForRentExemptionWithConfig(
						context.TODO(),
						100,
						GetMinimumBalanceForRentExemptionConfig{
							Commitment: CommitmentProcessed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  1586880,
				},
				ExpectedError: nil,
			},
		},
	)
}
