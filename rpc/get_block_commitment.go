package rpc

import (
	"context"
)

// GetBlockCommitmentResponse is a full raw rpc response of `getBlockCommitment`
type GetBlockCommitmentResponse struct {
	GeneralResponse
	Result GetBlockCommitmentResult `json:"result"`
}

// GetBlockCommitmentResult is a part of raw rpc response of `getBlockCommitment`
type GetBlockCommitmentResult struct {
	Commitment *[]uint64 `json:"commitment"`
	TotalStake uint64    `json:"totalStake"`
}

// GetBlockCommitment returns commitment for particular block
func (c *RpcClient) GetBlockCommitment(ctx context.Context, slot uint64) (GetBlockCommitmentResponse, error) {
	return c.processGetBlockCommitment(c.Call(ctx, "getBlockCommitment", slot))
}

func (c *RpcClient) processGetBlockCommitment(body []byte, rpcErr error) (res GetBlockCommitmentResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
