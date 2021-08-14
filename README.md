
<h1 align="center">Solana Go SDK</h1>
<div align="center">
	<img src="https://github.com/portto/solana-go-sdk/actions/workflows/go.yml/badge.svg?branch=main"></img>
	<img src="https://goreportcard.com/badge/github.com/portto/solana-go-sdk"></img>
	<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/portto/solana-go-sdk">
	<img alt="GitHub release (latest SemVer)" src="https://img.shields.io/github/v/release/portto/solana-go-sdk">
	<a href="https://yihau.gitbook.io/solana-development-with-go/">
		<img src="https://img.shields.io/badge/docs-gitbook-green"></img>
	</a>
</div>

# Guide

## Tutorial

There is a little tour in the [gitbook](https://yihau.gitbook.io/solana-development-with-go/) for newer to acquaint with Solana

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
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
)

func main() {
	c := client.NewClient(rpc.MainnetRPCEndpoint)

	resp, err := c.GetVersion(context.TODO())
	if err != nil {
		log.Fatalf("failed to version info, err: %v", err)
	}

	fmt.Println("version", resp.SolanaCore)
}

```

## RPC

All interfaces of rpc follow the [solana's json-rpc docs](https://docs.solana.com/developing/clients/jsonrpc-api).

The implementation of client in this project separate into two parts, rpc and wrapped. The wrapped only returns main result value and the rpc returns whole rpc response. You can switch it by yourself for different situation. Take `getBalance` as example:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// get balance
	balance, err := c.GetBalance(
		context.TODO(),
		"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
	)
	if err != nil {
		log.Fatalf("failed to get balance, err: %v", err)
	}
	fmt.Printf("balance: %v\n", balance) // balance: 6999995000

	// for advanced usage. fetch full rpc response
	res, err := c.RpcClient.GetBalance(
		context.TODO(),
		"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
		rpc.GetBalanceConfig{},
	)
	if err != nil {
		log.Fatalf("failed to get balance via rpc client, err: %v", err)
	}
	fmt.Printf("response: %+v\n", res) // response: {GeneralResponse:{JsonRPC:2.0 ID:1 Error:<nil>} Result:{Context:{Slot:73962152} Value:6999995000}}
}
```

## Programing model & Program

There are some important tpyes in solana.

- Program

resides in the `program/` folder.

- Pubkey (a basic identity of key)

resides in the `common/` folder.

- Insturciton (contain many pubkeys and program ID)
- Message (contain many instructions)
- Transaction (contain a message and many signatures)
- Account (a pub/pri keypair )

reside in the `types/` folder.

### More Example

#### New Account & Get Some Airdrop

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/types"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	account := types.NewAccount()
	fmt.Println("account address:", account.PublicKey.ToBase58())
	fmt.Println("account private key:", account.PrivateKey)

	airdropTxHash, err := c.RequestAirdrop(
		context.Background(),
		account.PublicKey.ToBase58(),
		1e9, // 1 SOL = 10^9 lamports
	)
	if err != nil {
		log.Fatalln("request airdrop error", err)
	}

	fmt.Println("airdrop txhash:", airdropTxHash)
	// you can lookup this txhash on https://explorer.solana.com/?cluster=devnet
}

```

#### Send Transaction

There are two ways to compose transaction

```go
package main

import (
	"context"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	feePayer := types.AccountFromPrivateKeyBytes([]byte{128, 146, 1, 80, 86, 97, 143, 62, 20, 136, 245, 33, 79, 63, 34, 54, 115, 6, 9, 77, 99, 157, 156, 100, 177, 229, 245, 8, 25, 25, 68, 165, 38, 28, 93, 198, 46, 101, 158, 208, 135, 126, 226, 94, 66, 153, 164, 162, 19, 231, 38, 240, 114, 74, 116, 32, 178, 61, 64, 95, 187, 211, 239, 180})

	// create a random receiver
	to := types.NewAccount()

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			// use system program's instruction, transfer
			sysprog.Transfer(
				feePayer.PublicKey,
				to.PublicKey,
				1, // 1 lamports
			),
		},
		Signers:         []types.Account{feePayer},
		FeePayer:        feePayer.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txSig, err := c.SendRawTransaction(context.Background(), rawTx)
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
	"context"
	"crypto/ed25519"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	feePayer := types.AccountFromPrivateKeyBytes([]byte{128, 146, 1, 80, 86, 97, 143, 62, 20, 136, 245, 33, 79, 63, 34, 54, 115, 6, 9, 77, 99, 157, 156, 100, 177, 229, 245, 8, 25, 25, 68, 165, 38, 28, 93, 198, 46, 101, 158, 208, 135, 126, 226, 94, 66, 153, 164, 162, 19, 231, 38, 240, 114, 74, 116, 32, 178, 61, 64, 95, 187, 211, 239, 180})

	// create a random receiver
	to := types.NewAccount()

	// prepare message
	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	message := types.NewMessage(
		feePayer.PublicKey,
		[]types.Instruction{
			sysprog.Transfer(feePayer.PublicKey, to.PublicKey, 1),
		},
		res.Blockhash,
	)
	serializeMessage, err := message.Serialize()
	if err != nil {
		log.Fatalf("serialize message error, err: %v\n", err)
	}

	// message + signature = tx
	tx, err := types.CreateTransaction(
		message,
		map[common.PublicKey]types.Signature{
			feePayer.PublicKey: ed25519.Sign(feePayer.PrivateKey, serializeMessage),
		},
	)
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}
	rawTx, err := tx.Serialize()
	if err != nil {
		log.Fatalf("serialize tx error, err: %v\n", err)
	}

	txSig, err := c.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		log.Fatalf("send tx error, err: %v\n", err)
	}

	log.Println("txHash:", txSig)
}

```

for more examples, I'm going to update in `examples/` folder
