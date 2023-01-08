package client

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/rpc"
)

func TestClient_GetAccountInfo(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187539624},"value":{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfo(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					)
				},
				ExpectedValue: AccountInfo{
					Lamports:   1461600,
					Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
					Executable: false,
					RentEpoch:  371,
					Data:       []byte{1, 0, 0, 0, 6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetAccountInfoWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				Name:         "empty",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187539624},"value":{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfoWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetAccountInfoConfig{},
					)
				},
				ExpectedValue: AccountInfo{
					Lamports:   1461600,
					Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
					Executable: false,
					RentEpoch:  371,
					Data:       []byte{1, 0, 0, 0, 6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
				ExpectedError: nil,
			},
			{
				Name:         "with commitment",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"commitment": "processed", "encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187537245},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfoWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetAccountInfoConfig{
							Commitment: rpc.CommitmentProcessed,
						},
					)
				},
				ExpectedValue: AccountInfo{
					Lamports:   1461600,
					Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
					Executable: false,
					RentEpoch:  371,
					Data:       []byte{6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222},
				},
				ExpectedError: nil,
			},
			{
				Name:         "with data slice",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187537245},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfoWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetAccountInfoConfig{
							DataSlice: &rpc.DataSlice{
								Offset: 4,
								Length: 32,
							},
						},
					)
				},
				ExpectedValue: AccountInfo{
					Lamports:   1461600,
					Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
					Executable: false,
					RentEpoch:  371,
					Data:       []byte{6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222},
				},
				ExpectedError: nil,
			},
			{
				Name:         "with all",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"commitment": "confirmed", "encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187537245},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfoWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetAccountInfoConfig{
							Commitment: rpc.CommitmentConfirmed,
							DataSlice: &rpc.DataSlice{
								Offset: 4,
								Length: 32,
							},
						},
					)
				},
				ExpectedValue: AccountInfo{
					Lamports:   1461600,
					Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
					Executable: false,
					RentEpoch:  371,
					Data:       []byte{6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetAccountInfoAndContext(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187537245},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfoAndContext(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					)
				},
				ExpectedValue: rpc.ValueWithContext[AccountInfo]{
					Context: rpc.Context{
						Slot:       187537245,
						ApiVersion: "1.14.10",
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						Executable: false,
						RentEpoch:  371,
						Data:       []byte{6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetAccountInfoAndContextWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				Name:         "empty",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187539624},"value":{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfoAndContextWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetAccountInfoConfig{},
					)
				},
				ExpectedValue: rpc.ValueWithContext[AccountInfo]{
					Context: rpc.Context{
						Slot:       187539624,
						ApiVersion: "1.14.10",
					},
					Value: AccountInfo{
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
				Name:         "with commitment",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"commitment": "processed", "encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187537245},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfoAndContextWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetAccountInfoConfig{
							Commitment: rpc.CommitmentProcessed,
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[AccountInfo]{
					Context: rpc.Context{
						Slot:       187537245,
						ApiVersion: "1.14.10",
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						Executable: false,
						RentEpoch:  371,
						Data:       []byte{6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "with data slice",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187537245},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfoAndContextWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetAccountInfoConfig{
							DataSlice: &rpc.DataSlice{
								Offset: 4,
								Length: 32,
							},
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[AccountInfo]{
					Context: rpc.Context{
						Slot:       187537245,
						ApiVersion: "1.14.10",
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						Executable: false,
						RentEpoch:  371,
						Data:       []byte{6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "with all",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"commitment": "confirmed", "encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187537245},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetAccountInfoAndContextWithConfig(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
						GetAccountInfoConfig{
							Commitment: rpc.CommitmentConfirmed,
							DataSlice: &rpc.DataSlice{
								Offset: 4,
								Length: 32,
							},
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[AccountInfo]{
					Context: rpc.Context{
						Slot:       187537245,
						ApiVersion: "1.14.10",
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						Executable: false,
						RentEpoch:  371,
						Data:       []byte{6, 62, 112, 217, 145, 87, 152, 220, 163, 238, 219, 145, 55, 207, 100, 199, 15, 250, 251, 98, 205, 211, 130, 129, 5, 37, 11, 215, 107, 29, 108, 222},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
