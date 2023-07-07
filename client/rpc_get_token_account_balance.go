package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

type GetTokenAccountBalanceConfig struct {
	Commitment rpc.Commitment
}

func (c GetTokenAccountBalanceConfig) toRpc() rpc.GetTokenAccountBalanceConfig {
	return rpc.GetTokenAccountBalanceConfig{
		Commitment: c.Commitment,
	}
}

func (c *Client) GetTokenAccountBalance(ctx context.Context, addr string) (TokenAmount, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.TokenAccountBalance]], error) {
			return c.RpcClient.GetTokenAccountBalance(ctx, addr)
		},
		convertGetTokenAccountBalance,
	)
}

func (c *Client) GetTokenAccountBalanceWithConfig(ctx context.Context, addr string, cfg GetTokenAccountBalanceConfig) (TokenAmount, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.TokenAccountBalance]], error) {
			return c.RpcClient.GetTokenAccountBalanceWithConfig(ctx, addr, cfg.toRpc())
		},
		convertGetTokenAccountBalance,
	)
}

func (c *Client) GetTokenAccountBalanceAndContext(ctx context.Context, addr string) (rpc.ValueWithContext[TokenAmount], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.TokenAccountBalance]], error) {
			return c.RpcClient.GetTokenAccountBalance(ctx, addr)
		},
		convertGetTokenAccountBalanceAndContext,
	)
}

func (c *Client) GetTokenAccountBalanceAndContextWithConfig(ctx context.Context, addr string, cfg GetTokenAccountBalanceConfig) (rpc.ValueWithContext[TokenAmount], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.TokenAccountBalance]], error) {
			return c.RpcClient.GetTokenAccountBalanceWithConfig(ctx, addr, cfg.toRpc())
		},
		convertGetTokenAccountBalanceAndContext,
	)
}

func convertGetTokenAccountBalance(v rpc.ValueWithContext[rpc.TokenAccountBalance]) (TokenAmount, error) {
	return newTokenAmount(v.Value.Amount, v.Value.Decimals, v.Value.UIAmountString)
}

func convertGetTokenAccountBalanceAndContext(v rpc.ValueWithContext[rpc.TokenAccountBalance]) (rpc.ValueWithContext[TokenAmount], error) {
	tokenAmount, err := convertGetTokenAccountBalance(v)
	if err != nil {
		return rpc.ValueWithContext[TokenAmount]{}, err
	}
	return rpc.ValueWithContext[TokenAmount]{
		Context: v.Context,
		Value:   tokenAmount,
	}, nil
}
