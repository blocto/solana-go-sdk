package client

type GetEpochInfoResponse struct {
	AbsoluteSlot int `json:"absoluteSlot"`
	BlockHeight  int `json:"blockHeight"`
	Epoch        int `json:"epoch"`
	SlotIndex    int `json:"slotIndex"`
	SlotsInEpoch int `json:"slotsInEpoch"`
}

func (s *Client) GetEpochInfo(commitment Commitment) (GetEpochInfoResponse, error) {
	res := struct {
		GeneralResponse
		Result GetEpochInfoResponse `json:"result"`
	}{}
	err := s.request("getEpochInfo", []interface{}{map[string]interface{}{"commitment": commitment}}, &res)
	if err != nil {
		return GetEpochInfoResponse{}, err
	}
	return res.Result, nil
}
