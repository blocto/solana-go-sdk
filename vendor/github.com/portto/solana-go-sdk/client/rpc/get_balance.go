package rpc

import (
	"context"
)

// GetBalanceResponse is a full raw rpc response of `getBalance`
type GetBalanceResponse struct {
	GeneralResponse
	Result GetBalanceResult `json:"result"`
}

// GetBalanceResult is a part of raw rpc response of `getBalance`
type GetBalanceResult struct {
	Context Context `json:"context"`
	Value   uint64  `json:"value"`
}

// GetBalanceConfig is a option config for `getBalance`
type GetBalanceConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBalance returns the SOL balance
func (c *RpcClient) GetBalance(ctx context.Context, base58Addr string) (GetBalanceResponse, error) {
	return c.processGetBalance(c.Call(ctx, "getBalance", base58Addr))
}

// GetBalanceWithCfg returns the SOL balance
func (c *RpcClient) GetBalanceWithCfg(ctx context.Context, base58Addr string, cfg GetBalanceConfig) (GetBalanceResponse, error) {
	return c.processGetBalance(c.Call(ctx, "getBalance", base58Addr, cfg))
}

func (c *RpcClient) processGetBalance(body []byte, rpcErr error) (res GetBalanceResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
