package rpc

import (
	"context"
	"encoding/json"
	"fmt"
)

// GetBalanceResponse is a full raw rpc response of `getBalance`
type GetBalanceResponse struct {
	GeneralResponse
	Result GetBalanceResult `json:"result"`
}

// GetBalanceResult is a part of raw rpc response of `getBalance`
type GetBalanceResult struct {
	Context Context `json:"context"`
	Value   uint64  `json:"value"`
}

// GetBalanceConfig is a option config for `getBalance`
type GetBalanceConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetBalance returns the SOL balance
func (c *RpcClient) GetBalance(ctx context.Context, base58Addr string) (GetBalanceResponse, error) {
	body, err := c.Call(ctx, "getBalance", base58Addr)
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

// GetBalanceWithCfg returns the SOL balance
func (c *RpcClient) GetBalanceWithCfg(ctx context.Context, base58Addr string, cfg GetBalanceConfig) (GetBalanceResponse, error) {
	body, err := c.Call(ctx, "getBalance", base58Addr, cfg)
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
