package client

import (
	"context"
	"errors"
)

// GetBlockTime returns the estimated production time of a block.
func (s *Client) GetBlockTime(ctx context.Context, slot uint64) (int64, error) {
	res := struct {
		GeneralResponse
		Result int64 `json:"result"`
	}{}
	err := s.request(ctx, "getBlockTime", []interface{}{slot}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
