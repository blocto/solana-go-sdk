package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
)

func TestClient_GetEpochInfo(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getEpochInfo"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"absoluteSlot":86715160,"blockHeight":84901536,"epoch":200,"slotIndex":315160,"slotsInEpoch":432000,"transactionCount":2265984079},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetEpochInfo(context.TODO())
				},
				ExpectedValue: GetEpochInfo{
					AbsoluteSlot:     86715160,
					BlockHeight:      84901536,
					Epoch:            200,
					SlotIndex:        315160,
					SlotsInEpoch:     432000,
					TransactionCount: pointer.Get[uint64](2265984079),
				},
				ExpectedError: nil,
			},
		},
	)
}
