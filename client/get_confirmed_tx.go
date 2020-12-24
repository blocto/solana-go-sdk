package client

import (
	"fmt"
)

type GetConfirmedTransactionResponse struct {
	Slot        uint64          `json:"slot"`
	Meta        TransactionMeta `json:"meta"`
	Transaction Transaction     `json:"transaction"`
}

func (s *Client) GetConfirmedTransaction(txhash string) (GetConfirmedTransactionResponse, error) {
	res := struct {
		GeneralResponse
		Result GetConfirmedTransactionResponse `json:"result"`
	}{}
	fmt.Println(s.request("getConfirmedTransaction", []interface{}{txhash, "json"}, &res))
	return res.Result, nil
}
