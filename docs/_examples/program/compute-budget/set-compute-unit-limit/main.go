package main

import (
	"context"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/program/cmptbdgprog"
	"github.com/portto/solana-go-sdk/program/memoprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

func main() {
	c := client.NewClient(rpc.LocalnetRPCEndpoint)

	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get latest blockhash, err: %v", err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				cmptbdgprog.SetComputeUnitLimit(cmptbdgprog.SetComputeUnitLimitParam{
					Units: 100000,
				}),
				memoprog.BuildMemo(memoprog.BuildMemoParam{
					Memo: []byte("👻"),
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to new a transaction, err: %v", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	log.Println("txhash:", txhash)
}
