package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/rpc"
)

type ClusterNode struct {
	Pubkey       common.PublicKey
	Gossip       *string
	Tpu          *string
	Rpc          *string
	Version      *string
	FeatureSet   *uint32
	ShredVersion *uint16
}

// GetClusterNodes returns information about all the nodes participating in the cluster
func (c *Client) GetClusterNodes(ctx context.Context) ([]ClusterNode, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.GetClusterNodes], error) {
			return c.RpcClient.GetClusterNodes(ctx)
		},
		convertGetClusterNodes,
	)
}

func convertGetClusterNodes(v rpc.GetClusterNodes) ([]ClusterNode, error) {
	output := make([]ClusterNode, 0, len(v))
	for _, info := range v {
		output = append(output, ClusterNode{
			Pubkey:       common.PublicKeyFromString(info.Pubkey),
			Gossip:       info.Gossip,
			Tpu:          info.Tpu,
			Rpc:          info.Rpc,
			Version:      info.Version,
			FeatureSet:   info.FeatureSet,
			ShredVersion: info.ShredVersion,
		})
	}
	return output, nil
}
