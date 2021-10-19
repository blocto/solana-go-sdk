package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mr-tron/base58"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

type Client struct {
	RpcClient rpc.RpcClient
}

func NewClient(endpoint string) *Client {
	return &Client{rpc.NewRpcClient(endpoint)}
}

// GetBalance fetch users lamports(SOL) balance
func (c *Client) GetBalance(ctx context.Context, base58Addr string) (uint64, error) {
	res, err := c.RpcClient.GetBalance(ctx, base58Addr)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result.Value, nil
}

// GetBalance fetch users lamports(SOL) balance with specific commitment
func (c *Client) GetBalanceWithConfig(ctx context.Context, base58Addr string, cfg rpc.GetBalanceConfig) (uint64, error) {
	res, err := c.RpcClient.GetBalanceWithConfig(ctx, base58Addr, cfg)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result.Value, nil
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *Client) GetTokenAccountBalance(ctx context.Context, base58Addr string) (uint64, uint8, error) {
	res, err := c.RpcClient.GetTokenAccountBalance(ctx, base58Addr)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, 0, err
	}
	balance, err := strconv.ParseUint(res.Result.Value.Amount, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to cast token amount, err: %v", err)
	}
	return balance, res.Result.Value.Decimals, nil
}

// GetTokenAccountBalance returns the token balance of an SPL Token account
func (c *Client) GetTokenAccountBalanceWithConfig(ctx context.Context, base58Addr string, cfg rpc.GetTokenAccountBalanceConfig) (uint64, uint8, error) {
	res, err := c.RpcClient.GetTokenAccountBalanceWithConfig(ctx, base58Addr, cfg)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, 0, err
	}
	balance, err := strconv.ParseUint(res.Result.Value.Amount, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to cast token amount, err: %v", err)
	}
	return balance, res.Result.Value.Decimals, nil
}

// GetTokenSupply returns the total supply of an SPL Token type.
func (c *Client) GetTokenSupply(ctx context.Context, mintAddr string) (uint64, uint8, error) {
	res, err := c.RpcClient.GetTokenSupply(ctx, mintAddr)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, 0, err
	}
	balance, err := strconv.ParseUint(res.Result.Value.Amount, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to cast token amount, err: %v", err)
	}
	return balance, res.Result.Value.Decimals, nil
}

// GetTokenSupply returns the total supply of an SPL Token type.
func (c *Client) GetTokenSupplyWithConfig(ctx context.Context, mintAddr string, cfg rpc.GetTokenSupplyConfig) (uint64, uint8, error) {
	res, err := c.RpcClient.GetTokenSupplyWithConfig(ctx, mintAddr, cfg)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, 0, err
	}
	balance, err := strconv.ParseUint(res.Result.Value.Amount, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to cast token amount, err: %v", err)
	}
	return balance, res.Result.Value.Decimals, nil
}

type AccountInfo struct {
	Lamports  uint64
	Owner     string
	Excutable bool
	RentEpoch uint64
	Data      []byte
}

// GetAccountInfo return account's info
func (c *Client) GetAccountInfo(ctx context.Context, base58Addr string) (AccountInfo, error) {
	res, err := c.RpcClient.GetAccountInfoWithConfig(ctx, base58Addr, rpc.GetAccountInfoConfig{
		Encoding: rpc.GetAccountInfoConfigEncodingBase64,
	})
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return AccountInfo{}, err
	}
	if res.Result.Value == (rpc.GetAccountInfoResultValue{}) {
		return AccountInfo{}, nil
	}

	data, ok := res.Result.Value.Data.([]interface{})
	if !ok {
		return AccountInfo{}, fmt.Errorf("failed to cast raw response to []interface{}")
	}
	if data[1] != string(rpc.GetAccountInfoConfigEncodingBase64) {
		return AccountInfo{}, fmt.Errorf("encoding mistmatch")
	}
	rawData, err := base64.StdEncoding.DecodeString(data[0].(string))
	if err != nil {
		return AccountInfo{}, fmt.Errorf("failed to base64 decode data")
	}
	return AccountInfo{
		Lamports:  res.Result.Value.Lamports,
		Owner:     res.Result.Value.Owner,
		Excutable: res.Result.Value.Excutable,
		RentEpoch: res.Result.Value.RentEpoch,
		Data:      rawData,
	}, nil
}

// GetRecentBlockhash return recent blockhash information
func (c *Client) GetRecentBlockhash(ctx context.Context) (rpc.GetRecentBlockHashResultValue, error) {
	res, err := c.RpcClient.GetRecentBlockhash(ctx)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return rpc.GetRecentBlockHashResultValue{}, err
	}
	return res.Result.Value, nil
}

type QuickSendTransactionParam struct {
	Instructions []types.Instruction
	Signers      []types.Account
	FeePayer     common.PublicKey
}

// QuickSendTransaction is a quick way to send tx
func (c *Client) QuickSendTransaction(ctx context.Context, param QuickSendTransactionParam) (string, error) {
	recentBlockhashRes, err := c.GetRecentBlockhash(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get recent blockhash, err: %v", err)
	}
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions:    param.Instructions,
		Signers:         param.Signers,
		FeePayer:        param.FeePayer,
		RecentBlockHash: recentBlockhashRes.Blockhash,
	})
	if err != nil {
		return "", fmt.Errorf("failed to build tx, err: %v", err)
	}
	res, err := c.RpcClient.SendTransactionWithConfig(
		ctx,
		base64.StdEncoding.EncodeToString(rawTx),
		rpc.SendTransactionConfig{Encoding: rpc.SendTransactionConfigEncodingBase64},
	)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// SendTransaction send transaction struct directly
func (c *Client) SendTransaction(ctx context.Context, tx types.Transaction) (string, error) {
	rawTx, err := tx.Serialize()
	if err != nil {
		return "", fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	res, err := c.RpcClient.SendTransactionWithConfig(
		ctx,
		base64.StdEncoding.EncodeToString(rawTx),
		rpc.SendTransactionConfig{
			Encoding: rpc.SendTransactionConfigEncodingBase64,
		},
	)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// GetSlot get current slot (finalized)
func (c *Client) GetSlot(ctx context.Context) (uint64, error) {
	res, err := c.RpcClient.GetSlot(ctx)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

// GetSlotWithConfig get slot by commitment
func (c *Client) GetSlotWithConfig(ctx context.Context, cfg rpc.GetSlotConfig) (uint64, error) {
	res, err := c.RpcClient.GetSlotWithConfig(ctx, cfg)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

type GetTransactionResponse struct {
	Slot        uint64
	Meta        *TransactionMeta
	Transaction types.Transaction
	BlockTime   *int64
}

type TransactionMeta struct {
	Err               interface{}
	Fee               uint64
	PreBalances       []int64
	PostBalances      []int64
	PreTokenBalances  []rpc.TransactionMetaTokenBalance
	PostTokenBalances []rpc.TransactionMetaTokenBalance
	LogMessages       []string
	InnerInstructions []TransactionMetaInnerInstruction
}

type TransactionMetaInnerInstruction struct {
	Index        uint64
	Instructions []types.CompiledInstruction
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedTransaction for solana-core v1.6
// GetTransaction returns transaction details for a confirmed transaction
func (c *Client) GetTransaction(ctx context.Context, txhash string) (GetTransactionResponse, error) {
	res, err := c.RpcClient.GetTransactionWithConfig(
		ctx,
		txhash,
		rpc.GetTransactionConfig{
			Encoding: rpc.GetTransactionConfigEncodingBase64,
		},
	)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return GetTransactionResponse{}, err
	}
	return getTransaction(res)
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedTransaction for solana-core v1.6
// GetTransactionWithConfig returns transaction details for a confirmed transaction
// will ignore encoding
func (c *Client) GetTransactionWithConfig(ctx context.Context, txhash string, cfg rpc.GetTransactionConfig) (GetTransactionResponse, error) {
	res, err := c.RpcClient.GetTransactionWithConfig(
		ctx,
		txhash,
		rpc.GetTransactionConfig{
			Encoding:   rpc.GetTransactionConfigEncodingBase64,
			Commitment: cfg.Commitment,
		},
	)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return GetTransactionResponse{}, err
	}
	return getTransaction(res)
}

func getTransaction(res rpc.GetTransactionResponse) (GetTransactionResponse, error) {
	data, ok := res.Result.Transaction.([]interface{})
	if !ok {
		return GetTransactionResponse{}, fmt.Errorf("failed to cast raw response to []interface{}")
	}
	if data[1] != string(rpc.GetTransactionConfigEncodingBase64) {
		return GetTransactionResponse{}, fmt.Errorf("encoding mistmatch")
	}
	rawTx, err := base64.StdEncoding.DecodeString(data[0].(string))
	if err != nil {
		return GetTransactionResponse{}, fmt.Errorf("failed to base64 decode data, err: %v", err)
	}
	tx, err := types.TransactionDeserialize(rawTx)
	if err != nil {
		return GetTransactionResponse{}, fmt.Errorf("failed to deserialize transaction, err: %v", err)
	}

	var transactionMeta *TransactionMeta
	if res.Result.Meta != nil {
		innerInstructions := make([]TransactionMetaInnerInstruction, 0, len(res.Result.Meta.InnerInstructions))
		for _, metaInnerInstruction := range res.Result.Meta.InnerInstructions {
			compiledInstructions := make([]types.CompiledInstruction, 0, len(metaInnerInstruction.Instructions))
			for _, innerInstruction := range metaInnerInstruction.Instructions {
				data, err := base58.Decode(innerInstruction.Data)
				if err != nil {
					return GetTransactionResponse{}, fmt.Errorf("failed to base58 decode data, data: %v, err: %v", innerInstruction.Data, err)
				}
				compiledInstructions = append(compiledInstructions, types.CompiledInstruction{
					ProgramIDIndex: innerInstruction.ProgramIDIndex,
					Accounts:       innerInstruction.Accounts,
					Data:           data,
				})
			}
			innerInstructions = append(innerInstructions, TransactionMetaInnerInstruction{
				Index:        metaInnerInstruction.Index,
				Instructions: compiledInstructions,
			})
		}
		transactionMeta = &TransactionMeta{
			Err:               res.Result.Meta.Err,
			Fee:               res.Result.Meta.Fee,
			PreBalances:       res.Result.Meta.PreBalances,
			PostBalances:      res.Result.Meta.PostBalances,
			PreTokenBalances:  res.Result.Meta.PreTokenBalances,
			PostTokenBalances: res.Result.Meta.PostTokenBalances,
			LogMessages:       res.Result.Meta.LogMessages,
			InnerInstructions: innerInstructions,
		}
	}

	return GetTransactionResponse{
		Slot:        res.Result.Slot,
		BlockTime:   res.Result.BlockTime,
		Transaction: tx,
		Meta:        transactionMeta,
	}, nil
}

type GetBlockResponse struct {
	Blockhash         string
	BlockTime         *int64
	BlockHeight       *int64
	PreviousBlockhash string
	ParentSLot        uint64
	Transactions      []GetBlockTransaction
	Rewards           []rpc.GetBlockReward
}

type GetBlockTransaction struct {
	Meta        *TransactionMeta
	Transaction types.Transaction
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedBlock for solana-core v1.6
// GetBlock returns identity and transaction information about a confirmed block in the ledger
func (c *Client) GetBlock(ctx context.Context, slot uint64) (GetBlockResponse, error) {
	res, err := c.RpcClient.GetBlockWithConfig(
		ctx,
		slot,
		rpc.GetBlockConfig{
			Encoding: rpc.GetBlockConfigEncodingBase64,
		},
	)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return GetBlockResponse{}, err
	}
	return getBlock(res)
}

// add test and get block
func getBlock(res rpc.GetBlockResponse) (GetBlockResponse, error) {
	txs := make([]GetBlockTransaction, 0, len(res.Result.Transactions))
	for _, rTx := range res.Result.Transactions {
		data, ok := rTx.Transaction.([]interface{})
		if !ok {
			return GetBlockResponse{}, fmt.Errorf("failed to cast raw response to []interface{}")
		}
		if data[1] != string(rpc.GetTransactionConfigEncodingBase64) {
			return GetBlockResponse{}, fmt.Errorf("encoding mistmatch")
		}
		rawTx, err := base64.StdEncoding.DecodeString(data[0].(string))
		if err != nil {
			return GetBlockResponse{}, fmt.Errorf("failed to base64 decode data, err: %v", err)
		}
		tx, err := types.TransactionDeserialize(rawTx)
		if err != nil {
			return GetBlockResponse{}, fmt.Errorf("failed to deserialize transaction, err: %v", err)
		}

		var transactionMeta *TransactionMeta
		if rTx.Meta != nil {
			innerInstructions := make([]TransactionMetaInnerInstruction, 0, len(rTx.Meta.InnerInstructions))
			for _, metaInnerInstruction := range rTx.Meta.InnerInstructions {
				compiledInstructions := make([]types.CompiledInstruction, 0, len(metaInnerInstruction.Instructions))
				for _, innerInstruction := range metaInnerInstruction.Instructions {
					data, err := base58.Decode(innerInstruction.Data)
					if err != nil {
						return GetBlockResponse{}, fmt.Errorf("failed to base58 decode data, data: %v, err: %v", innerInstruction.Data, err)
					}
					compiledInstructions = append(compiledInstructions, types.CompiledInstruction{
						ProgramIDIndex: innerInstruction.ProgramIDIndex,
						Accounts:       innerInstruction.Accounts,
						Data:           data,
					})
				}
				innerInstructions = append(innerInstructions, TransactionMetaInnerInstruction{
					Index:        metaInnerInstruction.Index,
					Instructions: compiledInstructions,
				})
			}
			transactionMeta = &TransactionMeta{
				Err:               rTx.Meta.Err,
				Fee:               rTx.Meta.Fee,
				PreBalances:       rTx.Meta.PreBalances,
				PostBalances:      rTx.Meta.PostBalances,
				PreTokenBalances:  rTx.Meta.PreTokenBalances,
				PostTokenBalances: rTx.Meta.PostTokenBalances,
				LogMessages:       rTx.Meta.LogMessages,
				InnerInstructions: innerInstructions,
			}
		}

		txs = append(txs,
			GetBlockTransaction{
				Meta:        transactionMeta,
				Transaction: tx,
			},
		)
	}
	return GetBlockResponse{
		Blockhash:         res.Result.Blockhash,
		BlockTime:         res.Result.BlockTime,
		BlockHeight:       res.Result.BlockHeight,
		PreviousBlockhash: res.Result.PreviousBlockhash,
		ParentSLot:        res.Result.ParentSLot,
		Rewards:           res.Result.Rewards,
		Transactions:      txs,
	}, nil
}

// GetMinimumBalanceForRentExemption returns minimum balance required to make account rent exempt
func (c *Client) GetMinimumBalanceForRentExemption(ctx context.Context, dataLen uint64) (uint64, error) {
	res, err := c.RpcClient.GetMinimumBalanceForRentExemption(ctx, dataLen)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

// GetBlockTime returns the estimated production time of a block.
func (c *Client) GetBlockTime(ctx context.Context, slot uint64) (int64, error) {
	res, err := c.RpcClient.GetBlockTime(ctx, slot)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

// GetIdentity returns the identity pubkey for the current node
func (c *Client) GetIdentity(ctx context.Context) (string, error) {
	res, err := c.RpcClient.GetIdentity(ctx)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return "", err
	}
	return res.Result.Identity, nil
}

// GetGenesisHash returns the genesis hash
func (c *Client) GetGenesisHash(ctx context.Context) (string, error) {
	res, err := c.RpcClient.GetGenesisHash(ctx)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// GetFirstAvailableBlock returns the slot of the lowest confirmed block that has not been purged from the ledger
func (c *Client) GetFirstAvailableBlock(ctx context.Context) (uint64, error) {
	res, err := c.RpcClient.GetFirstAvailableBlock(ctx)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

// GetVersion returns the current solana versions running on the node
func (c *Client) GetVersion(ctx context.Context) (rpc.GetVersionResult, error) {
	res, err := c.RpcClient.GetVersion(ctx)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return rpc.GetVersionResult{}, err
	}
	return res.Result, nil
}

// RequestAirdrop requests an airdrop of lamports to a Pubkey
func (c *Client) RequestAirdrop(ctx context.Context, base58Addr string, lamports uint64) (string, error) {
	res, err := c.RpcClient.RequestAirdrop(ctx, base58Addr, lamports)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// MinimumLedgerSlot returns the lowest slot that the node has information about in its ledger.
// This value may increase over time if the node is configured to purge older ledger data
func (c *Client) MinimumLedgerSlot(ctx context.Context) (uint64, error) {
	res, err := c.RpcClient.MinimumLedgerSlot(ctx)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

// GetTransactionCount returns the current Transaction count from the ledger
func (c *Client) GetTransactionCount(ctx context.Context) (uint64, error) {
	res, err := c.RpcClient.GetTransactionCount(ctx)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

// GetTransactionCountWithConfig returns the current Transaction count from the ledger
func (c *Client) GetTransactionCountWithConfig(ctx context.Context, cfg rpc.GetTransactionCountConfig) (uint64, error) {
	res, err := c.RpcClient.GetTransactionCountWithConfig(ctx, cfg)
	err = checkRpcResult(res.GeneralResponse, err)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}

func checkRpcResult(res rpc.GeneralResponse, err error) error {
	if err != nil {
		return err
	}
	if res.Error != nil {
		errRes, err := json.Marshal(res.Error)
		if err != nil {
			return fmt.Errorf("rpc response error: %v", res.Error)
		}
		return fmt.Errorf("rpc response error: %v", string(errRes))
	}
	return nil
}
