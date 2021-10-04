package rpc

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

// GetClusterNodes returns information about all the nodes participating in the cluster
func (s *RpcClient) GetClusterNodes(ctx context.Context) ([]GetClusterNodesResponse, error) {
	res := struct {
		GeneralResponse
		Result []GetClusterNodesResponse `json:"result"`
	}{}
	err := s.request(ctx, "getClusterNodes", []interface{}{}, &res)
	if err != nil {
		return []GetClusterNodesResponse{}, err
	}
	if res.Error != nil {
		return []GetClusterNodesResponse{}, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
