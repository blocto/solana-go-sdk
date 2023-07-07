package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	res, err := c.RpcClient.GetProgramAccountsWithConfig(
		context.Background(),
		common.SystemProgramID.ToBase58(),
		rpc.GetProgramAccountsConfig{
			Encoding: rpc.AccountEncodingBase64,
			Filters: []rpc.GetProgramAccountsConfigFilter{
				{
					DataSize: system.NonceAccountSize,
				},
				{
					MemCmp: &rpc.GetProgramAccountsConfigFilterMemCmp{
						Offset: 8,
						Bytes:  "9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde", // owner address
					},
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("failed to get program accounts, err: %v", err)
	}

	for _, a := range res.Result {
		fmt.Println("pubkey", a.Pubkey)
		data, err := base64.StdEncoding.DecodeString((a.Account.Data.([]any))[0].(string))
		if err != nil {
			log.Fatalf("failed to decode data, err: %v", err)
		}
		nonceAccount, err := system.NonceAccountDeserialize(data)
		if err != nil {
			log.Fatalf("failed to parse nonce account, err: %v", err)
		}
		fmt.Printf("%+v\n", nonceAccount)
	}
}
