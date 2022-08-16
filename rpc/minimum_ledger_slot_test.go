package rpc

import (
	"context"
	"testing"
)

func TestMinimumLedgerSlot(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"minimumLedgerSlot"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":84044778,"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.MinimumLedgerSlot(
					context.TODO(),
				)
			},
			ExpectedResponse: MinimumLedgerSlotResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: 84044778,
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
