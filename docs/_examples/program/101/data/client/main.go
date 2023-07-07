// data is the most powerful part in an instruction
// we can pack everything into data, like number, pubkey ... whatever you want.
// we need to make them become a u8 array when we try to pack it in to a tx.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

var programId = common.PublicKeyFromString("c6vyXkJqgA85rYnLiMqxXd39fusJWbRJkoF3jXTd96H")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get latest blockhash, err: %v\n", err)
	}

	// (our example prgoram will parse the first byte as the selector then print remaining data.)
	{
		tx, err := types.NewTransaction(types.NewTransactionParam{
			Signers: []types.Account{feePayer},
			Message: types.NewMessage(types.NewMessageParam{
				FeePayer:        feePayer.PublicKey,
				RecentBlockhash: res.Blockhash,
				Instructions: []types.Instruction{
					{
						ProgramID: programId,
						Accounts:  []types.AccountMeta{},
						Data:      []byte{0, 1, 2, 3, 4},
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

		// 5X3qhwXJcjSZ3KqY8cTs5YbswBTS7yyqnfTn2diGwGERPrMNUjh9efc6Y9ABanfDUzaQN1n6BHHyMRjDJk2tfy1i
		fmt.Println(sig)
	}

	{
		tx, err := types.NewTransaction(types.NewTransactionParam{
			Signers: []types.Account{feePayer},
			Message: types.NewMessage(types.NewMessageParam{
				FeePayer:        feePayer.PublicKey,
				RecentBlockhash: res.Blockhash,
				Instructions: []types.Instruction{
					{
						ProgramID: programId,
						Accounts:  []types.AccountMeta{},
						Data:      []byte{1, 5, 6, 7, 8},
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

		// 5GETq1uLwMGdmsAH79pPByzGxQxXS1LqZ6Y9T6dBQ5qiMHFo5EgwDAHccFVEcc9hTYyj5zGfLX6j1uSz5NX7HZ8Q
		fmt.Println(sig)
	}
}
