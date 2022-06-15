package main

import (
	"context"
	"fmt"
	"log"

	"github.com/OldSmokeGun/solana-go-sdk/client"
	"github.com/OldSmokeGun/solana-go-sdk/common"
	"github.com/OldSmokeGun/solana-go-sdk/pkg/pointer"
	"github.com/OldSmokeGun/solana-go-sdk/program/assotokenprog"
	"github.com/OldSmokeGun/solana-go-sdk/program/metaplex/tokenmeta"
	"github.com/OldSmokeGun/solana-go-sdk/program/sysprog"
	"github.com/OldSmokeGun/solana-go-sdk/program/tokenprog"
	"github.com/OldSmokeGun/solana-go-sdk/rpc"
	"github.com/OldSmokeGun/solana-go-sdk/types"
	"github.com/mr-tron/base58"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

// 9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde
var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	mint := types.NewAccount()
	fmt.Printf("NFT: %v\n", mint.PublicKey.ToBase58())

	collection := types.NewAccount()
	fmt.Println(base58.Encode(collection.PrivateKey))
	fmt.Printf("collection: %v\n", collection.PublicKey.ToBase58())

	ata, _, err := common.FindAssociatedTokenAddress(feePayer.PublicKey, mint.PublicKey)
	if err != nil {
		log.Fatalf("failed to find a valid ata, err: %v", err)
	}

	tokenMetadataPubkey, err := tokenmeta.GetTokenMetaPubkey(mint.PublicKey)
	if err != nil {
		log.Fatalf("failed to find a valid token metadata, err: %v", err)

	}
	tokenMasterEditionPubkey, err := tokenmeta.GetMasterEdition(mint.PublicKey)
	if err != nil {
		log.Fatalf("failed to find a valid master edition, err: %v", err)
	}

	mintAccountRent, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenprog.MintAccountSize)
	if err != nil {
		log.Fatalf("failed to get mint account rent, err: %v", err)
	}

	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{mint, feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				sysprog.CreateAccount(sysprog.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      mint.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: mintAccountRent,
					Space:    tokenprog.MintAccountSize,
				}),
				tokenprog.InitializeMint(tokenprog.InitializeMintParam{
					Decimals: 0,
					Mint:     mint.PublicKey,
					MintAuth: feePayer.PublicKey,
				}),
				tokenmeta.CreateMetadataAccountV2(tokenmeta.CreateMetadataAccountV2Param{
					Metadata:                tokenMetadataPubkey,
					Mint:                    mint.PublicKey,
					MintAuthority:           feePayer.PublicKey,
					Payer:                   feePayer.PublicKey,
					UpdateAuthority:         feePayer.PublicKey,
					UpdateAuthorityIsSigner: true,
					IsMutable:               true,
					Data: tokenmeta.DataV2{
						Name:                 "Fake SMS #1355",
						Symbol:               "FSMB",
						Uri:                  "https://34c7ef24f4v2aejh75xhxy5z6ars4xv47gpsdrei6fiowptk2nqq.arweave.net/3wXyF1wvK6ARJ_9ue-O58CMuXrz5nyHEiPFQ6z5q02E",
						SellerFeeBasisPoints: 100,
						Creators: &[]tokenmeta.Creator{
							{
								Address:  feePayer.PublicKey,
								Verified: true,
								Share:    100,
							},
						},
						Collection: &tokenmeta.Collection{
							Verified: false,
							Key:      collection.PublicKey,
						},
						Uses: &tokenmeta.Uses{
							UseMethod: tokenmeta.Burn,
							Remaining: 10,
							Total:     10,
						},
					},
				}),
				assotokenprog.CreateAssociatedTokenAccount(assotokenprog.CreateAssociatedTokenAccountParam{
					Funder:                 feePayer.PublicKey,
					Owner:                  feePayer.PublicKey,
					Mint:                   mint.PublicKey,
					AssociatedTokenAccount: ata,
				}),
				tokenprog.MintTo(tokenprog.MintToParam{
					Mint:   mint.PublicKey,
					To:     ata,
					Auth:   feePayer.PublicKey,
					Amount: 1,
				}),
				tokenmeta.CreateMasterEditionV3(tokenmeta.CreateMasterEditionParam{
					Edition:         tokenMasterEditionPubkey,
					Mint:            mint.PublicKey,
					UpdateAuthority: feePayer.PublicKey,
					MintAuthority:   feePayer.PublicKey,
					Metadata:        tokenMetadataPubkey,
					Payer:           feePayer.PublicKey,
					MaxSupply:       pointer.Uint64(0),
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to new a tx, err: %v", err)
	}

	sig, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	fmt.Println(sig)
}
