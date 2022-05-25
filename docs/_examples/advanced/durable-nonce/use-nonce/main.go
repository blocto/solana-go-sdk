package main

import (
	"context"
	"fmt"
	"log"

	"github.com/OldSmokeGun/solana-go-sdk/client"
	"github.com/OldSmokeGun/solana-go-sdk/common"
	"github.com/OldSmokeGun/solana-go-sdk/program/memoprog"
	"github.com/OldSmokeGun/solana-go-sdk/program/sysprog"
	"github.com/OldSmokeGun/solana-go-sdk/rpc"
	"github.com/OldSmokeGun/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

// 9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde
var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// get nonce account
	nonceAccountPubkey := common.PublicKeyFromString("DJyNpXgggw1WGgjTVzFsNjb3fuQZVMqhoakvSBfX9LYx")
	nonceAccount, err := c.GetNonceAccount(context.Background(), nonceAccountPubkey.ToBase58())
	if err != nil {
		log.Fatalf("failed to get nonce account, err: %v", err)
	}

	// create a tx
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer, alice},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: nonceAccount.Nonce.ToBase58(),
			Instructions: []types.Instruction{
				sysprog.AdvanceNonceAccount(sysprog.AdvanceNonceAccountParam{
					Nonce: nonceAccountPubkey,
					Auth:  alice.PublicKey,
				}),
				memoprog.BuildMemo(memoprog.BuildMemoParam{
					Memo: []byte("use nonce"),
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to new a transaction, err: %v", err)
	}

	sig, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	fmt.Println("txhash", sig)
}
