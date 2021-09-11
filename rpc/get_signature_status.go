package rpc

import (
	"context"
	"errors"
)

type GetSignatureStatusesResponse struct {
	Slot               uint64      `json:"slot"`
	Confirmations      *uint64     `json:"confirmations"`
	ConfirmationStatus *Commitment `json:"confirmationStatus"`
	Err                interface{} `json:"err"`
}

func (s *RpcClient) GetSignatureStatuses(ctx context.Context, signatures []string) ([]GetSignatureStatusesResponse, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context                        `json:"context"`
			Value   []GetSignatureStatusesResponse `json:"value"`
		} `json:"result"`
	}{}
	err := s.request(ctx, "getSignatureStatuses", []interface{}{signatures, map[string]interface{}{"searchTransactionHistory": true}}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
