package rpc

import (
	"context"
	"testing"
)

func TestGetRecentBlockhash(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentBlockhash"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77387537},"value":{"blockhash":"867JxboSVrJLWQNZfF2odbP1QVVsd3DHYxbhsRX85Tsj","feeCalculator":{"lamportsPerSignature":5000}}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetRecentBlockhash(context.Background())
			},
			ExpectedResponse: GetRecentBlockHashResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetRecentBlockHashResult{
					Context: Context{
						Slot: 77387537,
					},
					Value: GetRecentBlockHashResultValue{
						Blockhash: "867JxboSVrJLWQNZfF2odbP1QVVsd3DHYxbhsRX85Tsj",
						FeeCalculator: FeeCalculator{
							LamportsPerSignature: 5000,
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentBlockhash", "params":[{"commitment": "finalized"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77387538},"value":{"blockhash":"5nNRmBkGM7CwtD9LUtd3pjHe33viBVjdGA1coq2Lz22E","feeCalculator":{"lamportsPerSignature":5000}}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetRecentBlockhashWithConfig(context.Background(), GetRecentBlockhashConfig{Commitment: CommitmentFinalized})
			},
			ExpectedResponse: GetRecentBlockHashResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetRecentBlockHashResult{
					Context: Context{
						Slot: 77387538,
					},
					Value: GetRecentBlockHashResultValue{
						Blockhash: "5nNRmBkGM7CwtD9LUtd3pjHe33viBVjdGA1coq2Lz22E",
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
