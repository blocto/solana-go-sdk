package client

import (
	"context"
	"fmt"
	"time"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

type GetBlockConfig struct {
	Commitment         rpc.Commitment
	TransactionDetails rpc.GetBlockConfigTransactionDetails
	Rewards            *bool
}

func (c GetBlockConfig) toRpc() rpc.GetBlockConfig {
	return rpc.GetBlockConfig{
		Commitment:                     c.Commitment,
		TransactionDetails:             c.TransactionDetails,
		Rewards:                        c.Rewards,
		Encoding:                       rpc.GetBlockConfigEncodingBase64,
		MaxSupportedTransactionVersion: pointer.Get[uint8](0),
	}
}

type Block struct {
	Blockhash         string
	BlockTime         *time.Time
	BlockHeight       *int64
	PreviousBlockhash string
	ParentSlot        uint64
	Transactions      []BlockTransaction
	Signatures        []string
	Rewards           []Reward
}

type BlockTransaction struct {
	// rpc fields
	Meta        *TransactionMeta
	Transaction types.Transaction

	// custom fileds
	AccountKeys []common.PublicKey
}

func (c *Client) GetBlock(ctx context.Context, slot uint64) (*Block, error) {
	return process(
		func() (rpc.JsonRpcResponse[*rpc.GetBlock], error) {
			return c.RpcClient.GetBlockWithConfig(ctx, slot, GetBlockConfig{}.toRpc())
		},
		convertBlock,
	)
}

func (c *Client) GetBlockWithConfig(ctx context.Context, slot uint64, cfg GetBlockConfig) (*Block, error) {
	return process(
		func() (rpc.JsonRpcResponse[*rpc.GetBlock], error) {
			return c.RpcClient.GetBlockWithConfig(ctx, slot, cfg.toRpc())
		},
		convertBlock,
	)
}

func convertBlock(v *rpc.GetBlock) (*Block, error) {
	if v == nil {
		return nil, nil
	}

	var blockTime *time.Time
	if v.BlockTime != nil {
		t := time.Unix(*v.BlockTime, 0)
		blockTime = &t
	}

	var txs []BlockTransaction
	if len(v.Transactions) > 0 {
		txs = make([]BlockTransaction, 0, len(v.Transactions))
		for _, vtx := range v.Transactions {
			transactionMeta, err := convertTransactionMeta(vtx.Meta)
			if err != nil {
				return nil, fmt.Errorf("failed to convert transaction meta, err: %v", err)
			}

			tx, accountKeys, err := parseBase64Tx(vtx.Transaction, transactionMeta)
			if err != nil {
				return nil, fmt.Errorf("failed to parse tx, err: %v", err)
			}

			txs = append(txs, BlockTransaction{
				Meta:        transactionMeta,
				Transaction: tx,
				AccountKeys: accountKeys,
			})
		}
	}

	var rewards []Reward
	if len(v.Rewards) > 0 {
		rewards = convertRewards(v.Rewards)
	}

	return &Block{
		Blockhash:         v.Blockhash,
		BlockTime:         blockTime,
		BlockHeight:       v.BlockHeight,
		PreviousBlockhash: v.PreviousBlockhash,
		ParentSlot:        v.ParentSlot,
		Transactions:      txs,
		Signatures:        v.Signatures,
		Rewards:           rewards,
	}, nil
}
