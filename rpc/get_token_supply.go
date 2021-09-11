package rpc

import (
	"context"
	"errors"
)

type GetTokenSupply struct {
	Amount         string `json:"amount"`
	Decimals       int64  `json:"decimals"`
	UIAmountString string `json:"uiAmountString"`
}

// GetTokenSupply returns the total supply of an SPL Token type.
func (s *RpcClient) GetTokenSupply(ctx context.Context, mintBase58Addr string, commitment Commitment) (GetTokenSupply, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context        `json:"context"`
			Value   GetTokenSupply `json:"value"`
		} `json:"result"`
	}{}
	err := s.request(ctx, "getTokenSupply", []interface{}{mintBase58Addr, map[string]interface{}{"commitment": commitment}}, &res)
	if err != nil {
		return GetTokenSupply{}, err
	}
	if res.Error != nil {
		return GetTokenSupply{}, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
