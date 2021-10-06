package rpc

import (
	"context"
)

// GetTransactionCountResponse is a full raw rpc response of `getTransactionCount`
type GetTransactionCountResponse struct {
	GeneralResponse
	Result uint64 `json:"result"`
}

// GetTransactionCountConfig is a option config for `getTransactionCount`
type GetTransactionCountConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetTransactionCount returns the current Transaction count from the ledger
func (c *RpcClient) GetTransactionCount(ctx context.Context) (GetTransactionCountResponse, error) {
	return c.processGetTransactionCount(c.Call(ctx, "getTransactionCount"))
}

// GetTransactionCountWithConfig returns the current Transaction count from the ledger
func (c *RpcClient) GetTransactionCountWithConfig(ctx context.Context, cfg GetTransactionCountConfig) (GetTransactionCountResponse, error) {
	return c.processGetTransactionCount(c.Call(ctx, "getTransactionCount", cfg))
}

func (c *RpcClient) processGetTransactionCount(body []byte, rpcErr error) (res GetTransactionCountResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
