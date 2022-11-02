package rpc

import (
	"context"
)

type GetFirstAvailableBlockResponse JsonRpcResponse[uint64]

// GetFirstAvailableBlock returns the slot of the lowest confirmed block that has not been purged from the ledger
func (c *RpcClient) GetFirstAvailableBlock(ctx context.Context) (JsonRpcResponse[uint64], error) {
	return call[JsonRpcResponse[uint64]](c, ctx, "getFirstAvailableBlock")
}
