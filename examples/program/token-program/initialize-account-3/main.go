package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/program/tokenprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

// testUGLTMrKMpQuTLdgbQMgeHSCpJFgRbuV3ei4Cz45
var feePayer, _ = types.AccountFromBytes([]byte{191, 179, 234, 188, 179, 210, 40, 194, 39, 0, 0, 226, 107, 244, 219, 182, 242, 52, 70, 152, 96, 149, 80, 29, 241, 98, 132, 115, 20, 231, 204, 225, 13, 59, 115, 7, 254, 205, 140, 164, 208, 187, 139, 169, 158, 52, 4, 190, 19, 23, 97, 76, 246, 117, 202, 180, 200, 77, 65, 9, 30, 211, 3, 126})
var mint = common.PublicKeyFromString("5ziPkdJ4wbeTST56JPbWHX28uxbY2McHhdtWxeqcXvTi")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	tokenAccount := types.NewAccount()

	minBalance, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenprog.TokenAccountSize)
	if err != nil {
		log.Fatalf("failed to get min balance, err: %v", err)
	}

	sig, err := c.QuickSendTransaction(context.Background(), client.QuickSendTransactionParam{
		Instructions: []types.Instruction{
			sysprog.CreateAccount(sysprog.CreateAccountParam{
				From:     feePayer.PublicKey,
				New:      tokenAccount.PublicKey,
				Owner:    common.TokenProgramID,
				Lamports: minBalance,
				Space:    tokenprog.TokenAccountSize,
			}),
			// tokenprog.InitializeAccount3(
			// 	tokenAccount.PublicKey,
			// 	mint,
			// 	feePayer.PublicKey,
			// ),
		},
		Signers:  []types.Account{feePayer, tokenAccount},
		FeePayer: feePayer.PublicKey,
	})
	if err != nil {
		log.Fatalf("failed to get send tx, err: %v", err)
	}

	fmt.Printf("https://explorer.solana.com/tx/%v?cluster=devnet\n", sig)
}
