package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/pkg/pointer"
)

func TestGetVersion(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVersion"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"feature-set":1824749018,"solana-core":"1.7.14"},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetVersion(
					context.TODO(),
				)
			},
			ExpectedResponse: JsonRpcResponse[GetVersion]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetVersion{
					SolanaCore: "1.7.14",
					FeatureSet: pointer.Uint32(1824749018),
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
