package client

import "context"

type GetInflationRate struct {
	Epoch      int64   `json:"epoch"`
	Foundation float64 `json:"foundation"`
	Total      float64 `json:"total"`
	Validator  float64 `json:"validator"`
}

// GetInflationRate returns the specific inflation values for the current epoch
func (s *Client) GetInflationRate(ctx context.Context, commitment Commitment) (GetInflationRate, error) {
	res := struct {
		GeneralResponse
		Result GetInflationRate `json:"result"`
	}{}
	err := s.request(ctx, "getEpochInfo", []interface{}{}, &res)
	if err != nil {
		return GetInflationRate{}, err
	}
	return res.Result, nil
}
