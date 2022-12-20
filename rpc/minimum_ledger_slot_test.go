package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestMinimumLedgerSlot(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"minimumLedgerSlot"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":84044778,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.MinimumLedgerSlot(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  84044778,
				},
				ExpectedError: nil,
			},
		},
	)
}
