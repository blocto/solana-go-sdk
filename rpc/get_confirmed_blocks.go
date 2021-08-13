package client

import "context"

// DEPRECATED: Please use getBlocks instead This method is expected to be removed in solana-core v1.8
// GetConfirmedBlocks returns a list of confirmed blocks between two slots
func (s *RpcClient) GetConfirmedBlocks(ctx context.Context, startSlot uint64, endSlot uint64) ([]uint64, error) {
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
