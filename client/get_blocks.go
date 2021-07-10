package client

import "context"

type GetBlocksConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedBlocks for solana-core v1.6
// GetBlocks returns a list of confirmed blocks between two slots
// Max range allowed is 500,000 slots
func (s *Client) GetBlocks(ctx context.Context, startSlot uint64, endSlot uint64, cfg GetBlocksConfig) ([]uint64, error) {
	res := struct {
		GeneralResponse
		Result []uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getBlocks", []interface{}{startSlot, endSlot, cfg}, &res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}
