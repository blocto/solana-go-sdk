package rpc

import (
	"context"
)

// MinimumLedgerSlotResponse is a full raw rpc response of `minimumLedgerSlot`
type MinimumLedgerSlotResponse struct {
	GeneralResponse
	Result uint64 `json:"result"`
}

// MinimumLedgerSlot returns the lowest slot that the node has information about in its ledger.
// This value may increase over time if the node is configured to purge older ledger data
func (c *RpcClient) MinimumLedgerSlot(ctx context.Context) (MinimumLedgerSlotResponse, error) {
	return c.processMinimumLedgerSlot(c.Call(ctx, "minimumLedgerSlot"))
}

func (c *RpcClient) processMinimumLedgerSlot(body []byte, rpcErr error) (res MinimumLedgerSlotResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
