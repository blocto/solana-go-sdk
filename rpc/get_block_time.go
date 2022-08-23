package rpc

import (
	"context"
)

type GetBlockTimeResponse JsonRpcResponse[int64]

// GetBlockTime returns the estimated production time of a block.
func (c *RpcClient) GetBlockTime(ctx context.Context, slot uint64) (JsonRpcResponse[int64], error) {
	return c.processGetBlockTime(c.Call(ctx, "getBlockTime", slot))
}

func (c *RpcClient) processGetBlockTime(body []byte, rpcErr error) (res JsonRpcResponse[int64], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
