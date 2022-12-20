package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetInflationRate(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getInflationRate"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"epoch":200,"foundation":0.0,"total":0.06956826778571164,"validator":0.06956826778571164},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetInflationRate(context.TODO())
				},
				ExpectedValue: JsonRpcResponse[GetInflationRate]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetInflationRate{
						Epoch:      200,
						Foundation: 0.0,
						Total:      0.06956826778571164,
						Validator:  0.06956826778571164,
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
