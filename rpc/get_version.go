package rpc

import "context"

// GetVersionResponse is a full raw rpc response of `getVersion`
type GetVersionResponse struct {
	GeneralResponse
	Result GetVersionResult `json:"result"`
}

// GetVersionResult is a part of raw rpc response of `getVersion`
type GetVersionResult struct {
	SolanaCore string  `json:"solana-core"`
	FeatureSet *uint32 `json:"feature-set"`
}

// GetVersion returns the current solana versions running on the node
func (c *RpcClient) GetVersion(ctx context.Context) (GetVersionResponse, error) {
	return c.processGetVersion(c.Call(ctx, "getVersion"))
}

func (c *RpcClient) processGetVersion(body []byte, rpcErr error) (res GetVersionResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
