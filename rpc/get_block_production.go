package rpc

import (
	"context"
)

// GetBlockProductionResponse is a rpc response of `getBlockProduction`
type GetBlockProductionResponse struct {
	GeneralResponse
	Result GetBlockProductionResponseResult `json:"result"`
}

// GetBlockProductionResponseResult is a part of raw rpc response of `getBlockProduction`
type GetBlockProductionResponseResult struct {
	Context Context                               `json:"context"`
	Value   GetBlockProductionResponseResultValue `json:"value"`
}

type GetBlockProductionResponseResultValue struct {
	ByIdentity map[string][]uint64     `json:"byIdentity"`
	Range      GetBlockProductionRange `json:"range"`
}

// GetBlockProductionConfig is a option config for `getBlockProduction`
type GetBlockProductionConfig struct {
	Commitment Commitment               `json:"commitment,omitempty"`
	Range      *GetBlockProductionRange `json:"range,omitempty"`
	Identity   string                   `json:"identity,omitempty"`
}

type GetBlockProductionRange struct {
	FirstSlot uint64 `json:"firstSlot"`
	LastSlot  uint64 `json:"lastSlot,omitempty"`
}

// GetBlockProduction returns the current block height of the node
func (c *RpcClient) GetBlockProduction(ctx context.Context) (GetBlockProductionResponse, error) {
	return c.processGetBlockProduction(c.Call(ctx, "getBlockProduction"))
}

// GetBlockProductionWithConfig returns the current block height of the node
func (c *RpcClient) GetBlockProductionWithConfig(ctx context.Context, cfg GetBlockProductionConfig) (GetBlockProductionResponse, error) {
	return c.processGetBlockProduction(c.Call(ctx, "getBlockProduction", cfg))
}

func (c *RpcClient) processGetBlockProduction(body []byte, rpcErr error) (res GetBlockProductionResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
