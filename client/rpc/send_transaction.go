package rpc

import (
	"context"
	"encoding/json"
	"fmt"
)

type SendTransactionResponse struct {
	GeneralResponse
	Result string `json:"result"`
}

type SendTransactionConfigEncoding string

const (
	SendTransactionConfigEncodingBase58 SendTransactionConfigEncoding = "base58"
	SendTransactionConfigEncodingBase64 SendTransactionConfigEncoding = "base64"
)

type SendTransactionConfig struct {
	SkipPreflight       bool                          `json:"skipPreflight,omitempty"`       // default: false
	PreflightCommitment Commitment                    `json:"preflightCommitment,omitempty"` // default: finalized
	Encoding            SendTransactionConfigEncoding `json:"encoding,omitempty"`            // default: base58
	MaxRetries          uint64                        `json:"maxRetries,omitempty"`
}

// SendTransaction submits a signed transaction to the cluster for processing
func (c *RpcClient) SendTransaction(ctx context.Context, tx string) (SendTransactionResponse, error) {
	return c.processSendTransaction(c.Call(ctx, "sendTransaction", tx))
}

// SendTransaction submits a signed transaction to the cluster for processing
func (c *RpcClient) SendTransactionWithConfig(ctx context.Context, tx string, cfg SendTransactionConfig) (SendTransactionResponse, error) {
	return c.processSendTransaction(c.Call(ctx, "sendTransaction", tx, cfg))
}

func (c *RpcClient) processSendTransaction(body []byte, err error) (SendTransactionResponse, error) {
	if err != nil {
		return SendTransactionResponse{}, fmt.Errorf("rpc: call error, err: %v", err)
	}
	var res SendTransactionResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return SendTransactionResponse{}, fmt.Errorf("rpc: failed to json decode body, err: %v", err)
	}
	return res, nil
}
