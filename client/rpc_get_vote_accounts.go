package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/rpc"
)

type VoteAccountStatus struct {
	Current    []VoteAccountInfo
	Delinquent []VoteAccountInfo
}

type VoteAccountInfo struct {
	VotePubkey       common.PublicKey
	NodePubkey       common.PublicKey
	ActivatedStake   uint64
	Commission       uint8
	EpochVoteAccount bool
	LastVote         uint64
	EpochCredits     []EpochCredits
	RootSlot         uint64
}

type EpochCredits struct {
	Epoch           uint64
	Credits         uint64
	PreviousCredits uint64
}

// GetVoteAccounts returns the account info and associated stake for all the voting accounts in the current bank.
func (c *Client) GetVoteAccounts(ctx context.Context) (VoteAccountStatus, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.GetVoteAccounts], error) {
			return c.RpcClient.GetVoteAccounts(ctx)
		},
		convertGetVoteAccounts,
	)
}

func convertGetVoteAccounts(v rpc.GetVoteAccounts) (VoteAccountStatus, error) {
	convertVoteAccount := func(rpcVoteAccountInfo rpc.VoteAccount) VoteAccountInfo {
		epochCredits := make([]EpochCredits, 0, len(rpcVoteAccountInfo.EpochCredits))
		for _, e := range rpcVoteAccountInfo.EpochCredits {
			epochCredits = append(epochCredits, EpochCredits{
				Epoch:           e[0],
				Credits:         e[1],
				PreviousCredits: e[2],
			})
		}

		return VoteAccountInfo{
			VotePubkey:       common.PublicKeyFromString(rpcVoteAccountInfo.VotePubkey),
			NodePubkey:       common.PublicKeyFromString(rpcVoteAccountInfo.NodePubkey),
			ActivatedStake:   rpcVoteAccountInfo.ActivatedStake,
			Commission:       rpcVoteAccountInfo.Commission,
			EpochVoteAccount: rpcVoteAccountInfo.EpochVoteAccount,
			LastVote:         rpcVoteAccountInfo.LastVote,
			EpochCredits:     epochCredits,
			RootSlot:         rpcVoteAccountInfo.RootSlot,
		}
	}

	current := make([]VoteAccountInfo, 0, len(v.Current))
	for _, voteAccount := range v.Current {
		current = append(current, convertVoteAccount(voteAccount))
	}

	deliquent := make([]VoteAccountInfo, 0, len(v.Deliquent))
	for _, voteAccount := range v.Deliquent {
		deliquent = append(deliquent, convertVoteAccount(voteAccount))
	}

	return VoteAccountStatus{
		Current:    current,
		Delinquent: deliquent,
	}, nil
}
