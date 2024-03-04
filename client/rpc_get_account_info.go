package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/rpc"
)

type AccountInfo struct {
	Lamports   uint64
	Owner      common.PublicKey
	Executable bool
	RentEpoch  uint64
	Data       []byte
}

func convertAccountInfo(v rpc.AccountInfo) (AccountInfo, error) {
	if v == (rpc.AccountInfo{}) {
		return AccountInfo{}, nil
	}
	data, ok := v.Data.([]any)
	if !ok {
		return AccountInfo{}, fmt.Errorf("failed to cast raw response to []any")
	}
	if data[1] != string(rpc.AccountEncodingBase64) {
		return AccountInfo{}, fmt.Errorf("return value should be base64 encoded")
	}
	rawData, err := base64.StdEncoding.DecodeString(data[0].(string))
	if err != nil {
		return AccountInfo{}, fmt.Errorf("failed to base64 decode data")
	}
	return AccountInfo{
		Lamports:   v.Lamports,
		Owner:      common.PublicKeyFromString(v.Owner),
		Executable: v.Executable,
		RentEpoch:  v.RentEpoch,
		Data:       rawData,
	}, nil
}

func convertAccountInfoWithData(v rpc.AccountInfoWithData) (AccountInfo, error) {
	if v == (rpc.AccountInfoWithData{}) {
		return AccountInfo{}, nil
	}
	rawData, err := json.Marshal(v.Data)
	if err != nil {
		return AccountInfo{}, fmt.Errorf("failed to base64 decode data")
	}
	return AccountInfo{
		Lamports:   v.Lamports,
		Owner:      common.PublicKeyFromString(v.Owner),
		Executable: v.Executable,
		RentEpoch:  v.RentEpoch,
		Data:       rawData,
	}, nil
}

type GetAccountInfoConfig struct {
	Encoding   rpc.AccountEncoding
	Commitment rpc.Commitment
	DataSlice  *rpc.DataSlice
}

func (c GetAccountInfoConfig) toRpc() rpc.GetAccountInfoConfig {
	if c.Encoding == "" {
		c.Encoding = rpc.AccountEncodingBase64
	}
	return rpc.GetAccountInfoConfig{
		Encoding:   c.Encoding,
		Commitment: c.Commitment,
		DataSlice:  c.DataSlice,
	}
}

// GetAccountInfo return account's info
func (c *Client) GetAccountInfo(ctx context.Context, base58Addr string) (AccountInfo, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.AccountInfo]], error) {
			return c.RpcClient.GetAccountInfoWithConfig(ctx, base58Addr, GetAccountInfoConfig{}.toRpc())
		},
		convertGetAccountInfo,
	)
}

// GetAccountInfoWithConfig return account's info
func (c *Client) GetAccountInfoWithConfig(ctx context.Context, base58Addr string, cfg GetAccountInfoConfig) (AccountInfo, error) {
	if cfg.Encoding == rpc.AccountEncodingJsonParsed {
		return process(
			func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.AccountInfo]], error) {
				return c.RpcClient.GetAccountInfoWithConfig(ctx, base58Addr, cfg.toRpc())
			},
			convertGetAccountInfoWithData,
		)
	}
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.AccountInfo]], error) {
			return c.RpcClient.GetAccountInfoWithConfig(ctx, base58Addr, cfg.toRpc())
		},
		convertGetAccountInfo,
	)
}

// GetAccountInfoAndContext return account's info
func (c *Client) GetAccountInfoAndContext(ctx context.Context, base58Addr string) (rpc.ValueWithContext[AccountInfo], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.AccountInfo]], error) {
			return c.RpcClient.GetAccountInfoWithConfig(ctx, base58Addr, GetAccountInfoConfig{Encoding: rpc.AccountEncodingBase64}.toRpc())
		},
		convertGetAccountInfoAndContext,
	)
}

// GetAccountInfoAndContextWithConfig return account's info
func (c *Client) GetAccountInfoAndContextWithConfig(ctx context.Context, base58Addr string, cfg GetAccountInfoConfig) (rpc.ValueWithContext[AccountInfo], error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.AccountInfo]], error) {
			return c.RpcClient.GetAccountInfoWithConfig(ctx, base58Addr, cfg.toRpc())
		},
		convertGetAccountInfoAndContext,
	)
}

func convertGetAccountInfo(v rpc.ValueWithContext[rpc.AccountInfo]) (AccountInfo, error) {
	return convertAccountInfo(v.Value)
}

func convertGetAccountInfoWithData(v rpc.ValueWithContext[rpc.AccountInfoWithData]) (AccountInfo, error) {
	return convertAccountInfoWithData(v.Value)
}

func convertGetAccountInfoAndContext(v rpc.ValueWithContext[rpc.AccountInfo]) (rpc.ValueWithContext[AccountInfo], error) {
	accountInfo, err := convertGetAccountInfo(v)
	if err != nil {
		return rpc.ValueWithContext[AccountInfo]{}, err
	}
	return rpc.ValueWithContext[AccountInfo]{
		Context: v.Context,
		Value:   accountInfo,
	}, nil
}
