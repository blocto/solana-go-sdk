package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

// GetHealthResponse returns the current health of the node. A healthy node is one that is within HEALTH_CHECK_SLOT_DISTANCE slots of the latest cluster confirmed slot.
func (c *Client) GetHealth(ctx context.Context) (bool, error) {
	res, err := process(
		func() (rpc.JsonRpcResponse[string], error) {
			return c.RpcClient.GetHealth(ctx)
		},
		forward[string],
	)
	if err != nil {
		return false, err
	}
	if res != "ok" {
		return false, err
	}
	return true, nil
}
