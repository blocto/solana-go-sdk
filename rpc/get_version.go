package rpc

import "context"

type GetVersionResponse JsonRpcResponse[GetVersion]

// GetVersionResult is a part of raw rpc response of `getVersion`
type GetVersion struct {
	SolanaCore string  `json:"solana-core"`
	FeatureSet *uint32 `json:"feature-set"`
}

// GetVersion returns the current solana versions running on the node
func (c *RpcClient) GetVersion(ctx context.Context) (JsonRpcResponse[GetVersion], error) {
	return c.processGetVersion(c.Call(ctx, "getVersion"))
}

func (c *RpcClient) processGetVersion(body []byte, rpcErr error) (res JsonRpcResponse[GetVersion], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
