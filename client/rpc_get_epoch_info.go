package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

type GetEpochInfo struct {
	AbsoluteSlot     uint64
	BlockHeight      uint64
	Epoch            uint64
	SlotIndex        uint64
	SlotsInEpoch     uint64
	TransactionCount *uint64
}

// GetEpochInfo returns information about the current epoch
func (c *Client) GetEpochInfo(ctx context.Context) (GetEpochInfo, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.GetEpochInfo], error) {
			return c.RpcClient.GetEpochInfo(ctx)
		},
		convertGetEpochInfo,
	)
}

func convertGetEpochInfo(v rpc.GetEpochInfo) (GetEpochInfo, error) {
	return GetEpochInfo{
		AbsoluteSlot:     v.AbsoluteSlot,
		BlockHeight:      v.BlockHeight,
		Epoch:            v.Epoch,
		SlotIndex:        v.SlotIndex,
		SlotsInEpoch:     v.SlotsInEpoch,
		TransactionCount: v.TransactionCount,
	}, nil
}
