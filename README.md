# solana-go-sdk
[![Go](https://github.com/portto/solana-go-sdk/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/portto/solana-go-sdk/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/portto/solana-go-sdk)](https://goreportcard.com/report/github.com/portto/solana-go-sdk)

Solana Golang SDK

## Getting Started

### Installing

```sh
go get -v github.com/portto/solana-go-sdk
```

### Example

#### Hello World

```go
package main

import (
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
)

func main() {
	c := client.NewClient(client.TestnetRPCEndpoint)
  
	res, err := c.GetVersion()
	if err != nil {
		log.Fatalf("get version error, err: %v", err)
	}
  
	fmt.Println("solana version:", res.SolanaCore)
	
	// solana version: 1.6.1
}
```

#### New Account

```go
package main

import (
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/types"
)

func main() {
	c := client.NewClient(client.TestnetRPCEndpoint)

	account := types.NewAccount()
	fmt.Println("account address:", account.PublicKey.ToBase58())
	fmt.Println("account private key:", account.PrivateKey)

	airdropTxHash, err := c.RequestAirdrop(account.PublicKey.ToBase58(), 1000000000) // 1 SOL = 10e9 lamports
	if err != nil {
		log.Fatalln("request airdrop error", err)
	}
	fmt.Println("airdrop txhash:", airdropTxHash)
}
```

#### Send Transaction

There are two ways to generate tx.

You can use `CreateRawTransaction` to generate raw tx

```go
package main

import (
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

func main() {
	c := client.NewClient(client.TestnetRPCEndpoint)

	res, err := c.GetRecentBlockhash()
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}

	feePayer := types.AccountFromPrivateKeyBytes([]byte{57, 17, 193, 142, 252, 221, 81, 90, 60, 28, 93, 237, 212, 51, 95, 95, 41, 104, 221, 59, 13, 244, 54, 1, 79, 180, 120, 178, 81, 45, 46, 193, 142, 11, 237, 209, 82, 24, 36, 72, 7, 76, 66, 215, 44, 116, 17, 132, 252, 205, 47, 74, 57, 230, 36, 98, 119, 86, 11, 40, 71, 195, 47, 254})

	accountA := types.AccountFromPrivateKeyBytes([]byte{185, 195, 153, 239, 225, 24, 36, 241, 184, 42, 46, 216, 48, 6, 157, 169, 66, 255, 174, 87, 189, 1, 255, 106, 202, 38, 57, 214, 26, 188, 154, 136, 119, 16, 250, 24, 46, 183, 154, 255, 79, 143, 141, 38, 53, 77, 142, 212, 243, 8, 185, 215, 169, 217, 96, 191, 139, 190, 157, 96, 101, 88, 27, 79})

	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			sysprog.Transfer(
				accountA.PublicKey, // from
				common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"), // to
				1000000000, // 1 SOL
			),
		},
		Signers:         []types.Account{feePayer, accountA},
		FeePayer:        feePayer.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txSig, err := c.SendRawTransaction(rawTx)
	if err != nil {
		log.Fatalf("send tx error, err: %v\n", err)
	}

	log.Println("txHash:", txSig)
}
```

or you can create raw message then fill signatures by yourself.

```go
package main

import (
	"crypto/ed25519"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

func main() {
	c := client.NewClient(client.TestnetRPCEndpoint)

	res, err := c.GetRecentBlockhash()
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}

	feePayer := types.AccountFromPrivateKeyBytes([]byte{57, 17, 193, 142, 252, 221, 81, 90, 60, 28, 93, 237, 212, 51, 95, 95, 41, 104, 221, 59, 13, 244, 54, 1, 79, 180, 120, 178, 81, 45, 46, 193, 142, 11, 237, 209, 82, 24, 36, 72, 7, 76, 66, 215, 44, 116, 17, 132, 252, 205, 47, 74, 57, 230, 36, 98, 119, 86, 11, 40, 71, 195, 47, 254})

	accountA := types.AccountFromPrivateKeyBytes([]byte{185, 195, 153, 239, 225, 24, 36, 241, 184, 42, 46, 216, 48, 6, 157, 169, 66, 255, 174, 87, 189, 1, 255, 106, 202, 38, 57, 214, 26, 188, 154, 136, 119, 16, 250, 24, 46, 183, 154, 255, 79, 143, 141, 38, 53, 77, 142, 212, 243, 8, 185, 215, 169, 217, 96, 191, 139, 190, 157, 96, 101, 88, 27, 79})

	message := types.NewMessage(
		feePayer.PublicKey,
		[]types.Instruction{
			sysprog.Transfer(
				accountA.PublicKey, // from
				common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"), // to
				1000000000, // 1 SOL
			),
		},
		res.Blockhash,
	)

	serializeMessage, err := message.Serialize()
	if err != nil {
		log.Fatalf("serialize message error, err: %v\n", err)
	}

	tx, err := types.CreateTransaction(message, map[common.PublicKey]types.Signature{
		feePayer.PublicKey: ed25519.Sign(feePayer.PrivateKey, serializeMessage),
		accountA.PublicKey: ed25519.Sign(accountA.PrivateKey, serializeMessage),
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	rawTx, err := tx.Serialize()
	if err != nil {
		log.Fatalf("serialize tx error, err: %v\n", err)
	}

	txSig, err := c.SendRawTransaction(rawTx)
	if err != nil {
		log.Fatalf("send tx error, err: %v\n", err)
	}

	log.Println("txHash:", txSig)
}
```
