// an account meta list indicates which accounts will be used in a program.
// we can only read/write accounts in progarms via this list.
// (except some official vars)

// an account meta includs
//   - isSigner
//     when an account is a signer. it needs to sign the tx.
//   - isWritable
//     when an account is writable. its data can be modified in this tx.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

var programId = common.PublicKeyFromString("CDKLz6tftV4kSD8sPVBD6ACqZpDY4Zuxf8rgSEYzR4M2")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get latest blockhash, err: %v\n", err)
	}

	firstAccount := types.NewAccount()
	fmt.Printf("first account: %v\n", firstAccount.PublicKey)

	secondAccount := types.NewAccount()
	fmt.Printf("second account: %v\n", secondAccount.PublicKey)

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer, firstAccount},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				{
					ProgramID: programId,
					Accounts: []types.AccountMeta{
						{
							PubKey:     firstAccount.PublicKey,
							IsSigner:   true,
							IsWritable: false,
						},
						{
							PubKey:     secondAccount.PublicKey,
							IsSigner:   false,
							IsWritable: true,
						},
					},
					Data: []byte{},
				},
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to new a tx, err: %v", err)
	}

	sig, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send the tx, err: %v", err)
	}

	// 4jYHKXhZMoDL3HsRuYhFPhCiQJhtNjDzPv8FhSnH6cMi9mwBjgW649uoqvfBjpGbkdFB53NEUux6oq3GUV8e9YQA
	fmt.Println(sig)
}
