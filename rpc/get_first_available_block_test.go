package rpc

import (
	"context"
	"testing"
)

func TestGetFirstAvailableBlock(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFirstAvailableBlock"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":0,"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetFirstAvailableBlock(
					context.TODO(),
				)
			},
			ExpectedResponse: JsonRpcResponse[uint64]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result:  0,
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
