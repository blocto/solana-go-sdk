package rpc

import (
	"context"
)

// GetAccountInfoConfigEncoding is account's data encode format
type GetAccountInfoConfigEncoding string

const (
	// GetAccountInfoConfigEncodingBase58 limited to Account data of less than 128 bytes
	GetAccountInfoConfigEncodingBase58     GetAccountInfoConfigEncoding = "base58"
	GetAccountInfoConfigEncodingJsonParsed GetAccountInfoConfigEncoding = "jsonParsed"
	GetAccountInfoConfigEncodingBase64     GetAccountInfoConfigEncoding = "base64"
	GetAccountInfoConfigEncodingBase64Zstd GetAccountInfoConfigEncoding = "base64+zstd"
)

// GetAccountInfoConfig is an option config for `getAccountInfo`
type GetAccountInfoConfig struct {
	Commitment Commitment                     `json:"commitment,omitempty"`
	Encoding   GetAccountInfoConfigEncoding   `json:"encoding,omitempty"`
	DataSlice  *GetAccountInfoConfigDataSlice `json:"dataSlice,omitempty"`
}

// GetAccountInfoResponse is a full raw rpc response of `getAccountInfo`
type GetAccountInfoResponse struct {
	GeneralResponse
	Result GetAccountInfoResult `json:"result"`
}

// GetAccountInfoConfigDataSlice is a part of GetAccountInfoConfig
type GetAccountInfoConfigDataSlice struct {
	Offset uint64 `json:"offset,omitempty"`
	Length uint64 `json:"length,omitempty"`
}

// GetAccountInfoResult is rpc result of `getAccountInfo`
type GetAccountInfoResult struct {
	Context Context                   `json:"context"`
	Value   GetAccountInfoResultValue `json:"value"`
}

// GetAccountInfoResultValue is rpc result of `getAccountInfo`
type GetAccountInfoResultValue struct {
	Lamports  uint64      `json:"lamports"`
	Owner     string      `json:"owner"`
	Excutable bool        `json:"excutable"`
	RentEpoch uint64      `json:"rentEpoch"`
	Data      interface{} `json:"data"`
}

// GetAccountInfo returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetAccountInfo(ctx context.Context, base58Addr string) (GetAccountInfoResponse, error) {
	return c.processGetAccountInfo(c.Call(ctx, "getAccountInfo", base58Addr))
}

// GetAccountInfo returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetAccountInfoWithConfig(ctx context.Context, base58Addr string, cfg GetAccountInfoConfig) (GetAccountInfoResponse, error) {
	return c.processGetAccountInfo(c.Call(ctx, "getAccountInfo", base58Addr, cfg))
}

func (c *RpcClient) processGetAccountInfo(body []byte, rpcErr error) (res GetAccountInfoResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
