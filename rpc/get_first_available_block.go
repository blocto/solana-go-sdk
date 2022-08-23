package rpc

import (
	"context"
)

type GetFirstAvailableBlockResponse JsonRpcResponse[uint64]

// GetFirstAvailableBlock returns the slot of the lowest confirmed block that has not been purged from the ledger
func (c *RpcClient) GetFirstAvailableBlock(ctx context.Context) (JsonRpcResponse[uint64], error) {
	return c.processGetFirstAvailableBlock(c.Call(ctx, "getFirstAvailableBlock"))
}

func (c *RpcClient) processGetFirstAvailableBlock(body []byte, rpcErr error) (res JsonRpcResponse[uint64], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
