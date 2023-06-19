package client

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/rpc"
)

func process[A any, B any](fetch func() (rpc.JsonRpcResponse[A], error), convert func(A) (B, error)) (B, error) {
	var output B
	res, err := fetch()
	if err != nil {
		return output, err
	}
	if err = res.GetError(); err != nil {
		return output, err
	}
	return convert(res.GetResult())
}

func value[T any](v rpc.ValueWithContext[T]) (T, error) {
	return v.Value, nil
}

func forward[T any](v T) (T, error) {
	return v, nil
}

type TokenAmount struct {
	Amount         uint64
	Decimals       uint8
	UIAmountString string
}

func newTokenAmount(amount string, decimals uint8, uiAmountString string) (TokenAmount, error) {
	u64Amount, err := strconv.ParseUint(amount, 10, 64)
	if err != nil {
		return TokenAmount{}, fmt.Errorf("failed to convert amount to u64")
	}
	return TokenAmount{
		Amount:         u64Amount,
		Decimals:       decimals,
		UIAmountString: uiAmountString,
	}, nil
}

type ReturnData struct {
	ProgramId common.PublicKey
	Data      []byte
}

func convertReturnData(d rpc.ReturnData) (ReturnData, error) {
	programId := common.PublicKeyFromString(d.ProgramId)
	s, ok := d.Data.([]any)
	if !ok {
		return ReturnData{}, fmt.Errorf("failed to get data")
	}
	if len(s) != 2 {
		return ReturnData{}, fmt.Errorf("unexpected slice lentgh")
	}
	if s[1].(string) != "base64" {
		return ReturnData{}, fmt.Errorf("unexpected encoding method")
	}
	data, err := base64.StdEncoding.DecodeString(s[0].(string))
	if err != nil {
		return ReturnData{}, fmt.Errorf("failed to decode data")
	}

	return ReturnData{
		ProgramId: programId,
		Data:      data,
	}, nil
}

type Reward struct {
	Pubkey       common.PublicKey `json:"pubkey"`
	Lamports     int64            `json:"lamports"`
	PostBalances uint64           `json:"postBalance"`
	RewardType   *rpc.RewardType  `json:"rewardType"`
	Commission   *uint8           `json:"commission"`
}

func convertReward(r rpc.Reward) Reward {
	return Reward{
		Pubkey:       common.PublicKeyFromString(r.Pubkey),
		Lamports:     r.Lamports,
		PostBalances: r.PostBalances,
		RewardType:   r.RewardType,
		Commission:   r.Commission,
	}
}

func convertRewards(rs []rpc.Reward) []Reward {
	output := make([]Reward, 0, len(rs))
	for _, r := range rs {
		output = append(output, convertReward(r))
	}
	return output
}
