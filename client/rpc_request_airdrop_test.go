package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/rpc"
)

func TestClient_RequestAirdrop(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"requestAirdrop", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", 1]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"2HsNt2iPgHVKzbYxsivEUWutMAzFkJL1YBs7phaBTtKY82sbDLuhzEBqbmGwxBAWTRSiPwkqop8vqWxezkxcuaVW","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.RequestAirdrop(
						context.TODO(),
						"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
						1,
					)
				},
				ExpectedValue: "2HsNt2iPgHVKzbYxsivEUWutMAzFkJL1YBs7phaBTtKY82sbDLuhzEBqbmGwxBAWTRSiPwkqop8vqWxezkxcuaVW",
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_RequestAirdropWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"requestAirdrop", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7", 1, {"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"4eAWQLipk6hA7AcRuLUw2VXbiHRVR6HABVcQQqaBuukNNMohWg4ToAn4Qh2RaiFnK1LiUxrGnVgm1n4kpUbB7Yt9","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.RequestAirdropWithConfig(
						context.TODO(),
						"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
						1,
						RequestAirdropConfig{
							Commitment: rpc.CommitmentProcessed,
						},
					)
				},
				ExpectedValue: "4eAWQLipk6hA7AcRuLUw2VXbiHRVR6HABVcQQqaBuukNNMohWg4ToAn4Qh2RaiFnK1LiUxrGnVgm1n4kpUbB7Yt9",
				ExpectedError: nil,
			},
		},
	)
}
