package rpc

import (
	"context"
)

type GetInflationRewardResponse JsonRpcResponse[[]*GetInflationReward]

// GetInflationRewardResult is a part of raw rpc response of `getInflationReward`
type GetInflationReward struct {
	Epoch         uint64 `json:"epoch"`
	EffectiveSlot uint64 `json:"effectiveSlot"`
	Amount        uint64 `json:"amount"`
	PostBalance   uint64 `json:"postBalance"`
	Commission    *uint8 `json:"commission"`
}

// GetInflationRewardConfig is a option config for `getInflationReward`
type GetInflationRewardConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
	Epoch      uint64     `json:"epoch,omitempty"`
}

// GetInflationReward returns the inflation reward for a list of addresses for an epoch
func (c *RpcClient) GetInflationReward(ctx context.Context, stakeAccountAddrs []string) (JsonRpcResponse[[]*GetInflationReward], error) {
	return call[JsonRpcResponse[[]*GetInflationReward]](c, ctx, "getInflationReward", stakeAccountAddrs)
}

// GetInflationRewardWithConfig returns the inflation reward for a list of addresses for an epoch
func (c *RpcClient) GetInflationRewardWithConfig(ctx context.Context, stakeAccountAddrs []string, cfg GetInflationRewardConfig) (JsonRpcResponse[[]*GetInflationReward], error) {
	return call[JsonRpcResponse[[]*GetInflationReward]](c, ctx, "getInflationReward", stakeAccountAddrs, cfg)
}
