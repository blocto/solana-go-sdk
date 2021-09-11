package rpc

import "context"

type GetBlockConfig struct {
	// TODO custom
	// Encoding           string     `json:"encoding"`             // default: "json", either "json", "jsonParsed", "base58" (slow), "base64"
	// TransactionDetails string     `json:"transactionDetails"`   // default: "full", either "full", "signatures", "none"
	Commitment Commitment `json:"commitment,omitempty"` // "processed" is not supported. If parameter not provided, the default is "finalized".
}

type GetBlockResponse struct {
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

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedBlock for solana-core v1.6
// GetBlock returns identity and transaction information about a confirmed block in the ledger
func (s *RpcClient) GetBlock(ctx context.Context, slot uint64, cfg GetBlockConfig) (GetBlockResponse, error) {
	res := struct {
		GeneralResponse
		Result GetBlockResponse `json:"result"`
	}{}
	err := s.request(ctx, "getBlock", []interface{}{slot, cfg}, &res)
	if err != nil {
		return GetBlockResponse{}, err
	}
	return res.Result, nil
}
