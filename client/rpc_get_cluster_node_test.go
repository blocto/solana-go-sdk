package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
)

func TestClient_GetClusterNodes(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getClusterNodes"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"featureSet":1797267350,"gossip":"127.0.0.1:1024","pubkey":"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn","rpc":"127.0.0.1:8899","shredVersion":23492,"tpu":"127.0.0.1:1027","version":"1.8.1"}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetClusterNodes(
						context.TODO(),
					)
				},
				ExpectedValue: []ClusterNode{
					{
						Pubkey:       common.PublicKeyFromString("8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn"),
						Gossip:       pointer.Get("127.0.0.1:1024"),
						Tpu:          pointer.Get("127.0.0.1:1027"),
						Rpc:          pointer.Get("127.0.0.1:8899"),
						Version:      pointer.Get("1.8.1"),
						FeatureSet:   pointer.Get[uint32](1797267350),
						ShredVersion: pointer.Get[uint16](23492),
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
