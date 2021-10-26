package rpc

import (
	"context"
)

// GetFeesResponse is a full raw rpc response of `getFees`
type GetFeesResponse struct {
	GeneralResponse
	Result GetFeesResponseResult `json:"result"`
}

// GetFeesResult is a part of raw rpc response of `getFees`
type GetFeesResponseResult struct {
	Context Context                    `json:"context"`
	Value   GetFeesResponseResultValue `json:"value"`
}

type GetFeesResponseResultValue struct {
	Blockhash            string        `json:"blockhash"`
	FeeCalculator        FeeCalculator `json:"feeCalculator"`
	LastValidSlot        uint64        `json:"lastValidSlot"`
	LastValidBlockHeight uint64        `json:"lastValidBlockHeight"`
}

// GetFeesConfig is a option config for `getFees`
type GetFeesConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetFees returns a recent block hash from the ledger,
// a fee schedule that can be used to compute the cost of submitting a transaction using it,
// and the last slot in which the blockhash will be valid
func (c *RpcClient) GetFees(ctx context.Context) (GetFeesResponse, error) {
	return c.processGetFees(c.Call(ctx, "getFees"))
}

// GetFeesWithConfig returns a recent block hash from the ledger,
// a fee schedule that can be used to compute the cost of submitting a transaction using it,
// and the last slot in which the blockhash will be valid
func (c *RpcClient) GetFeesWithConfig(ctx context.Context, cfg GetFeesConfig) (GetFeesResponse, error) {
	return c.processGetFees(c.Call(ctx, "getFees", cfg))
}

func (c *RpcClient) processGetFees(body []byte, rpcErr error) (res GetFeesResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
