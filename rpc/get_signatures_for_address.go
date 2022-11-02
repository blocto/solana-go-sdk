package rpc

import (
	"context"
)

type GetSignaturesForAddressResponse JsonRpcResponse[GetSignaturesForAddress]

type GetSignaturesForAddress []SignatureWithStatus

type SignatureWithStatus struct {
	Signature string  `json:"signature"`
	Slot      uint64  `json:"slot"`
	BlockTime *int64  `json:"blockTime"`
	Err       any     `json:"err"`
	Memo      *string `json:"memo"`
}

// GetSignaturesForAddressConfig is option config of `getSignaturesForAddress`
type GetSignaturesForAddressConfig struct {
	Limit      int        `json:"limit,omitempty"` // between 1 and 1000, default: 1000
	Before     string     `json:"before,omitempty"`
	Until      string     `json:"until,omitempty"`
	Commitment Commitment `json:"commitment,omitempty"` // "processed" is not supported, default is "finalized"
}

// GetSignaturesForAddress returns confirmed signatures for transactions involving an address backwards
// in time from the provided signature or most recent confirmed block
func (c *RpcClient) GetSignaturesForAddress(ctx context.Context, base58Addr string) (JsonRpcResponse[GetSignaturesForAddress], error) {
	return call[JsonRpcResponse[GetSignaturesForAddress]](c, ctx, "getSignaturesForAddress", base58Addr)
}

// GetSignaturesForAddressWithConfig returns confirmed signatures for transactions involving an address backwards
// in time from the provided signature or most recent confirmed block
func (c *RpcClient) GetSignaturesForAddressWithConfig(ctx context.Context, base58Addr string, cfg GetSignaturesForAddressConfig) (JsonRpcResponse[GetSignaturesForAddress], error) {
	return call[JsonRpcResponse[GetSignaturesForAddress]](c, ctx, "getSignaturesForAddress", base58Addr, cfg)
}
