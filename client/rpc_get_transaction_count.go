package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

type GetTransactionCountConfig struct {
	Commitment rpc.Commitment
}

func (c GetTransactionCountConfig) toRpc() rpc.GetTransactionCountConfig {
	return rpc.GetTransactionCountConfig{
		Commitment: c.Commitment,
	}
}

// GetTransactionCount returns the current Transaction count from the ledger
func (c *Client) GetTransactionCount(ctx context.Context) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetTransactionCount(ctx)
		},
		forward[uint64],
	)
}

// GetTransactionCount returns the current Transaction count from the ledger
func (c *Client) GetTransactionCountWithConfig(ctx context.Context, cfg GetTransactionCountConfig) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetTransactionCountWithConfig(ctx, cfg.toRpc())
		},
		forward[uint64],
	)
}
