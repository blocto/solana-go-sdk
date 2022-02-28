package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.LocalnetRPCEndpoint)
	accountInfo, err := c.GetAccountInfo(
		context.Background(),
		"CJBP7wJcYbPqfhvtSmLBUf4VzBqJbgC776Wr7CzUCd1m",
	)
	if err != nil {
		log.Fatalf("failed to get account info, err: %v", err)
	}

	nonceAccount, err := sysprog.NonceAccountDeserialize(accountInfo.Data)
	if err != nil {
		log.Fatalf("failed to deserialize nonce account, err: %v", err)
	}

	/*
		type NonceAccount struct {
			Version          uint32
			State            uint32
			AuthorizedPubkey common.PublicKey
			Nonce            common.PublicKey
			FeeCalculator    FeeCalculator
		}
	*/
	fmt.Printf("%+v\n", nonceAccount)
}
