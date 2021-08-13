package rpc

import (
	"context"
	"errors"
)

type GetConfirmedSignaturesForAddress struct {
	BlockTime *int64      `json:"blockTime"`
	Err       interface{} `json:"err"`
	Memo      *string     `json:"memo"`
	Signature string      `json:"signature"`
	Slot      int64       `json:"slot"`
}

type GetConfirmedSignaturesForAddressConfig struct {
	Limit      int        `json:"limit,omitempty"`
	Before     string     `json:"before,omitempty"`
	Until      string     `json:"until,omitempty"`
	Commitment Commitment `json:"commitment,omitempty"`
}

// DEPRECATED: Please use getSignaturesForAddress instead This method is expected to be removed in solana-core v1.8
// GetConfirmedSignaturesForAddress returns confirmed signatures for transactions involving an address
// backwards in time from the provided signature or most recent confirmed block
func (s *RpcClient) GetConfirmedSignaturesForAddress(ctx context.Context, base58Addr string, config GetConfirmedSignaturesForAddressConfig) ([]GetConfirmedSignaturesForAddress, error) {
	res := struct {
		GeneralResponse
		Result []GetConfirmedSignaturesForAddress `json:"result"`
	}{}
	err := s.request(ctx, "getConfirmedSignaturesForAddress2", []interface{}{base58Addr, config}, &res)
	if err != nil {
		return []GetConfirmedSignaturesForAddress{}, err
	}
	if res.Error != nil {
		return []GetConfirmedSignaturesForAddress{}, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
