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
	Err                  any                               `json:"err"`
	Fee                  uint64                            `json:"fee"`
	PreBalances          []int64                           `json:"preBalances"`
	PostBalances         []int64                           `json:"postBalances"`
	PreTokenBalances     []TransactionMetaTokenBalance     `json:"preTokenBalances"`
	PostTokenBalances    []TransactionMetaTokenBalance     `json:"postTokenBalances"`
	Rewards              []Reward                          `json:"rewards"`
	LogMessages          []string                          `json:"logMessages"`
	InnerInstructions    []TransactionMetaInnerInstruction `json:"innerInstructions"`
	LoadedAddresses      TransactionLoadedAddresses        `json:"loadedAddresses"`
	ReturnData           *ReturnData                       `json:"returnData"`
	ComputeUnitsConsumed *uint64                           `json:"computeUnitsConsumed"`
}

// TransactionMetaTokenBalance is a part of TransactionMeta
type TransactionMetaTokenBalance struct {
	AccountIndex  uint64              `json:"accountIndex"`
	Mint          string              `json:"mint"`
	Owner         string              `json:"owner,omitempty"`
	ProgramId     string              `json:"programId,omitempty"`
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
	return call[JsonRpcResponse[*GetTransaction]](c, ctx, "getTransaction", txhash)
}

// GetTransactionWithConfig returns transaction details for a confirmed transaction
func (c *RpcClient) GetTransactionWithConfig(ctx context.Context, txhash string, cfg GetTransactionConfig) (JsonRpcResponse[*GetTransaction], error) {
	return call[JsonRpcResponse[*GetTransaction]](c, ctx, "getTransaction", txhash, cfg)
}
