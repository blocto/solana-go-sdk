package rpc

import (
	"context"
)

type GetLatestBlockhashResponse JsonRpcResponse[GetLatestBlockhash]

type GetLatestBlockhash ValueWithContext[GetLatestBlockhashValue]

// GetLatestBlockhashResult is a part of raw rpc response of `getLatestBlockhash`
type GetLatestBlockhashValue struct {
	Blockhash              string `json:"blockhash"`
	LatestValidBlockHeight uint64 `json:"lastValidBlockHeight"`
}

// GetLatestBlockhashConfig is a option config for `getLatestBlockhash`
type GetLatestBlockhashConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getRecentBlockhash for solana-core v1.8
// GetLatestBlockhash returns the latest blockhash
func (c *RpcClient) GetLatestBlockhash(ctx context.Context) (JsonRpcResponse[ValueWithContext[GetLatestBlockhashValue]], error) {
	return call[JsonRpcResponse[ValueWithContext[GetLatestBlockhashValue]]](c, ctx, "getLatestBlockhash")
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getRecentBlockhash for solana-core v1.8
// GetLatestBlockhashWithConfig returns the latest blockhash
func (c *RpcClient) GetLatestBlockhashWithConfig(ctx context.Context, cfg GetLatestBlockhashConfig) (JsonRpcResponse[ValueWithContext[GetLatestBlockhashValue]], error) {
	return call[JsonRpcResponse[ValueWithContext[GetLatestBlockhashValue]]](c, ctx, "getLatestBlockhash", cfg)
}
