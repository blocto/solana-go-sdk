package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

type GetMinimumBalanceForRentExemptionConfig struct {
	Commitment rpc.Commitment
}

func (c GetMinimumBalanceForRentExemptionConfig) toRpc() rpc.GetMinimumBalanceForRentExemptionConfig {
	return rpc.GetMinimumBalanceForRentExemptionConfig{
		Commitment: c.Commitment,
	}
}

// GetMinimumBalanceForRentExemption returns minimum balance required to make account rent exempt
func (c *Client) GetMinimumBalanceForRentExemption(ctx context.Context, dataLen uint64) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetMinimumBalanceForRentExemption(ctx, dataLen)
		},
		forward[uint64],
	)
}

// GetMinimumBalanceForRentExemption returns minimum balance required to make account rent exempt
func (c *Client) GetMinimumBalanceForRentExemptionWithConfig(ctx context.Context, dataLen uint64, cfg GetMinimumBalanceForRentExemptionConfig) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetMinimumBalanceForRentExemptionWithConfig(ctx, dataLen, cfg.toRpc())
		},
		forward[uint64],
	)
}
