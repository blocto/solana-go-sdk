package client

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestClient_GetGenesisHash(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getGenesisHash"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetGenesisHash(
						context.TODO(),
					)
				},
				ExpectedValue: "EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG",
				ExpectedError: nil,
			},
		},
	)
}
