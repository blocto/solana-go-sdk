package rpc

import (
	"context"
	"testing"
)

func TestGetBlockCommitment(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockCommitment", "params":[86708800]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"commitment":null,"totalStake":156502861915805458},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockCommitment(
					context.TODO(),
					86708800,
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlockCommitment]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlockCommitment{
					Commitment: nil,
					TotalStake: 156502861915805458,
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockCommitment", "params":[86708895]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"commitment":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,140000814436952564],"totalStake":156502861915805458},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockCommitment(
					context.TODO(),
					86708895,
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlockCommitment]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlockCommitment{
					Commitment: &[]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 140000814436952564},
					TotalStake: 156502861915805458,
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
