package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

type GetSignaturesForAddressConfig struct {
	Limit      int
	Before     string
	Until      string
	Commitment rpc.Commitment
}

func (c GetSignaturesForAddressConfig) toRpc() rpc.GetSignaturesForAddressConfig {
	return rpc.GetSignaturesForAddressConfig{
		Limit:      c.Limit,
		Before:     c.Before,
		Until:      c.Until,
		Commitment: c.Commitment,
	}
}

func (c *Client) GetSignaturesForAddress(ctx context.Context, addr string) (rpc.GetSignaturesForAddress, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.GetSignaturesForAddress], error) {
			return c.RpcClient.GetSignaturesForAddress(ctx, addr)
		},
		forward[rpc.GetSignaturesForAddress],
	)
}

func (c *Client) GetSignaturesForAddressWithConfig(ctx context.Context, addr string, cfg GetSignaturesForAddressConfig) (rpc.GetSignaturesForAddress, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.GetSignaturesForAddress], error) {
			return c.RpcClient.GetSignaturesForAddressWithConfig(ctx, addr, cfg.toRpc())
		},
		forward[rpc.GetSignaturesForAddress],
	)
}
