package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetHealth(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getHealth"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"ok","id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetHealth(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  "ok",
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getHealth"}`,
				ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32005,"message":"Node is behind by 42 slots","data":{"numSlotsBehind":42}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetHealth(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error: &JsonRpcError{
						Code:    -32005,
						Message: `Node is behind by 42 slots`,
						Data: map[string]any{
							"numSlotsBehind": float64(42),
						},
					},
					Result: "",
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getHealth"}`,
				ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32005,"message":"Node is unhealthy","data":{}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetHealth(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error: &JsonRpcError{
						Code:    -32005,
						Message: `Node is unhealthy`,
						Data:    map[string]any{},
					},
					Result: "",
				},
				ExpectedError: nil,
			},
		},
	)
}
