package client

func (s *Client) GetConfirmedBlocksWithLimit(startSlot uint64, limit uint64) ([]uint64, error) {
	res := struct {
		GeneralResponse
		Result []uint64 `json:"result"`
	}{}
	err := s.request("getConfirmedBlocksWithLimit", []interface{}{startSlot, limit}, &res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}
