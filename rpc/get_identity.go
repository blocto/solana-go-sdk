package rpc

import (
	"context"
)

type GetIdentityResponse JsonRpcResponse[GetIdentity]

type GetIdentity struct {
	Identity string `json:"identity"`
}

// GetIdentity returns the identity pubkey for the current node
func (c *RpcClient) GetIdentity(ctx context.Context) (JsonRpcResponse[GetIdentity], error) {
	return call[JsonRpcResponse[GetIdentity]](c, ctx, "getIdentity")
}
