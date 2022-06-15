package main

import (
	"context"
	"fmt"
	"log"

	"github.com/OldSmokeGun/solana-go-sdk/client"
	"github.com/OldSmokeGun/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// get balance
	balance, err := c.GetBalance(
		context.TODO(),
		"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
	)
	if err != nil {
		log.Fatalf("failed to get balance, err: %v", err)
	}
	fmt.Printf("balance: %v\n", balance)

	// get balance with sepcific commitment
	balance, err = c.GetBalanceWithConfig(
		context.TODO(),
		"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
		rpc.GetBalanceConfig{
			Commitment: rpc.CommitmentProcessed,
		},
	)
	if err != nil {
		log.Fatalf("failed to get balance with cfg, err: %v", err)
	}
	fmt.Printf("balance: %v\n", balance)

	// for advanced usage. fetch full rpc response
	res, err := c.RpcClient.GetBalance(
		context.TODO(),
		"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
	)
	if err != nil {
		log.Fatalf("failed to get balance via rpc client, err: %v", err)
	}
	fmt.Printf("response: %+v\n", res)
}
