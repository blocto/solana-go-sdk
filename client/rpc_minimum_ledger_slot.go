package client

import (
	"context"

	"github.com/portto/solana-go-sdk/rpc"
)

// MinimumLedgerSlot returns the lowest slot that the node has information about in its ledger.
// This value may increase over time if the node is configured to purge older ledger data
func (c *Client) MinimumLedgerSlot(ctx context.Context) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.MinimumLedgerSlot(ctx)
		},
		forward[uint64],
	)
}
