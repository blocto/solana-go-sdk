package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

// GetSlotLeaders returns the slot leaders for a given slot range
// (limit: 1~5000)
func (c *Client) GetSlotLeaders(ctx context.Context, startSlot uint64, limit uint64) ([]string, error) {
	return process(
		func() (rpc.JsonRpcResponse[[]string], error) {
			return c.RpcClient.GetSlotLeaders(ctx, startSlot, limit)
		},
		forward[[]string],
	)
}
