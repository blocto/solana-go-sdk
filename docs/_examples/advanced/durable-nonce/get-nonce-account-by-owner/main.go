package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	res, err := c.RpcClient.GetProgramAccountsWithConfig(
		context.Background(),
		common.SystemProgramID.ToBase58(),
		rpc.GetProgramAccountsConfig{
			Encoding: rpc.GetProgramAccountsConfigEncodingBase64,
			Filters: []rpc.GetProgramAccountsConfigFilter{
				{
					DataSize: sysprog.NonceAccountSize,
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
		data, err := base64.StdEncoding.DecodeString((a.Account.Data.([]interface{}))[0].(string))
		if err != nil {
			log.Fatalf("failed to decode data, err: %v", err)
		}
		nonceAccount, err := sysprog.NonceAccountDeserialize(data)
		if err != nil {
			log.Fatalf("failed to parse nonce account, err: %v", err)
		}
		fmt.Printf("%+v\n", nonceAccount)
	}
}
