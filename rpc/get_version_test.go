package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
)

func TestGetVersion(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVersion"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"feature-set":1824749018,"solana-core":"1.7.14"},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetVersion(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[GetVersion]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetVersion{
						SolanaCore: "1.7.14",
						FeatureSet: pointer.Get[uint32](1824749018),
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
