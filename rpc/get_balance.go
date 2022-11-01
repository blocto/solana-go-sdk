package rpc

import (
	"context"
)

type GetBalanceResponse JsonRpcResponse[GetBalance]

// GetBalance is a part of raw rpc response of `getBalance`
type GetBalance struct {
	Context Context `json:"context"`
	Value   uint64  `json:"value"`
}

// GetBalanceConfig is a option config for `getBalance`
type GetBalanceConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBalance returns the SOL balance
func (c *RpcClient) GetBalance(ctx context.Context, base58Addr string) (JsonRpcResponse[GetBalance], error) {
	return call[JsonRpcResponse[GetBalance]](c, ctx, "getBalance", base58Addr)
}

// GetBalanceWithConfig returns the SOL balance
func (c *RpcClient) GetBalanceWithConfig(ctx context.Context, base58Addr string, cfg GetBalanceConfig) (JsonRpcResponse[GetBalance], error) {
	return call[JsonRpcResponse[GetBalance]](c, ctx, "getBalance", base58Addr, cfg)
}
