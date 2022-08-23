package rpc

import "context"

type GetInflationRateResponse JsonRpcResponse[GetInflationRate]

type GetInflationRate struct {
	Epoch      uint64  `json:"epoch"`
	Foundation float64 `json:"foundation"`
	Total      float64 `json:"total"`
	Validator  float64 `json:"validator"`
}

// GetInflationRate returns the specific inflation values for the current epoch
func (c *RpcClient) GetInflationRate(ctx context.Context) (JsonRpcResponse[GetInflationRate], error) {
	return c.processGetInflationRate(c.Call(ctx, "getInflationRate"))
}

func (c *RpcClient) processGetInflationRate(body []byte, rpcErr error) (res JsonRpcResponse[GetInflationRate], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
