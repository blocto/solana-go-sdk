package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

// There are many ways you can send tx
var feePayer = types.AccountFromPrivateKeyBytes([]byte{}) // fill your private key here (u8 array)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// 1. use wrapped client to send tx (pros: easy to get started, cons: cannot use durable nonce machanism)
	sig, err := c.SendTransaction(context.Background(), client.SendTransactionParam{
		Instructions: []types.Instruction{
			// your instruction here
		},
		Signers:  []types.Account{feePayer},
		FeePayer: feePayer.PublicKey,
	})
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}
	fmt.Println(sig)

	// 2. send raw tx (pros: more custom tx you can send, cons: build tx steps are more complex)
	resp, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			sysprog.Transfer(feePayer.PublicKey, feePayer.PublicKey, 1),
		},
		Signers:         []types.Account{feePayer},
		FeePayer:        feePayer.PublicKey,
		RecentBlockHash: resp.Blockhash,
	})
	if err != nil {
		log.Fatalf("failed to build raw tx, err: %v", err)
	}
	sig, err = c.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}
	fmt.Println(sig)

	// 3. use raw rpc to send your tx (pros: the most customable, cons: the most complex)
	// build tx like 2.
	// use c.RpcClient.SendTransaction() or c.RpcClient.SendTransactionWithConfig()
}
