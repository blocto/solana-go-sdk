package rpc

import (
	"context"
)

type GetMinimumBalanceForRentExemptionResponse JsonRpcResponse[uint64]

type GetMinimumBalanceForRentExemptionConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetMinimumBalanceForRentExemption returns minimum balance required to make account rent exempt
func (c *RpcClient) GetMinimumBalanceForRentExemption(ctx context.Context, dataLen uint64) (JsonRpcResponse[uint64], error) {
	return call[JsonRpcResponse[uint64]](c, ctx, "getMinimumBalanceForRentExemption", dataLen)
}

// GetMinimumBalanceForRentExemptionWithConfig returns minimum balance required to make account rent exempt
func (c *RpcClient) GetMinimumBalanceForRentExemptionWithConfig(ctx context.Context, dataLen uint64, cfg GetMinimumBalanceForRentExemptionConfig) (JsonRpcResponse[uint64], error) {
	return call[JsonRpcResponse[uint64]](c, ctx, "getMinimumBalanceForRentExemption", dataLen, cfg)
}
