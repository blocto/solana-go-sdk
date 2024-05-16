package client

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/program/token"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

// Ensure Client implements the ClientInterface
var _ ClientInterface = (*Client)(nil)

// ClientInterface is an interface for various clients implementations and mocks generation
type ClientInterface interface {
	QuickSendTransaction(ctx context.Context, param QuickSendTransactionParam) (string, error)
	GetAccountInfo(ctx context.Context, base58Addr string) (AccountInfo, error)
	GetAccountInfoWithConfig(ctx context.Context, base58Addr string, cfg GetAccountInfoConfig) (AccountInfo, error)
	GetAccountInfoAndContext(ctx context.Context, base58Addr string) (rpc.ValueWithContext[AccountInfo], error)
	GetAccountInfoAndContextWithConfig(ctx context.Context, base58Addr string, cfg GetAccountInfoConfig) (rpc.ValueWithContext[AccountInfo], error)
	GetBalance(ctx context.Context, base58Addr string) (uint64, error)
	GetBalanceWithConfig(ctx context.Context, base58Addr string, cfg GetBalanceConfig) (uint64, error)
	GetBalanceAndContext(ctx context.Context, base58Addr string) (rpc.ValueWithContext[uint64], error)
	GetBalanceAndContextWithConfig(ctx context.Context, base58Addr string, cfg GetBalanceConfig) (rpc.ValueWithContext[uint64], error)
	GetBlockTime(ctx context.Context, slot uint64) (*int64, error)
	GetBlock(ctx context.Context, slot uint64) (*Block, error)
	GetBlockWithConfig(ctx context.Context, slot uint64, cfg GetBlockConfig) (*Block, error)
	GetClusterNodes(ctx context.Context) ([]ClusterNode, error)
	GetEpochInfo(ctx context.Context) (GetEpochInfo, error)
	GetFeeForMessage(ctx context.Context, message types.Message) (*uint64, error)
	GetFeeForMessageWithConfig(ctx context.Context, message types.Message, cfg GetFeeForMessageConfig) (*uint64, error)
	GetFeeForMessageAndContext(ctx context.Context, message types.Message) (rpc.ValueWithContext[*uint64], error)
	GetFeeForMessageAndContextWithConfig(ctx context.Context, message types.Message, cfg GetFeeForMessageConfig) (rpc.ValueWithContext[*uint64], error)
	GetFirstAvailableBlock(ctx context.Context) (uint64, error)
	GetGenesisHash(ctx context.Context) (string, error)
	GetHealth(ctx context.Context) (bool, error)
	GetIdentity(ctx context.Context) (rpc.GetIdentity, error)
	GetLatestBlockhash(ctx context.Context) (rpc.GetLatestBlockhashValue, error)
	GetLatestBlockhashWithConfig(ctx context.Context, cfg GetLatestBlockhashConfig) (rpc.GetLatestBlockhashValue, error)
	GetLatestBlockhashAndContext(ctx context.Context) (rpc.ValueWithContext[rpc.GetLatestBlockhashValue], error)
	GetLatestBlockhashAndContextWithConfig(ctx context.Context, cfg GetLatestBlockhashConfig) (rpc.ValueWithContext[rpc.GetLatestBlockhashValue], error)
	GetMinimumBalanceForRentExemption(ctx context.Context, dataLen uint64) (uint64, error)
	GetMinimumBalanceForRentExemptionWithConfig(ctx context.Context, dataLen uint64, cfg GetMinimumBalanceForRentExemptionConfig) (uint64, error)
	GetMultipleAccounts(ctx context.Context, addrs []string) ([]AccountInfo, error)
	GetMultipleAccountsWithConfig(ctx context.Context, addrs []string, cfg GetMultipleAccountsConfig) ([]AccountInfo, error)
	GetMultipleAccountsAndContext(ctx context.Context, addrs []string) (rpc.ValueWithContext[[]AccountInfo], error)
	GetMultipleAccountsAndContextWithConfig(ctx context.Context, addrs []string, cfg GetMultipleAccountsConfig) (rpc.ValueWithContext[[]AccountInfo], error)
	GetRecentPrioritizationFees(ctx context.Context, addresses []common.PublicKey) (rpc.PrioritizationFees, error)
	GetSignatureStatus(ctx context.Context, signature string) (*rpc.SignatureStatus, error)
	GetSignatureStatusWithConfig(ctx context.Context, signature string, cfg GetSignatureStatusesConfig) (*rpc.SignatureStatus, error)
	GetSignatureStatuses(ctx context.Context, signatures []string) (rpc.SignatureStatuses, error)
	GetSignatureStatusesWithConfig(ctx context.Context, signatures []string, cfg GetSignatureStatusesConfig) (rpc.SignatureStatuses, error)
	GetSignaturesForAddress(ctx context.Context, addr string) (rpc.GetSignaturesForAddress, error)
	GetSignaturesForAddressWithConfig(ctx context.Context, addr string, cfg GetSignaturesForAddressConfig) (rpc.GetSignaturesForAddress, error)
	GetSlot(ctx context.Context) (uint64, error)
	GetSlotWithConfig(ctx context.Context, cfg GetSlotConfig) (uint64, error)
	GetTokenAccountBalance(ctx context.Context, addr string) (TokenAmount, error)
	GetTokenAccountBalanceWithConfig(ctx context.Context, addr string, cfg GetTokenAccountBalanceConfig) (TokenAmount, error)
	GetTokenAccountBalanceAndContext(ctx context.Context, addr string) (rpc.ValueWithContext[TokenAmount], error)
	GetTokenAccountBalanceAndContextWithConfig(ctx context.Context, addr string, cfg GetTokenAccountBalanceConfig) (rpc.ValueWithContext[TokenAmount], error)
	GetTokenAccountsByOwnerByMint(ctx context.Context, owner, mintAddr string) ([]TokenAccount, error)
	GetTokenAccountsByOwnerByProgram(ctx context.Context, owner, programId string) ([]TokenAccount, error)
	GetTokenAccountsByOwnerWithContextByMint(ctx context.Context, owner, mintAddr string) (rpc.ValueWithContext[[]TokenAccount], error)
	GetTokenAccountsByOwnerWithContextByProgram(ctx context.Context, owner, programId string) (rpc.ValueWithContext[[]TokenAccount], error)
	GetTokenSupply(ctx context.Context, mintAddr string) (TokenAmount, error)
	GetTokenSupplyWithConfig(ctx context.Context, mintAddr string, cfg GetTokenSupplyConfig) (TokenAmount, error)
	GetTokenSupplyAndContext(ctx context.Context, mintAddr string) (rpc.ValueWithContext[TokenAmount], error)
	GetTokenSupplyAndContextWithConfig(ctx context.Context, mintAddr string, cfg GetTokenSupplyConfig) (rpc.ValueWithContext[TokenAmount], error)
	GetTransactionCount(ctx context.Context) (uint64, error)
	GetTransactionCountWithConfig(ctx context.Context, cfg GetTransactionCountConfig) (uint64, error)
	GetTransaction(ctx context.Context, txhash string) (*Transaction, error)
	GetTransactionWithConfig(ctx context.Context, txhash string, cfg GetTransactionConfig) (*Transaction, error)
	GetVersion(ctx context.Context) (rpc.GetVersion, error)
	GetVoteAccounts(ctx context.Context) (VoteAccountStatus, error)
	IsBlockhashValid(ctx context.Context, blockhash string) (bool, error)
	IsBlockhashValidWithConfig(ctx context.Context, blockhash string, cfg IsBlockhashValidConfig) (bool, error)
	IsBlockhashValidAndContext(ctx context.Context, blockhash string) (rpc.ValueWithContext[bool], error)
	IsBlockhashValidAndContextWithConfig(ctx context.Context, blockhash string, cfg IsBlockhashValidConfig) (rpc.ValueWithContext[bool], error)
	MinimumLedgerSlot(ctx context.Context) (uint64, error)
	RequestAirdrop(ctx context.Context, base58Addr string, lamports uint64) (string, error)
	RequestAirdropWithConfig(ctx context.Context, base58Addr string, lamports uint64, cfg RequestAirdropConfig) (string, error)
	SendTransaction(ctx context.Context, tx types.Transaction) (string, error)
	SendTransactionWithConfig(ctx context.Context, tx types.Transaction, cfg SendTransactionConfig) (string, error)
	SimulateTransaction(ctx context.Context, tx types.Transaction) (SimulateTransaction, error)
	SimulateTransactionWithConfig(ctx context.Context, tx types.Transaction, cfg SimulateTransactionConfig) (SimulateTransaction, error)
	SimulateTransactionAndContext(ctx context.Context, tx types.Transaction) (rpc.ValueWithContext[SimulateTransaction], error)
	SimulateTransactionAndContextWithConfig(ctx context.Context, tx types.Transaction, cfg SimulateTransactionConfig) (rpc.ValueWithContext[SimulateTransaction], error)
	GetNonceAccount(ctx context.Context, base58Addr string) (system.NonceAccount, error)
	GetNonceFromNonceAccount(ctx context.Context, base58Addr string) (string, error)
	GetTokenAccount(ctx context.Context, base58Addr string) (token.TokenAccount, error)
}

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

func checkJsonRpcResponse[T any](res rpc.JsonRpcResponse[T], err error) error {
	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}
