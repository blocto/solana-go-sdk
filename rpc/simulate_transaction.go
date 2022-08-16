package rpc

import (
	"context"
)

// SimulateTransactionResponse is a complete rpc response of `simulateTransaction`
type SimulateTransactionResponse struct {
	GeneralResponse
	Result SimulateTransactionResponseResult `json:"result"`
}

// SimulateTransactionResponseResult is a part of SimulateTransactionResponse
type SimulateTransactionResponseResult struct {
	Context Context                                `json:"context"`
	Value   SimulateTransactionResponseResultValue `json:"value"`
}

// SimulateTransactionResponseResultValue is a part of SimulateTransactionResponseResult
type SimulateTransactionResponseResultValue struct {
	Err      any            `json:"err"`
	Logs     []string       `json:"logs,omitempty"`
	Accounts []*AccountInfo `json:"accounts,omitempty"`
}

type SimulateTransactionConfig struct {
	SigVerify              bool                               `json:"sigVerify,omitempty"`              // default: false, conflicts with replace blockhash
	Commitment             Commitment                         `json:"commitment,omitempty"`             // default: finalized
	Encoding               SimulateTransactionConfigEncoding  `json:"encoding,omitempty"`               // default: "base58"
	ReplaceRecentBlockhash bool                               `json:"replaceRecentBlockhash,omitempty"` // default: false, conflicts with sigVerify
	Accounts               *SimulateTransactionConfigAccounts `json:"accounts,omitempty"`
}

type SimulateTransactionConfigEncoding string

const (
	SimulateTransactionConfigEncodingBase58 SimulateTransactionConfigEncoding = "base58" // (slow, DEPRECATED)
	SimulateTransactionConfigEncodingBase64 SimulateTransactionConfigEncoding = "base64"
)

type SimulateTransactionConfigAccounts struct {
	Encoding AccountEncoding `json:"encoding,omitempty"`
	// An array of accounts to return, as base-58 encoded strings
	Addresses []string `json:"addresses"`
}

// SimulateTransaction simulate sending a transaction
func (c *RpcClient) SimulateTransaction(ctx context.Context, rawTx string) (SimulateTransactionResponse, error) {
	return c.processSimulateTransaction(c.Call(ctx, "simulateTransaction", rawTx))
}

// SimulateTransaction simulate sending a transaction
func (c *RpcClient) SimulateTransactionWithConfig(ctx context.Context, rawTx string, cfg SimulateTransactionConfig) (SimulateTransactionResponse, error) {
	return c.processSimulateTransaction(c.Call(ctx, "simulateTransaction", rawTx, cfg))
}

func (c *RpcClient) processSimulateTransaction(body []byte, rpcErr error) (res SimulateTransactionResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
