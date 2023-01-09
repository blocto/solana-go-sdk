package client

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/rpc"
)

func TestClient_GetIdentity(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getIdentity"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"identity":"BjHeMczor9oycGJHLepRTCU2LpkZNtpy2mdQKianx1EJ"},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetIdentity(
						context.TODO(),
					)
				},
				ExpectedValue: rpc.GetIdentity{
					Identity: "BjHeMczor9oycGJHLepRTCU2LpkZNtpy2mdQKianx1EJ",
				},
				ExpectedError: nil,
			},
		},
	)
}
