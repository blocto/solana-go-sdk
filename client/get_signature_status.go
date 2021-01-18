package client

import "errors"

type GetSignatureStatusesResponse struct {
	Slot          uint64      `json:"slot"`
	Confirmations *uint64     `json:"confirmations"`
	Err           interface{} `json:"err"`
}

func (s *Client) GetSignatureStatuses(signatures []string) ([]GetSignatureStatusesResponse, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context                        `json:"context"`
			Value   []GetSignatureStatusesResponse `json:"value"`
		} `json:"result"`
	}{}
	err := s.request("getSignatureStatuses", []interface{}{signatures, map[string]interface{}{"searchTransactionHistory": true}}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != (ErrorResponse{}) {
		return nil, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
