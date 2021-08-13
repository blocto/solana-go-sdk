package client

import (
	"context"
	"errors"
)

type GetSignaturesForAddress struct {
	Signature string      `json:"signature"`
	Slot      uint64      `json:"slot"`
	BlockTime *int64      `json:"blockTime"`
	Err       interface{} `json:"err"`
	Memo      *string     `json:"memo"`
}

type GetSignaturesForAddressConfig struct {
	Limit      int        `json:"limit,omitempty"` // between 1 and 1000, default: 1000
	Before     string     `json:"before,omitempty"`
	Until      string     `json:"until,omitempty"`
	Commitment Commitment `json:"commitment,omitempty"` // "processed" is not supported, default is "finalized"
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use "getConfirmedSignaturesForAddress2" for solana-core v1.6
// GetSignaturesForAddress Returns confirmed signatures for transactions involving an address backwards
// in time from the provided signature or most recent confirmed block
func (s *RpcClient) GetSignaturesForAddress(ctx context.Context, base58Addr string, config GetConfirmedSignaturesForAddressConfig) ([]GetConfirmedSignaturesForAddress, error) {
	res := struct {
		GeneralResponse
		Result []GetConfirmedSignaturesForAddress `json:"result"`
	}{}
	err := s.request(ctx, "getSignaturesForAddress", []interface{}{base58Addr, config}, &res)
	if err != nil {
		return []GetConfirmedSignaturesForAddress{}, err
	}
	if res.Error != (ErrorResponse{}) {
		return []GetConfirmedSignaturesForAddress{}, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
