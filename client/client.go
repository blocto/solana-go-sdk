package client

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

type Client struct {
	RpcClient rpc.RpcClient
}

func New(opts ...rpc.Option) *Client {
	return &Client{
		RpcClient: rpc.New(opts...),
	}
}

func NewClient(endpoint string) *Client {
	return &Client{rpc.New(rpc.WithEndpoint(endpoint))}
}

type QuickSendTransactionParam struct {
	Instructions []types.Instruction
	Signers      []types.Account
	FeePayer     common.PublicKey
}

// Deprecated: please use sendTransaction
// QuickSendTransaction is a quick way to send tx
func (c *Client) QuickSendTransaction(ctx context.Context, param QuickSendTransactionParam) (string, error) {
	recentBlockhashRes, err := c.GetLatestBlockhash(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get recent blockhash, err: %v", err)
	}
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			Instructions:    param.Instructions,
			FeePayer:        param.FeePayer,
			RecentBlockhash: recentBlockhashRes.Blockhash,
		}),
		Signers: param.Signers,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create new tx, err: %v", err)
	}
	rawTx, err := tx.Serialize()
	if err != nil {
		return "", fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	res, err := c.RpcClient.SendTransactionWithConfig(
		ctx,
		base64.StdEncoding.EncodeToString(rawTx),
		rpc.SendTransactionConfig{Encoding: rpc.SendTransactionConfigEncodingBase64},
	)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

func (c *Client) GetSignatureStatus(ctx context.Context, signature string) (*rpc.SignatureStatus, error) {
	res, err := c.RpcClient.GetSignatureStatuses(ctx, []string{signature})
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	return res.Result.Value[0], nil
}

func (c *Client) GetSignatureStatusWithConfig(ctx context.Context, signature string, cfg rpc.GetSignatureStatusesConfig) (*rpc.SignatureStatus, error) {
	res, err := c.RpcClient.GetSignatureStatusesWithConfig(ctx, []string{signature}, cfg)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	return res.Result.Value[0], nil
}

func (c *Client) GetSignatureStatuses(ctx context.Context, signatures []string) (rpc.SignatureStatuses, error) {
	res, err := c.RpcClient.GetSignatureStatuses(ctx, signatures)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	return res.Result.Value, nil
}

func (c *Client) GetSignatureStatusesWithConfig(ctx context.Context, signatures []string, cfg rpc.GetSignatureStatusesConfig) (rpc.SignatureStatuses, error) {
	res, err := c.RpcClient.GetSignatureStatusesWithConfig(ctx, signatures, cfg)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	return res.Result.Value, nil
}

func checkJsonRpcResponse[T any](res rpc.JsonRpcResponse[T], err error) error {
	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// helper function
func convertToIntSlice(input []any) ([]int, error) {
	output := make([]int, 0, len(input))
	for _, v := range input {
		output = append(output, int(v.(float64)))
	}
	return output, nil
}
