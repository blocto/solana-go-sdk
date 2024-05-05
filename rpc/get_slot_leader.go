package rpc

import (
	"context"
)

type GetSlotLeaderResponse JsonRpcResponse[string]

// GetSlotLeaderConfig is a option config for `getSlotLeader`
type GetSlotLeaderConfig struct {
	Commitment     *Commitment `json:"commitment,omitempty"`
	MinContextSlot *uint64     `json:"minContextSlot,omitempty"`
}

// GetSlotLeader returns the current slot leader
func (c *RpcClient) GetSlotLeader(ctx context.Context) (JsonRpcResponse[string], error) {
	return call[JsonRpcResponse[string]](c, ctx, "getSlotLeader")
}

// GetSlotWithConfig returns the current slot leader
func (c *RpcClient) GetSlotLeaderWithConfig(ctx context.Context, cfg GetSlotLeaderConfig) (JsonRpcResponse[string], error) {
	return call[JsonRpcResponse[string]](c, ctx, "getSlotLeader", cfg)
}
