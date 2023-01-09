package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

type RequestAirdropConfig struct {
	Commitment rpc.Commitment
}

func (c RequestAirdropConfig) toRpc() rpc.RequestAirdropConfig {
	return rpc.RequestAirdropConfig{
		Commitment: c.Commitment,
	}
}

// RequestAirdrop requests an airdrop of lamports to a Pubkey
func (c *Client) RequestAirdrop(ctx context.Context, base58Addr string, lamports uint64) (string, error) {
	return process(
		func() (rpc.JsonRpcResponse[string], error) {
			return c.RpcClient.RequestAirdrop(ctx, base58Addr, lamports)
		},
		forward[string],
	)
}

// RequestAirdrop requests an airdrop of lamports to a Pubkey
func (c *Client) RequestAirdropWithConfig(ctx context.Context, base58Addr string, lamports uint64, cfg RequestAirdropConfig) (string, error) {
	return process(
		func() (rpc.JsonRpcResponse[string], error) {
			return c.RpcClient.RequestAirdropWithConfig(
				ctx,
				base58Addr,
				lamports,
				cfg.toRpc(),
			)
		},
		forward[string],
	)
}
