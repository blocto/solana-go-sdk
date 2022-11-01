package rpc

import (
	"context"
)

type GetGenesisHashResponse JsonRpcResponse[string]

// GetGenesisHash returns the genesis hash
func (c *RpcClient) GetGenesisHash(ctx context.Context) (JsonRpcResponse[string], error) {
	return call[JsonRpcResponse[string]](c, ctx, "getGenesisHash")
}
