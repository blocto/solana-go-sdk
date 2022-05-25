package main

import (
	"context"
	"fmt"
	"log"

	"github.com/OldSmokeGun/solana-go-sdk/client"
	"github.com/OldSmokeGun/solana-go-sdk/common"
	"github.com/OldSmokeGun/solana-go-sdk/program/tokenprog"
	"github.com/OldSmokeGun/solana-go-sdk/rpc"
)

var mintPubkey = common.PublicKeyFromString("F6tecPzBMF47yJ2EN6j2aGtE68yR5jehXcZYVZa6ZETo")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	getAccountInfoResponse, err := c.GetAccountInfo(context.TODO(), mintPubkey.ToBase58())
	if err != nil {
		log.Fatalf("failed to get account info, err: %v", err)
	}

	mintAccount, err := tokenprog.MintAccountFromData(getAccountInfoResponse.Data)
	if err != nil {
		log.Fatalf("failed to parse data to a mint account, err: %v", err)
	}

	fmt.Printf("%+v\n", mintAccount)
	// {MintAuthority:9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde Supply:0 Decimals:8 IsInitialized:true FreezeAuthority:<nil>}
}
