package rpc

import "context"

type GetEpochInfoResponse struct {
	AbsoluteSlot uint64 `json:"absoluteSlot"`
	BlockHeight  uint64 `json:"blockHeight"`
	Epoch        uint64 `json:"epoch"`
	SlotIndex    uint64 `json:"slotIndex"`
	SlotsInEpoch uint64 `json:"slotsInEpoch"`
}

// GetEpochInfo returns information about the current epoch
func (s *RpcClient) GetEpochInfo(ctx context.Context, commitment Commitment) (GetEpochInfoResponse, error) {
	res := struct {
		GeneralResponse
		Result GetEpochInfoResponse `json:"result"`
	}{}
	err := s.request(ctx, "getEpochInfo", []interface{}{map[string]interface{}{"commitment": commitment}}, &res)
	if err != nil {
		return GetEpochInfoResponse{}, err
	}
	return res.Result, nil
}
