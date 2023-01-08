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

// IsBlockhashValid get the fee the network will charge for a particular Message
func (c *RpcClient) IsBlockhashValid(ctx context.Context, message string) (JsonRpcResponse[ValueWithContext[bool]], error) {
	return call[JsonRpcResponse[ValueWithContext[bool]]](c, ctx, "isBlockhashValid", message)
}

// IsBlockhashValidWithConfig get the fee the network will charge for a particular Message
func (c *RpcClient) IsBlockhashValidWithConfig(ctx context.Context, message string, cfg IsBlockhashValidConfig) (JsonRpcResponse[ValueWithContext[bool]], error) {
	return call[JsonRpcResponse[ValueWithContext[bool]]](c, ctx, "isBlockhashValid", message, cfg)
}
