package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/program/token"
	"github.com/blocto/solana-go-sdk/rpc"
)

func TestClient_GetTokenAccountsByOwnerByMint(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountsByOwner", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"mint": "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"}, {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.17","slot":219416878},"value":[{"account":{"data":["M72Y4VtywPCapPDIhmN7Y+l309jqFamd0HPBVhiGx5AQllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3AAacRgCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371},"pubkey":"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenAccountsByOwnerByMint(
						context.Background(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
					)
				},
				ExpectedValue: []TokenAccount{
					{
						TokenAccount: token.TokenAccount{
							Mint:            common.PublicKeyFromString("4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"),
							Owner:           common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ"),
							Amount:          9000000000,
							Delegate:        nil,
							State:           token.TokenAccountStateInitialized,
							IsNative:        nil,
							DelegatedAmount: 0,
							CloseAuthority:  nil,
						},
						PublicKey: common.PublicKeyFromString("AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"),
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTokenAccountsByOwnerByProgram(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountsByOwner", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"programId": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"}, {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.17","slot":219416878},"value":[{"account":{"data":["M72Y4VtywPCapPDIhmN7Y+l309jqFamd0HPBVhiGx5AQllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3AAacRgCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371},"pubkey":"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenAccountsByOwnerByProgram(
						context.Background(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						common.TokenProgramID.ToBase58(),
					)
				},
				ExpectedValue: []TokenAccount{
					{
						TokenAccount: token.TokenAccount{
							Mint:            common.PublicKeyFromString("4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"),
							Owner:           common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ"),
							Amount:          9000000000,
							Delegate:        nil,
							State:           token.TokenAccountStateInitialized,
							IsNative:        nil,
							DelegatedAmount: 0,
							CloseAuthority:  nil,
						},
						PublicKey: common.PublicKeyFromString("AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"),
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTokenAccountsByOwnerWithContextByMint(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountsByOwner", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"mint": "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"}, {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.17","slot":219416878},"value":[{"account":{"data":["M72Y4VtywPCapPDIhmN7Y+l309jqFamd0HPBVhiGx5AQllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3AAacRgCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371},"pubkey":"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenAccountsByOwnerWithContextByMint(
						context.Background(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
					)
				},
				ExpectedValue: rpc.ValueWithContext[[]TokenAccount]{
					Context: rpc.Context{
						ApiVersion: "1.14.17",
						Slot:       219416878,
					},
					Value: []TokenAccount{
						{
							TokenAccount: token.TokenAccount{
								Mint:            common.PublicKeyFromString("4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"),
								Owner:           common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ"),
								Amount:          9000000000,
								Delegate:        nil,
								State:           token.TokenAccountStateInitialized,
								IsNative:        nil,
								DelegatedAmount: 0,
								CloseAuthority:  nil,
							},
							PublicKey: common.PublicKeyFromString("AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"),
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetTokenAccountsByOwnerWithContextByProgram(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTokenAccountsByOwner", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"programId": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"}, {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.17","slot":219416878},"value":[{"account":{"data":["M72Y4VtywPCapPDIhmN7Y+l309jqFamd0HPBVhiGx5AQllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3AAacRgCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371},"pubkey":"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTokenAccountsByOwnerWithContextByProgram(
						context.Background(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						common.TokenProgramID.ToBase58(),
					)
				},
				ExpectedValue: rpc.ValueWithContext[[]TokenAccount]{
					Context: rpc.Context{
						ApiVersion: "1.14.17",
						Slot:       219416878,
					},
					Value: []TokenAccount{
						{
							TokenAccount: token.TokenAccount{
								Mint:            common.PublicKeyFromString("4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"),
								Owner:           common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ"),
								Amount:          9000000000,
								Delegate:        nil,
								State:           token.TokenAccountStateInitialized,
								IsNative:        nil,
								DelegatedAmount: 0,
								CloseAuthority:  nil,
							},
							PublicKey: common.PublicKeyFromString("AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"),
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
