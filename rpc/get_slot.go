package rpc

import (
	"context"
)

type GetSlotResponse JsonRpcResponse[uint64]

// GetSlotConfig is a option config for `getSlot`
type GetSlotConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetSlot returns the SOL balance
func (c *RpcClient) GetSlot(ctx context.Context) (JsonRpcResponse[uint64], error) {
	return call[JsonRpcResponse[uint64]](c, ctx, "getSlot")
}

// GetSlotWithConfig returns the SOL balance
func (c *RpcClient) GetSlotWithConfig(ctx context.Context, cfg GetSlotConfig) (JsonRpcResponse[uint64], error) {
	return call[JsonRpcResponse[uint64]](c, ctx, "getSlot", cfg)
}
