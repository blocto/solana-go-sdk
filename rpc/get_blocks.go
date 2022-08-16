package rpc

import "context"

type GetBlocksResponse JsonRpcResponse[[]uint64]

// GetBlocksConfig is a option config for `getBlocks`
type GetBlocksConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBlocks returns a list of confirmed blocks between two slots
// Max range allowed is 500,000 slot
func (c *RpcClient) GetBlocks(ctx context.Context, startSlot uint64, endSlot uint64) (JsonRpcResponse[[]uint64], error) {
	return c.processGetBlocks(c.Call(ctx, "getBlocks", startSlot, endSlot))
}

// GetBlocks returns a list of confirmed blocks between two slots
// Max range allowed is 500,000 slot
func (c *RpcClient) GetBlocksWithConfig(ctx context.Context, startSlot uint64, endSlot uint64, cfg GetBlocksConfig) (JsonRpcResponse[[]uint64], error) {
	return c.processGetBlocks(c.Call(ctx, "getBlocks", startSlot, endSlot, cfg))
}

func (c *RpcClient) processGetBlocks(body []byte, rpcErr error) (res JsonRpcResponse[[]uint64], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
