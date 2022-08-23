package rpc

import (
	"context"
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317716},"value":{"data":"","executable":false,"lamports":21474700400,"owner":"11111111111111111111111111111111","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfo(
					context.Background(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot: 77317716,
					},
					Value: AccountInfo{
						Lamports:   21474700400,
						Owner:      "11111111111111111111111111111111",
						Executable: false,
						RentEpoch:  178,
						Data:       "",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["FaTGhPTgKeZZzQwLenoxn2VZXPWV1FpjQ1AQe77JUeJw"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77382573},"value":null},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfo(
					context.Background(),
					"FaTGhPTgKeZZzQwLenoxn2VZXPWV1FpjQ1AQe77JUeJw",
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot: 77382573,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317716},"value":{"data":"DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj","executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfo(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot: 77317716,
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Executable: false,
						RentEpoch:  178,
						Data:       "DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D"]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32600,"message":"Encoded binary (base 58) data should be less than 128 bytes, please use Base64 encoding."},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfo(
					context.Background(),
					"9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D",
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error: &JsonRpcError{
					Code:    -32600,
					Message: "Encoded binary (base 58) data should be less than 128 bytes, please use Base64 encoding.",
				},
				Result: GetAccountInfo{},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"commitment": "finalized"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317716},"value":{"data":"DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj","executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfoWithConfig(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						Commitment: CommitmentFinalized,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot: 77317716,
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Executable: false,
						RentEpoch:  178,
						Data:       "DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317717},"value":{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfoWithConfig(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						Encoding: AccountEncodingBase64,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot: 77317717,
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Executable: false,
						RentEpoch:  178,
						Data:       []any{"AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", "base64"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64+zstd"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317717},"value":{"data":["KLUv/QBYjQEAhAIBAAAABj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4ACQEAAgAAGXXBEw==","base64+zstd"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfoWithConfig(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						Encoding: AccountEncodingBase64Zstd,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot: 77317717,
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Executable: false,
						RentEpoch:  178,
						Data:       []any{"KLUv/QBYjQEAhAIBAAAABj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4ACQEAAgAAGXXBEw==", "base64+zstd"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"dataSlice": {"length": 32}}]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32602,"message":"Invalid params: missing field` + "`offset`" + `."},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfoWithConfig(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						DataSlice: &DataSlice{
							Length: 32,
						},
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error: &JsonRpcError{
					Code:    -32602,
					Message: `Invalid params: missing field` + "`offset`" + `.`,
				},
				Result: GetAccountInfo{},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"dataSlice": {"offset": 4}}]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32602,"message":"Invalid params: missing field` + "`length`" + `."},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfoWithConfig(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						DataSlice: &DataSlice{
							Offset: 4,
						},
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error: &JsonRpcError{
					Code:    -32602,
					Message: `Invalid params: missing field` + "`length`" + `.`,
				},
				Result: GetAccountInfo{},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"dataSlice": {"offset": 4, "length": 32}}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77322439},"value":{"data":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfoWithConfig(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						DataSlice: &DataSlice{
							Offset: 4,
							Length: 32,
						},
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot: 77322439,
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Executable: false,
						RentEpoch:  178,
						Data:       "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317718},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfoWithConfig(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						Encoding: AccountEncodingBase64,
						DataSlice: &DataSlice{
							Offset: 4,
							Length: 32,
						},
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot: 77317718,
					},
					Value: AccountInfo{
						Lamports:   1461600,
						Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Executable: false,
						RentEpoch:  178,
						Data:       []any{"Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=", "base64"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["5xtKiHGFfhK6ynJwWrApoVkVHeTJ25czqnezDwJiT86N", {"encoding": "base64+zstd"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":121172974},"value":{"data":["KLUv/QBYIQEAAgAAADLAWbd6nbiA1th1E/yH93WWg4hgfgn3t2BoTE9m2JGH","base64+zstd"],"executable":true,"lamports":1141440,"owner":"BPFLoaderUpgradeab1e11111111111111111111111","rentEpoch":280}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfoWithConfig(
					context.Background(),
					"5xtKiHGFfhK6ynJwWrApoVkVHeTJ25czqnezDwJiT86N",
					GetAccountInfoConfig{
						Encoding: AccountEncodingBase64Zstd,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot: 121172974,
					},
					Value: AccountInfo{
						Lamports:   1141440,
						Owner:      "BPFLoaderUpgradeab1e11111111111111111111111",
						Executable: true,
						RentEpoch:  280,
						Data:       []any{"KLUv/QBYIQEAAgAAADLAWbd6nbiA1th1E/yH93WWg4hgfgn3t2BoTE9m2JGH", "base64+zstd"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.10.34","slot":155401599},"value":{"data":"","executable":false,"lamports":114638463277,"owner":"11111111111111111111111111111111","rentEpoch":359}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetAccountInfo(
					context.Background(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
				)
			},
			ExpectedResponse: JsonRpcResponse[GetAccountInfo]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetAccountInfo{
					Context: Context{
						Slot:       155401599,
						ApiVersion: "1.10.34",
					},
					Value: AccountInfo{
						Lamports:   114638463277,
						Owner:      "11111111111111111111111111111111",
						Executable: false,
						RentEpoch:  359,
						Data:       "",
					},
				},
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			testRpcCall(t, tt)
		})
	}
}
