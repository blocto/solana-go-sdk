// an instruction is the smallest unit in a tx.
// each instruction represent an action with a program.
// there are three basic fields in an instruction:
//   - program id
//     the program which you would like to interact with
//   - account meta list
//     accounts which are used in the program
//   - data
//     a u8 array.

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

var programId = common.PublicKeyFromString("EGz5CDh7dG7BwzqL7y5ePpZNvrw7ehr4E4oGRhCCpiEK")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get latest blockhash, err: %v\n", err)
	}

	// our first program won't use any accounts and parse any data. we leave them empty atm.
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				{
					ProgramID: programId,
					Accounts:  []types.AccountMeta{},
					Data:      []byte{},
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

	// 5eCfcdKGzn8yLGfgX6V1AfjAavTQRGQ936Zyuqejf9W1tTsjABjHTrc66nv5g9qKStmRPr3FeCVznADuCnJ8Zfbq
	fmt.Println(sig)
}
