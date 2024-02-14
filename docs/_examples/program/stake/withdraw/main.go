package main

import (
	"context"
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

	aliceStakeAccountInfo, err := c.GetAccountInfo(context.Background(), aliceStakeAccountPubkey.String())
	if err != nil {
		log.Fatalf("failed to get stake account info, err: %v", err)
	}

	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				stake.Withdraw(stake.WithdrawParam{
					Stake:    aliceStakeAccountPubkey,
					Auth:     alice.PublicKey,
					To:       alice.PublicKey,
					Lamports: aliceStakeAccountInfo.Lamports, // withdraw all
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
