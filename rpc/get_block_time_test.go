package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/pkg/pointer"
)

func TestGetBlockTime(t *testing.T) {
	var nilValue *int64
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockTime", "params":[100000]}`,
				ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32009,"message":"Slot 100000 was skipped, or missing in long-term storage"},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockTime(
						context.TODO(),
						100000,
					)
				},
				ExpectedValue: JsonRpcResponse[*int64]{
					JsonRpc: "2.0",
					Id:      1,
					Error: &JsonRpcError{
						Code:    -32009,
						Message: "Slot 100000 was skipped, or missing in long-term storage",
					},
					Result: nilValue,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockTime", "params":[100048426]}`,
				ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32004,"message":"Block not available for slot 100048426"},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockTime(
						context.TODO(),
						100048426,
					)
				},
				ExpectedValue: JsonRpcResponse[*int64]{
					JsonRpc: "2.0",
					Id:      1,
					Error: &JsonRpcError{
						Code:    -32004,
						Message: "Block not available for slot 100048426",
					},
					Result: nilValue,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockTime", "params":[85588104]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":1633531934,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockTime(
						context.TODO(),
						85588104,
					)
				},
				ExpectedValue: JsonRpcResponse[*int64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  pointer.Get[int64](1633531934),
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockTime", "params":[85588104]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":null,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockTime(
						context.TODO(),
						85588104,
					)
				},
				ExpectedValue: JsonRpcResponse[*int64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  nilValue,
				},
				ExpectedError: nil,
			},
		},
	)
}
