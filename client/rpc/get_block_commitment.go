package rpc

import (
	"context"
	"errors"
)

type GetBlockCommitmentResponse struct {
	Commitment []uint64 `json:"commitment"`
	TotalStake uint64   `json:"totalStake"`
}

// GetBlockCommitment returns commitment for particular block
func (s *RpcClient) GetBlockCommitment(ctx context.Context, slot uint64) (GetBlockCommitmentResponse, error) {
	res := struct {
		GeneralResponse
		Result GetBlockCommitmentResponse `json:"result"`
	}{}
	err := s.request(ctx, "getBlockCommitment", []interface{}{slot}, &res)
	if err != nil {
		return GetBlockCommitmentResponse{}, err
	}
	if res.Error != nil {
		return GetBlockCommitmentResponse{}, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
