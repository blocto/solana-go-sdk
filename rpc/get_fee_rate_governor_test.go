package rpc

import (
	"context"
	"testing"
)

func TestGetFeeRateGovernor(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeRateGovernor"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":15230},"value":{"feeRateGovernor":{"burnPercent":50,"maxLamportsPerSignature":100000,"minLamportsPerSignature":5000,"targetLamportsPerSignature":10000,"targetSignaturesPerSlot":20000}}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetFeeRateGovernor(
					context.TODO(),
				)
			},
			ExpectedResponse: GetFeeRateGovernorResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetFeeRateGovernorResponseResult{
					Context: Context{
						Slot: 15230,
					},
					Value: GetFeeRateGovernorResponseResultValue{
						FeeRateGovernor: FeeRateGovernor{
							MaxLamportsPerSignature:    100000,
							MinLamportsPerSignature:    5000,
							TargetLamportsPerSignature: 10000,
							TargetSignaturesPerSlot:    20000,
							BurnPercent:                50,
						},
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
