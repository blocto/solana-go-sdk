package client

import (
	"context"
	"errors"
)

type SimulateTransactionConfig struct {
	SigVerify           bool       `json:"sigVerify"`           // default: false
	PreflightCommitment Commitment `json:"preflightCommitment"` // default: max
	Encoding            string     `json:"encoding"`            // base58 or base64
}

type SimulateTransactionResponse struct {
	Err  interface{} `json:"err"`
	Logs []string    `json:"logs"`
}

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
