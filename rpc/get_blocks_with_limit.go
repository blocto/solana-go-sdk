package rpc

import "context"

// GetBlocksWithLimitResponse is a full raw rpc response of `getBlocksWithLimit`
type GetBlocksWithLimitResponse struct {
	GeneralResponse
	Result []uint64 `json:"result"`
}

// GetBlocksWithLimitConfig is a option config for `getBlocksWithLimit`
type GetBlocksWithLimitConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// NEW: This method is only available in solana-core v1.7 or newer.
// GetBlocksWithLimit eturns a list of confirmed blocks starting at the given slot
func (c *RpcClient) GetBlocksWithLimit(ctx context.Context, startSlot uint64, limit uint64) (GetBlocksWithLimitResponse, error) {
	return c.processGetBlocksWithLimit(c.Call(ctx, "getBlocksWithLimit", startSlot, limit))
}

// NEW: This method is only available in solana-core v1.7 or newer.
// GetBlocksWithLimit eturns a list of confirmed blocks starting at the given slot
func (c *RpcClient) GetBlocksWithLimitWithConfig(ctx context.Context, startSlot uint64, limit uint64, cfg GetBlocksWithLimitConfig) (GetBlocksWithLimitResponse, error) {
	return c.processGetBlocksWithLimit(c.Call(ctx, "getBlocksWithLimit", startSlot, limit, cfg))
}

func (c *RpcClient) processGetBlocksWithLimit(body []byte, rpcErr error) (res GetBlocksWithLimitResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
