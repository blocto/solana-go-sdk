package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetSlot(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlot"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":78413497,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSlot(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  78413497,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlot", "params":[{"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":78478796,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSlotWithConfig(
						context.TODO(),
						GetSlotConfig{
							Commitment: CommitmentProcessed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  78478796,
				},
				ExpectedError: nil,
			},
		},
	)
}
