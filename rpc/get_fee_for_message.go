package rpc

import (
	"context"
)

type GetFeeForMessageResponse JsonRpcResponse[ValueWithContext[*uint64]]

type GetFeeForMessage ValueWithContext[*uint64]

// GetFeeForMessageConfig is a option config for `GetFeeForMessage`
type GetFeeForMessageConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8
// GetFeeForMessage get the fee the network will charge for a particular Message
func (c *RpcClient) GetFeeForMessage(ctx context.Context, message string) (JsonRpcResponse[ValueWithContext[*uint64]], error) {
	return call[JsonRpcResponse[ValueWithContext[*uint64]]](c, ctx, "getFeeForMessage", message)
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8
// GetFeeForMessageWithConfig get the fee the network will charge for a particular Message
func (c *RpcClient) GetFeeForMessageWithConfig(ctx context.Context, message string, cfg GetFeeForMessageConfig) (JsonRpcResponse[ValueWithContext[*uint64]], error) {
	return call[JsonRpcResponse[ValueWithContext[*uint64]]](c, ctx, "getFeeForMessage", message, cfg)
}
