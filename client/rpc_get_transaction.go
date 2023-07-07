package client

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/mr-tron/base58"
)

type GetTransactionConfig struct {
	Commitment rpc.Commitment
}

func (c GetTransactionConfig) toRpc() rpc.GetTransactionConfig {
	return rpc.GetTransactionConfig{
		Commitment:                     c.Commitment,
		Encoding:                       rpc.TransactionEncodingBase64,
		MaxSupportedTransactionVersion: pointer.Get[uint8](0),
	}
}

type Transaction struct {
	// rpc
	Slot        uint64
	Meta        *TransactionMeta
	Transaction types.Transaction
	BlockTime   *int64

	// custom
	AccountKeys []common.PublicKey
}

func (t Transaction) Version() types.MessageVersion {
	return t.Transaction.Message.Version
}

type TransactionMeta struct {
	Err                  any
	Fee                  uint64
	PreBalances          []int64
	PostBalances         []int64
	PreTokenBalances     []rpc.TransactionMetaTokenBalance
	PostTokenBalances    []rpc.TransactionMetaTokenBalance
	LogMessages          []string
	InnerInstructions    []InnerInstruction
	LoadedAddresses      rpc.TransactionLoadedAddresses
	ReturnData           *ReturnData
	ComputeUnitsConsumed *uint64
}

type InnerInstruction struct {
	Index        uint64
	Instructions []types.CompiledInstruction
}

// GetTransaction returns transaction details for a confirmed transaction
func (c *Client) GetTransaction(ctx context.Context, txhash string) (*Transaction, error) {
	return process(
		func() (rpc.JsonRpcResponse[*rpc.GetTransaction], error) {
			return c.RpcClient.GetTransactionWithConfig(ctx, txhash, GetTransactionConfig{}.toRpc())
		},
		convertTransaction,
	)
}

// GetTransaction returns transaction details for a confirmed transaction
func (c *Client) GetTransactionWithConfig(ctx context.Context, txhash string, cfg GetTransactionConfig) (*Transaction, error) {
	return process(
		func() (rpc.JsonRpcResponse[*rpc.GetTransaction], error) {
			return c.RpcClient.GetTransactionWithConfig(ctx, txhash, cfg.toRpc())
		},
		convertTransaction,
	)
}

func convertTransaction(v *rpc.GetTransaction) (*Transaction, error) {
	if v == nil {
		return nil, nil
	}

	// transaction meta
	transactionMeta, err := convertTransactionMeta(v.Meta)
	if err != nil {
		return nil, fmt.Errorf("failed to convert transaction meta, err: %v", err)
	}

	tx, accountKeys, err := parseBase64Tx(v.Transaction, transactionMeta)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tx, err: %v", err)
	}

	return &Transaction{
		Slot:        v.Slot,
		BlockTime:   v.BlockTime,
		Transaction: tx,
		Meta:        transactionMeta,
		AccountKeys: accountKeys,
	}, nil
}

func convertTransactionMeta(meta *rpc.TransactionMeta) (*TransactionMeta, error) {
	if meta == nil {
		return nil, nil
	}

	innerInstructions := make([]InnerInstruction, 0, len(meta.InnerInstructions))
	for _, metaInnerInstruction := range meta.InnerInstructions {
		compiledInstructions := make([]types.CompiledInstruction, 0, len(metaInnerInstruction.Instructions))
		for _, innerInstruction := range metaInnerInstruction.Instructions {
			parsedInstruction, ok := innerInstruction.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("failed to convert inner instruction type. value: %v", innerInstruction)
			}

			rawAccounts, ok := parsedInstruction["accounts"].([]any)
			if !ok {
				return nil, fmt.Errorf("failed to parse instruction accounts")
			}
			accounts := make([]int, 0, len(rawAccounts))
			for _, v := range rawAccounts {
				accounts = append(accounts, int(v.(float64)))
			}

			var data []byte
			var err error
			if dataString := parsedInstruction["data"].(string); len(dataString) > 0 {
				data, err = base58.Decode(dataString)
				if err != nil {
					return nil, fmt.Errorf("failed to base58 decode data, data: %v, err: %v", parsedInstruction["data"], err)
				}
			}

			compiledInstructions = append(compiledInstructions, types.CompiledInstruction{
				ProgramIDIndex: int(parsedInstruction["programIdIndex"].(float64)),
				Accounts:       accounts,
				Data:           data,
			})
		}

		innerInstructions = append(innerInstructions, InnerInstruction{
			Index:        metaInnerInstruction.Index,
			Instructions: compiledInstructions,
		})
	}

	var returnData *ReturnData
	if v := meta.ReturnData; v != nil {
		d, err := convertReturnData(*v)
		if err != nil {
			return nil, fmt.Errorf("failed to process return data, err: %v", err)
		}
		returnData = &d
	}

	return &TransactionMeta{
		Err:                  meta.Err,
		Fee:                  meta.Fee,
		PreBalances:          meta.PreBalances,
		PostBalances:         meta.PostBalances,
		PreTokenBalances:     meta.PreTokenBalances,
		PostTokenBalances:    meta.PostTokenBalances,
		LogMessages:          meta.LogMessages,
		InnerInstructions:    innerInstructions,
		LoadedAddresses:      meta.LoadedAddresses,
		ReturnData:           returnData,
		ComputeUnitsConsumed: meta.ComputeUnitsConsumed,
	}, nil
}

func parseBase64Tx(raw any, transactionMeta *TransactionMeta) (types.Transaction, []common.PublicKey, error) {
	// transaction
	data, ok := raw.([]any)
	if !ok {
		return types.Transaction{}, nil, fmt.Errorf("failed to cast raw response to []any")
	}
	if data[1] != string(rpc.TransactionEncodingBase64) {
		return types.Transaction{}, nil, fmt.Errorf("encoding mistmatch")
	}
	rawTx, err := base64.StdEncoding.DecodeString(data[0].(string))
	if err != nil {
		return types.Transaction{}, nil, fmt.Errorf("failed to base64 decode data, err: %v", err)
	}
	tx, err := types.TransactionDeserialize(rawTx)
	if err != nil {
		return types.Transaction{}, nil, fmt.Errorf("failed to deserialize transaction, err: %v", err)
	}

	// account keys
	l := len(tx.Message.Accounts)
	for _, alt := range tx.Message.AddressLookupTables {
		l += (len(alt.WritableIndexes) + len(alt.ReadonlyIndexes))
	}
	accountKeys := make([]common.PublicKey, 0, l)
	accountKeys = append(accountKeys, tx.Message.Accounts...)
	if transactionMeta != nil {
		for _, s := range transactionMeta.LoadedAddresses.Writable {
			accountKeys = append(accountKeys, common.PublicKeyFromString(s))
		}
		for _, s := range transactionMeta.LoadedAddresses.Readonly {
			accountKeys = append(accountKeys, common.PublicKeyFromString(s))
		}
	}

	return tx, accountKeys, nil
}
