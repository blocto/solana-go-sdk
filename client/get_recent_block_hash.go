package client

import (
	"context"
	"errors"
)

type GetRecentBlockHashResponse struct {
	Blockhash     string `json:"blockhash"`
	FeeCalculator struct {
		LamportsPerSignature uint64 `json:"lamportsPerSignature"`
	} `json:"feeCalculator"`
}

// getRecentBlockhash returns a recent block hash from the ledger, and a fee schedule that can be used to compute
// the cost of submitting a transaction using it.
func (s *Client) GetRecentBlockhash(ctx context.Context) (GetRecentBlockHashResponse, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context                    `json:"context"`
			Value   GetRecentBlockHashResponse `json:"value"`
		} `json:"result"`
	}{}
	err := s.request(ctx, "getRecentBlockhash", []interface{}{}, &res)
	if err != nil {
		return GetRecentBlockHashResponse{}, err
	}
	if res.Error != (ErrorResponse{}) {
		return GetRecentBlockHashResponse{}, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
