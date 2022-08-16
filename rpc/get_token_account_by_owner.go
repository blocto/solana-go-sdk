package rpc

import (
	"context"
)

type GetTokenAccountsByOwnerResponse JsonRpcResponse[GetTokenAccountsByOwner]

type GetTokenAccountsByOwner struct {
	Context Context            `json:"context"`
	Value   GetProgramAccounts `json:"value"`
}

// GetTokenAccountsByOwnerConfig is a option config for `GetTokenAccountsByOwner`
type GetTokenAccountsByOwnerConfig struct {
	Commitment Commitment      `json:"commitment,omitempty"`
	Encoding   AccountEncoding `json:"encoding,omitempty"`
	DataSlice  *DataSlice      `json:"dataSlice,omitempty"`
}

// GetTokenAccountsByOwnerConfigFilter either mint or programId
type GetTokenAccountsByOwnerConfigFilter struct {
	Mint      string `json:"mint,omitempty"`
	ProgramId string `json:"programId,omitempty"`
}

func (c *RpcClient) GetTokenAccountsByOwner(ctx context.Context, base58Addr string, filter GetTokenAccountsByOwnerConfigFilter) (JsonRpcResponse[GetTokenAccountsByOwner], error) {
	return c.processGetTokenAccountsByOwner(c.Call(ctx, "getTokenAccountsByOwner", base58Addr, filter))
}

func (c *RpcClient) GetTokenAccountsByOwnerWithConfig(ctx context.Context, base58Addr string, filter GetTokenAccountsByOwnerConfigFilter, cfg GetTokenAccountsByOwnerConfig) (JsonRpcResponse[GetTokenAccountsByOwner], error) {
	return c.processGetTokenAccountsByOwner(c.Call(ctx, "getTokenAccountsByOwner", base58Addr, filter, cfg))
}

func (c *RpcClient) processGetTokenAccountsByOwner(body []byte, rpcErr error) (res JsonRpcResponse[GetTokenAccountsByOwner], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
