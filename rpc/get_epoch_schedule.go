package rpc

import "context"

type GetEpochScheduleResponse JsonRpcResponse[GetEpochSchedule]

type GetEpochSchedule struct {
	FirstNormalEpoch         uint64 `json:"firstNormalEpoch"`
	FirstNormalSlot          uint64 `json:"firstNormalSlot"`
	LeaderScheduleSlotOffset uint64 `json:"leaderScheduleSlotOffset"`
	SlotsPerEpoch            uint64 `json:"slotsPerEpoch"`
	Warmup                   bool   `json:"warmup"`
}

// GetEpochSchedule returns epoch schedule information from this cluster's genesis config
func (c *RpcClient) GetEpochSchedule(ctx context.Context) (JsonRpcResponse[GetEpochSchedule], error) {
	return call[JsonRpcResponse[GetEpochSchedule]](c, ctx, "getEpochSchedule")
}
