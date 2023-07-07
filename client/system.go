package client

import (
	"context"
	"errors"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/mr-tron/base58"
)

func (c *Client) GetNonceAccount(ctx context.Context, base58Addr string) (system.NonceAccount, error) {
	accu, err := c.GetAccountInfo(ctx, base58Addr)
	if err != nil {
		return system.NonceAccount{}, err
	}
	if accu.Owner != common.SystemProgramID {
		return system.NonceAccount{}, errors.New("owner mismatch")
	}
	return system.NonceAccountDeserialize(accu.Data)
}

func (c *Client) GetNonceFromNonceAccount(ctx context.Context, base58Addr string) (string, error) {
	accuInfo, err := c.GetAccountInfoWithConfig(ctx, base58Addr, GetAccountInfoConfig{
		DataSlice: &rpc.DataSlice{
			Offset: 40,
			Length: 32,
		},
	})
	if err != nil {
		return "", err
	}
	return base58.Encode(accuInfo.Data), nil
}
