package main

import (
	"context"
	"fmt"
	"log"

	"github.com/OldSmokeGun/solana-go-sdk/client"
	"github.com/OldSmokeGun/solana-go-sdk/program/sysprog"
	"github.com/OldSmokeGun/solana-go-sdk/rpc"
	"github.com/OldSmokeGun/solana-go-sdk/types"
)

// There are many ways you can send tx
var feePayer, _ = types.AccountFromBytes([]byte{}) // fill your private key here (u8 array)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// 1. use wrapped client to send tx (pros: easy to get started, cons: cannot use durable nonce machanism)
	sig, err := c.QuickSendTransaction(context.Background(), client.QuickSendTransactionParam{
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
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: resp.Blockhash,
			Instructions: []types.Instruction{
				sysprog.Transfer(sysprog.TransferParam{
					From:   feePayer.PublicKey,
					To:     feePayer.PublicKey,
					Amount: 1,
				}),
			},
		}),
		Signers: []types.Account{feePayer},
	})
	if err != nil {
		log.Fatalf("failed to build raw tx, err: %v", err)
	}
	sig, err = c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}
	fmt.Println(sig)

	// 3. use raw rpc to send your tx (pros: the most customable, cons: the most complex)
	// build tx like 2.
	// use c.RpcClient.SendTransaction() or c.RpcClient.SendTransactionWithConfig()
}
