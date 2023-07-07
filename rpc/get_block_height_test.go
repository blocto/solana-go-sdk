package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetBlockHeight(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockHeight"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":83518197,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockHeight(context.TODO())
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  83518197,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockHeight", "params":[{"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":83518231,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockHeightWithConfig(
						context.Background(),
						GetBlockHeightConfig{
							Commitment: CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  83518231,
				},
				ExpectedError: nil,
			},
		},
	)
}
