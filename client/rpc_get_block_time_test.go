package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
)

func TestClient_GetBlockTime(t *testing.T) {
	var nilValue *int64
	client_test.TestAll(
		t,
		[]client_test.Param{

			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockTime", "params":[85588104]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":1633531934,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetBlockTime(
						context.TODO(),
						85588104,
					)
				},
				ExpectedValue: pointer.Get[int64](1633531934),
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockTime", "params":[85588104]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":null,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetBlockTime(
						context.TODO(),
						85588104,
					)
				},
				ExpectedValue: nilValue,
				ExpectedError: nil,
			},
		},
	)
}
