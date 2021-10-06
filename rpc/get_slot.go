package rpc

import (
	"context"
)

// GetSlotResponse is a full raw rpc response of `getSlot`
type GetSlotResponse struct {
	GeneralResponse
	Result uint64 `json:"result"`
}

// GetSlotConfig is a option config for `getSlot`
type GetSlotConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetSlot returns the SOL balance
func (c *RpcClient) GetSlot(ctx context.Context) (GetSlotResponse, error) {
	return c.processGetSlot(c.Call(ctx, "getSlot"))
}

// GetSlotWithConfig returns the SOL balance
func (c *RpcClient) GetSlotWithConfig(ctx context.Context, cfg GetSlotConfig) (GetSlotResponse, error) {
	return c.processGetSlot(c.Call(ctx, "getSlot", cfg))
}

func (c *RpcClient) processGetSlot(body []byte, rpcErr error) (res GetSlotResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
