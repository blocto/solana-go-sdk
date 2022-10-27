package rpc

import (
	"context"
)

type GetTransactionResponse JsonRpcResponse[*GetTransaction]

// GetTransaction is a part of GetTransactionResponse
type GetTransaction struct {
	Slot        uint64           `json:"slot"`
	Meta        *TransactionMeta `json:"meta"`
	Transaction any              `json:"transaction"`
	BlockTime   *int64           `json:"blockTime"`
	Version     any              `json:"version,omitempty"`
}

// TransactionMeta is a part of GetTransactionResult
type TransactionMeta struct {
	Err               any                               `json:"err"`
	Fee               uint64                            `json:"fee"`
	PreBalances       []int64                           `json:"preBalances"`
	PostBalances      []int64                           `json:"postBalances"`
	PreTokenBalances  []TransactionMetaTokenBalance     `json:"preTokenBalances"`
	PostTokenBalances []TransactionMetaTokenBalance     `json:"postTokenBalances"`
	LogMessages       []string                          `json:"logMessages"`
	InnerInstructions []TransactionMetaInnerInstruction `json:"innerInstructions"`
	LoadedAddresses   TransactionLoadedAddresses        `json:"loadedAddresses"`
}

// TransactionMetaTokenBalance is a part of TransactionMeta
type TransactionMetaTokenBalance struct {
	AccountIndex  uint64              `json:"accountIndex"`
	Mint          string              `json:"mint"`
	Owner         string              `json:"owner,omitempty"`
	UITokenAmount TokenAccountBalance `json:"uiTokenAmount"`
}

// TransactionMetaInnerInstruction is a part of TransactionMeta
type TransactionMetaInnerInstruction struct {
	Index        uint64 `json:"index"`
	Instructions []any  `json:"instructions"`
}

// Instruction is a part of TransactionMetaInnerInstruction
type Instruction struct {
	ProgramIDIndex int    `json:"programIdIndex"`
	Accounts       []int  `json:"accounts"`
	Data           string `json:"data"`
}

// TransactionMetaReward is a part of TransactionMeta
type TransactionMetaReward struct {
	Pubkey       string                    `json:"pubkey"`
	Lamports     int64                     `json:"lamports"`
	PostBalances uint64                    `json:"postBalance"`
	RewardType   TransactionMetaRewardType `json:"rewardType"`
	Commission   *uint8                    `json:"commission"`
}

type TransactionMetaRewardType string

const (
	// currently only "rent", other types may be added in the future
	TransactionMetaRewardTypeRent TransactionMetaRewardType = "rent"
)

type TransactionLoadedAddresses struct {
	Writable []string `json:"writable"`
	Readonly []string `json:"readonly"`
}

// GetTransactionConfig is a option config for `getTransaction`
type GetTransactionConfig struct {
	Encoding                       TransactionEncoding `json:"encoding,omitempty"`
	Commitment                     Commitment          `json:"commitment,omitempty"`                     // "processed" is not supported
	MaxSupportedTransactionVersion *uint8              `json:"maxSupportedTransactionVersion,omitempty"` // default: nil legacy only
}

// GetTransaction returns transaction details for a confirmed transaction
func (c *RpcClient) GetTransaction(ctx context.Context, txhash string) (JsonRpcResponse[*GetTransaction], error) {
	return c.processGetTransaction(c.Call(ctx, "getTransaction", txhash))
}

// GetTransactionWithConfig returns transaction details for a confirmed transaction
func (c *RpcClient) GetTransactionWithConfig(ctx context.Context, txhash string, cfg GetTransactionConfig) (JsonRpcResponse[*GetTransaction], error) {
	return c.processGetTransaction(c.Call(ctx, "getTransaction", txhash, cfg))
}

func (c *RpcClient) processGetTransaction(body []byte, rpcErr error) (res JsonRpcResponse[*GetTransaction], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
