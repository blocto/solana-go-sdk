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
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfo(
					context.Background(),
					"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetAccountInfoResult{
					Context: Context{
						Slot: 77317716,
					},
					Value: GetAccountInfoResultValue{
						Lamports:  21474700400,
						Owner:     "11111111111111111111111111111111",
						Excutable: false,
						RentEpoch: 178,
						Data:      "",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["FaTGhPTgKeZZzQwLenoxn2VZXPWV1FpjQ1AQe77JUeJw"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77382573},"value":null},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfo(
					context.Background(),
					"FaTGhPTgKeZZzQwLenoxn2VZXPWV1FpjQ1AQe77JUeJw",
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetAccountInfoResult{
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
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfo(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetAccountInfoResult{
					Context: Context{
						Slot: 77317716,
					},
					Value: GetAccountInfoResultValue{
						Lamports:  1461600,
						Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Excutable: false,
						RentEpoch: 178,
						Data:      "DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D"]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32600,"message":"Encoded binary (base 58) data should be less than 128 bytes, please use Base64 encoding."},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfo(
					context.Background(),
					"9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D",
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error: &ErrorResponse{
						Code:    -32600,
						Message: "Encoded binary (base 58) data should be less than 128 bytes, please use Base64 encoding.",
					},
				},
				Result: GetAccountInfoResult{},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"commitment": "finalized"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317716},"value":{"data":"DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj","executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfoWithCfg(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						Commitment: CommitmentFinalized,
					},
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetAccountInfoResult{
					Context: Context{
						Slot: 77317716,
					},
					Value: GetAccountInfoResultValue{
						Lamports:  1461600,
						Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Excutable: false,
						RentEpoch: 178,
						Data:      "DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317717},"value":{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfoWithCfg(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						Encoding: GetAccountInfoConfigEncodingBase64,
					},
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetAccountInfoResult{
					Context: Context{
						Slot: 77317717,
					},
					Value: GetAccountInfoResultValue{
						Lamports:  1461600,
						Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Excutable: false,
						RentEpoch: 178,
						Data:      []interface{}{"AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", "base64"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64+zstd"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317717},"value":{"data":["KLUv/QBYjQEAhAIBAAAABj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4ACQEAAgAAGXXBEw==","base64+zstd"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfoWithCfg(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						Encoding: GetAccountInfoConfigEncodingBase64Zstd,
					},
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetAccountInfoResult{
					Context: Context{
						Slot: 77317717,
					},
					Value: GetAccountInfoResultValue{
						Lamports:  1461600,
						Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Excutable: false,
						RentEpoch: 178,
						Data:      []interface{}{"KLUv/QBYjQEAhAIBAAAABj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4ACQEAAgAAGXXBEw==", "base64+zstd"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"dataSlice": {"length": 32}}]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32602,"message":"Invalid params: missing field` + "`offset`" + `."},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfoWithCfg(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						DataSlice: &GetAccountInfoConfigDataSlice{
							Length: 32,
						},
					},
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error: &ErrorResponse{
						Code:    -32602,
						Message: `Invalid params: missing field` + "`offset`" + `.`,
					},
				},
				Result: GetAccountInfoResult{},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"dataSlice": {"offset": 4}}]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32602,"message":"Invalid params: missing field` + "`length`" + `."},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfoWithCfg(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						DataSlice: &GetAccountInfoConfigDataSlice{
							Offset: 4,
						},
					},
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error: &ErrorResponse{
						Code:    -32602,
						Message: `Invalid params: missing field` + "`length`" + `.`,
					},
				},
				Result: GetAccountInfoResult{},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"dataSlice": {"offset": 4, "length": 32}}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77322439},"value":{"data":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfoWithCfg(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						DataSlice: &GetAccountInfoConfigDataSlice{
							Offset: 4,
							Length: 32,
						},
					},
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetAccountInfoResult{
					Context: Context{
						Slot: 77322439,
					},
					Value: GetAccountInfoResultValue{
						Lamports:  1461600,
						Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Excutable: false,
						RentEpoch: 178,
						Data:      "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317718},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetAccountInfoWithCfg(
					context.Background(),
					"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					GetAccountInfoConfig{
						Encoding: GetAccountInfoConfigEncodingBase64,
						DataSlice: &GetAccountInfoConfigDataSlice{
							Offset: 4,
							Length: 32,
						},
					},
				)
			},
			ExpectedResponse: GetAccountInfoResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetAccountInfoResult{
					Context: Context{
						Slot: 77317718,
					},
					Value: GetAccountInfoResultValue{
						Lamports:  1461600,
						Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						Excutable: false,
						RentEpoch: 178,
						Data:      []interface{}{"Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=", "base64"},
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
