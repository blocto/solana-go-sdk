package client

import "context"

type GetConfirmedTransactionResponse struct {
	Slot        uint64          `json:"slot"`
	Meta        TransactionMeta `json:"meta"`
	Transaction Transaction     `json:"transaction"`
}

// DEPRECATED: Please use getTransaction instead This method is expected to be removed in solana-core v1.8
// GetConfirmedTransaction returns transaction details for a confirmed transaction
func (s *RpcClient) GetConfirmedTransaction(ctx context.Context, txhash string) (GetConfirmedTransactionResponse, error) {
	res := struct {
		GeneralResponse
		Result GetConfirmedTransactionResponse `json:"result"`
	}{}
	err := s.request(ctx, "getConfirmedTransaction", []interface{}{txhash, "json"}, &res)
	if err != nil {
		return GetConfirmedTransactionResponse{}, err
	}
	return res.Result, nil
}
