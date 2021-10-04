package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

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
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result.Value, nil
}

// GetBalance fetch users lamports(SOL) balance with specific commitment
func (c *Client) GetBalanceWithCfg(ctx context.Context, base58Addr string, cfg rpc.GetBalanceConfig) (uint64, error) {
	res, err := c.RpcClient.GetBalanceWithCfg(ctx, base58Addr, cfg)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result.Value, nil
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *Client) GetTokenAccountBalance(ctx context.Context, base58Addr string) (uint64, uint8, error) {
	res, err := c.RpcClient.GetTokenAccountBalance(ctx, base58Addr)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, 0, err
	}
	balance, err := strconv.ParseUint(res.Result.Value.Amount, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to cast token amount, err: %v", err)
	}
	return balance, res.Result.Value.Decimals, nil
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *Client) GetTokenAccountBalanceWithCfg(ctx context.Context, base58Addr string, cfg rpc.GetTokenAccountBalanceConfig) (uint64, uint8, error) {
	res, err := c.RpcClient.GetTokenAccountBalanceWithCfg(ctx, base58Addr, cfg)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, 0, err
	}
	balance, err := strconv.ParseUint(res.Result.Value.Amount, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to cast token amount, err: %v", err)
	}
	return balance, res.Result.Value.Decimals, nil
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
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return AccountInfo{}, err
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
	err = checkRpcResult(res.GeneralResponse, err)
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
	err = checkRpcResult(res.GeneralResponse, err)
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
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// SendTransaction2 send transaction struct directly
func (c *Client) SendTransaction2(ctx context.Context, tx types.Transaction) (string, error) {
	rawTx, err := tx.Serialize()
	if err != nil {
		return "", fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	res, err := c.RpcClient.SendTransactionWithConfig(
		ctx,
		base64.StdEncoding.EncodeToString(rawTx),
		rpc.SendTransactionConfig{
			Encoding: rpc.SendTransactionConfigEncodingBase64,
		},
	)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// GetSlot get current slot (finalized)
func (c *Client) GetSlot(ctx context.Context) (uint64, error) {
	res, err := c.RpcClient.GetSlot(ctx)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

// GetSlotWithCfg get slot by commitment
func (c *Client) GetSlotWithCfg(ctx context.Context, cfg rpc.GetSlotConfig) (uint64, error) {
	res, err := c.RpcClient.GetSlotWithCfg(ctx, cfg)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

func checkRpcResult(res rpc.GeneralResponse, err error) error {
	if err != nil {
		return err
	}
	if res.Error != nil {
		errRes, err := json.Marshal(res.Error)
		if err != nil {
			return fmt.Errorf("rpc response error: %v", res.Error)
		}
		return fmt.Errorf("rpc response error: %v", string(errRes))
	}
	return nil
}
