package rpc

import (
	"context"
)

// GetSignaturesForAddressResponse is full `getSignaturesForAddress` raw response
type GetSignaturesForAddressResponse struct {
	GeneralResponse
	Result []GetSignaturesForAddressResult `json:"result"`
}

// GetSignaturesForAddressResult is a part of `getSignaturesForAddress` raw response
type GetSignaturesForAddressResult struct {
	Signature string      `json:"signature"`
	Slot      uint64      `json:"slot"`
	BlockTime *int64      `json:"blockTime"`
	Err       interface{} `json:"err"`
	Memo      *string     `json:"memo"`
}

// GetSignaturesForAddressConfig is option config of `getSignaturesForAddress`
type GetSignaturesForAddressConfig struct {
	Limit      int        `json:"limit,omitempty"` // between 1 and 1000, default: 1000
	Before     string     `json:"before,omitempty"`
	Until      string     `json:"until,omitempty"`
	Commitment Commitment `json:"commitment,omitempty"` // "processed" is not supported, default is "finalized"
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use "getConfirmedSignaturesForAddress2" for solana-core v1.6
// GetSignaturesForAddress returns confirmed signatures for transactions involving an address backwards
// in time from the provided signature or most recent confirmed block
func (c *RpcClient) GetSignaturesForAddress(ctx context.Context, base58Addr string) (GetSignaturesForAddressResponse, error) {
	return c.processGetSignaturesForAddress(c.Call(ctx, "getSignaturesForAddress", base58Addr))
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use "getConfirmedSignaturesForAddress2" for solana-core v1.6
// GetSignaturesForAddressWithConfig returns confirmed signatures for transactions involving an address backwards
// in time from the provided signature or most recent confirmed block
func (c *RpcClient) GetSignaturesForAddressWithConfig(ctx context.Context, base58Addr string, cfg GetSignaturesForAddressConfig) (GetSignaturesForAddressResponse, error) {
	return c.processGetSignaturesForAddress(c.Call(ctx, "getSignaturesForAddress", base58Addr, cfg))
}

func (c *RpcClient) processGetSignaturesForAddress(body []byte, rpcErr error) (res GetSignaturesForAddressResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
