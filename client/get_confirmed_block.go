package client

import (
	"fmt"
)

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

func (s *Client) GetConfirmedBlock(slot uint64) (GetConfirmBlockResponse, error) {
	res := struct {
		GeneralResponse
		Result GetConfirmBlockResponse `json:"result"`
	}{}
	fmt.Println(s.request("getConfirmedBlock", []interface{}{slot, "json"}, &res))
	return res.Result, nil
}
