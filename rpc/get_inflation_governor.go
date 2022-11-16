package rpc

import (
	"context"
)

type GetInflationGovernorResponse JsonRpcResponse[GetInflationGovernor]

type GetInflationGovernor struct {
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
func (c *RpcClient) GetInflationGovernor(ctx context.Context) (JsonRpcResponse[GetInflationGovernor], error) {
	return call[JsonRpcResponse[GetInflationGovernor]](c, ctx, "getInflationGovernor")
}

// GetInflationGovernorWithConfig returns the current inflation governor
func (c *RpcClient) GetInflationGovernorWithConfig(ctx context.Context, cfg GetInflationGovernorConfig) (JsonRpcResponse[GetInflationGovernor], error) {
	return call[JsonRpcResponse[GetInflationGovernor]](c, ctx, "getInflationGovernor", cfg)
}
