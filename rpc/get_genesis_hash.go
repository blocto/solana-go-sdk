package client

import (
	"context"
	"errors"
)

// GetGenesisHash returns the genesis hash
func (s *RpcClient) GetGenesisHash(ctx context.Context) (string, error) {
	res := struct {
		GeneralResponse
		Result string `json:"result"`
	}{}
	err := s.request(ctx, "getGenesisHash", []interface{}{}, &res)
	if err != nil {
		return "", err
	}
	if res.Error != (ErrorResponse{}) {
		return "", errors.New(res.Error.Message)
	}
	return res.Result, nil
}
