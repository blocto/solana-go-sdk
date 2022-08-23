package rpc

import (
	"context"
)

type GetSlotResponse JsonRpcResponse[uint64]

// GetSlotConfig is a option config for `getSlot`
type GetSlotConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetSlot returns the SOL balance
func (c *RpcClient) GetSlot(ctx context.Context) (JsonRpcResponse[uint64], error) {
	return c.processGetSlot(c.Call(ctx, "getSlot"))
}

// GetSlotWithConfig returns the SOL balance
func (c *RpcClient) GetSlotWithConfig(ctx context.Context, cfg GetSlotConfig) (JsonRpcResponse[uint64], error) {
	return c.processGetSlot(c.Call(ctx, "getSlot", cfg))
}

func (c *RpcClient) processGetSlot(body []byte, rpcErr error) (res JsonRpcResponse[uint64], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
