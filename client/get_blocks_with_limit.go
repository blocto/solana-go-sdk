package client

import "context"

type GetBlocksWithLimitConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedBlocksWithLimit for solana-core v1.6
// GetBlocksWithLimit returns a list of confirmed blocks starting at the given slot
func (s *Client) GetBlocksWithLimit(ctx context.Context, startSlot uint64, limit uint64, cfg GetBlocksWithLimitConfig) ([]uint64, error) {
	res := struct {
		GeneralResponse
		Result []uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getBlocksWithLimit", []interface{}{startSlot, limit, cfg}, &res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}
