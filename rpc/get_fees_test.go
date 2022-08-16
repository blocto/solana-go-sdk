package rpc

import (
	"context"
	"testing"
)

func TestGetFees(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFees"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":16039},"value":{"blockhash":"3nWZN5T3JeQzxcAJrm6U9pXUxvyQA7nwn3ChucpYBhzo","feeCalculator":{"lamportsPerSignature":5000},"lastValidBlockHeight":16339,"lastValidSlot":16339}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetFees(
					context.TODO(),
				)
			},
			ExpectedResponse: GetFeesResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetFeesResponseResult{
					Context: Context{
						Slot: 16039,
					},
					Value: GetFeesResponseResultValue{
						Blockhash: "3nWZN5T3JeQzxcAJrm6U9pXUxvyQA7nwn3ChucpYBhzo",
						FeeCalculator: FeeCalculator{
							LamportsPerSignature: 5000,
						},
						LastValidSlot:        16339,
						LastValidBlockHeight: 16339,
					},
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
