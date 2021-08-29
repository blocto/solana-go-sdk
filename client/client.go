package client

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/portto/solana-go-sdk/client/rpc"
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
