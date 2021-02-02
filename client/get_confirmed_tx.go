package client

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
	err := s.request("getConfirmedTransaction", []interface{}{txhash, "json"}, &res)
	if err != nil {
		return GetConfirmedTransactionResponse{}, err
	}
	return res.Result, nil
}
