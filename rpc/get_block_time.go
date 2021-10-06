package rpc

import (
	"context"
)

// GetBlockTimeResponse is a full raw rpc response of `getBlockTime`
type GetBlockTimeResponse struct {
	GeneralResponse
	Result int64 `json:"result"`
}

// GetBlockTime returns the estimated production time of a block.
func (c *RpcClient) GetBlockTime(ctx context.Context, slot uint64) (GetBlockTimeResponse, error) {
	return c.processGetBlockTime(c.Call(ctx, "getBlockTime", slot))
}

func (c *RpcClient) processGetBlockTime(body []byte, rpcErr error) (res GetBlockTimeResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
