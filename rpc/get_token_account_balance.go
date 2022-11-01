package rpc

import "context"

type GetTokenAccountBalanceResponse JsonRpcResponse[GetTokenAccountBalance]

type GetTokenAccountBalance struct {
	Context Context             `json:"context"`
	Value   TokenAccountBalance `json:"value"`
}

type GetTokenAccountBalanceConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *RpcClient) GetTokenAccountBalance(ctx context.Context, base58Addr string) (JsonRpcResponse[GetTokenAccountBalance], error) {
	return call[JsonRpcResponse[GetTokenAccountBalance]](c, ctx, "getTokenAccountBalance", base58Addr)
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *RpcClient) GetTokenAccountBalanceWithConfig(ctx context.Context, base58Addr string, cfg GetTokenAccountBalanceConfig) (JsonRpcResponse[GetTokenAccountBalance], error) {
	return call[JsonRpcResponse[GetTokenAccountBalance]](c, ctx, "getTokenAccountBalance", base58Addr, cfg)
}
