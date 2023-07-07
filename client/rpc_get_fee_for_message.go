package client

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

type GetFeeForMessageConfig struct {
	Commitment rpc.Commitment
}

func (c GetFeeForMessageConfig) toRpc() rpc.GetFeeForMessageConfig {
	return rpc.GetFeeForMessageConfig{
		Commitment: c.Commitment,
	}
}

func (c *Client) GetFeeForMessage(ctx context.Context, message types.Message) (*uint64, error) {
	rawMessage, err := message.Serialize()
	if err != nil {
		return nil, fmt.Errorf("failed to serialize message, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[*uint64]], error) {
			return c.RpcClient.GetFeeForMessage(
				ctx,
				base64.StdEncoding.EncodeToString(rawMessage),
			)
		},
		value[*uint64],
	)
}

func (c *Client) GetFeeForMessageWithConfig(ctx context.Context, message types.Message, cfg GetFeeForMessageConfig) (*uint64, error) {
	rawMessage, err := message.Serialize()
	if err != nil {
		return nil, fmt.Errorf("failed to serialize message, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[*uint64]], error) {
			return c.RpcClient.GetFeeForMessageWithConfig(
				ctx,
				base64.StdEncoding.EncodeToString(rawMessage),
				cfg.toRpc(),
			)
		},
		value[*uint64],
	)
}

func (c *Client) GetFeeForMessageAndContext(ctx context.Context, message types.Message) (rpc.ValueWithContext[*uint64], error) {
	rawMessage, err := message.Serialize()
	if err != nil {
		return rpc.ValueWithContext[*uint64]{}, fmt.Errorf("failed to serialize message, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[*uint64]], error) {
			return c.RpcClient.GetFeeForMessage(
				ctx,
				base64.StdEncoding.EncodeToString(rawMessage),
			)
		},
		forward[rpc.ValueWithContext[*uint64]],
	)
}

func (c *Client) GetFeeForMessageAndContextWithConfig(ctx context.Context, message types.Message, cfg GetFeeForMessageConfig) (rpc.ValueWithContext[*uint64], error) {
	rawMessage, err := message.Serialize()
	if err != nil {
		return rpc.ValueWithContext[*uint64]{}, fmt.Errorf("failed to serialize message, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[*uint64]], error) {
			return c.RpcClient.GetFeeForMessageWithConfig(
				ctx,
				base64.StdEncoding.EncodeToString(rawMessage),
				cfg.toRpc(),
			)
		},
		forward[rpc.ValueWithContext[*uint64]],
	)
}
