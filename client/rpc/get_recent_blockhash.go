package rpc

import (
	"context"
	"encoding/json"
	"fmt"
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

func (c *RpcClient) processGetRecentBlockhash(body []byte, err error) (GetRecentBlockHashResponse, error) {
	if err != nil {
		return GetRecentBlockHashResponse{}, fmt.Errorf("rpc: call error, err: %v", err)
	}
	var res GetRecentBlockHashResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return GetRecentBlockHashResponse{}, fmt.Errorf("rpc: failed to json decode body, err: %v", err)
	}
	return res, nil
}
