package rpc

import (
	"context"
)

type GetTransactionCountResponse JsonRpcResponse[uint64]

// GetTransactionCountConfig is a option config for `getTransactionCount`
type GetTransactionCountConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetTransactionCount returns the current Transaction count from the ledger
func (c *RpcClient) GetTransactionCount(ctx context.Context) (JsonRpcResponse[uint64], error) {
	return call[JsonRpcResponse[uint64]](c, ctx, "getTransactionCount")
}

// GetTransactionCountWithConfig returns the current Transaction count from the ledger
func (c *RpcClient) GetTransactionCountWithConfig(ctx context.Context, cfg GetTransactionCountConfig) (JsonRpcResponse[uint64], error) {
	return call[JsonRpcResponse[uint64]](c, ctx, "getTransactionCount", cfg)
}
