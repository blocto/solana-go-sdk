package client

import (
	"context"
	"errors"
)

type SimulateTransactionEncoding string

const (
	SimulateTransactionEncodingBase58 SimulateTransactionEncoding = "base58" // (slow, DEPRECATED)
	SimulateTransactionEncodingBase64 SimulateTransactionEncoding = "base64"
)

type SimulateTransactionConfig struct {
	SigVerify              bool                               `json:"sigVerify,omitempty"`              // default: false, conflicts with replace blockhash
	PreflightCommitment    Commitment                         `json:"preflightCommitment,omitempty"`    // default: finalized
	Encoding               SimulateTransactionEncoding        `json:"encoding,omitempty"`               // default: "base58", either "base58" or "base64"
	ReplaceRecentBlockhash bool                               `json:"replaceRecentBlockhash,omitempty"` // default: false, conflicts with sigVerify
	Accounts               *SimulateTransactionConfigAccounts `json:"accounts,omitempty"`
}

type SimulateTransactionConfigAccountsEncoding string

type SimulateTransactionConfigAccounts struct {
	Encoding  string   `json:"encoding,omitempty"`
	Addresses []string `json:"addresses"`
}

type SimulateTransactionResponse struct {
	Err      interface{}              `json:"err"`
	Logs     []string                 `json:"logs"`
	Accounts []GetAccountInfoResponse `json:"accounts"`
}

// SimulateTransaction simulate sending a transaction
func (s *Client) SimulateTransaction(ctx context.Context, rawTx string, cfg SimulateTransactionConfig) (SimulateTransactionResponse, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context                     `json:"context"`
			Value   SimulateTransactionResponse `json:"value"`
		} `json:"result"`
	}{}
	err := s.request(ctx, "simulateTransaction", []interface{}{rawTx, cfg}, &res)
	if err != nil {
		return SimulateTransactionResponse{}, err
	}
	if res.Error != (ErrorResponse{}) {
		return SimulateTransactionResponse{}, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
