package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

type GetBalanceConfig struct {
	Commitment rpc.Commitment
}

func (c GetBalanceConfig) toRpc() rpc.GetBalanceConfig {
	return rpc.GetBalanceConfig{
		Commitment: c.Commitment,
	}
}

// GetBalance fetch users lamports(SOL) balance
func (c *Client) GetBalance(ctx context.Context, base58Addr string) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[uint64]], error) {
			return c.RpcClient.GetBalance(ctx, base58Addr)
		},
		value[uint64],
	)
}

// GetBalanceWithConfig fetch users lamports(SOL) balance with specific commitment
func (c *Client) GetBalanceWithConfig(ctx context.Context, base58Addr string, cfg GetBalanceConfig) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[uint64]], error) {
			return c.RpcClient.GetBalanceWithConfig(ctx, base58Addr, cfg.toRpc())
		},
		value[uint64],
	)
}

// GetBalanceAndContext fetch users lamports(SOL) balance
func (c *Client) GetBalanceAndContext(ctx context.Context, base58Addr string) (rpc.ValueWithContext[uint64], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[uint64]], error) {
			return c.RpcClient.GetBalance(ctx, base58Addr)
		},
		forward[rpc.ValueWithContext[uint64]],
	)
}

// GetBalanceAndContextWithConfig fetch users lamports(SOL) balance with specific commitment
func (c *Client) GetBalanceAndContextWithConfig(ctx context.Context, base58Addr string, cfg GetBalanceConfig) (rpc.ValueWithContext[uint64], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[uint64]], error) {
			return c.RpcClient.GetBalanceWithConfig(ctx, base58Addr, cfg.toRpc())
		},
		forward[rpc.ValueWithContext[uint64]],
	)
}
