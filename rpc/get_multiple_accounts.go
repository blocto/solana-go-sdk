package rpc

import (
	"context"
)

// GetMultipleAccountsConfigEncoding is account's data encode format
type GetMultipleAccountsConfigEncoding string

const (
	// GetMultipleAccountConfigEncodingBase58 limited to Account data of less than 128 bytes
	GetMultipleAccountsConfigEncodingBase58     GetMultipleAccountsConfigEncoding = "base58"
	GetMultipleAccountsConfigEncodingJsonParsed GetMultipleAccountsConfigEncoding = "jsonParsed"
	GetMultipleAccountsConfigEncodingBase64     GetMultipleAccountsConfigEncoding = "base64"
	GetMultipleAccountsConfigEncodingBase64Zstd GetMultipleAccountsConfigEncoding = "base64+zstd"
)

// GetMultipleAccountsConfig is an option config for `getAccountInfo`
type GetMultipleAccountsConfig struct {
	Commitment Commitment                          `json:"commitment,omitempty"`
	Encoding   GetMultipleAccountsConfigEncoding   `json:"encoding,omitempty"`
	DataSlice  *GetMultipleAccountsConfigDataSlice `json:"dataSlice,omitempty"`
}

// GetMultipleAccountsResponse is a full raw rpc response of `getAccountInfo`
type GetMultipleAccountsResponse struct {
	GeneralResponse
	Result GetMultipleAccountsResult `json:"result"`
}

// GetMultipleAccountsConfigDataSlice is a part of GetAccountInfoConfig
type GetMultipleAccountsConfigDataSlice struct {
	Offset uint64 `json:"offset,omitempty"`
	Length uint64 `json:"length,omitempty"`
}

// GetMultipleAccountsResult is rpc result of `getAccountInfo`
type GetMultipleAccountsResult struct {
	Context Context                          `json:"context"`
	Value   []GetMultipleAccountsResultValue `json:"value"`
}

// GetMultipleAccountsResultValue is rpc result of `getAccountInfo`
type GetMultipleAccountsResultValue struct {
	Lamports   uint64      `json:"lamports"`
	Owner      string      `json:"owner"`
	Executable bool        `json:"executable"`
	RentEpoch  uint64      `json:"rentEpoch"`
	Data       interface{} `json:"data"`
}

// GetMultipleAccounts returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetMultipleAccounts(ctx context.Context, base58Addrs []string) (GetMultipleAccountsResponse, error) {
	return c.processGetMultipleAccounts(c.Call(ctx, "getMultipleAccounts", base58Addrs))
}

// GetMultipleAccounts returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetMultipleAccountsWithConfig(ctx context.Context, base58Addrs []string, cfg GetMultipleAccountsConfig) (GetMultipleAccountsResponse, error) {
	return c.processGetMultipleAccounts(c.Call(ctx, "getMultipleAccounts", base58Addrs, cfg))
}

func (c *RpcClient) processGetMultipleAccounts(body []byte, rpcErr error) (res GetMultipleAccountsResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
