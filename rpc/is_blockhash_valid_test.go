package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestIsBlockhashValid(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params":["14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112890169},"value":false},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.IsBlockhashValid(context.TODO(), "14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N")
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[bool]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[bool]{
						Context: Context{
							Slot: 112890169,
						},
						Value: false,
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params":["14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N", {"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112890231},"value":true},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.IsBlockhashValidWithConfig(
						context.TODO(),
						"14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N",
						IsBlockhashValidConfig{
							Commitment: CommitmentProcessed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[bool]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[bool]{
						Context: Context{
							Slot: 112890231,
						},
						Value: true,
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
