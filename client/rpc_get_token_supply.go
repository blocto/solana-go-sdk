package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

type GetTokenSupplyConfig struct {
	Commitment rpc.Commitment
}

func (c GetTokenSupplyConfig) toRpc() rpc.GetTokenSupplyConfig {
	return rpc.GetTokenSupplyConfig{
		Commitment: c.Commitment,
	}
}

func (c *Client) GetTokenSupply(ctx context.Context, mintAddr string) (TokenAmount, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetTokenSupplyResultValue]], error) {
			return c.RpcClient.GetTokenSupply(ctx, mintAddr)
		},
		convertGetTokenSupply,
	)
}

func (c *Client) GetTokenSupplyWithConfig(ctx context.Context, mintAddr string, cfg GetTokenSupplyConfig) (TokenAmount, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetTokenSupplyResultValue]], error) {
			return c.RpcClient.GetTokenSupplyWithConfig(ctx, mintAddr, cfg.toRpc())
		},
		convertGetTokenSupply,
	)
}

func (c *Client) GetTokenSupplyAndContext(ctx context.Context, mintAddr string) (rpc.ValueWithContext[TokenAmount], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetTokenSupplyResultValue]], error) {
			return c.RpcClient.GetTokenSupply(ctx, mintAddr)
		},
		convertGetTokenSupplyAndContext,
	)
}

func (c *Client) GetTokenSupplyAndContextWithConfig(ctx context.Context, mintAddr string, cfg GetTokenSupplyConfig) (rpc.ValueWithContext[TokenAmount], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetTokenSupplyResultValue]], error) {
			return c.RpcClient.GetTokenSupplyWithConfig(ctx, mintAddr, cfg.toRpc())
		},
		convertGetTokenSupplyAndContext,
	)
}

func convertGetTokenSupply(v rpc.ValueWithContext[rpc.GetTokenSupplyResultValue]) (TokenAmount, error) {
	return newTokenAmount(v.Value.Amount, v.Value.Decimals, v.Value.UIAmountString)
}

func convertGetTokenSupplyAndContext(v rpc.ValueWithContext[rpc.GetTokenSupplyResultValue]) (rpc.ValueWithContext[TokenAmount], error) {
	tokenSupply, err := convertGetTokenSupply(v)
	if err != nil {
		return rpc.ValueWithContext[TokenAmount]{}, err
	}
	return rpc.ValueWithContext[TokenAmount]{
		Context: v.Context,
		Value:   tokenSupply,
	}, nil
}
