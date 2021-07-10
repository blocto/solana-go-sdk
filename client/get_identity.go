package client

import (
	"context"
	"errors"
)

// GetIdentity returns the identity pubkey for the current node
func (s *Client) GetIdentity(ctx context.Context) (string, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Identity string `json:"identity"`
		} `json:"result"`
	}{}
	err := s.request(ctx, "getIdentity", []interface{}{}, &res)
	if err != nil {
		return "", err
	}
	if res.Error != (ErrorResponse{}) {
		return "", errors.New(res.Error.Message)
	}
	return res.Result.Identity, nil
}
