package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetTokenAccountsByOwner(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountsByOwner", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"mint": "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":88024144},"value":[{"account":{"data":"error: data too large for bs58 encoding","executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":203},"pubkey":"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTokenAccountsByOwner(
						context.TODO(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						GetTokenAccountsByOwnerConfigFilter{
							Mint: "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetProgramAccounts]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetProgramAccounts]{
						Context: Context{
							Slot: 88024144,
						},
						Value: GetProgramAccounts{
							{
								Pubkey: "AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
								Account: AccountInfo{
									Lamports:   2039280,
									Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									RentEpoch:  203,
									Data:       "error: data too large for bs58 encoding",
									Executable: false,
								},
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountsByOwner", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"programId": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":88024144},"value":[{"account":{"data":"error: data too large for bs58 encoding","executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":203},"pubkey":"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTokenAccountsByOwner(
						context.TODO(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						GetTokenAccountsByOwnerConfigFilter{
							ProgramId: common.TokenProgramID.ToBase58(),
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetProgramAccounts]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetProgramAccounts]{
						Context: Context{
							Slot: 88024144,
						},
						Value: GetProgramAccounts{
							{
								Pubkey: "AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
								Account: AccountInfo{
									Lamports:   2039280,
									Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									RentEpoch:  203,
									Data:       "error: data too large for bs58 encoding",
									Executable: false,
								},
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountsByOwner", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"mint": "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"}, {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":88024145},"value":[{"account":{"data":["M72Y4VtywPCapPDIhmN7Y+l309jqFamd0HPBVhiGx5AQllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3ADkC1QCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":203},"pubkey":"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTokenAccountsByOwnerWithConfig(
						context.TODO(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						GetTokenAccountsByOwnerConfigFilter{
							Mint: "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
						},
						GetTokenAccountsByOwnerConfig{
							Encoding: AccountEncodingBase64,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetProgramAccounts]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetProgramAccounts]{
						Context: Context{
							Slot: 88024145,
						},
						Value: GetProgramAccounts{
							{
								Pubkey: "AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
								Account: AccountInfo{
									Lamports:  2039280,
									Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									RentEpoch: 203,
									Data: []any{
										"M72Y4VtywPCapPDIhmN7Y+l309jqFamd0HPBVhiGx5AQllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3ADkC1QCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
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
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountsByOwner", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"mint": "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"}, {"encoding":"base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":88024145},"value":[{"account":{"data":["M72Y4VtywPCapPDIhmN7Y+l309jqFamd0HPBVhiGx5AQllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3ADkC1QCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":203},"pubkey":"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTokenAccountsByOwnerWithConfig(
						context.TODO(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						GetTokenAccountsByOwnerConfigFilter{
							Mint: "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
						},
						GetTokenAccountsByOwnerConfig{
							Encoding: AccountEncodingBase64,
							DataSlice: &DataSlice{
								Offset: 4,
								Length: 32,
							},
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetProgramAccounts]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetProgramAccounts]{
						Context: Context{
							Slot: 88024145,
						},
						Value: GetProgramAccounts{
							{
								Pubkey: "AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
								Account: AccountInfo{
									Lamports:  2039280,
									Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									RentEpoch: 203,
									Data: []any{
										"M72Y4VtywPCapPDIhmN7Y+l309jqFamd0HPBVhiGx5AQllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3ADkC1QCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
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
		},
	)
}
