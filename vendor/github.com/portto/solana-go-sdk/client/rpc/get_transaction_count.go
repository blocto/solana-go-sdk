package rpc

import (
	"context"
	"errors"
)

// GetTransactionCount returns the current transaction count from the ledger
func (s *RpcClient) GetTransactionCount(ctx context.Context) (uint64, error) {
	res := struct {
		GeneralResponse
		Result uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getTransactionCount", []interface{}{}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != nil {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
