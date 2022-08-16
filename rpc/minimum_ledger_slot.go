package rpc

import (
	"context"
)

type MinimumLedgerSlotResponse JsonRpcResponse[uint64]

// MinimumLedgerSlot returns the lowest slot that the node has information about in its ledger.
// This value may increase over time if the node is configured to purge older ledger data
func (c *RpcClient) MinimumLedgerSlot(ctx context.Context) (JsonRpcResponse[uint64], error) {
	return c.processMinimumLedgerSlot(c.Call(ctx, "minimumLedgerSlot"))
}

func (c *RpcClient) processMinimumLedgerSlot(body []byte, rpcErr error) (res JsonRpcResponse[uint64], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
