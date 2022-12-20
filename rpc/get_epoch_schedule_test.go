package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetEpochSchedule(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getEpochSchedule"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"firstNormalEpoch":0,"firstNormalSlot":0,"leaderScheduleSlotOffset":432000,"slotsPerEpoch":432000,"warmup":false},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetEpochSchedule(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[GetEpochSchedule]{
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
		},
	)
}
