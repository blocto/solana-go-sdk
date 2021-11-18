package rpc

import (
	"context"
)

// GetTransactionResponse is a complete rpc response of `getTransaction`
type GetTransactionResponse struct {
	GeneralResponse
	Result *GetTransactionResult `json:"result"`
}

// GetTransactionResult is a part of GetTransactionResponse
type GetTransactionResult struct {
	Slot        uint64           `json:"slot"`
	Meta        *TransactionMeta `json:"meta"`
	Transaction interface{}      `json:"transaction"`
	BlockTime   *int64           `json:"blockTime"`
}

// TransactionMeta is a part of GetTransactionResult
type TransactionMeta struct {
	Err               interface{}                       `json:"err"`
	Fee               uint64                            `json:"fee"`
	PreBalances       []int64                           `json:"preBalances"`
	PostBalances      []int64                           `json:"postBalances"`
	PreTokenBalances  []TransactionMetaTokenBalance     `json:"preTokenBalances"`
	PostTokenBalances []TransactionMetaTokenBalance     `json:"postTokenBalances"`
	LogMessages       []string                          `json:"logMessages"`
	InnerInstructions []TransactionMetaInnerInstruction `json:"innerInstructions"`
}

// TransactionMetaTokenBalance is a part of TransactionMeta
type TransactionMetaTokenBalance struct {
	AccountIndex  uint64                            `json:"accountIndex"`
	Mint          string                            `json:"mint"`
	Owner         string                            `json:"owner,omitempty"`
	UITokenAmount GetTokenAccountBalanceResultValue `json:"uiTokenAmount"`
}

// TransactionMetaInnerInstruction is a part of TransactionMeta
type TransactionMetaInnerInstruction struct {
	Index        uint64        `json:"index"`
	Instructions []Instruction `json:"instructions"`
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

// GetTransactionConfig is a option config for `getTransaction`
type GetTransactionConfig struct {
	Encoding   GetTransactionConfigEncoding `json:"encoding,omitempty"`
	Commitment Commitment                   `json:"commitment,omitempty"` // "processed" is not supported
}

type GetTransactionConfigEncoding string

const (
	GetTransactionConfigEncodingJson       GetTransactionConfigEncoding = "json"
	GetTransactionConfigEncodingJsonParsed GetTransactionConfigEncoding = "jsonParsed"
	GetTransactionConfigEncodingBase58     GetTransactionConfigEncoding = "base58"
	GetTransactionConfigEncodingBase64     GetTransactionConfigEncoding = "base64"
)

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedTransaction for solana-core v1.6
// GetTransaction returns transaction details for a confirmed transaction
func (c *RpcClient) GetTransaction(ctx context.Context, txhash string) (GetTransactionResponse, error) {
	return c.processGetTransaction(c.Call(ctx, "getTransaction", txhash))
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedTransaction for solana-core v1.6
// GetTransactionWithConfig returns transaction details for a confirmed transaction
func (c *RpcClient) GetTransactionWithConfig(ctx context.Context, txhash string, cfg GetTransactionConfig) (GetTransactionResponse, error) {
	return c.processGetTransaction(c.Call(ctx, "getTransaction", txhash, cfg))
}

func (c *RpcClient) processGetTransaction(body []byte, rpcErr error) (res GetTransactionResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
