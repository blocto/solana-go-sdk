package rpc

import "context"

// GetFeeRateGovernorResponse is a full raw rpc response of `getFeeRateGovernor`
type GetFeeRateGovernorResponse struct {
	GeneralResponse
	Result GetFeeRateGovernorResponseResult `json:"result"`
}

type GetFeeRateGovernorResponseResult struct {
	Context Context                               `json:"context"`
	Value   GetFeeRateGovernorResponseResultValue `json:"value"`
}

type GetFeeRateGovernorResponseResultValue struct {
	FeeRateGovernor FeeRateGovernor `json:"feeRateGovernor"`
}

type FeeRateGovernor struct {
	MaxLamportsPerSignature    uint64 `json:"maxLamportsPerSignature"`
	MinLamportsPerSignature    uint64 `json:"minLamportsPerSignature"`
	TargetLamportsPerSignature uint64 `json:"targetLamportsPerSignature"`
	TargetSignaturesPerSlot    uint64 `json:"targetSignaturesPerSlot"`
	BurnPercent                uint8  `json:"burnPercent"`
}

// DEPRECATED
// GetFeeRateGovernor returns the fee rate governor information from the root bank
func (c *RpcClient) GetFeeRateGovernor(ctx context.Context) (GetFeeRateGovernorResponse, error) {
	return c.processGetFeeRateGovernor(c.Call(ctx, "getFeeRateGovernor"))
}

func (c *RpcClient) processGetFeeRateGovernor(body []byte, rpcErr error) (res GetFeeRateGovernorResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
