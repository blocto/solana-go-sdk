package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetVoteAccounts(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,104,0]],"epochVoteAccount":true,"lastVote":134,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":103,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetVoteAccounts(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[GetVoteAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetVoteAccounts{
						Current: VoteAccounts{
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
						Deliquent: VoteAccounts{},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,2932,0]],"epochVoteAccount":true,"lastVote":2962,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":2931,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetVoteAccountsWithConfig(
						context.TODO(),
						GetVoteAccountsConfig{
							Commitment: CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetVoteAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetVoteAccounts{
						Current: VoteAccounts{
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
						Deliquent: VoteAccounts{},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"votePubkey": "5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,2900,0]],"epochVoteAccount":true,"lastVote":2930,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":2899,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetVoteAccountsWithConfig(
						context.TODO(),
						GetVoteAccountsConfig{
							VotePubkey: "5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp",
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetVoteAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetVoteAccounts{
						Current: VoteAccounts{
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
						Deliquent: VoteAccounts{},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"votePubkey": "2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"current":[],"delinquent":[]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetVoteAccountsWithConfig(
						context.TODO(),
						GetVoteAccountsConfig{
							VotePubkey: "2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV",
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetVoteAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetVoteAccounts{
						Current:   VoteAccounts{},
						Deliquent: VoteAccounts{},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"keepUnstakedDelinquents": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,2900,0]],"epochVoteAccount":true,"lastVote":2930,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":2899,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetVoteAccountsWithConfig(
						context.TODO(),
						GetVoteAccountsConfig{
							KeepUnstakedDelinquents: true,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetVoteAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetVoteAccounts{
						Current: VoteAccounts{
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
						Deliquent: VoteAccounts{},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts", "params":[{"delinquentSlotDistance": 100}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,2900,0]],"epochVoteAccount":true,"lastVote":2930,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":2899,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetVoteAccountsWithConfig(
						context.TODO(),
						GetVoteAccountsConfig{
							DelinquentSlotDistance: 100,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetVoteAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetVoteAccounts{
						Current: VoteAccounts{
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
						Deliquent: VoteAccounts{},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
