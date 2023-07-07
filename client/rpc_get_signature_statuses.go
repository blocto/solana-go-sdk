package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

type GetSignatureStatusesConfig struct {
	SearchTransactionHistory bool
}

func (c GetSignatureStatusesConfig) toRpc() rpc.GetSignatureStatusesConfig {
	return rpc.GetSignatureStatusesConfig{
		SearchTransactionHistory: c.SearchTransactionHistory,
	}
}

func (c *Client) GetSignatureStatus(ctx context.Context, signature string) (*rpc.SignatureStatus, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.SignatureStatuses]], error) {
			return c.RpcClient.GetSignatureStatuses(ctx, []string{signature})
		},
		func(v rpc.ValueWithContext[rpc.SignatureStatuses]) (*rpc.SignatureStatus, error) {
			return v.Value[0], nil
		},
	)
}

func (c *Client) GetSignatureStatusWithConfig(ctx context.Context, signature string, cfg GetSignatureStatusesConfig) (*rpc.SignatureStatus, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.SignatureStatuses]], error) {
			return c.RpcClient.GetSignatureStatusesWithConfig(ctx, []string{signature}, cfg.toRpc())
		},
		func(v rpc.ValueWithContext[rpc.SignatureStatuses]) (*rpc.SignatureStatus, error) {
			return v.Value[0], nil
		},
	)
}

func (c *Client) GetSignatureStatuses(ctx context.Context, signatures []string) (rpc.SignatureStatuses, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.SignatureStatuses]], error) {
			return c.RpcClient.GetSignatureStatuses(ctx, signatures)
		},
		value[rpc.SignatureStatuses],
	)
}

func (c *Client) GetSignatureStatusesWithConfig(ctx context.Context, signatures []string, cfg GetSignatureStatusesConfig) (rpc.SignatureStatuses, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.SignatureStatuses]], error) {
			return c.RpcClient.GetSignatureStatusesWithConfig(ctx, signatures, cfg.toRpc())
		},
		value[rpc.SignatureStatuses],
	)
}
