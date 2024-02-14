package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/stake"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

// 9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde
var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

var aliceStakeAccountPubkey = common.PublicKeyFromString("oyRPx4Ejo11J6b4AGaCx9UXUvGzkEmZQoGxKqx4Yp4B")

func main() {
	c := client.NewClient(rpc.LocalnetRPCEndpoint)

	// obtain a random voting account here, or you can use your own. please note that a voting account is required here, rather than an identity.
	voteAccountStatus, err := c.GetVoteAccounts(context.Background())
	if err != nil {
		log.Fatalf("failed to get vote account status, err: %v", err)
	}
	if len(voteAccountStatus.Current) == 0 {
		log.Fatalf("there are no decent voting accounts")
	}
	delegatedVotePubkey := voteAccountStatus.Current[0].VotePubkey
	fmt.Println("delegated vote pubkey:", delegatedVotePubkey.String())

	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				stake.DelegateStake(stake.DelegateStakeParam{
					Stake: aliceStakeAccountPubkey,
					Auth:  alice.PublicKey,
					Vote:  delegatedVotePubkey,
				}),
			},
		}),
		Signers: []types.Account{feePayer, alice},
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("send tx error, err: %v\n", err)
	}

	log.Println("txhash:", txhash)
}
