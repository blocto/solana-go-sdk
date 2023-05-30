package client

import (
	"context"

	"github.com/portto/solana-go-sdk/program/token"
	"github.com/portto/solana-go-sdk/rpc"
)

func (c *Client) GetTokenAccountsByOwnerByMint(ctx context.Context, owner, mintAddr string) ([]token.TokenAccount, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetProgramAccounts]], error) {
			return c.RpcClient.GetTokenAccountsByOwnerWithConfig(
				ctx,
				owner,
				rpc.GetTokenAccountsByOwnerConfigFilter{
					Mint: mintAddr,
				},
				rpc.GetTokenAccountsByOwnerConfig{
					Encoding: rpc.AccountEncodingBase64,
				},
			)
		},
		convertGetTokenAccountsByOwner,
	)
}

func (c *Client) GetTokenAccountsByOwnerByProgram(ctx context.Context, owner, programId string) ([]token.TokenAccount, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetProgramAccounts]], error) {
			return c.RpcClient.GetTokenAccountsByOwnerWithConfig(
				ctx,
				owner,
				rpc.GetTokenAccountsByOwnerConfigFilter{
					ProgramId: programId,
				},
				rpc.GetTokenAccountsByOwnerConfig{
					Encoding: rpc.AccountEncodingBase64,
				},
			)
		},
		convertGetTokenAccountsByOwner,
	)
}

func (c *Client) GetTokenAccountsByOwnerWithContextByMint(ctx context.Context, owner, mintAddr string) (rpc.ValueWithContext[[]token.TokenAccount], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetProgramAccounts]], error) {
			return c.RpcClient.GetTokenAccountsByOwnerWithConfig(
				ctx,
				owner,
				rpc.GetTokenAccountsByOwnerConfigFilter{
					Mint: mintAddr,
				},
				rpc.GetTokenAccountsByOwnerConfig{
					Encoding: rpc.AccountEncodingBase64,
				},
			)
		},
		convertGetTokenAccountsByOwnerAndContext,
	)
}

func (c *Client) GetTokenAccountsByOwnerWithContextByProgram(ctx context.Context, owner, programId string) (rpc.ValueWithContext[[]token.TokenAccount], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.GetProgramAccounts]], error) {
			return c.RpcClient.GetTokenAccountsByOwnerWithConfig(
				ctx,
				owner,
				rpc.GetTokenAccountsByOwnerConfigFilter{
					ProgramId: programId,
				},
				rpc.GetTokenAccountsByOwnerConfig{
					Encoding: rpc.AccountEncodingBase64,
				},
			)
		},
		convertGetTokenAccountsByOwnerAndContext,
	)
}

func convertGetTokenAccountsByOwner(v rpc.ValueWithContext[rpc.GetProgramAccounts]) ([]token.TokenAccount, error) {
	tokenAccounts := make([]token.TokenAccount, 0, len(v.Value))
	for _, v := range v.Value {
		accountInfo, err := convertAccountInfo(v.Account)
		if err != nil {
			return nil, err
		}
		tokenAccount, err := token.DeserializeTokenAccount(accountInfo.Data, accountInfo.Owner)
		if err != nil {
			return nil, err
		}
		tokenAccounts = append(tokenAccounts, tokenAccount)
	}
	return tokenAccounts, nil
}

func convertGetTokenAccountsByOwnerAndContext(v rpc.ValueWithContext[rpc.GetProgramAccounts]) (rpc.ValueWithContext[[]token.TokenAccount], error) {
	tokenAccounts, err := convertGetTokenAccountsByOwner(v)
	if err != nil {
		return rpc.ValueWithContext[[]token.TokenAccount]{}, err
	}
	return rpc.ValueWithContext[[]token.TokenAccount]{
		Context: v.Context,
		Value:   tokenAccounts,
	}, nil
}
