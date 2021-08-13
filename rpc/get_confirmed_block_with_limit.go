package client

import "context"

// DEPRECATED: Please use getBlocksWithLimit instead This method is expected to be removed in solana-core v1.8
// GetConfirmedBlocksWithLimit returns a list of confirmed blocks starting at the given slot
func (s *RpcClient) GetConfirmedBlocksWithLimit(ctx context.Context, startSlot uint64, limit uint64) ([]uint64, error) {
	res := struct {
		GeneralResponse
		Result []uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getConfirmedBlocksWithLimit", []interface{}{startSlot, limit}, &res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}
