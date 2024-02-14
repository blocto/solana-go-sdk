package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/stake"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

// 9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde
var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

var stakeAmountInLamports uint64 = 1_000_000_000 // 1 SOL

func main() {
	c := client.NewClient(rpc.LocalnetRPCEndpoint)

	// create an stake account
	stakeAccount := types.NewAccount()
	fmt.Println("stake account:", stakeAccount.PublicKey.ToBase58())

	// get rent
	rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(
		context.Background(),
		stake.AccountSize,
	)
	if err != nil {
		log.Fatalf("get min balacne for rent exemption, err: %v", err)
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
				system.CreateAccount(system.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      stakeAccount.PublicKey,
					Owner:    common.StakeProgramID,
					Lamports: rentExemptionBalance + stakeAmountInLamports,
					Space:    stake.AccountSize,
				}),
				stake.Initialize(stake.InitializeParam{
					Stake: stakeAccount.PublicKey,
					Auth: stake.Authorized{
						Staker:     alice.PublicKey,
						Withdrawer: alice.PublicKey,
					},
					Lockup: stake.Lockup{},
				}),
			},
		}),
		Signers: []types.Account{feePayer, stakeAccount},
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
