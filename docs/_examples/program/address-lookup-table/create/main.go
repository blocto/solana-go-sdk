package main

import (
	"context"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/program/address_lookup_table"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get latest blockhash, err: %v", err)
	}

	slot, err := c.GetSlotWithConfig(context.Background(), rpc.GetSlotConfig{rpc.CommitmentFinalized})
	if err != nil {
		log.Fatalf("failed to get slot", err)
	}

	lookupTablePubkey, bumpSeed := address_lookup_table.DeriveLookupTableAddress(
		feePayer.PublicKey,
		slot,
	)
	log.Printf("account lookup address: %v\n", lookupTablePubkey)

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				address_lookup_table.CreateLookupTable(address_lookup_table.CreateLookupTableParams{
					LookupTable: lookupTablePubkey,
					Authority:   feePayer.PublicKey,
					Payer:       feePayer.PublicKey,
					RecentSlot:  slot,
					BumpSeed:    bumpSeed,
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
