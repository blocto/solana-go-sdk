package main

import (
	"fmt"

	"github.com/portto/solana-go-sdk/types"
)

func main() {
	// create new account
	newAccount := types.NewAccount()
	fmt.Println(newAccount.PublicKey.ToBase58())
	fmt.Println(newAccount.PrivateKey)

	// recover account by its private key
	recoverAccount := types.AccountFromPrivateKeyBytes(
		newAccount.PrivateKey,
	)
	fmt.Println(recoverAccount.PublicKey.ToBase58())
}
