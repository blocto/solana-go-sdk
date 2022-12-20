package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetTransactionCount(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransactionCount"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":2168509541,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransactionCount(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  2168509541,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransactionCount", "params":[{"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":2168514398,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransactionCountWithConfig(
						context.TODO(),
						GetTransactionCountConfig{
							Commitment: CommitmentProcessed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  2168514398,
				},
				ExpectedError: nil,
			},
		},
	)
}
