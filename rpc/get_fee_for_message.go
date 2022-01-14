package rpc

import (
	"context"
)

// GetFeeForMessageResponse is a full raw rpc response of `getFeeForMessage`
type GetFeeForMessageResponse struct {
	GeneralResponse
	Result GetFeeForMessageResult `json:"result"`
}

// GetFeeForMessageResult is a part of raw rpc response of `getFeeForMessage`
type GetFeeForMessageResult struct {
	Context Context `json:"context"`
	Value   *uint64 `json:"value"`
}

// GetFeeForMessageConfig is a option config for `GetFeeForMessage`
type GetFeeForMessageConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8
// GetFeeForMessage get the fee the network will charge for a particular Message
func (c *RpcClient) GetFeeForMessage(ctx context.Context, message string) (GetFeeForMessageResponse, error) {
	return c.processGetFeeForMessage(c.Call(ctx, "getFeeForMessage", message))
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8
// GetFeeForMessageWithConfig get the fee the network will charge for a particular Message
func (c *RpcClient) GetFeeForMessageWithConfig(ctx context.Context, message string, cfg GetFeeForMessageConfig) (GetFeeForMessageResponse, error) {
	return c.processGetFeeForMessage(c.Call(ctx, "getFeeForMessage", message, cfg))
}

func (c *RpcClient) processGetFeeForMessage(body []byte, rpcErr error) (res GetFeeForMessageResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
