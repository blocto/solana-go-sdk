package rpc

import (
	"context"
)

type GetHealthResponse JsonRpcResponse[string]

// GetHealthResponse returns the current health of the node. A healthy node is one that is within HEALTH_CHECK_SLOT_DISTANCE slots of the latest cluster confirmed slot.
func (c *RpcClient) GetHealth(ctx context.Context) (JsonRpcResponse[string], error) {
	return call[JsonRpcResponse[string]](c, ctx, "getHealth")
}
