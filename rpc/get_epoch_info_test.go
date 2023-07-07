package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
)

func TestGetEpochInfo(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getEpochInfo"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"absoluteSlot":86715160,"blockHeight":84901536,"epoch":200,"slotIndex":315160,"slotsInEpoch":432000,"transactionCount":2265984079},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetEpochInfo(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[GetEpochInfo]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetEpochInfo{
						AbsoluteSlot:     86715160,
						BlockHeight:      84901536,
						Epoch:            200,
						SlotIndex:        315160,
						SlotsInEpoch:     432000,
						TransactionCount: pointer.Get[uint64](2265984079),
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getEpochInfo", "params":[{"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"absoluteSlot":86715194,"blockHeight":84901570,"epoch":200,"slotIndex":315194,"slotsInEpoch":432000,"transactionCount":2265987458},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetEpochInfoWithConfig(
						context.TODO(),
						GetEpochInfoConfig{
							Commitment: CommitmentProcessed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetEpochInfo]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetEpochInfo{
						AbsoluteSlot:     86715194,
						BlockHeight:      84901570,
						Epoch:            200,
						SlotIndex:        315194,
						SlotsInEpoch:     432000,
						TransactionCount: pointer.Get[uint64](2265987458),
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
