package rpc

import (
	"context"
)

// GetTokenAccountsByOwnerResponse is a full rpc response for `GetTokenAccountsByOwner`
type GetTokenAccountsByOwnerResponse struct {
	GeneralResponse
	Result GetTokenAccountsByOwnerResponseResult `json:"result"`
}

type GetTokenAccountsByOwnerResponseResult struct {
	Context Context              `json:"context"`
	Value   []GetProgramAccounts `json:"value"`
}

type GetTokenAccountsByOwnerConfigEncoding string

const (
	// GetTokenAccountsByOwnerConfigEncodingBase58 limited to Account data of less than 128 bytes
	GetTokenAccountsByOwnerConfigEncodingBase58     GetTokenAccountsByOwnerConfigEncoding = "base58"
	GetTokenAccountsByOwnerConfigEncodingJsonParsed GetTokenAccountsByOwnerConfigEncoding = "jsonParsed"
	GetTokenAccountsByOwnerConfigEncodingBase64     GetTokenAccountsByOwnerConfigEncoding = "base64"
	GetTokenAccountsByOwnerConfigEncodingBase64Zstd GetTokenAccountsByOwnerConfigEncoding = "base64+zstd"
)

// GetTokenAccountsByOwnerConfig is a option config for `GetTokenAccountsByOwner`
type GetTokenAccountsByOwnerConfig struct {
	Commitment Commitment                              `json:"commitment,omitempty"`
	Encoding   GetTokenAccountsByOwnerConfigEncoding   `json:"encoding,omitempty"`
	DataSlice  *GetTokenAccountsByOwnerConfigDataSlice `json:"dataSlice,omitempty"`
}

// GetTokenAccountsByOwnerConfigDataSlice is a part of GetTokenAccountsByOwnerConfig
type GetTokenAccountsByOwnerConfigDataSlice struct {
	Offset uint64 `json:"offset"`
	Length uint64 `json:"length"`
}

// GetTokenAccountsByOwnerConfigFilter either mint or programId
type GetTokenAccountsByOwnerConfigFilter struct {
	Mint      string `json:"mint,omitempty"`
	ProgramId string `json:"programId,omitempty"`
}

func (c *RpcClient) GetTokenAccountsByOwner(ctx context.Context, base58Addr string, filter GetTokenAccountsByOwnerConfigFilter) (GetTokenAccountsByOwnerResponse, error) {
	return c.processGetTokenAccountsByOwner(c.Call(ctx, "getTokenAccountsByOwner", base58Addr, filter))
}

func (c *RpcClient) GetTokenAccountsByOwnerWithConfig(ctx context.Context, base58Addr string, filter GetTokenAccountsByOwnerConfigFilter, cfg GetTokenAccountsByOwnerConfig) (GetTokenAccountsByOwnerResponse, error) {
	return c.processGetTokenAccountsByOwner(c.Call(ctx, "getTokenAccountsByOwner", base58Addr, filter, cfg))
}

func (c *RpcClient) processGetTokenAccountsByOwner(body []byte, rpcErr error) (res GetTokenAccountsByOwnerResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
