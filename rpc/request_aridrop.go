package client

import (
	"context"
	"errors"
)

// RequestAirdrop Requests an airdrop of lamports to a Pubkey, return string is Transaction Signature of airdrop, as base-58 encoded
func (s *RpcClient) RequestAirdrop(ctx context.Context, base58Addr string, lamport uint64) (string, error) {
	res := struct {
		GeneralResponse
		Result string `json:"result"`
	}{}
	err := s.request(ctx, "requestAirdrop", []interface{}{base58Addr, lamport}, &res)
	if err != nil {
		return "", err
	}
	if res.Error != (ErrorResponse{}) {
		return "", errors.New(res.Error.Message)
	}
	return res.Result, nil
}
