package rpc

import (
	"context"
)

type GetTokenSupplyResponse JsonRpcResponse[GetTokenSupply]

type GetTokenSupply struct {
	Context Context                   `json:"context"`
	Value   GetTokenSupplyResultValue `json:"value"`
}

// GetTokenSupplyResultValue is a part of `getTokenSupply` raw response
type GetTokenSupplyResultValue struct {
	Amount         string `json:"amount"`
	Decimals       uint8  `json:"decimals"`
	UIAmountString string `json:"uiAmountString"`
}

// GetTokenSupplyConfig is option config of `getTokenSupply`
type GetTokenSupplyConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetTokenSupply returns the token balance of an SPL Token account
func (c *RpcClient) GetTokenSupply(ctx context.Context, mintAddr string) (JsonRpcResponse[GetTokenSupply], error) {
	return call[JsonRpcResponse[GetTokenSupply]](c, ctx, "getTokenSupply", mintAddr)
}

// GetTokenSupply returns the token balance of an SPL Token account
func (c *RpcClient) GetTokenSupplyWithConfig(ctx context.Context, mintAddr string, cfg GetTokenSupplyConfig) (JsonRpcResponse[GetTokenSupply], error) {
	return call[JsonRpcResponse[GetTokenSupply]](c, ctx, "getTokenSupply", mintAddr, cfg)
}
