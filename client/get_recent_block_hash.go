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
