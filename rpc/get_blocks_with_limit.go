package rpc

import "context"

type GetBlocksWithLimitResponse JsonRpcResponse[[]uint64]

// GetBlocksWithLimitConfig is a option config for `getBlocksWithLimit`
type GetBlocksWithLimitConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBlocksWithLimit eturns a list of confirmed blocks starting at the given slot
func (c *RpcClient) GetBlocksWithLimit(ctx context.Context, startSlot uint64, limit uint64) (JsonRpcResponse[[]uint64], error) {
	return call[JsonRpcResponse[[]uint64]](c, ctx, "getBlocksWithLimit", startSlot, limit)
}

// GetBlocksWithLimit eturns a list of confirmed blocks starting at the given slot
func (c *RpcClient) GetBlocksWithLimitWithConfig(ctx context.Context, startSlot uint64, limit uint64, cfg GetBlocksWithLimitConfig) (JsonRpcResponse[[]uint64], error) {
	return call[JsonRpcResponse[[]uint64]](c, ctx, "getBlocksWithLimit", startSlot, limit, cfg)
}
