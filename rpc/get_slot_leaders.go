package rpc

import "context"

type GetSlotLeadersResponse JsonRpcResponse[[]string]

// GetSlotLeaders returns the slot leaders for a given slot range
// (limit: 1~5000)
func (c *RpcClient) GetSlotLeaders(ctx context.Context, startSlot uint64, limit uint64) (JsonRpcResponse[[]string], error) {
	return call[JsonRpcResponse[[]string]](c, ctx, "getSlotLeaders", startSlot, limit)
}
