package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	nonceAccountAddr := "DJyNpXgggw1WGgjTVzFsNjb3fuQZVMqhoakvSBfX9LYx"
	nonce, err := c.GetNonceFromNonceAccount(context.Background(), nonceAccountAddr)
	if err != nil {
		log.Fatalf("failed to get nonce account, err: %v", err)
	}

	fmt.Println("nonce", nonce)
}
