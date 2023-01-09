package rpc

import (
	"context"
)

type GetBlockTimeResponse JsonRpcResponse[*int64]

// GetBlockTime returns the estimated production time of a block.
func (c *RpcClient) GetBlockTime(ctx context.Context, slot uint64) (JsonRpcResponse[*int64], error) {
	return call[JsonRpcResponse[*int64]](c, ctx, "getBlockTime", slot)
}
