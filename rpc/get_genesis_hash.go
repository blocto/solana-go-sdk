package rpc

import (
	"context"
)

type GetGenesisHashResponse JsonRpcResponse[string]

// GetGenesisHash returns the genesis hash
func (c *RpcClient) GetGenesisHash(ctx context.Context) (JsonRpcResponse[string], error) {
	return c.processGetGenesisHash(c.Call(ctx, "getGenesisHash"))
}

func (c *RpcClient) processGetGenesisHash(body []byte, rpcErr error) (res JsonRpcResponse[string], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
