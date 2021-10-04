package rpc

import (
	"context"
	"errors"
)

// GetMinimumBalanceForRentExemption returns minimum balance required to make account rent exempt.
func (s *RpcClient) GetMinimumBalanceForRentExemption(ctx context.Context, accountDataLen uint64) (uint64, error) {
	res := struct {
		GeneralResponse
		Result uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getMinimumBalanceForRentExemption", []interface{}{accountDataLen}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != nil {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
