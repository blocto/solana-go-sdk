package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestRequestAirdrop(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"requestAirdrop", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", 1]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"2HsNt2iPgHVKzbYxsivEUWutMAzFkJL1YBs7phaBTtKY82sbDLuhzEBqbmGwxBAWTRSiPwkqop8vqWxezkxcuaVW","id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.RequestAirdrop(
						context.TODO(),
						"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
						1,
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  "2HsNt2iPgHVKzbYxsivEUWutMAzFkJL1YBs7phaBTtKY82sbDLuhzEBqbmGwxBAWTRSiPwkqop8vqWxezkxcuaVW",
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"requestAirdrop", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", 1, {"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"4eAWQLipk6hA7AcRuLUw2VXbiHRVR6HABVcQQqaBuukNNMohWg4ToAn4Qh2RaiFnK1LiUxrGnVgm1n4kpUbB7Yt9","id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.RequestAirdropWithConfig(
						context.TODO(),
						"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
						1,
						RequestAirdropConfig{
							Commitment: CommitmentProcessed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  "4eAWQLipk6hA7AcRuLUw2VXbiHRVR6HABVcQQqaBuukNNMohWg4ToAn4Qh2RaiFnK1LiUxrGnVgm1n4kpUbB7Yt9",
				},
				ExpectedError: nil,
			},
		},
	)
}
