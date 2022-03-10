package client

import (
	"context"
	"errors"

	"github.com/mr-tron/base58"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/rpc"
)

func (c *Client) GetNonceAccount(ctx context.Context, base58Addr string) (sysprog.NonceAccount, error) {
	accu, err := c.GetAccountInfo(ctx, base58Addr)
	if err != nil {
		return sysprog.NonceAccount{}, err
	}
	if accu.Owner != common.SystemProgramID.ToBase58() {
		return sysprog.NonceAccount{}, errors.New("owner mismatch")
	}
	return sysprog.NonceAccountDeserialize(accu.Data)
}

func (c *Client) GetNonceFromNonceAccount(ctx context.Context, base58Addr string) (string, error) {
	accuInfo, err := c.GetAccountInfoWithConfig(ctx, base58Addr, GetAccountInfoConfig{
		DataSlice: &rpc.GetAccountInfoConfigDataSlice{
			Offset: 40,
			Length: 32,
		},
	})
	if err != nil {
		return "", err
	}
	return base58.Encode(accuInfo.Data), nil
}
