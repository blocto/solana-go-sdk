package rpc

import (
	"context"
)

type GetBlockCommitmentResponse JsonRpcResponse[GetBlockCommitment]

// GetBlockCommitmentResult is a part of raw rpc response of `getBlockCommitment`
type GetBlockCommitment struct {
	Commitment *[]uint64 `json:"commitment"`
	TotalStake uint64    `json:"totalStake"`
}

// GetBlockCommitment returns commitment for particular block
func (c *RpcClient) GetBlockCommitment(ctx context.Context, slot uint64) (JsonRpcResponse[GetBlockCommitment], error) {
	return c.processGetBlockCommitment(c.Call(ctx, "getBlockCommitment", slot))
}

func (c *RpcClient) processGetBlockCommitment(body []byte, rpcErr error) (res JsonRpcResponse[GetBlockCommitment], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
