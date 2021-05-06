package client

import (
	"context"
	"errors"
)

func (s *Client) GetBalance(ctx context.Context, base58Addr string) (uint64, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context `json:"context"`
			Value   uint64  `json:"value"`
		} `json:"result"`
	}{}
	err := s.request(ctx, "getBalance", []interface{}{base58Addr}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
