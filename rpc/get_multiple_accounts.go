package rpc

import (
	"context"
)

type GetMultipleAccountsResponse JsonRpcResponse[GetMultipleAccounts]

type GetMultipleAccounts ValueWithContext[[]AccountInfo]

// GetMultipleAccountsConfigEncoding is account's data encode format
type GetMultipleAccountsConfigEncoding string

// GetMultipleAccountsConfig is an option config for `getAccountInfo`
type GetMultipleAccountsConfig struct {
	Commitment Commitment      `json:"commitment,omitempty"`
	Encoding   AccountEncoding `json:"encoding,omitempty"`
	DataSlice  *DataSlice      `json:"dataSlice,omitempty"`
}

// GetMultipleAccounts returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetMultipleAccounts(ctx context.Context, base58Addrs []string) (JsonRpcResponse[ValueWithContext[[]AccountInfo]], error) {
	return call[JsonRpcResponse[ValueWithContext[[]AccountInfo]]](c, ctx, "getMultipleAccounts", base58Addrs)
}

// GetMultipleAccounts returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetMultipleAccountsWithConfig(ctx context.Context, base58Addrs []string, cfg GetMultipleAccountsConfig) (JsonRpcResponse[ValueWithContext[[]AccountInfo]], error) {
	return call[JsonRpcResponse[ValueWithContext[[]AccountInfo]]](c, ctx, "getMultipleAccounts", base58Addrs, cfg)
}
