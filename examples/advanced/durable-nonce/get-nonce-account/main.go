package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/program/sysprog"
)

func main() {
	c := client.NewClient("http://localhost:8899")
	accountInfo, err := c.GetAccountInfo(
		context.Background(),
		"CjJWxNi3j8PyxSuTwSiJnLSbKuzV5JgRi8WpdPz1LzPX",
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
