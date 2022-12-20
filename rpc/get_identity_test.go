package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetIdentity(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getIdentity"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"identity":"BjHeMczor9oycGJHLepRTCU2LpkZNtpy2mdQKianx1EJ"},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetIdentity(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[GetIdentity]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetIdentity{
						Identity: "BjHeMczor9oycGJHLepRTCU2LpkZNtpy2mdQKianx1EJ",
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
