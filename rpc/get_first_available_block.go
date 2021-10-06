package rpc

import (
	"context"
)

// GetFirstAvailableBlockResponse is a full raw rpc response of `getFirstAvailableBlock`
type GetFirstAvailableBlockResponse struct {
	GeneralResponse
	Result uint64 `json:"result"`
}

// GetFirstAvailableBlock returns the slot of the lowest confirmed block that has not been purged from the ledger
func (c *RpcClient) GetFirstAvailableBlock(ctx context.Context) (GetFirstAvailableBlockResponse, error) {
	return c.processGetFirstAvailableBlock(c.Call(ctx, "getFirstAvailableBlock"))
}

func (c *RpcClient) processGetFirstAvailableBlock(body []byte, rpcErr error) (res GetFirstAvailableBlockResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
