package rpc

import (
	"context"
)

type GetMinimumBalanceForRentExemptionResponse struct {
	GeneralResponse
	Result uint64 `json:"result"`
}

type GetMinimumBalanceForRentExemptionConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetMinimumBalanceForRentExemption returns minimum balance required to make account rent exempt
func (c *RpcClient) GetMinimumBalanceForRentExemption(ctx context.Context, dataLen uint64) (GetMinimumBalanceForRentExemptionResponse, error) {
	return c.processGetMinimumBalanceForRentExemption(c.Call(ctx, "getMinimumBalanceForRentExemption", dataLen))
}

// GetMinimumBalanceForRentExemptionWithCfg returns minimum balance required to make account rent exempt
func (c *RpcClient) GetMinimumBalanceForRentExemptionWithCfg(ctx context.Context, dataLen uint64, cfg GetMinimumBalanceForRentExemptionConfig) (GetMinimumBalanceForRentExemptionResponse, error) {
	return c.processGetMinimumBalanceForRentExemption(c.Call(ctx, "getMinimumBalanceForRentExemption", dataLen, cfg))
}

func (c *RpcClient) processGetMinimumBalanceForRentExemption(body []byte, rpcErr error) (res GetMinimumBalanceForRentExemptionResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
