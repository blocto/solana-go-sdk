package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

// GetVersion returns the current solana versions running on the node
func (c *Client) GetVersion(ctx context.Context) (rpc.GetVersion, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.GetVersion], error) {
			return c.RpcClient.GetVersion(ctx)
		},
		forward[rpc.GetVersion],
	)
}
