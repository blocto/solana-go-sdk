package rpc

import (
	"context"
)

type GetAccountResponse JsonRpcResponse[GetAccountInfo]

type GetAccountInfo ValueWithContext[AccountInfo]

// GetAccountInfoConfig is an option config for `getAccountInfo`
type GetAccountInfoConfig struct {
	Commitment Commitment      `json:"commitment,omitempty"`
	Encoding   AccountEncoding `json:"encoding,omitempty"`
	DataSlice  *DataSlice      `json:"dataSlice,omitempty"`
}

// GetAccountInfo returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetAccountInfo(ctx context.Context, base58Addr string) (JsonRpcResponse[ValueWithContext[AccountInfo]], error) {
	return call[JsonRpcResponse[ValueWithContext[AccountInfo]]](c, ctx, "getAccountInfo", base58Addr)
}

// GetAccountInfo returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetAccountInfoWithConfig(ctx context.Context, base58Addr string, cfg GetAccountInfoConfig) (JsonRpcResponse[ValueWithContext[AccountInfo]], error) {
	return call[JsonRpcResponse[ValueWithContext[AccountInfo]]](c, ctx, "getAccountInfo", base58Addr, cfg)
}
