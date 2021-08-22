package rpc

import (
	"context"
	"encoding/json"
	"fmt"
)

type GetBalanceResponse struct {
	GeneralResponse
	Result GetBalanceResult `json:"result"`
}

type GetBalanceResult struct {
	Context Context `json:"context"`
	Value   uint64  `json:"value"`
}

type GetBalanceConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBalance returns the balance of the account of provided Pubkey
func (c *RpcClient) GetBalance(ctx context.Context, base58Addr string, cfg GetBalanceConfig) (GetBalanceResponse, error) {
	var body []byte
	var err error
	if cfg == (GetBalanceConfig{}) {
		body, err = c.Call(ctx, "getBalance", base58Addr)
	} else {
		body, err = c.Call(ctx, "getBalance", base58Addr, cfg)
	}
	if err != nil {
		return GetBalanceResponse{}, fmt.Errorf("rpc: call error, err: %v", err)
	}

	var res GetBalanceResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return GetBalanceResponse{}, fmt.Errorf("rpc: failed to json decode body, err: %v", err)
	}
	return res, nil
}
