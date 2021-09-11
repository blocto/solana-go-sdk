package rpc

import (
	"context"
)

// GetProgramAccountsResponse is a full rpc response for `getProgramAccounts`
type GetProgramAccountsResponse struct {
	GeneralResponse
	Result []GetProgramAccounts `json:"result"`
}

// GetProgramAccountsResponse is a full rpc response for `getProgramAccounts` with context
type GetProgramAccountsWithContextResponse struct {
	GeneralResponse
	Result GetProgramAccountsWithContextResult `json:"result"`
}

type GetProgramAccountsConfigEncoding string

const (
	// GetProgramAccountsConfigEncodingBase58 limited to Account data of less than 128 bytes
	GetProgramAccountsConfigEncodingBase58     GetProgramAccountsConfigEncoding = "base58"
	GetProgramAccountsConfigEncodingJsonParsed GetProgramAccountsConfigEncoding = "jsonParsed"
	GetProgramAccountsConfigEncodingBase64     GetProgramAccountsConfigEncoding = "base64"
	GetProgramAccountsConfigEncodingBase64Zstd GetProgramAccountsConfigEncoding = "base64+zstd"
)

// GetProgramAccountsConfig is a option config for `getProgramAccounts`
type GetProgramAccountsConfig struct {
	Encoding   GetProgramAccountsConfigEncoding   `json:"encoding,omitempty"`
	Commitment Commitment                         `json:"commitment,omitempty"`
	DataSlice  *GetProgramAccountsConfigDataSlice `json:"dataSlice,omitempty"`
	Filters    []GetProgramAccountsConfigFilter   `json:"filters,omitempty"`
}

type getProgramAccountsConfig struct {
	GetProgramAccountsConfig
	WithContext bool `json:"withContext,omitempty"`
}

// GetProgramAccountsConfigDataSlice is a part of GetProgramAccountsConfig
type GetProgramAccountsConfigDataSlice struct {
	Offset uint64 `json:"offset"`
	Length uint64 `json:"length"`
}

// GetProgramAccountsConfigFilter you can set either MemCmp or DataSize but can be both, if needed, separate them into two
type GetProgramAccountsConfigFilter struct {
	MemCmp   *GetProgramAccountsConfigFilterMemCmp `json:"memcmp,omitempty"`
	DataSize uint64                                `json:"dataSize,omitempty"`
}

type GetProgramAccountsConfigFilterMemCmp struct {
	Offset uint64 `json:"offset"`
	Bytes  string `json:"bytes"`
}

type GetProgramAccounts struct {
	Pubkey  string                    `json:"pubkey"`
	Account GetProgramAccountsAccount `json:"account"`
}

type GetProgramAccountsAccount struct {
	Lamports   uint64      `json:"lamports"`
	Owner      string      `json:"owner"`
	RentEpoch  uint64      `json:"rentEpoch"`
	Data       interface{} `json:"data"`
	Executable bool        `json:"executable"`
}

type GetProgramAccountsWithContextResult struct {
	Context Context              `json:"context"`
	Value   []GetProgramAccounts `json:"value"`
}

func (c *RpcClient) GetProgramAccounts(ctx context.Context, programId string) (GetProgramAccountsResponse, error) {
	return c.processGetProgramAccounts(c.Call(ctx, "getProgramAccounts", programId))
}

func (c *RpcClient) GetProgramAccountsWithConfig(ctx context.Context, programId string, cfg GetProgramAccountsConfig) (GetProgramAccountsResponse, error) {
	return c.processGetProgramAccounts(c.Call(ctx, "getProgramAccounts", programId, c.toInternalGetProgramAccountsConfig(cfg, false)))
}

func (c *RpcClient) processGetProgramAccounts(body []byte, rpcErr error) (res GetProgramAccountsResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}

func (c *RpcClient) GetProgramAccountsWithContext(ctx context.Context, programId string) (GetProgramAccountsWithContextResponse, error) {
	return c.processGetProgramAccountsWithContext(c.Call(ctx, "getProgramAccounts", programId, c.toInternalGetProgramAccountsConfig(GetProgramAccountsConfig{}, true)))
}

func (c *RpcClient) GetProgramAccountsWithContextAndConfig(ctx context.Context, programId string, cfg GetProgramAccountsConfig) (GetProgramAccountsWithContextResponse, error) {
	return c.processGetProgramAccountsWithContext(c.Call(ctx, "getProgramAccounts", programId, c.toInternalGetProgramAccountsConfig(cfg, true)))
}

func (c *RpcClient) processGetProgramAccountsWithContext(body []byte, rpcErr error) (res GetProgramAccountsWithContextResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}

func (c *RpcClient) toInternalGetProgramAccountsConfig(cfg GetProgramAccountsConfig, withContext bool) getProgramAccountsConfig {
	return getProgramAccountsConfig{
		GetProgramAccountsConfig: cfg,
		WithContext:              withContext,
	}
}
