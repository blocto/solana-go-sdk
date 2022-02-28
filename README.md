
<h1 align="center">Solana Go SDK</h1>
<div align="center">
	<img src="https://github.com/portto/solana-go-sdk/actions/workflows/go.yml/badge.svg?branch=main"></img>
	<img src="https://goreportcard.com/badge/github.com/portto/solana-go-sdk"></img>
	<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/portto/solana-go-sdk">
	<img alt="GitHub release (latest SemVer)" src="https://img.shields.io/github/v/release/portto/solana-go-sdk?display_name=tag">
	<a href="https://yihau.gitbook.io/solana-go/">
		<img src="https://img.shields.io/badge/docs-gitbook-green"></img>
	</a>
</div>

# Guide

## Document

I have wrote a [gitbook](https://yihau.gitbook.io/solana-go). It includes many examples like, create account, create token, transfer SOL/Token, fetch NFT ...

You will find more info in the book!

Also feel free to post a issue for missing content.

## Getting Started

### Installing

```sh
go get -v github.com/portto/solana-go-sdk@v1.14.0
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
	"github.com/portto/solana-go-sdk/rpc"
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
	"github.com/portto/solana-go-sdk/rpc"
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
	fmt.Printf("balance: %v\n", balance)

	// get balance with sepcific commitment
	balance, err = c.GetBalanceWithConfig(
		context.TODO(),
		"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
		rpc.GetBalanceConfig{
			Commitment: rpc.CommitmentProcessed,
		},
	)
	if err != nil {
		log.Fatalf("failed to get balance with cfg, err: %v", err)
	}
	fmt.Printf("balance: %v\n", balance)

	// for advanced usage. fetch full rpc response
	res, err := c.RpcClient.GetBalance(
		context.TODO(),
		"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
	)
	if err != nil {
		log.Fatalf("failed to get balance via rpc client, err: %v", err)
	}
	fmt.Printf("response: %+v\n", res)
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

for more examples, follow `examples/` folder
