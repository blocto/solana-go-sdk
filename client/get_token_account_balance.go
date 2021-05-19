package client

import (
	"context"
	"errors"
)

type GetTokenAccountBalanceResponse struct {
	Amount   string `json:"amount"`
	Decimals int    `json:"decimals"`
}

func (s *Client) GetTokenAccountBalance(
	ctx context.Context, base58Addr string) (GetTokenAccountBalanceResponse, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context                        `json:"context"`
			Value   GetTokenAccountBalanceResponse `json:"value"`
		} `json:"result"`
	}{}
	err := s.request(ctx, "getTokenAccountBalance", []interface{}{base58Addr}, &res)
	if err != nil {
		return GetTokenAccountBalanceResponse{}, err
	}
	if res.Error != (ErrorResponse{}) {
		return GetTokenAccountBalanceResponse{}, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
