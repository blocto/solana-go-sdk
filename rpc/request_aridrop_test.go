package rpc

import (
	"context"
	"testing"
)

func TestRequestAirdrop(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"requestAirdrop", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", 1]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":"2HsNt2iPgHVKzbYxsivEUWutMAzFkJL1YBs7phaBTtKY82sbDLuhzEBqbmGwxBAWTRSiPwkqop8vqWxezkxcuaVW","id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.RequestAirdrop(
					context.TODO(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
					1,
				)
			},
			ExpectedResponse: RequestAirdropResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: "2HsNt2iPgHVKzbYxsivEUWutMAzFkJL1YBs7phaBTtKY82sbDLuhzEBqbmGwxBAWTRSiPwkqop8vqWxezkxcuaVW",
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"requestAirdrop", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", 1, {"commitment": "processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":"4eAWQLipk6hA7AcRuLUw2VXbiHRVR6HABVcQQqaBuukNNMohWg4ToAn4Qh2RaiFnK1LiUxrGnVgm1n4kpUbB7Yt9","id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.RequestAirdropWithConfig(
					context.TODO(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
					1,
					RequestAirdropConfig{
						Commitment: CommitmentProcessed,
					},
				)
			},
			ExpectedResponse: RequestAirdropResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: "4eAWQLipk6hA7AcRuLUw2VXbiHRVR6HABVcQQqaBuukNNMohWg4ToAn4Qh2RaiFnK1LiUxrGnVgm1n4kpUbB7Yt9",
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
