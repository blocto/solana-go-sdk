package rpc

import (
	"context"
)

// GetFeeCalculatorForBlockhashResponse is a full raw rpc response of `getFeeCalculatorForBlockhash`
type GetFeeCalculatorForBlockhashResponse struct {
	GeneralResponse
	Result GetFeeCalculatorForBlockhashResponseResult `json:"result"`
}

// GetFeeCalculatorForBlockhashResult is a part of raw rpc response of `getFeeCalculatorForBlockhash`
type GetFeeCalculatorForBlockhashResponseResult struct {
	Context Context                                          `json:"context"`
	Value   *GetFeeCalculatorForBlockhashResponseResultValue `json:"value"`
}

type GetFeeCalculatorForBlockhashResponseResultValue struct {
	FeeCalculator FeeCalculator `json:"feeCalculator"`
}

// GetFeeCalculatorForBlockhashConfig is a option config for `getFeeCalculatorForBlockhash`
type GetFeeCalculatorForBlockhashConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// GetFeeCalculatorForBlockhash returns the SOL balance
func (c *RpcClient) GetFeeCalculatorForBlockhash(ctx context.Context, base58Addr string) (GetFeeCalculatorForBlockhashResponse, error) {
	return c.processGetFeeCalculatorForBlockhash(c.Call(ctx, "getFeeCalculatorForBlockhash", base58Addr))
}

// GetFeeCalculatorForBlockhashWithConfig returns the SOL balance
func (c *RpcClient) GetFeeCalculatorForBlockhashWithConfig(ctx context.Context, base58Addr string, cfg GetFeeCalculatorForBlockhashConfig) (GetFeeCalculatorForBlockhashResponse, error) {
	return c.processGetFeeCalculatorForBlockhash(c.Call(ctx, "getFeeCalculatorForBlockhash", base58Addr, cfg))
}

func (c *RpcClient) processGetFeeCalculatorForBlockhash(body []byte, rpcErr error) (res GetFeeCalculatorForBlockhashResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
