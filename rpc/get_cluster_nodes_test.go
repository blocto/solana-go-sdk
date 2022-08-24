package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/pkg/pointer"
)

func TestGetClusterNodes(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getClusterNodes"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":[{"featureSet":1797267350,"gossip":"127.0.0.1:1024","pubkey":"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn","rpc":"127.0.0.1:8899","shredVersion":23492,"tpu":"127.0.0.1:1027","version":"1.8.1"}],"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetClusterNodes(
					context.TODO(),
				)
			},
			ExpectedResponse: JsonRpcResponse[GetClusterNodes]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetClusterNodes{
					{
						Pubkey:       "8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn",
						Gossip:       pointer.Get("127.0.0.1:1024"),
						Tpu:          pointer.Get("127.0.0.1:1027"),
						Rpc:          pointer.Get("127.0.0.1:8899"),
						Version:      pointer.Get("1.8.1"),
						FeatureSet:   pointer.Get[uint32](1797267350),
						ShredVersion: pointer.Get[uint16](23492),
					},
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
