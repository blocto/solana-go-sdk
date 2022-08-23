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
	return c.processGetIdentity(c.Call(ctx, "getIdentity"))
}

func (c *RpcClient) processGetIdentity(body []byte, rpcErr error) (res JsonRpcResponse[GetIdentity], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
