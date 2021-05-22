package client

import (
	"context"
	"errors"
)

type GetClusterNodesResponse struct {
	FeatureSet   *uint64 `json:"featureSet"`
	Gossip       string  `json:"gossip"`
	Pubkey       string  `json:"pubkey"`
	RPC          string  `json:"rpc"`
	ShredVersion int     `json:"shredVersion"`
	Tpu          string  `json:"tpu"`
	Version      *string `json:"version"`
}

func (s *Client) GetClusterNodes(ctx context.Context) ([]GetClusterNodesResponse, error) {
	res := struct {
		GeneralResponse
		Result []GetClusterNodesResponse `json:"result"`
	}{}
	err := s.request(ctx, "getClusterNodes", []interface{}{}, &res)
	if err != nil {
		return []GetClusterNodesResponse{}, err
	}
	if res.Error != (ErrorResponse{}) {
		return []GetClusterNodesResponse{}, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
