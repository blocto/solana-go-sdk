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

	var slot uint64
	var err error

	// get slot
	slot, err = c.GetSlot(context.TODO())
	if err != nil {
		log.Fatalf("failed to GetSlot, err: %v", err)
	}
	fmt.Println("GetSlot", slot)

	// get slot by commitment
	slot, err = c.GetSlotWithConfig(context.TODO(), rpc.GetSlotConfig{Commitment: rpc.CommitmentProcessed})
	if err != nil {
		log.Fatalf("failed to GetSlotWithConfig, err: %v", err)
	}
	fmt.Println("GetSlotWithConfig", slot)

	rpcRes, err := c.RpcClient.GetSlot(context.TODO())
	if err != nil {
		log.Fatalf("failed to RpcClient.GetSlot, err: %v", err)
	}
	fmt.Printf("res: %+v\n", rpcRes)

	rpcRes, err = c.RpcClient.GetSlotWithConfig(context.TODO(), rpc.GetSlotConfig{Commitment: rpc.CommitmentProcessed})
	if err != nil {
		log.Fatalf("failed to RpcClient.GetSlotWithConfig, err: %v", err)
	}
	fmt.Printf("res: %+v\n", rpcRes)
}
