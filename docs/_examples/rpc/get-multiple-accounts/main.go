package main

import (
	"context"
	"log"

	"github.com/OldSmokeGun/solana-go-sdk/client"
	"github.com/OldSmokeGun/solana-go-sdk/rpc"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	accountInfos, err := c.GetMultipleAccounts(
		context.Background(),
		[]string{
			"AngPj3DQh1zNW68E1sa2yWfy71D6mcd7iwzhbLiLMxcR",
			"4T2BWArMHpFiwAMkMsJUCxa5ux9vYZBnVigJznxzbeVx",
		},
	)
	if err != nil {
		log.Fatalf("failed to get multiple account infos, err: %v", err)
	}

	spew.Dump(accountInfos)
}
