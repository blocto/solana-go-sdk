package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

// GetFirstAvailableBlock returns the slot of the lowest confirmed block that has not been purged from the ledger
func (c *Client) GetFirstAvailableBlock(ctx context.Context) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetFirstAvailableBlock(ctx)
		},
		forward[uint64],
	)
}
