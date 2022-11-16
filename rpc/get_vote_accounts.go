package rpc

import "context"

type GetVoteAccountsResponse JsonRpcResponse[GetVoteAccounts]

type GetVoteAccounts struct {
	Current   VoteAccounts `json:"current"`
	Deliquent VoteAccounts `json:"delinquent"`
}

type VoteAccounts []VoteAccount

type VoteAccount struct {
	VotePubkey       string      `json:"votePubkey"`
	NodePubkey       string      `json:"nodePubkey"`
	ActivatedStake   uint64      `json:"activatedStake"`
	Commission       uint8       `json:"commission"`
	EpochVoteAccount bool        `json:"epochVoteAccount"`
	LastVote         uint64      `json:"lastVote"`
	EpochCredits     [][3]uint64 `json:"epochCredits"`
	RootSlot         uint64      `json:"rootSlot"`
}

type GetVoteAccountsConfig struct {
	Commitment              Commitment `json:"commitment,omitempty"`
	VotePubkey              string     `json:"votePubkey,omitempty"`
	KeepUnstakedDelinquents bool       `json:"keepUnstakedDelinquents,omitempty"`
	DelinquentSlotDistance  uint64     `json:"delinquentSlotDistance,omitempty"`
}

// GetVoteAccounts returns the account info and associated stake for all the voting accounts in the current bank.
func (c *RpcClient) GetVoteAccounts(ctx context.Context) (JsonRpcResponse[GetVoteAccounts], error) {
	return call[JsonRpcResponse[GetVoteAccounts]](c, ctx, "getVoteAccounts")
}

// GetVoteAccountsWithConfig returns the account info and associated stake for all the voting accounts in the current bank.
func (c *RpcClient) GetVoteAccountsWithConfig(ctx context.Context, cfg GetVoteAccountsConfig) (JsonRpcResponse[GetVoteAccounts], error) {
	return call[JsonRpcResponse[GetVoteAccounts]](c, ctx, "getVoteAccounts", cfg)
}
