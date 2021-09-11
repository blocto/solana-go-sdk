package rpc

import "context"

type GetConfirmBlockResponse struct {
	Blockhash         string `json:"blockhash"`
	PreviousBlockhash string `json:"previousBlockhash"`
	ParentSLot        uint64 `json:"parentSlot"`
	BlockTime         int64  `json:"blockTime"`
	Transactions      []struct {
		Meta        TransactionMeta `json:"meta"`
		Transaction Transaction     `json:"transaction"`
	} `json:"transactions"`
	Rewards []struct {
		Pubkey      string `json:"pubkey"`
		Lamports    int64  `json:"lamports"`
		PostBalance uint64 `json:"postBalance"`
		RewardType  string `json:"rewardType"` // type of reward: "fee", "rent", "voting", "staking"
	} `json:"rewards"`
}

// DEPRECATED: Please use getBlock instead This method is expected to be removed in solana-core v1.8
// GetConfirmedBlock returns identity and transaction information about a confirmed block in the ledger
func (s *RpcClient) GetConfirmedBlock(ctx context.Context, slot uint64) (GetConfirmBlockResponse, error) {
	res := struct {
		GeneralResponse
		Result GetConfirmBlockResponse `json:"result"`
	}{}
	err := s.request(ctx, "getConfirmedBlock", []interface{}{slot, "json"}, &res)
	if err != nil {
		return GetConfirmBlockResponse{}, err
	}
	return res.Result, nil
}
