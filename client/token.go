package client

import (
	"context"

	"github.com/portto/solana-go-sdk/program/tokenprog"
)

func (c *Client) GetTokenAccount(ctx context.Context, base58Addr string) (tokenprog.TokenAccount, error) {
	accountInfo, err := c.GetAccountInfo(ctx, base58Addr)
	if err != nil {
		return tokenprog.TokenAccount{}, err
	}
	return tokenprog.DeserializeTokenAccount(accountInfo.Data, accountInfo.Owner)
}
