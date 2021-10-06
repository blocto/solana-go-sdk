package rpc

import (
	"context"
)

// GetBlockHeightResponse is a rpc response of `getBlockHeight`
type GetBlockHeightResponse struct {
	GeneralResponse
	Result uint64 `json:"result"`
}

// GetBlockHeightConfig is a option config for `getBlockHeight`
type GetBlockHeightConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBlockHeight returns the current block height of the node
func (c *RpcClient) GetBlockHeight(ctx context.Context) (GetBlockHeightResponse, error) {
	return c.processGetBlockHeight(c.Call(ctx, "getBlockHeight"))
}

// GetBlockHeightWithConfig returns the current block height of the node
func (c *RpcClient) GetBlockHeightWithConfig(ctx context.Context, cfg GetBlockHeightConfig) (GetBlockHeightResponse, error) {
	return c.processGetBlockHeight(c.Call(ctx, "getBlockHeight", cfg))
}

func (c *RpcClient) processGetBlockHeight(body []byte, rpcErr error) (res GetBlockHeightResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
