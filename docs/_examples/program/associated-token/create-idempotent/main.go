package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/associated_token_account"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

// feeEo6y4CVaGzpxBqU5EyxtkTNUBrgKqkaNgT2mfWqs
var feePayer, _ = types.AccountFromBase58("znndxWePEfyktMfQP1AGP7ve47MhzxJj8A9S5GPQyo1pnumo6QvC1ci2XNeVGJgjhZjs2cUZmUu5Nii29P4ypxu")

// aceMErudbPawZPpwBr5tj28jS749rbFDWii9QMMgLUW
var alice, _ = types.AccountFromBase58("2y7PhFMP6ynvohRTWFkCocced7oDCT3x8AyvP5th9NZFN12CMq1LEocuXnNsES7AbvDZu8wS58aUMkigYCEn9uHY")

var mint = common.PublicKeyFromString("ntZA8ZQcKZwfC3ChVaJWRKN65mxKv9Cwqkpe6fvFZxs")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get latest blockhash, err: %v", err)
	}

	ata, _, err := common.FindAssociatedTokenAddress(alice.PublicKey, mint)
	if err != nil {
		log.Fatalf("failed to find associated token address, err: %v", err)
	}
	fmt.Println("ata", ata.ToBase58())

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				associated_token_account.CreateIdempotent(associated_token_account.CreateIdempotentParam{
					Funder:                 feePayer.PublicKey,
					Owner:                  alice.PublicKey,
					Mint:                   mint,
					AssociatedTokenAccount: ata,
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to create a transaction, err: %v", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	log.Println("txhash:", txhash)
}
