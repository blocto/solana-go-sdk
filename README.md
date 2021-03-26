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
}
```

```sh
go run main.go
# solana version: 1.6.1
```

you can find more examples in [wiki page](https://github.com/portto/solana-go-sdk/wiki)

## Testing

```sh
go test -v ./...
```
