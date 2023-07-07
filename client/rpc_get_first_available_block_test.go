package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestClient_GetFirstAvailableBlock(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFirstAvailableBlock"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":100,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetFirstAvailableBlock(
						context.TODO(),
					)
				},
				ExpectedValue: uint64(100),
				ExpectedError: nil,
			},
		},
	)
}
