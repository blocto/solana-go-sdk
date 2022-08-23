package client

import (
	"context"

	"github.com/portto/solana-go-sdk/program/token"
)

func (c *Client) GetTokenAccount(ctx context.Context, base58Addr string) (token.TokenAccount, error) {
	accountInfo, err := c.GetAccountInfo(ctx, base58Addr)
	if err != nil {
		return token.TokenAccount{}, err
	}
	return token.DeserializeTokenAccount(accountInfo.Data, accountInfo.Owner)
}
