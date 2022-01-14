package rpc

import "context"

// GetBlocksResponse is a full raw rpc response of `getBlocks`
type GetBlocksResponse struct {
	GeneralResponse
	Result []uint64 `json:"result"`
}

// GetBlocksConfig is a option config for `getBlocks`
type GetBlocksConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBlocks returns a list of confirmed blocks between two slots
// Max range allowed is 500,000 slot
func (c *RpcClient) GetBlocks(ctx context.Context, startSlot uint64, endSlot uint64) (GetBlocksResponse, error) {
	return c.processGetBlocks(c.Call(ctx, "getBlocks", startSlot, endSlot))
}

// GetBlocks returns a list of confirmed blocks between two slots
// Max range allowed is 500,000 slot
func (c *RpcClient) GetBlocksWithConfig(ctx context.Context, startSlot uint64, endSlot uint64, cfg GetBlocksConfig) (GetBlocksResponse, error) {
	return c.processGetBlocks(c.Call(ctx, "getBlocks", startSlot, endSlot, cfg))
}

func (c *RpcClient) processGetBlocks(body []byte, rpcErr error) (res GetBlocksResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
