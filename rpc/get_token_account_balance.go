package rpc

import "context"

type GetTokenAccountBalanceResponse JsonRpcResponse[ValueWithContext[TokenAccountBalance]]

type GetTokenAccountBalance ValueWithContext[TokenAccountBalance]

type GetTokenAccountBalanceConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *RpcClient) GetTokenAccountBalance(ctx context.Context, base58Addr string) (JsonRpcResponse[ValueWithContext[TokenAccountBalance]], error) {
	return call[JsonRpcResponse[ValueWithContext[TokenAccountBalance]]](c, ctx, "getTokenAccountBalance", base58Addr)
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *RpcClient) GetTokenAccountBalanceWithConfig(ctx context.Context, base58Addr string, cfg GetTokenAccountBalanceConfig) (JsonRpcResponse[ValueWithContext[TokenAccountBalance]], error) {
	return call[JsonRpcResponse[ValueWithContext[TokenAccountBalance]]](c, ctx, "getTokenAccountBalance", base58Addr, cfg)
}
