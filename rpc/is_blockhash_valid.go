package rpc

import (
	"context"
)

// IsBlockhashValidResponse is a full raw rpc response of `isBlockhashValid`
type IsBlockhashValidResponse struct {
	GeneralResponse
	Result IsBlockhashValidResult `json:"result"`
}

// IsBlockhashValidResult is a part of raw rpc response of `isBlockhashValid`
type IsBlockhashValidResult struct {
	Context Context `json:"context"`
	Value   bool    `json:"value"`
}

// IsBlockhashValidConfig is a option config for `IsBlockhashValid`
type IsBlockhashValidConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8
// IsBlockhashValid get the fee the network will charge for a particular Message
func (c *RpcClient) IsBlockhashValid(ctx context.Context, message string) (IsBlockhashValidResponse, error) {
	return c.processIsBlockhashValid(c.Call(ctx, "isBlockhashValid", message))
}

// NEW: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8
// IsBlockhashValidWithConfig get the fee the network will charge for a particular Message
func (c *RpcClient) IsBlockhashValidWithConfig(ctx context.Context, message string, cfg IsBlockhashValidConfig) (IsBlockhashValidResponse, error) {
	return c.processIsBlockhashValid(c.Call(ctx, "isBlockhashValid", message, cfg))
}

func (c *RpcClient) processIsBlockhashValid(body []byte, rpcErr error) (res IsBlockhashValidResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
