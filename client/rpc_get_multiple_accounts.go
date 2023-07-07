package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

type GetMultipleAccountsConfig struct {
	Commitment rpc.Commitment
	DataSlice  *rpc.DataSlice
}

func (c GetMultipleAccountsConfig) toRpc() rpc.GetMultipleAccountsConfig {
	return rpc.GetMultipleAccountsConfig{
		Encoding:   rpc.AccountEncodingBase64,
		Commitment: c.Commitment,
		DataSlice:  c.DataSlice,
	}
}

// GetMultipleAccounts returns multiple accounts info
func (c *Client) GetMultipleAccounts(ctx context.Context, addrs []string) ([]AccountInfo, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[[]rpc.AccountInfo]], error) {
			return c.RpcClient.GetMultipleAccountsWithConfig(
				ctx,
				addrs,
				GetMultipleAccountsConfig{}.toRpc(),
			)
		},
		convertGetMultipleAccounts,
	)
}

// GetMultipleAccountsWithConfig return account's info
func (c *Client) GetMultipleAccountsWithConfig(ctx context.Context, addrs []string, cfg GetMultipleAccountsConfig) ([]AccountInfo, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[[]rpc.AccountInfo]], error) {
			return c.RpcClient.GetMultipleAccountsWithConfig(
				ctx,
				addrs,
				cfg.toRpc(),
			)
		},
		convertGetMultipleAccounts,
	)
}

// GetMultipleAccounts returns multiple accounts info
func (c *Client) GetMultipleAccountsAndContext(ctx context.Context, addrs []string) (rpc.ValueWithContext[[]AccountInfo], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[[]rpc.AccountInfo]], error) {
			return c.RpcClient.GetMultipleAccountsWithConfig(
				ctx,
				addrs,
				GetMultipleAccountsConfig{}.toRpc(),
			)
		},
		convertGetMultipleAccountsAndContext,
	)
}

// GetMultipleAccountsWithConfig return account's info
func (c *Client) GetMultipleAccountsAndContextWithConfig(ctx context.Context, addrs []string, cfg GetMultipleAccountsConfig) (rpc.ValueWithContext[[]AccountInfo], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[[]rpc.AccountInfo]], error) {
			return c.RpcClient.GetMultipleAccountsWithConfig(
				ctx,
				addrs,
				cfg.toRpc(),
			)
		},
		convertGetMultipleAccountsAndContext,
	)
}

func convertGetMultipleAccounts(v rpc.ValueWithContext[[]rpc.AccountInfo]) ([]AccountInfo, error) {
	output := make([]AccountInfo, 0, len(v.Value))
	for _, rac := range v.Value {
		ac, err := convertAccountInfo(rac)
		if err != nil {
			return nil, err
		}
		output = append(output, ac)
	}
	return output, nil
}

func convertGetMultipleAccountsAndContext(v rpc.ValueWithContext[[]rpc.AccountInfo]) (rpc.ValueWithContext[[]AccountInfo], error) {
	accountInfos, err := convertGetMultipleAccounts(v)
	if err != nil {
		return rpc.ValueWithContext[[]AccountInfo]{}, err
	}
	return rpc.ValueWithContext[[]AccountInfo]{
		Context: v.Context,
		Value:   accountInfos,
	}, nil
}
