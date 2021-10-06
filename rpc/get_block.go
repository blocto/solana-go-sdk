package rpc

import "context"

type GetBlockResponse struct {
	GeneralResponse
	Result GetBlockResponseResult `json:"result"`
}

type GetBlockResponseResult struct {
	Blockhash         string                `json:"blockhash"`
	BlockTime         *int64                `json:"blockTime"`
	BlockHeight       *int64                `json:"blockHeight"`
	PreviousBlockhash string                `json:"previousBlockhash"`
	ParentSLot        uint64                `json:"parentSlot"`
	Transactions      []GetBlockTransaction `json:"transactions"`
	Signatures        []string              `json:"signatures"`
	Rewards           []GetBlockReward      `json:"rewards"`
}

type GetBlockReward struct {
	Pubkey       string             `json:"pubkey"`
	Lamports     int64              `json:"lamports"`
	PostBalances uint64             `json:"postBalance"`
	RewardType   GetBlockRewardType `json:"rewardType"`
	Commission   *uint8             `json:"commission"`
}

type GetBlockRewardType string

const (
	GetBlockRewardTypeNone    GetBlockRewardType = ""
	GetBlockRewardTypeFee     GetBlockRewardType = "fee"
	GetBlockRewardTypeRent    GetBlockRewardType = "rent"
	GetBlockRewardTypeVoting  GetBlockRewardType = "voting"
	GetBlockRewardTypeStaking GetBlockRewardType = "staking"
)

type GetBlockTransaction struct {
	Transaction interface{}      `json:"transaction"`
	Meta        *TransactionMeta `json:"meta"`
}

type GetBlockConfig struct {
	Encoding           GetBlockConfigEncoding           `json:"encoding,omitempty"`           // default: "json"
	TransactionDetails GetBlockConfigTransactionDetails `json:"transactionDetails,omitempty"` // default: "full", either "full", "signatures", "none"
	Rewards            *bool                            `json:"rewards,omitempty"`            // default: true
	Commitment         Commitment                       `json:"commitment,omitempty"`         // "processed" is not supported
}

type GetBlockConfigEncoding string

const (
	GetBlockConfigEncodingJson       GetBlockConfigEncoding = "json"
	GetBlockConfigEncodingJsonParsed GetBlockConfigEncoding = "jsonParsed"
	GetBlockConfigEncodingBase58     GetBlockConfigEncoding = "base58"
	GetBlockConfigEncodingBase64     GetBlockConfigEncoding = "base64"
)

type GetBlockConfigTransactionDetails string

const (
	GetBlockConfigTransactionDetailsFull       GetBlockConfigTransactionDetails = "full"
	GetBlockConfigTransactionDetailsSignatures GetBlockConfigTransactionDetails = "signatures"
	GetBlockConfigTransactionDetailsNone       GetBlockConfigTransactionDetails = "none"
)

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedBlock for solana-core v1.6
// GetBlock returns identity and transaction information about a confirmed block in the ledger
func (c *RpcClient) GetBlock(ctx context.Context, slot uint64) (GetBlockResponse, error) {
	return c.processGetBlock(c.Call(ctx, "getBlock", slot))
}

// GetAccountInfo returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetBlockWithConfig(ctx context.Context, slot uint64, cfg GetBlockConfig) (GetBlockResponse, error) {
	return c.processGetBlock(c.Call(ctx, "getBlock", slot, cfg))
}

func (c *RpcClient) processGetBlock(body []byte, rpcErr error) (res GetBlockResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
