package rpc

import (
	"context"
)

type GetSignatureStatusesResponse JsonRpcResponse[GetSignatureStatuses]

type GetSignatureStatuses ValueWithContext[SignatureStatuses]

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
func (c *RpcClient) GetSignatureStatuses(ctx context.Context, signatures []string) (JsonRpcResponse[ValueWithContext[SignatureStatuses]], error) {
	return call[JsonRpcResponse[ValueWithContext[SignatureStatuses]]](c, ctx, "getSignatureStatuses", signatures)
}

// GetSignatureStatusesWithConfig returns the SOL balance
func (c *RpcClient) GetSignatureStatusesWithConfig(ctx context.Context, signatures []string, cfg GetSignatureStatusesConfig) (JsonRpcResponse[ValueWithContext[SignatureStatuses]], error) {
	return call[JsonRpcResponse[ValueWithContext[SignatureStatuses]]](c, ctx, "getSignatureStatuses", signatures, cfg)
}
