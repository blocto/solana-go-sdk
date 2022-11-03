package rpc

import (
	"context"
	"testing"
)

func TestGetVoteAccounts(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,104,0]],"epochVoteAccount":true,"lastVote":134,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":103,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetVoteAccounts(
					context.TODO(),
				)
			},
			ExpectedResponse: JsonRpcResponse[GetVoteAccounts]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetVoteAccounts{
					Current: VoteAccountInfos{
						{
							ActivatedStake:   999999997717120,
							Commission:       0,
							EpochVoteAccount: true,
							EpochCredits: [][3]uint64{
								{
									0, 104, 0,
								},
							},
							NodePubkey: "2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV",
							VotePubkey: "5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp",
							LastVote:   134,
							RootSlot:   103,
						},
					},
					Deliquent: VoteAccountInfos{},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"commitment": "confirmed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,2932,0]],"epochVoteAccount":true,"lastVote":2962,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":2931,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetVoteAccountsWithConfig(
					context.TODO(),
					GetVoteAccountsConfig{
						Commitment: CommitmentConfirmed,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetVoteAccounts]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetVoteAccounts{
					Current: VoteAccountInfos{
						{
							ActivatedStake:   999999997717120,
							Commission:       0,
							EpochVoteAccount: true,
							EpochCredits: [][3]uint64{
								{
									0, 2932, 0,
								},
							},
							NodePubkey: "2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV",
							VotePubkey: "5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp",
							LastVote:   2962,
							RootSlot:   2931,
						},
					},
					Deliquent: VoteAccountInfos{},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"votePubkey": "5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,2900,0]],"epochVoteAccount":true,"lastVote":2930,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":2899,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetVoteAccountsWithConfig(
					context.TODO(),
					GetVoteAccountsConfig{
						VotePubkey: "5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp",
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetVoteAccounts]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetVoteAccounts{
					Current: VoteAccountInfos{
						{
							ActivatedStake:   999999997717120,
							Commission:       0,
							EpochVoteAccount: true,
							EpochCredits: [][3]uint64{
								{
									0, 2900, 0,
								},
							},
							NodePubkey: "2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV",
							VotePubkey: "5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp",
							LastVote:   2930,
							RootSlot:   2899,
						},
					},
					Deliquent: VoteAccountInfos{},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"votePubkey": "2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"current":[],"delinquent":[]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetVoteAccountsWithConfig(
					context.TODO(),
					GetVoteAccountsConfig{
						VotePubkey: "2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV",
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetVoteAccounts]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetVoteAccounts{
					Current:   VoteAccountInfos{},
					Deliquent: VoteAccountInfos{},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"keepUnstakedDelinquents": true}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,2900,0]],"epochVoteAccount":true,"lastVote":2930,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":2899,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetVoteAccountsWithConfig(
					context.TODO(),
					GetVoteAccountsConfig{
						KeepUnstakedDelinquents: true,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetVoteAccounts]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetVoteAccounts{
					Current: VoteAccountInfos{
						{
							ActivatedStake:   999999997717120,
							Commission:       0,
							EpochVoteAccount: true,
							EpochCredits: [][3]uint64{
								{
									0, 2900, 0,
								},
							},
							NodePubkey: "2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV",
							VotePubkey: "5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp",
							LastVote:   2930,
							RootSlot:   2899,
						},
					},
					Deliquent: VoteAccountInfos{},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"delinquentSlotDistance": 100}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,2900,0]],"epochVoteAccount":true,"lastVote":2930,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":2899,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetVoteAccountsWithConfig(
					context.TODO(),
					GetVoteAccountsConfig{
						DelinquentSlotDistance: 100,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetVoteAccounts]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetVoteAccounts{
					Current: VoteAccountInfos{
						{
							ActivatedStake:   999999997717120,
							Commission:       0,
							EpochVoteAccount: true,
							EpochCredits: [][3]uint64{
								{
									0, 2900, 0,
								},
							},
							NodePubkey: "2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV",
							VotePubkey: "5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp",
							LastVote:   2930,
							RootSlot:   2899,
						},
					},
					Deliquent: VoteAccountInfos{},
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
