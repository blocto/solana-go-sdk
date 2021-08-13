package client

import (
	"context"
	"errors"
)

type GetBlockHeightConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBlockHeight returns the current block height of the node
func (s *RpcClient) GetBlockHeight(ctx context.Context, cfg GetBlockHeightConfig) (uint64, error) {
	res := struct {
		GeneralResponse
		Result uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getBlockHeight", []interface{}{cfg}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
