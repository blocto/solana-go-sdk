package rpc

import (
	"context"
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

func (c *RpcClient) processSendTransaction(body []byte, rpcErr error) (res SendTransactionResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
