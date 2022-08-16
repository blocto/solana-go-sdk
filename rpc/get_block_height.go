package rpc

import (
	"context"
)

type GetBlockHeightResponse JsonRpcResponse[uint64]

// GetBlockHeightConfig is a option config for `getBlockHeight`
type GetBlockHeightConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBlockHeight returns the current block height of the node
func (c *RpcClient) GetBlockHeight(ctx context.Context) (JsonRpcResponse[uint64], error) {
	return c.processGetBlockHeight(c.Call(ctx, "getBlockHeight"))
}

// GetBlockHeightWithConfig returns the current block height of the node
func (c *RpcClient) GetBlockHeightWithConfig(ctx context.Context, cfg GetBlockHeightConfig) (JsonRpcResponse[uint64], error) {
	return c.processGetBlockHeight(c.Call(ctx, "getBlockHeight", cfg))
}

func (c *RpcClient) processGetBlockHeight(body []byte, rpcErr error) (res JsonRpcResponse[uint64], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
