package rpc

import (
	"context"
	"testing"
)

func TestGetInflationRate(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getInflationRate"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"epoch":200,"foundation":0.0,"total":0.06956826778571164,"validator":0.06956826778571164},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetInflationRate(context.TODO())
			},
			ExpectedResponse: GetInflationRateResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetInflationRateResult{
					Epoch:      200,
					Foundation: 0.0,
					Total:      0.06956826778571164,
					Validator:  0.06956826778571164,
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
