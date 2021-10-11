package rpc

import "context"

// GetInflationRateResponse is a full raw rpc response of `getInflationRate`
type GetInflationRateResponse struct {
	GeneralResponse
	Result GetInflationRateResult `json:"result"`
}

// GetInflationRateResult is a part of raw rpc response of `getInflationRate`
type GetInflationRateResult struct {
	Epoch      uint64  `json:"epoch"`
	Foundation float64 `json:"foundation"`
	Total      float64 `json:"total"`
	Validator  float64 `json:"validator"`
}

// GetInflationRate returns the specific inflation values for the current epoch
func (c *RpcClient) GetInflationRate(ctx context.Context) (GetInflationRateResponse, error) {
	return c.processGetInflationRate(c.Call(ctx, "getInflationRate"))
}

func (c *RpcClient) processGetInflationRate(body []byte, rpcErr error) (res GetInflationRateResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
