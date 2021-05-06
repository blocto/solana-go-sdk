package client

import (
	"context"
	"errors"
)

func (s *Client) GetMinimumBalanceForRentExemption(ctx context.Context, accountDataLen uint64) (uint64, error) {
	res := struct {
		GeneralResponse
		Result uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getMinimumBalanceForRentExemption", []interface{}{accountDataLen}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
