package rpc

import (
	"context"
)

type IsBlockhashValidResponse JsonRpcResponse[IsBlockhashValid]

type IsBlockhashValid ValueWithContext[bool]

// IsBlockhashValidConfig is a option config for `IsBlockhashValid`
type IsBlockhashValidConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8
// IsBlockhashValid get the fee the network will charge for a particular Message
func (c *RpcClient) IsBlockhashValid(ctx context.Context, message string) (JsonRpcResponse[IsBlockhashValid], error) {
	return call[JsonRpcResponse[IsBlockhashValid]](c, ctx, "isBlockhashValid", message)
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8
// IsBlockhashValidWithConfig get the fee the network will charge for a particular Message
func (c *RpcClient) IsBlockhashValidWithConfig(ctx context.Context, message string, cfg IsBlockhashValidConfig) (JsonRpcResponse[IsBlockhashValid], error) {
	return call[JsonRpcResponse[IsBlockhashValid]](c, ctx, "isBlockhashValid", message, cfg)
}
