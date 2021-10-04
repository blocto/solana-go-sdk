package rpc

import "context"

type GetTokenAccountBalanceResponse struct {
	GeneralResponse
	Result GetTokenAccountBalanceResult `json:"result"`
}

type GetTokenAccountBalanceResult struct {
	Context Context                           `json:"context"`
	Value   GetTokenAccountBalanceResultValue `json:"value"`
}

type GetTokenAccountBalanceResultValue struct {
	Amount         string `json:"amount"`
	Decimals       uint8  `json:"decimals"`
	UIAmountString string `json:"uiAmountString"`
}

type GetTokenAccountBalanceConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *RpcClient) GetTokenAccountBalance(ctx context.Context, base58Addr string) (GetTokenAccountBalanceResponse, error) {
	return c.processGetTokenAccountBalance(c.Call(ctx, "getTokenAccountBalance", base58Addr))
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *RpcClient) GetTokenAccountBalanceWithCfg(ctx context.Context, base58Addr string, cfg GetTokenAccountBalanceConfig) (GetTokenAccountBalanceResponse, error) {
	return c.processGetTokenAccountBalance(c.Call(ctx, "getTokenAccountBalance", base58Addr, cfg))
}

func (c *RpcClient) processGetTokenAccountBalance(body []byte, rpcErr error) (res GetTokenAccountBalanceResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
