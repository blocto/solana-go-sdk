package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetBlocksWithLimit(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlocksWithLimit", "params":[86686567, 15]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[86686567,86686572,86686573,86686574,86686575,86686576,86686577,86686578,86686579,86686580,86686581,86686582,86686583,86686584,86686585],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlocksWithLimit(
						context.TODO(),
						86686567,
						15,
					)
				},
				ExpectedValue: JsonRpcResponse[[]uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  []uint64{86686567, 86686572, 86686573, 86686574, 86686575, 86686576, 86686577, 86686578, 86686579, 86686580, 86686581, 86686582, 86686583, 86686584, 86686585},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlocksWithLimit", "params":[86686567, 15, {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[86686567,86686572,86686573,86686574,86686575,86686576,86686577,86686578,86686579,86686580,86686581,86686582,86686583,86686584,86686585],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlocksWithLimitWithConfig(
						context.TODO(),
						86686567,
						15,
						GetBlocksWithLimitConfig{
							Commitment: CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[[]uint64]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  []uint64{86686567, 86686572, 86686573, 86686574, 86686575, 86686576, 86686577, 86686578, 86686579, 86686580, 86686581, 86686582, 86686583, 86686584, 86686585},
				},
				ExpectedError: nil,
			},
		},
	)
}
