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
	return call[JsonRpcResponse[GetVersion]](c, ctx, "getVersion")
}
