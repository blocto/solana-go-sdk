package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetMultipleAccounts(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317716},"value":[{"data":"","executable":false,"lamports":21474700400,"owner":"11111111111111111111111111111111","rentEpoch":178}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMultipleAccounts(
						context.Background(),
						[]string{"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[[]AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[[]AccountInfo]{
						Context: Context{
							Slot: 77317716,
						},
						Value: []AccountInfo{
							{
								Lamports:   21474700400,
								Owner:      "11111111111111111111111111111111",
								Executable: false,
								RentEpoch:  178,
								Data:       "",
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["FaTGhPTgKeZZzQwLenoxn2VZXPWV1FpjQ1AQe77JUeJw"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77382573},"value":[null]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMultipleAccounts(
						context.Background(),
						[]string{"FaTGhPTgKeZZzQwLenoxn2VZXPWV1FpjQ1AQe77JUeJw"},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[[]AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[[]AccountInfo]{
						Context: Context{
							Slot: 77382573,
						},
						Value: []AccountInfo{
							{},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317716},"value":[{"data":"DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj","executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMultipleAccounts(
						context.Background(),
						[]string{"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[[]AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[[]AccountInfo]{
						Context: Context{
							Slot: 77317716,
						},
						Value: []AccountInfo{
							{
								Lamports:   1461600,
								Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								Executable: false,
								RentEpoch:  178,
								Data:       "DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj",
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32600,"message":"Encoded binary (base 58) data should be less than 128 bytes, please use Base64 encoding."},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMultipleAccounts(
						context.Background(),
						[]string{"9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D"},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[[]AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error: &JsonRpcError{
						Code:    -32600,
						Message: "Encoded binary (base 58) data should be less than 128 bytes, please use Base64 encoding.",
					},
					Result: ValueWithContext[[]AccountInfo]{},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"commitment": "finalized"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317716},"value":[{"data":"DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj","executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMultipleAccountsWithConfig(
						context.Background(),
						[]string{"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"},
						GetMultipleAccountsConfig{
							Commitment: CommitmentFinalized,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[[]AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[[]AccountInfo]{
						Context: Context{
							Slot: 77317716,
						},
						Value: []AccountInfo{
							{
								Lamports:   1461600,
								Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								Executable: false,
								RentEpoch:  178,
								Data:       "DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj",
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317717},"value":[{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMultipleAccountsWithConfig(
						context.Background(),
						[]string{"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"},
						GetMultipleAccountsConfig{
							Encoding: AccountEncodingBase64,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[[]AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[[]AccountInfo]{
						Context: Context{
							Slot: 77317717,
						},
						Value: []AccountInfo{
							{
								Lamports:   1461600,
								Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								Executable: false,
								RentEpoch:  178,
								Data:       []any{"AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", "base64"},
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"encoding": "base64+zstd"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317717},"value":[{"data":["KLUv/QBYjQEAhAIBAAAABj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4ACQEAAgAAGXXBEw==","base64+zstd"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMultipleAccountsWithConfig(
						context.Background(),
						[]string{"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"},
						GetMultipleAccountsConfig{
							Encoding: AccountEncodingBase64Zstd,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[[]AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[[]AccountInfo]{
						Context: Context{
							Slot: 77317717,
						},
						Value: []AccountInfo{
							{
								Lamports:   1461600,
								Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								Executable: false,
								RentEpoch:  178,
								Data:       []any{"KLUv/QBYjQEAhAIBAAAABj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4ACQEAAgAAGXXBEw==", "base64+zstd"},
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"dataSlice": {"offset": 4, "length": 32}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77322439},"value":[{"data":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMultipleAccountsWithConfig(
						context.Background(),
						[]string{"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"},
						GetMultipleAccountsConfig{
							DataSlice: &DataSlice{
								Offset: 4,
								Length: 32,
							},
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[[]AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[[]AccountInfo]{
						Context: Context{
							Slot: 77322439,
						},
						Value: []AccountInfo{
							{
								Lamports:   1461600,
								Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								Executable: false,
								RentEpoch:  178,
								Data:       "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMultipleAccounts", "params":[["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"], {"encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317718},"value":[{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetMultipleAccountsWithConfig(
						context.Background(),
						[]string{"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"},
						GetMultipleAccountsConfig{
							Encoding: AccountEncodingBase64,
							DataSlice: &DataSlice{
								Offset: 4,
								Length: 32,
							},
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[[]AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[[]AccountInfo]{
						Context: Context{
							Slot: 77317718,
						},
						Value: []AccountInfo{
							{
								Lamports:   1461600,
								Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								Executable: false,
								RentEpoch:  178,
								Data:       []any{"Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=", "base64"},
							},
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
