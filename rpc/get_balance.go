package rpc

import (
	"context"
)

type GetBalanceResponse JsonRpcResponse[GetBalance]

type GetBalance ValueWithContext[uint64]

// GetBalanceConfig is a option config for `getBalance`
type GetBalanceConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBalance returns the SOL balance
func (c *RpcClient) GetBalance(ctx context.Context, base58Addr string) (JsonRpcResponse[ValueWithContext[uint64]], error) {
	return call[JsonRpcResponse[ValueWithContext[uint64]]](c, ctx, "getBalance", base58Addr)
}

// GetBalanceWithConfig returns the SOL balance
func (c *RpcClient) GetBalanceWithConfig(ctx context.Context, base58Addr string, cfg GetBalanceConfig) (JsonRpcResponse[ValueWithContext[uint64]], error) {
	return call[JsonRpcResponse[ValueWithContext[uint64]]](c, ctx, "getBalance", base58Addr, cfg)
}
