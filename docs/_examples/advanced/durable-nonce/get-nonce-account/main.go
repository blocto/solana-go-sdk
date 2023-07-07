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
	nonceAccountAddr := "DJyNpXgggw1WGgjTVzFsNjb3fuQZVMqhoakvSBfX9LYx"
	nonceAccount, err := c.GetNonceAccount(context.Background(), nonceAccountAddr)
	if err != nil {
		log.Fatalf("failed to get nonce account, err: %v", err)
	}
	fmt.Printf("%+v\n", nonceAccount)
	/*
		type NonceAccount struct {
			Version          uint32
			State            uint32
			AuthorizedPubkey common.PublicKey
			Nonce            common.PublicKey
			FeeCalculator    FeeCalculator
		}
	*/
}
