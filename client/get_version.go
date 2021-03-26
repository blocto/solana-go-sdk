package client

type GetVersionResponse struct {
	SolanaCore string `json:"solana-core"`
	FeatureSet uint64 `json:"feature-set"`
}

func (s *Client) GetVersion() (GetVersionResponse, error) {
	res := struct {
		GeneralResponse
		Result GetVersionResponse `json:"result"`
	}{}
	err := s.request("getVersion", []interface{}{}, &res)
	if err != nil {
		return GetVersionResponse{}, err
	}
	return res.Result, nil
}
