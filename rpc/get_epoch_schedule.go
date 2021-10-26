package rpc

import "context"

// GetEpochScheduleResponse is a full raw rpc response of `getEpochSchedule`
type GetEpochScheduleResponse struct {
	GeneralResponse
	Result GetEpochScheduleResponseResult `json:"result"`
}

type GetEpochScheduleResponseResult struct {
	FirstNormalEpoch         uint64 `json:"firstNormalEpoch"`
	FirstNormalSlot          uint64 `json:"firstNormalSlot"`
	LeaderScheduleSlotOffset uint64 `json:"leaderScheduleSlotOffset"`
	SlotsPerEpoch            uint64 `json:"slotsPerEpoch"`
	Warmup                   bool   `json:"warmup"`
}

// GetEpochSchedule returns epoch schedule information from this cluster's genesis config
func (c *RpcClient) GetEpochSchedule(ctx context.Context) (GetEpochScheduleResponse, error) {
	return c.processGetEpochSchedule(c.Call(ctx, "getEpochSchedule"))
}

func (c *RpcClient) processGetEpochSchedule(body []byte, rpcErr error) (res GetEpochScheduleResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
