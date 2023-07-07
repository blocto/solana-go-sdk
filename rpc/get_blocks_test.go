package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetBlocks(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlocks", "params":[86686567, 86686578]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[86686567,86686572,86686573,86686574,86686575,86686576,86686577,86686578],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlocks(
						context.TODO(),
						86686567,
						86686578,
					)
				},
				ExpectedValue: JsonRpcResponse[[]uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  []uint64{86686567, 86686572, 86686573, 86686574, 86686575, 86686576, 86686577, 86686578},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlocks", "params":[86686567, 86686578, {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[86686567,86686572,86686573,86686574,86686575,86686576,86686577,86686578],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlocksWithConfig(
						context.TODO(),
						86686567,
						86686578,
						GetBlocksConfig{
							Commitment: CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[[]uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  []uint64{86686567, 86686572, 86686573, 86686574, 86686575, 86686576, 86686577, 86686578},
				},
				ExpectedError: nil,
			},
		},
	)
}
