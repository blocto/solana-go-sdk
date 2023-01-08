package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

type IsBlockhashValidConfig struct {
	Commitment rpc.Commitment
}

func (c IsBlockhashValidConfig) toRpc() rpc.IsBlockhashValidConfig {
	return rpc.IsBlockhashValidConfig{
		Commitment: c.Commitment,
	}
}

func (c *Client) IsBlockhashValid(ctx context.Context, blockhash string) (bool, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[bool]], error) {
			return c.RpcClient.IsBlockhashValid(ctx, blockhash)
		},
		value[bool],
	)
}

func (c *Client) IsBlockhashValidWithConfig(ctx context.Context, blockhash string, cfg IsBlockhashValidConfig) (bool, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[bool]], error) {
			return c.RpcClient.IsBlockhashValidWithConfig(ctx, blockhash, cfg.toRpc())
		},
		value[bool],
	)
}

func (c *Client) IsBlockhashValidAndContext(ctx context.Context, blockhash string) (rpc.ValueWithContext[bool], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[bool]], error) {
			return c.RpcClient.IsBlockhashValid(ctx, blockhash)
		},
		forward[rpc.ValueWithContext[bool]],
	)
}

func (c *Client) IsBlockhashValidAndContextWithConfig(ctx context.Context, blockhash string, cfg IsBlockhashValidConfig) (rpc.ValueWithContext[bool], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[bool]], error) {
			return c.RpcClient.IsBlockhashValidWithConfig(ctx, blockhash, cfg.toRpc())
		},
		forward[rpc.ValueWithContext[bool]],
	)
}
