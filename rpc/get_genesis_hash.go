package rpc

import (
	"context"
)

// GetGenesisHashResponse is a full raw rpc response of `getGenesisHash`
type GetGenesisHashResponse struct {
	GeneralResponse
	Result string `json:"result"`
}

// GetGenesisHash returns the genesis hash
func (c *RpcClient) GetGenesisHash(ctx context.Context) (GetGenesisHashResponse, error) {
	return c.processGetGenesisHash(c.Call(ctx, "getGenesisHash"))
}

func (c *RpcClient) processGetGenesisHash(body []byte, rpcErr error) (res GetGenesisHashResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
