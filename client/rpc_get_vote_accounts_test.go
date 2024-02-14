package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestClient_GetVoteAccounts(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVoteAccounts"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"current":[{"activatedStake":999999997717120,"commission":0,"epochCredits":[[0,104,0]],"epochVoteAccount":true,"lastVote":134,"nodePubkey":"2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV","rootSlot":103,"votePubkey":"5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"}],"delinquent":[]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetVoteAccounts(
						context.Background(),
					)
				},
				ExpectedValue: VoteAccountStatus{
					Current: []VoteAccountInfo{
						{
							ActivatedStake: 999999997717120,
							Commission:     0,
							EpochCredits: []EpochCredits{
								{
									Epoch:           0,
									Credits:         104,
									PreviousCredits: 0,
								},
							},
							EpochVoteAccount: true,
							LastVote:         134,
							NodePubkey:       common.PublicKeyFromString("2RcYr2dvjgdJsbfPfAonBTSi7yU3JwdkHqZJWMCJYFAV"),
							VotePubkey:       common.PublicKeyFromString("5wi1m4kquajfcVVavvTuFWoMD4Nri4BJEUjV9pfCrhsp"),
							RootSlot:         103,
						},
					},
					Delinquent: []VoteAccountInfo{},
				},
				ExpectedError: nil,
			},
		},
	)
}
