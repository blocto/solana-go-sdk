package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetGenesisHash(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getGenesisHash"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG","id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetGenesisHash(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  "EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG",
				},
				ExpectedError: nil,
			},
		},
	)
}
