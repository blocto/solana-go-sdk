package rpc

import (
	"context"
)

type SimulateTransactionResponse JsonRpcResponse[SimulateTransaction]

type SimulateTransaction ValueWithContext[SimulateTransactionValue]

// SimulateTransactionValue is a part of SimulateTransactionResponseResult
type SimulateTransactionValue struct {
	Err          any            `json:"err"`
	Logs         []string       `json:"logs,omitempty"`
	Accounts     []*AccountInfo `json:"accounts,omitempty"`
	ReturnData   *ReturnData    `json:"returnData,omitempty"`
	UnitConsumed *uint64        `json:"unitsConsumed,omitempty"`
}

type SimulateTransactionConfig struct {
	SigVerify              bool                               `json:"sigVerify,omitempty"`              // default: false, conflicts with replace blockhash
	Commitment             Commitment                         `json:"commitment,omitempty"`             // default: finalized
	Encoding               SimulateTransactionEncoding        `json:"encoding,omitempty"`               // default: "base58"
	ReplaceRecentBlockhash bool                               `json:"replaceRecentBlockhash,omitempty"` // default: false, conflicts with sigVerify
	Accounts               *SimulateTransactionConfigAccounts `json:"accounts,omitempty"`
}

type SimulateTransactionConfigAccounts struct {
	Encoding AccountEncoding `json:"encoding,omitempty"`
	// An array of accounts to return, as base-58 encoded strings
	Addresses []string `json:"addresses"`
}

type SimulateTransactionEncoding string

const (
	SimulateTransactionEncodingBase58 SimulateTransactionEncoding = "base58"
	SimulateTransactionEncodingBase64 SimulateTransactionEncoding = "base64"
)

// SimulateTransaction simulate sending a transaction
func (c *RpcClient) SimulateTransaction(ctx context.Context, rawTx string) (JsonRpcResponse[ValueWithContext[SimulateTransactionValue]], error) {
	return call[JsonRpcResponse[ValueWithContext[SimulateTransactionValue]]](c, ctx, "simulateTransaction", rawTx)
}

// SimulateTransaction simulate sending a transaction
func (c *RpcClient) SimulateTransactionWithConfig(ctx context.Context, rawTx string, cfg SimulateTransactionConfig) (JsonRpcResponse[ValueWithContext[SimulateTransactionValue]], error) {
	return call[JsonRpcResponse[ValueWithContext[SimulateTransactionValue]]](c, ctx, "simulateTransaction", rawTx, cfg)
}
