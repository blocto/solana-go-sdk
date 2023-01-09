package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

// GetIdentity returns the identity pubkey for the current node
func (c *Client) GetIdentity(ctx context.Context) (rpc.GetIdentity, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.GetIdentity], error) {
			return c.RpcClient.GetIdentity(ctx)
		},
		forward[rpc.GetIdentity],
	)
}
