package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/memoprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

var feePayer, _ = types.AccountFromBytes([]byte{178, 244, 76, 4, 247, 41, 113, 40, 111, 103, 12, 76, 195, 4, 100, 123, 88, 226, 37, 56, 209, 180, 92, 77, 39, 85, 78, 202, 121, 162, 88, 29, 125, 155, 223, 107, 139, 223, 229, 82, 89, 209, 27, 43, 108, 205, 144, 2, 74, 159, 215, 57, 198, 4, 193, 36, 161, 50, 160, 119, 89, 240, 102, 184})

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	r, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	alice := types.NewAccount()
	fmt.Println("signer", alice.PublicKey.ToBase58())

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer, alice},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: r.Blockhash,
			Instructions: []types.Instruction{
				memoprog.BuildMemo(memoprog.BuildMemoParam{
					SignerPubkeys: []common.PublicKey{
						alice.PublicKey,
					},
					Memo: []byte("🐳"),
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to build tx, err: %v", err)
	}

	sig, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	fmt.Println(sig)
}
