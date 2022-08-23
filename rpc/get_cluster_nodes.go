package rpc

import "context"

type GetClusterNodesResponse JsonRpcResponse[GetClusterNodes]

type GetClusterNodes []GetClusterNode

type GetClusterNode struct {
	Pubkey       string
	Gossip       *string
	Tpu          *string
	Rpc          *string
	Version      *string
	FeatureSet   *uint32
	ShredVersion *uint16
}

// GetClusterNodes returns information about all the nodes participating in the cluster
func (c *RpcClient) GetClusterNodes(ctx context.Context) (JsonRpcResponse[GetClusterNodes], error) {
	return c.processGetClusterNodes(c.Call(ctx, "getClusterNodes"))
}

func (c *RpcClient) processGetClusterNodes(body []byte, rpcErr error) (res JsonRpcResponse[GetClusterNodes], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
