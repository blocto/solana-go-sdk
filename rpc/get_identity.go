package rpc

import (
	"context"
)

// GetIdentityResponse is a full raw rpc response of `getIdentity`
type GetIdentityResponse struct {
	GeneralResponse
	Result GetIdentityResult `json:"result"`
}

// GetIdentityResult is a part of raw rpc response of `getIdentity`
type GetIdentityResult struct {
	Identity string `json:"identity"`
}

// GetIdentity returns the identity pubkey for the current node
func (c *RpcClient) GetIdentity(ctx context.Context) (GetIdentityResponse, error) {
	return c.processGetIdentity(c.Call(ctx, "getIdentity"))
}

func (c *RpcClient) processGetIdentity(body []byte, rpcErr error) (res GetIdentityResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
