package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

type GetSlotConfig struct {
	Commitment rpc.Commitment
}

func (c GetSlotConfig) toRpc() rpc.GetSlotConfig {
	return rpc.GetSlotConfig{
		Commitment: c.Commitment,
	}
}

// GetSlot get current slot (finalized)
func (c *Client) GetSlot(ctx context.Context) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetSlot(ctx)
		},
		forward[uint64],
	)
}

// GetSlotWithConfig get slot by commitment
func (c *Client) GetSlotWithConfig(ctx context.Context, cfg GetSlotConfig) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetSlotWithConfig(ctx, cfg.toRpc())
		},
		forward[uint64],
	)
}
