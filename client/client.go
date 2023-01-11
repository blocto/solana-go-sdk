package client

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/mr-tron/base58"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/token"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

type Client struct {
	RpcClient rpc.RpcClient
}

func New(opts ...rpc.Option) *Client {
	return &Client{
		RpcClient: rpc.New(opts...),
	}
}

func NewClient(endpoint string) *Client {
	return &Client{rpc.New(rpc.WithEndpoint(endpoint))}
}

func (c *Client) rpcAccountInfoToClientAccountInfo(v rpc.AccountInfo) (AccountInfo, error) {
	if v == (rpc.AccountInfo{}) {
		return AccountInfo{}, nil
	}

	data, ok := v.Data.([]any)
	if !ok {
		return AccountInfo{}, fmt.Errorf("failed to cast raw response to []any")
	}
	if data[1] != string(rpc.AccountEncodingBase64) {
		return AccountInfo{}, fmt.Errorf("encoding mistmatch")
	}
	rawData, err := base64.StdEncoding.DecodeString(data[0].(string))
	if err != nil {
		return AccountInfo{}, fmt.Errorf("failed to base64 decode data")
	}
	return AccountInfo{
		Lamports:   v.Lamports,
		Owner:      common.PublicKeyFromString(v.Owner),
		Executable: v.Executable,
		RentEpoch:  v.RentEpoch,
		Data:       rawData,
	}, nil
}

type QuickSendTransactionParam struct {
	Instructions []types.Instruction
	Signers      []types.Account
	FeePayer     common.PublicKey
}

// Deprecated: please use sendTransaction
// QuickSendTransaction is a quick way to send tx
func (c *Client) QuickSendTransaction(ctx context.Context, param QuickSendTransactionParam) (string, error) {
	recentBlockhashRes, err := c.GetLatestBlockhash(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get recent blockhash, err: %v", err)
	}
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			Instructions:    param.Instructions,
			FeePayer:        param.FeePayer,
			RecentBlockhash: recentBlockhashRes.Blockhash,
		}),
		Signers: param.Signers,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create new tx, err: %v", err)
	}
	rawTx, err := tx.Serialize()
	if err != nil {
		return "", fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	res, err := c.RpcClient.SendTransactionWithConfig(
		ctx,
		base64.StdEncoding.EncodeToString(rawTx),
		rpc.SendTransactionConfig{Encoding: rpc.SendTransactionConfigEncodingBase64},
	)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return "", err
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
	Err                  any
	Fee                  uint64
	PreBalances          []int64
	PostBalances         []int64
	PreTokenBalances     []rpc.TransactionMetaTokenBalance
	PostTokenBalances    []rpc.TransactionMetaTokenBalance
	LogMessages          []string
	InnerInstructions    []TransactionMetaInnerInstruction
	LoadedAddresses      rpc.TransactionLoadedAddresses
	ReturnData           *ReturnData
	ComputeUnitsConsumed *uint64
}

type TransactionMetaInnerInstruction struct {
	Index        uint64
	Instructions []types.CompiledInstruction
}

// GetTransaction returns transaction details for a confirmed transaction
func (c *Client) GetTransaction(ctx context.Context, txhash string) (*GetTransactionResponse, error) {
	res, err := c.RpcClient.GetTransactionWithConfig(
		ctx,
		txhash,
		rpc.GetTransactionConfig{
			Encoding: rpc.TransactionEncodingBase64,
		},
	)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	if res.Result == nil {
		return nil, nil
	}
	tx, err := getTransaction(res)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

// GetTransactionWithConfig returns transaction details for a confirmed transaction
// will ignore encoding
func (c *Client) GetTransactionWithConfig(ctx context.Context, txhash string, cfg rpc.GetTransactionConfig) (*GetTransactionResponse, error) {
	res, err := c.RpcClient.GetTransactionWithConfig(
		ctx,
		txhash,
		rpc.GetTransactionConfig{
			Encoding:                       rpc.TransactionEncodingBase64,
			Commitment:                     cfg.Commitment,
			MaxSupportedTransactionVersion: cfg.MaxSupportedTransactionVersion,
		},
	)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	if res.Result == nil {
		return nil, nil
	}
	tx, err := getTransaction(res)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func getTransaction(res rpc.JsonRpcResponse[*rpc.GetTransaction]) (GetTransactionResponse, error) {
	data, ok := res.Result.Transaction.([]any)
	if !ok {
		return GetTransactionResponse{}, fmt.Errorf("failed to cast raw response to []any")
	}
	if data[1] != string(rpc.TransactionEncodingBase64) {
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

	transactionMeta, err := convertTransactionMeta(res.Result.Meta)
	if err != nil {
		return GetTransactionResponse{}, fmt.Errorf("failed to convert transaction meta, err: %v", err)
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
	ParentSlot        uint64
	Transactions      []GetBlockTransaction
	Rewards           []rpc.GetBlockReward
}

type GetBlockTransaction struct {
	Meta        *TransactionMeta
	Transaction types.Transaction
}

// GetBlockWithConfig returns identity and transaction information about a confirmed block in the ledger
func (c *Client) GetBlockWithConfig(ctx context.Context, slot uint64, cfg rpc.GetBlockConfig) (GetBlockResponse, error) {
	res, err := c.RpcClient.GetBlockWithConfig(ctx, slot, cfg)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return GetBlockResponse{}, err
	}
	return getBlock(res)
}

// GetBlock returns identity and transaction information about a confirmed block in the ledger
func (c *Client) GetBlock(ctx context.Context, slot uint64) (GetBlockResponse, error) {
	return c.GetBlockWithConfig(
		ctx,
		slot,
		rpc.GetBlockConfig{
			Encoding: rpc.GetBlockConfigEncodingBase64,
		},
	)
}

// add test and get block
func getBlock(res rpc.JsonRpcResponse[rpc.GetBlock]) (GetBlockResponse, error) {
	txs := make([]GetBlockTransaction, 0, len(res.Result.Transactions))
	for _, rTx := range res.Result.Transactions {
		data, ok := rTx.Transaction.([]any)
		if !ok {
			return GetBlockResponse{}, fmt.Errorf("failed to cast raw response to []any")
		}
		if data[1] != string(rpc.TransactionEncodingBase64) {
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

		transactionMeta, err := convertTransactionMeta(rTx.Meta)
		if err != nil {
			return GetBlockResponse{}, fmt.Errorf("failed to convert transaction meta, err: %v", err)
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
		ParentSlot:        res.Result.ParentSlot,
		Rewards:           res.Result.Rewards,
		Transactions:      txs,
	}, nil
}

func (c *Client) GetSignatureStatus(ctx context.Context, signature string) (*rpc.SignatureStatus, error) {
	res, err := c.RpcClient.GetSignatureStatuses(ctx, []string{signature})
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	return res.Result.Value[0], nil
}

func (c *Client) GetSignatureStatusWithConfig(ctx context.Context, signature string, cfg rpc.GetSignatureStatusesConfig) (*rpc.SignatureStatus, error) {
	res, err := c.RpcClient.GetSignatureStatusesWithConfig(ctx, []string{signature}, cfg)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	return res.Result.Value[0], nil
}

func (c *Client) GetSignatureStatuses(ctx context.Context, signatures []string) (rpc.SignatureStatuses, error) {
	res, err := c.RpcClient.GetSignatureStatuses(ctx, signatures)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	return res.Result.Value, nil
}

func (c *Client) GetSignatureStatusesWithConfig(ctx context.Context, signatures []string, cfg rpc.GetSignatureStatusesConfig) (rpc.SignatureStatuses, error) {
	res, err := c.RpcClient.GetSignatureStatusesWithConfig(ctx, signatures, cfg)
	err = checkJsonRpcResponse(res, err)
	if err != nil {
		return nil, err
	}
	return res.Result.Value, nil
}

func checkJsonRpcResponse[T any](res rpc.JsonRpcResponse[T], err error) error {
	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *Client) GetTokenAccountsByOwner(ctx context.Context, base58Addr string) (map[common.PublicKey]token.TokenAccount, error) {
	getTokenAccountsByOwnerResponse, err := c.RpcClient.GetTokenAccountsByOwnerWithConfig(
		ctx,
		base58Addr,
		rpc.GetTokenAccountsByOwnerConfigFilter{
			ProgramId: common.TokenProgramID.ToBase58(),
		},
		rpc.GetTokenAccountsByOwnerConfig{
			Encoding: rpc.AccountEncodingBase64,
		},
	)
	if err != nil {
		return nil, err
	}

	m := map[common.PublicKey]token.TokenAccount{}
	for _, v := range getTokenAccountsByOwnerResponse.Result.Value {
		accountInfo, err := c.rpcAccountInfoToClientAccountInfo(v.Account)
		if err != nil {
			return nil, err
		}
		tokenAccount, err := token.DeserializeTokenAccount(accountInfo.Data, accountInfo.Owner)
		if err != nil {
			return nil, err
		}
		m[common.PublicKeyFromString(v.Pubkey)] = tokenAccount
	}
	return m, err
}

// helper function
func convertToIntSlice(input []any) ([]int, error) {
	output := make([]int, 0, len(input))
	for _, v := range input {
		output = append(output, int(v.(float64)))
	}
	return output, nil
}

func convertTransactionMeta(meta *rpc.TransactionMeta) (*TransactionMeta, error) {
	if meta == nil {
		return nil, nil
	}

	innerInstructions := make([]TransactionMetaInnerInstruction, 0, len(meta.InnerInstructions))
	for _, metaInnerInstruction := range meta.InnerInstructions {
		compiledInstructions := make([]types.CompiledInstruction, 0, len(metaInnerInstruction.Instructions))
		for _, innerInstruction := range metaInnerInstruction.Instructions {
			parsedInstruction, ok := innerInstruction.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("failed to convert inner instruction type. value: %v", innerInstruction)
			}

			accounts, err := convertToIntSlice(parsedInstruction["accounts"].([]any))
			if err != nil {
				return nil, fmt.Errorf("failed to cast instructions accounts, err: %v", err)
			}

			var data []byte
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

		innerInstructions = append(innerInstructions, TransactionMetaInnerInstruction{
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
