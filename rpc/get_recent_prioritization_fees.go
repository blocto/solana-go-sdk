package rpc

import (
	"context"
)

type GetRecentPrioritizationFeesResponse JsonRpcResponse[PrioritizationFees]

type PrioritizationFee struct {
	Slot              uint64 `json:"slot"`
	PrioritizationFee uint64 `json:"prioritizationFee"`
}

type PrioritizationFees []PrioritizationFee

// GetRecentPrioritizationFees returns a list of prioritization fees from recent blocks.
func (c *RpcClient) GetRecentPrioritizationFees(ctx context.Context, addresses []string) (JsonRpcResponse[PrioritizationFees], error) {
	return call[JsonRpcResponse[PrioritizationFees]](c, ctx, "getRecentPrioritizationFees", addresses)
}
