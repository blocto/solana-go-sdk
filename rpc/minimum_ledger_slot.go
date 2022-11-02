package rpc

import (
	"context"
)

type MinimumLedgerSlotResponse JsonRpcResponse[uint64]

// MinimumLedgerSlot returns the lowest slot that the node has information about in its ledger.
// This value may increase over time if the node is configured to purge older ledger data
func (c *RpcClient) MinimumLedgerSlot(ctx context.Context) (JsonRpcResponse[uint64], error) {
	return call[JsonRpcResponse[uint64]](c, ctx, "minimumLedgerSlot")
}
