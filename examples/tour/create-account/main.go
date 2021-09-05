package main

import (
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/types"
)

func main() {
	// create new account
	newAccount := types.NewAccount()
	fmt.Println(newAccount.PublicKey.ToBase58())
	fmt.Println(newAccount.PrivateKey)

	// recover account by its private key
	recoverAccount, err := types.AccountFromBytes(
		newAccount.PrivateKey,
	)
	if err != nil {
		log.Fatalf("failed to retrieve account from bytes, err: %v", err)
	}
	fmt.Println(recoverAccount.PublicKey.ToBase58())
}
