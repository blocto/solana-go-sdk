package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

type GetLatestBlockhashValue struct {
	Commitment rpc.Commitment
}

type GetLatestBlockhashConfig struct {
	Commitment rpc.Commitment
}

func (c GetLatestBlockhashConfig) toRpc() rpc.GetLatestBlockhashConfig {
	return rpc.GetLatestBlockhashConfig{
		Commitment: c.Commitment,
	}
}

// GetLatestBlockhash returns the latest blockhash
func (c *Client) GetLatestBlockhash(ctx context.Context) (rpc.GetLatestBlockhashValue, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetLatestBlockhashValue]], error) {
			return c.RpcClient.GetLatestBlockhash(ctx)
		},
		convertGetLatestBlockhash,
	)
}

// GetLatestBlockhash returns the latest blockhash
func (c *Client) GetLatestBlockhashWithConfig(ctx context.Context, cfg GetLatestBlockhashConfig) (rpc.GetLatestBlockhashValue, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetLatestBlockhashValue]], error) {
			return c.RpcClient.GetLatestBlockhashWithConfig(ctx, cfg.toRpc())
		},
		convertGetLatestBlockhash,
	)
}

// GetLatestBlockhashAndContext returns the latest blockhash
func (c *Client) GetLatestBlockhashAndContext(ctx context.Context) (rpc.ValueWithContext[rpc.GetLatestBlockhashValue], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetLatestBlockhashValue]], error) {
			return c.RpcClient.GetLatestBlockhash(ctx)
		},
		forward[rpc.ValueWithContext[rpc.GetLatestBlockhashValue]],
	)
}

// GetLatestBlockhashAndContextWithConfig returns the latest blockhash
func (c *Client) GetLatestBlockhashAndContextWithConfig(ctx context.Context, cfg GetLatestBlockhashConfig) (rpc.ValueWithContext[rpc.GetLatestBlockhashValue], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetLatestBlockhashValue]], error) {
			return c.RpcClient.GetLatestBlockhashWithConfig(ctx, cfg.toRpc())
		},
		forward[rpc.ValueWithContext[rpc.GetLatestBlockhashValue]],
	)
}

func convertGetLatestBlockhash(v rpc.ValueWithContext[rpc.GetLatestBlockhashValue]) (rpc.GetLatestBlockhashValue, error) {
	return v.Value, nil
}
