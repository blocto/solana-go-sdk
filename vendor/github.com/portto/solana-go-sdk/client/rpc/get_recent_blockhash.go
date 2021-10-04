package rpc

import (
	"context"
)

// GetRecentBlockHashResponse is full raw response of `getRecentBlockhash`
type GetRecentBlockHashResponse struct {
	GeneralResponse
	Result GetRecentBlockHashResult `json:"result"`
}

// GetRecentBlockHashResult is part of response of `getRecentBlockhash`
type GetRecentBlockHashResult struct {
	Context Context                       `json:"context"`
	Value   GetRecentBlockHashResultValue `json:"value"`
}

// GetRecentBlockHashResultValue is part of response of `getRecentBlockhash`
type GetRecentBlockHashResultValue struct {
	Blockhash     string        `json:"blockhash"`
	FeeCalculator FeeCalculator `json:"feeCalculator"`
}

// FeeCalculator is a list of fee
type FeeCalculator struct {
	LamportsPerSignature uint64 `json:"lamportsPerSignature"`
}

// GetRecentBlockhashConfig is a option config for `getRecentBlockhash`
type GetRecentBlockhashConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// getRecentBlockhash returns a recent block hash from the ledger, and a fee schedule that can be used to compute
// the cost of submitting a transaction using it.
func (c *RpcClient) GetRecentBlockhash(ctx context.Context) (GetRecentBlockHashResponse, error) {
	return c.processGetRecentBlockhash(c.Call(ctx, "getRecentBlockhash"))
}

// getRecentBlockhash returns a recent block hash from the ledger, and a fee schedule that can be used to compute
// the cost of submitting a transaction using it.
func (c *RpcClient) GetRecentBlockhashWithConfig(ctx context.Context, cfg GetRecentBlockhashConfig) (GetRecentBlockHashResponse, error) {
	return c.processGetRecentBlockhash(c.Call(ctx, "getRecentBlockhash", cfg))
}

func (c *RpcClient) processGetRecentBlockhash(body []byte, rpcErr error) (res GetRecentBlockHashResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
