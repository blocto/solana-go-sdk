package rpc

import "context"

// GetClusterNodesResponse is a full raw rpc response of `GetClusterNodes`
type GetClusterNodesResponse struct {
	GeneralResponse
	Result []GetClusterNodesResponseResult `json:"result"`
}

type GetClusterNodesResponseResult struct {
	Pubkey       string
	Gossip       *string
	Tpu          *string
	Rpc          *string
	Version      *string
	FeatureSet   *uint32
	ShredVersion *uint16
}

// GetClusterNodes returns information about all the nodes participating in the cluster
func (c *RpcClient) GetClusterNodes(ctx context.Context) (GetClusterNodesResponse, error) {
	return c.processGetClusterNodes(c.Call(ctx, "getClusterNodes"))
}

func (c *RpcClient) processGetClusterNodes(body []byte, rpcErr error) (res GetClusterNodesResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
