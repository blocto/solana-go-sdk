package client

import "context"

func (s *Client) GetConfirmedBlocks(ctx context.Context, startSlot uint64, endSlot uint64) ([]uint64, error) {
	res := struct {
		GeneralResponse
		Result []uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getConfirmedBlocks", []interface{}{startSlot, endSlot}, &res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}
