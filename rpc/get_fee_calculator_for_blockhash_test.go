package rpc

import (
	"context"
	"testing"
)

func TestGetFeeCalculatorForBlockhash(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeCalculatorForBlockhash", "params":["9EbVxv9NLnXVBUCdY5LEkv89iHVHVyFnfzHNZgP4eF5V"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":13752},"value":null},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetFeeCalculatorForBlockhash(
					context.TODO(),
					"9EbVxv9NLnXVBUCdY5LEkv89iHVHVyFnfzHNZgP4eF5V",
				)
			},
			ExpectedResponse: GetFeeCalculatorForBlockhashResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetFeeCalculatorForBlockhashResponseResult{
					Context: Context{
						Slot: 13752,
					},
					Value: nil,
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeCalculatorForBlockhash", "params":["BVvVdTXHSUMDnnCAMy53DHQ8zy27L7Ehm6WZsywbAPT7"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":13752},"value":{"feeCalculator":{"lamportsPerSignature":5000}}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetFeeCalculatorForBlockhash(
					context.TODO(),
					"BVvVdTXHSUMDnnCAMy53DHQ8zy27L7Ehm6WZsywbAPT7",
				)
			},
			ExpectedResponse: GetFeeCalculatorForBlockhashResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetFeeCalculatorForBlockhashResponseResult{
					Context: Context{
						Slot: 13752,
					},
					Value: &GetFeeCalculatorForBlockhashResponseResultValue{
						FeeCalculator: FeeCalculator{
							LamportsPerSignature: 5000,
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeCalculatorForBlockhash", "params":["BVvVdTXHSUMDnnCAMy53DHQ8zy27L7Ehm6WZsywbAPT7", {"commitment": "confirmed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":13752},"value":{"feeCalculator":{"lamportsPerSignature":5000}}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetFeeCalculatorForBlockhashWithConfig(
					context.TODO(),
					"BVvVdTXHSUMDnnCAMy53DHQ8zy27L7Ehm6WZsywbAPT7",
					GetFeeCalculatorForBlockhashConfig{
						Commitment: CommitmentConfirmed,
					},
				)
			},
			ExpectedResponse: GetFeeCalculatorForBlockhashResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetFeeCalculatorForBlockhashResponseResult{
					Context: Context{
						Slot: 13752,
					},
					Value: &GetFeeCalculatorForBlockhashResponseResultValue{
						FeeCalculator: FeeCalculator{
							LamportsPerSignature: 5000,
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
