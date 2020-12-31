package client

import "errors"

func (s *Client) GetMinimumBalanceForRentExemption(accountDataLen uint64) (uint64, error) {
	res := struct {
		GeneralResponse
		Result uint64 `json:"result"`
	}{}
	err := s.request("getMinimumBalanceForRentExemption", []interface{}{accountDataLen}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
