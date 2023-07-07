package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetFirstAvailableBlock(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFirstAvailableBlock"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":0,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetFirstAvailableBlock(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  0,
				},
				ExpectedError: nil,
			},
		},
	)
}
