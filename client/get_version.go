package client

import "context"

type GetVersionResponse struct {
	SolanaCore string `json:"solana-core"`
	FeatureSet uint64 `json:"feature-set"`
}

// GetVersion returns the current solana versions running on the node
func (s *Client) GetVersion(ctx context.Context) (GetVersionResponse, error) {
	res := struct {
		GeneralResponse
		Result GetVersionResponse `json:"result"`
	}{}
	err := s.request(ctx, "getVersion", []interface{}{}, &res)
	if err != nil {
		return GetVersionResponse{}, err
	}
	return res.Result, nil
}
