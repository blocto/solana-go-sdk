package rpc

import (
	"context"
	"testing"
)

func TestIsBlockhashValid(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params":["14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112890169},"value":false},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.IsBlockhashValid(context.TODO(), "14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N")
			},
			ExpectedResponse: JsonRpcResponse[IsBlockhashValid]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: IsBlockhashValid{
					Context: Context{
						Slot: 112890169,
					},
					Value: false,
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params":["14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N", {"commitment": "processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112890231},"value":true},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.IsBlockhashValidWithConfig(
					context.TODO(),
					"14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N",
					IsBlockhashValidConfig{
						Commitment: CommitmentProcessed,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[IsBlockhashValid]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: IsBlockhashValid{
					Context: Context{
						Slot: 112890231,
					},
					Value: true,
				},
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			testRpcCall(t, tt)
		})
	}
}
