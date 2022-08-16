package rpc

import (
	"context"
)

type GetSignatureStatusesResponse JsonRpcResponse[GetSignatureStatuses]

// GetSignatureStatusesResult is a part of raw rpc response of `getSignatureStatuses`
type GetSignatureStatuses struct {
	Context Context           `json:"context"`
	Value   SignatureStatuses `json:"value"`
}

type SignatureStatus struct {
	Slot               uint64      `json:"slot"`
	Confirmations      *uint64     `json:"confirmations"`
	ConfirmationStatus *Commitment `json:"confirmationStatus"`
	Err                any         `json:"err"`
}

type SignatureStatuses []*SignatureStatus

// GetSignatureStatusesConfig is a option config for `getSignatureStatuses`
type GetSignatureStatusesConfig struct {
	SearchTransactionHistory bool `json:"searchTransactionHistory,omitempty"`
}

// GetSignatureStatuses returns the SOL balance
func (c *RpcClient) GetSignatureStatuses(ctx context.Context, signatures []string) (JsonRpcResponse[GetSignatureStatuses], error) {
	return c.processGetSignatureStatuses(c.Call(ctx, "getSignatureStatuses", signatures))
}

// GetSignatureStatusesWithConfig returns the SOL balance
func (c *RpcClient) GetSignatureStatusesWithConfig(ctx context.Context, signatures []string, cfg GetSignatureStatusesConfig) (JsonRpcResponse[GetSignatureStatuses], error) {
	return c.processGetSignatureStatuses(c.Call(ctx, "getSignatureStatuses", signatures, cfg))
}

func (c *RpcClient) processGetSignatureStatuses(body []byte, rpcErr error) (res JsonRpcResponse[GetSignatureStatuses], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
