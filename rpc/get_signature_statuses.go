package rpc

import (
	"context"
)

// GetSignatureStatusesResponse is a full raw rpc response of `getSignatureStatuses`
type GetSignatureStatusesResponse struct {
	GeneralResponse
	Result GetSignatureStatusesResult `json:"result"`
}

// GetSignatureStatusesResult is a part of raw rpc response of `getSignatureStatuses`
type GetSignatureStatusesResult struct {
	Context Context                            `json:"context"`
	Value   []*GetSignatureStatusesResultValue `json:"value"`
}

type GetSignatureStatusesResultValue struct {
	Slot               uint64      `json:"slot"`
	Confirmations      *uint64     `json:"confirmations"`
	ConfirmationStatus *Commitment `json:"confirmationStatus"`
	Err                interface{} `json:"err"`
}

// GetSignatureStatusesConfig is a option config for `getSignatureStatuses`
type GetSignatureStatusesConfig struct {
	SearchTransactionHistory bool `json:"searchTransactionHistory,omitempty"`
}

// GetSignatureStatuses returns the SOL balance
func (c *RpcClient) GetSignatureStatuses(ctx context.Context, signatures []string) (GetSignatureStatusesResponse, error) {
	return c.processGetSignatureStatuses(c.Call(ctx, "getSignatureStatuses", signatures))
}

// GetSignatureStatusesWithConfig returns the SOL balance
func (c *RpcClient) GetSignatureStatusesWithConfig(ctx context.Context, signatures []string, cfg GetSignatureStatusesConfig) (GetSignatureStatusesResponse, error) {
	return c.processGetSignatureStatuses(c.Call(ctx, "getSignatureStatuses", signatures, cfg))
}

func (c *RpcClient) processGetSignatureStatuses(body []byte, rpcErr error) (res GetSignatureStatusesResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
