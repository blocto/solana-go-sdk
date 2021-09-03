package rpc

import (
	"context"
	"testing"
)

func TestGetProgramAccounts(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":[{"account":{"data":["0SWx++406Gemp6iqPyXxkCKUGHf4+wT/Ycjdf33fscwGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D"},{"account":{"data":["7OZLxMQO3fQ8HpH8SQe/BiLwnYeYtgyHcEZCLNFny4IGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"4MdGG2EnnAp4dhbMCDww1qLEEiW5SHfUAe9U9RtTqS8q"}],"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetProgramAccounts(
					context.TODO(),
					"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
				)
			},
			ExpectedResponse: GetProgramAccountsResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: []GetProgramAccounts{
					{
						Pubkey: "9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D",
						Account: GetProgramAccountsAccount{
							Lamports:  2039280,
							Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
							RentEpoch: 181,
							Data: []interface{}{
								"0SWx++406Gemp6iqPyXxkCKUGHf4+wT/Ycjdf33fscwGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
								"base64",
							},
							Executable: false,
						},
					},
					{
						Pubkey: "4MdGG2EnnAp4dhbMCDww1qLEEiW5SHfUAe9U9RtTqS8q",
						Account: GetProgramAccountsAccount{
							Lamports:  2039280,
							Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
							RentEpoch: 181,
							Data: []interface{}{
								"7OZLxMQO3fQ8HpH8SQe/BiLwnYeYtgyHcEZCLNFny4IGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
								"base64",
							},
							Executable: false,
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", {"encoding": "base64+zstd", "filters":[{"memcmp": {"offset": 0, "bytes": "GwktiL5dEA4mCGeVdGyUPDeWKzGC486KkPVTbwwnMGYq"}}]}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":[{"account":{"data":["KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=","base64+zstd"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6"}],"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetProgramAccountsWithConfig(
					context.TODO(),
					"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
					GetProgramAccountsConfig{
						Encoding: GetProgramAccountsConfigEncodingBase64Zstd,
						Filters: []GetProgramAccountsConfigFilter{
							{
								MemCmp: &GetProgramAccountsConfigFilterMemCmp{
									Offset: 0,
									Bytes:  "GwktiL5dEA4mCGeVdGyUPDeWKzGC486KkPVTbwwnMGYq",
								},
							},
						},
					},
				)
			},
			ExpectedResponse: GetProgramAccountsResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: []GetProgramAccounts{
					{
						Pubkey: "Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6",
						Account: GetProgramAccountsAccount{
							Lamports:  2039280,
							Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
							RentEpoch: 181,
							Data: []interface{}{
								"KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=",
								"base64+zstd",
							},
							Executable: false,
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", {"encoding": "base64+zstd", "filters":[{"dataSize": 165}]}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":[{"account":{"data":["KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=","base64+zstd"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6"}],"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetProgramAccountsWithConfig(
					context.TODO(),
					"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
					GetProgramAccountsConfig{
						Encoding: GetProgramAccountsConfigEncodingBase64Zstd,
						Filters: []GetProgramAccountsConfigFilter{
							{
								DataSize: 165,
							},
						},
					},
				)
			},
			ExpectedResponse: GetProgramAccountsResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: []GetProgramAccounts{
					{
						Pubkey: "Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6",
						Account: GetProgramAccountsAccount{
							Lamports:  2039280,
							Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
							RentEpoch: 181,
							Data: []interface{}{
								"KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=",
								"base64+zstd",
							},
							Executable: false,
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", {"encoding": "base64+zstd", "filters":[{"memcmp": {"offset": 0, "bytes": "GwktiL5dEA4mCGeVdGyUPDeWKzGC486KkPVTbwwnMGYq"}}, {"dataSize": 165}]}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":[{"account":{"data":["KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=","base64+zstd"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6"}],"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetProgramAccountsWithConfig(
					context.TODO(),
					"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
					GetProgramAccountsConfig{
						Encoding: GetProgramAccountsConfigEncodingBase64Zstd,
						Filters: []GetProgramAccountsConfigFilter{
							{
								MemCmp: &GetProgramAccountsConfigFilterMemCmp{
									Offset: 0,
									Bytes:  "GwktiL5dEA4mCGeVdGyUPDeWKzGC486KkPVTbwwnMGYq",
								},
							},
							{
								DataSize: 165,
							},
						},
					},
				)
			},
			ExpectedResponse: GetProgramAccountsResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: []GetProgramAccounts{
					{
						Pubkey: "Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6",
						Account: GetProgramAccountsAccount{
							Lamports:  2039280,
							Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
							RentEpoch: 181,
							Data: []interface{}{
								"KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=",
								"base64+zstd",
							},
							Executable: false,
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", {"withContext": true}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":78496860},"value":[{"account":{"data":["0SWx++406Gemp6iqPyXxkCKUGHf4+wT/Ycjdf33fscwGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D"},{"account":{"data":["7OZLxMQO3fQ8HpH8SQe/BiLwnYeYtgyHcEZCLNFny4IGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"4MdGG2EnnAp4dhbMCDww1qLEEiW5SHfUAe9U9RtTqS8q"}]},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetProgramAccountsWithContext(
					context.TODO(),
					"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
				)
			},
			ExpectedResponse: GetProgramAccountsWithContextResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetProgramAccountsWithContextResult{
					Context: Context{
						Slot: 78496860,
					},
					Value: []GetProgramAccounts{
						{
							Pubkey: "9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D",
							Account: GetProgramAccountsAccount{
								Lamports:  2039280,
								Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								RentEpoch: 181,
								Data: []interface{}{
									"0SWx++406Gemp6iqPyXxkCKUGHf4+wT/Ycjdf33fscwGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
									"base64",
								},
								Executable: false,
							},
						},
						{
							Pubkey: "4MdGG2EnnAp4dhbMCDww1qLEEiW5SHfUAe9U9RtTqS8q",
							Account: GetProgramAccountsAccount{
								Lamports:  2039280,
								Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								RentEpoch: 181,
								Data: []interface{}{
									"7OZLxMQO3fQ8HpH8SQe/BiLwnYeYtgyHcEZCLNFny4IGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
									"base64",
								},
								Executable: false,
							},
						},
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
