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

	res, err := c.GetSignaturesForAddress(context.Background(), "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	if err != nil {
		log.Fatalf("failed to GetSignaturesForAddress, err: %v", err)
	}

	spew.Dump(res)
}
