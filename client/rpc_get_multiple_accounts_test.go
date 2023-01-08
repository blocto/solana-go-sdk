package client

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/rpc"
)

func TestClient_GetMultipleAccounts(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8", "F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187635130},"value":[null,{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetMultipleAccounts(
						context.Background(),
						[]string{
							"3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8",
							"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						},
					)
				},
				ExpectedValue: []AccountInfo{
					{},
					{
						Lamports:   1461600,
						Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						Executable: false,
						RentEpoch:  371,
						Data:       []byte{1, 0, 0, 0, 6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetMultipleAccountsWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				Name:         "with commitment",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8", "F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"commitment": "processed", "encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187635130},"value":[null,{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetMultipleAccountsWithConfig(
						context.Background(),
						[]string{
							"3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8",
							"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						},
						GetMultipleAccountsConfig{
							Commitment: rpc.CommitmentProcessed,
						},
					)
				},
				ExpectedValue: []AccountInfo{
					{},
					{
						Lamports:   1461600,
						Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						Executable: false,
						RentEpoch:  371,
						Data:       []byte{1, 0, 0, 0, 6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "with data slice",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8", "F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187635130},"value":[null,{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetMultipleAccountsWithConfig(
						context.Background(),
						[]string{
							"3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8",
							"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						},
						GetMultipleAccountsConfig{
							DataSlice: &rpc.DataSlice{
								Offset: 4,
								Length: 32,
							},
						},
					)
				},
				ExpectedValue: []AccountInfo{
					{},
					{
						Lamports:   1461600,
						Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						Executable: false,
						RentEpoch:  371,
						Data:       []byte{1, 0, 0, 0, 6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetMultipleAccountsAndContext(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8", "F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187635130},"value":[null,{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetMultipleAccountsAndContext(
						context.Background(),
						[]string{
							"3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8",
							"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[[]AccountInfo]{
					Context: rpc.Context{
						Slot:       187635130,
						ApiVersion: "1.14.10",
					},
					Value: []AccountInfo{
						{},
						{
							Lamports:   1461600,
							Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
							Executable: false,
							RentEpoch:  371,
							Data:       []byte{1, 0, 0, 0, 6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetMultipleAccountsAndContextWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				Name:         "with commitment",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8", "F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"commitment": "processed", "encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187635130},"value":[null,{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetMultipleAccountsAndContextWithConfig(
						context.Background(),
						[]string{
							"3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8",
							"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						},
						GetMultipleAccountsConfig{
							Commitment: rpc.CommitmentProcessed,
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[[]AccountInfo]{
					Context: rpc.Context{
						Slot:       187635130,
						ApiVersion: "1.14.10",
					},
					Value: []AccountInfo{
						{},
						{
							Lamports:   1461600,
							Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
							Executable: false,
							RentEpoch:  371,
							Data:       []byte{1, 0, 0, 0, 6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "with data slice",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8", "F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187635130},"value":[null,{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetMultipleAccountsAndContextWithConfig(
						context.Background(),
						[]string{
							"3Af3cmANDdDcDPNiwNzDYxbwkVN6r5CELok3JjnSxcq8",
							"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						},
						GetMultipleAccountsConfig{
							DataSlice: &rpc.DataSlice{
								Offset: 4,
								Length: 32,
							},
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[[]AccountInfo]{
					Context: rpc.Context{
						Slot:       187635130,
						ApiVersion: "1.14.10",
					},
					Value: []AccountInfo{
						{},
						{
							Lamports:   1461600,
							Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
							Executable: false,
							RentEpoch:  371,
							Data:       []byte{1, 0, 0, 0, 6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
