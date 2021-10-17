package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

var feePayer, _ = types.AccountFromBytes([]byte{178, 244, 76, 4, 247, 41, 113, 40, 111, 103, 12, 76, 195, 4, 100, 123, 88, 226, 37, 56, 209, 180, 92, 77, 39, 85, 78, 202, 121, 162, 88, 29, 125, 155, 223, 107, 139, 223, 229, 82, 89, 209, 27, 43, 108, 205, 144, 2, 74, 159, 215, 57, 198, 4, 193, 36, 161, 50, 160, 119, 89, 240, 102, 184})

func main() {
	c := client.NewClient("http://localhost:8899")

	nonceAccountRentFreeBalance, err := c.GetMinimumBalanceForRentExemption(
		context.Background(),
		sysprog.NonceAccountSize,
	)
	if err != nil {
		log.Fatalf("failed to get min balance for nonce account, err: %v", err)
	}

	nonceAccount := types.NewAccount()
	fmt.Println("nonce account:", nonceAccount.PublicKey)

	sig, err := c.QuickSendTransaction(
		context.Background(),
		client.QuickSendTransactionParam{
			FeePayer: feePayer.PublicKey,
			Instructions: []types.Instruction{
				sysprog.CreateAccount(sysprog.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      nonceAccount.PublicKey,
					Owner:    common.SystemProgramID,
					Lamports: nonceAccountRentFreeBalance,
					Space:    sysprog.NonceAccountSize,
				}),
				sysprog.InitializeNonceAccount(
					// nonce account
					nonceAccount.PublicKey,
					// nonce account's owner
					feePayer.PublicKey,
				),
			},
			Signers: []types.Account{feePayer, nonceAccount},
		},
	)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	fmt.Println(sig)
}
