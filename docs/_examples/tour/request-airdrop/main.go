package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	sig, err := c.RequestAirdrop(
		context.TODO(),
		"9qeP9DmjXAmKQc4wy133XZrQ3Fo4ejsYteA7X4YFJ3an", // address
		1e9, // lamports (1 SOL = 10^9 lamports)
	)
	if err != nil {
		log.Fatalf("failed to request airdrop, err: %v", err)
	}
	fmt.Println(sig)
}
