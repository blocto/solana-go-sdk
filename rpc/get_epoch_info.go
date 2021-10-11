package rpc

import "context"

// GetEpochInfoResponse is a full raw rpc response of `getEpochInfo`
type GetEpochInfoResponse struct {
	GeneralResponse
	Result GetEpochInfoResponseResult `json:"result"`
}

type GetEpochInfoResponseResult struct {
	AbsoluteSlot     uint64  `json:"absoluteSlot"`
	BlockHeight      uint64  `json:"blockHeight"`
	Epoch            uint64  `json:"epoch"`
	SlotIndex        uint64  `json:"slotIndex"`
	SlotsInEpoch     uint64  `json:"slotsInEpoch"`
	TransactionCount *uint64 `json:"transactionCount"`
}

// GetEpochInfoConfig is a option config for `getEpochInfo`
type GetEpochInfoConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetEpochInfo returns the SOL balance
func (c *RpcClient) GetEpochInfo(ctx context.Context) (GetEpochInfoResponse, error) {
	return c.processGetEpochInfo(c.Call(ctx, "getEpochInfo"))
}

// GetEpochInfoWithConfig returns the SOL balance
func (c *RpcClient) GetEpochInfoWithConfig(ctx context.Context, cfg GetEpochInfoConfig) (GetEpochInfoResponse, error) {
	return c.processGetEpochInfo(c.Call(ctx, "getEpochInfo", cfg))
}

func (c *RpcClient) processGetEpochInfo(body []byte, rpcErr error) (res GetEpochInfoResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
