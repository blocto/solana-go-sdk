package rpc

import (
	"context"
)

// GetInflationGovernorResponse is a full raw rpc response of `getInflationGovernor`
type GetInflationGovernorResponse struct {
	GeneralResponse
	Result GetInflationGovernorResponseResult `json:"result"`
}

// GetInflationGovernorResult is a part of raw rpc response of `getInflationGovernor`
type GetInflationGovernorResponseResult struct {
	Foundation     float64 `json:"foundation"`
	FoundationTerm float64 `json:"foundationTerm"`
	Initial        float64 `json:"initial"`
	Taper          float64 `json:"taper"`
	Terminal       float64 `json:"terminal"`
}

// GetInflationGovernorConfig is a option config for `getInflationGovernor`
type GetInflationGovernorConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetInflationGovernor returns the current inflation governor
func (c *RpcClient) GetInflationGovernor(ctx context.Context) (GetInflationGovernorResponse, error) {
	return c.processGetInflationGovernor(c.Call(ctx, "getInflationGovernor"))
}

// GetInflationGovernorWithConfig returns the current inflation governor
func (c *RpcClient) GetInflationGovernorWithConfig(ctx context.Context, cfg GetInflationGovernorConfig) (GetInflationGovernorResponse, error) {
	return c.processGetInflationGovernor(c.Call(ctx, "getInflationGovernor", cfg))
}

func (c *RpcClient) processGetInflationGovernor(body []byte, rpcErr error) (res GetInflationGovernorResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
