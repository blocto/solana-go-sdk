package rpc

import (
	"context"
	"testing"
)

func TestGetEpochSchedule(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getEpochSchedule"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"firstNormalEpoch":0,"firstNormalSlot":0,"leaderScheduleSlotOffset":432000,"slotsPerEpoch":432000,"warmup":false},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetEpochSchedule(
					context.TODO(),
				)
			},
			ExpectedResponse: JsonRpcResponse[GetEpochSchedule]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetEpochSchedule{
					FirstNormalEpoch:         0,
					FirstNormalSlot:          0,
					LeaderScheduleSlotOffset: 432000,
					SlotsPerEpoch:            432000,
					Warmup:                   false,
				},
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			testRpcCall(t, tt)
		})
	}
}
