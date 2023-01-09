package client

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestClient_MinimumLedgerSlot(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"minimumLedgerSlot"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":84044778,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.MinimumLedgerSlot(
						context.TODO(),
					)
				},
				ExpectedValue: uint64(84044778),
				ExpectedError: nil,
			},
		},
	)
}
