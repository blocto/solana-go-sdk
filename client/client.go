package client

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Client struct {
	rpc.RpcClient
}

func NewClient(endpoint string) *Client {
	return &Client{rpc.NewRpcClient(endpoint)}
}

// GetBalance fetch users lamports(SOL) balance
func (c *Client) GetBalance(ctx context.Context, base58Addr string) (uint64, error) {
	res, err := c.RpcClient.GetBalance(ctx, base58Addr)
	if err != nil {
		return 0, err
	}
	return res.Result.Value, nil
}

// GetBalance fetch users lamports(SOL) balance with specific commitment
func (c *Client) GetBalanceWithCfg(ctx context.Context, base58Addr string, cfg rpc.GetBalanceConfig) (uint64, error) {
	res, err := c.RpcClient.GetBalanceWithCfg(ctx, base58Addr, cfg)
	if err != nil {
		return 0, err
	}
	return res.Result.Value, nil
}

type AccountInfo struct {
	Lamports  uint64
	Owner     string
	Excutable bool
	RentEpoch uint64
	Data      []byte
}

// GetAccountInfo return account's info
func (c *Client) GetAccountInfo(ctx context.Context, base58Addr string) (AccountInfo, error) {
	res, err := c.RpcClient.GetAccountInfoWithCfg(ctx, base58Addr, rpc.GetAccountInfoConfig{
		Encoding: rpc.GetAccountInfoConfigEncodingBase64,
	})
	if err != nil {
		return AccountInfo{}, err
	}

	if res.Error != nil {
		return AccountInfo{}, fmt.Errorf("%v", res.Error)
	}
	if res.Result.Value == (rpc.GetAccountInfoResultValue{}) {
		return AccountInfo{}, nil
	}

	data, ok := res.Result.Value.Data.([]interface{})
	if !ok {
		return AccountInfo{}, fmt.Errorf("failed to cast raw response to []interface{}")
	}
	if data[1] != string(rpc.GetAccountInfoConfigEncodingBase64) {
		return AccountInfo{}, fmt.Errorf("encoding mistmatch")
	}
	rawData, err := base64.StdEncoding.DecodeString(data[0].(string))
	if err != nil {
		return AccountInfo{}, fmt.Errorf("failed to base64 decode data")
	}
	return AccountInfo{
		Lamports:  res.Result.Value.Lamports,
		Owner:     res.Result.Value.Owner,
		Excutable: res.Result.Value.Excutable,
		RentEpoch: res.Result.Value.RentEpoch,
		Data:      rawData,
	}, nil
}

// GetRecentBlockhash return recent blockhash information
func (c *Client) GetRecentBlockhash(ctx context.Context) (rpc.GetRecentBlockHashResultValue, error) {
	res, err := c.RpcClient.GetRecentBlockhash(ctx)
	if err != nil {
		return rpc.GetRecentBlockHashResultValue{}, err
	}
	return res.Result.Value, nil
}

// SendRawTransaction will send your raw tx
func (c *Client) SendRawTransaction(ctx context.Context, tx []byte) (string, error) {
	res, err := c.RpcClient.SendTransactionWithConfig(
		ctx,
		base64.StdEncoding.EncodeToString(tx),
		rpc.SendTransactionConfig{
			Encoding: rpc.SendTransactionConfigEncodingBase64,
		},
	)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

type SendTransactionParam struct {
	Instructions []types.Instruction
	Signers      []types.Account
	FeePayer     common.PublicKey
}

// SendTransaction is a quick way to send tx
func (c *Client) SendTransaction(ctx context.Context, param SendTransactionParam) (string, error) {
	recentBlockhashRes, err := c.GetRecentBlockhash(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get recent blockhash, err: %v", err)
	}
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions:    param.Instructions,
		Signers:         param.Signers,
		FeePayer:        param.FeePayer,
		RecentBlockHash: recentBlockhashRes.Blockhash,
	})
	if err != nil {
		return "", fmt.Errorf("failed to build tx, err: %v", err)
	}
	res, err := c.RpcClient.SendTransactionWithConfig(
		ctx,
		base64.StdEncoding.EncodeToString(rawTx),
		rpc.SendTransactionConfig{Encoding: rpc.SendTransactionConfigEncodingBase64},
	)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}
