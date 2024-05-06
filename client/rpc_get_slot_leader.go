package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

type GetSlotLeaderConfig struct {
	Commitment     *rpc.Commitment
	MinContextSlot *uint64
}

func (c GetSlotLeaderConfig) toRpc() rpc.GetSlotLeaderConfig {
	return rpc.GetSlotLeaderConfig{
		Commitment:     c.Commitment,
		MinContextSlot: c.MinContextSlot,
	}
}

// GetSlotLeader returns the current slot leader
func (c *Client) GetSlotLeader(ctx context.Context) (string, error) {
	return process(
		func() (rpc.JsonRpcResponse[string], error) {
			return c.RpcClient.GetSlotLeader(ctx)
		},
		forward[string],
	)
}

// GetSlotWithConfig returns the current slot leader
func (c *Client) GetSlotLeaderWithConfig(ctx context.Context, cfg GetSlotLeaderConfig) (string, error) {
	return process(
		func() (rpc.JsonRpcResponse[string], error) {
			return c.RpcClient.GetSlotLeaderWithConfig(ctx, cfg.toRpc())
		},
		forward[string],
	)
}
