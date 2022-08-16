package rpc

import (
	"context"
)

type RequestAirdropResponse JsonRpcResponse[string]

// RequestAirdropConfig is a option config for `requestAirdrop`
type RequestAirdropConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// RequestAirdrop requests an airdrop of lamports to a Pubkey
func (c *RpcClient) RequestAirdrop(ctx context.Context, base58Addr string, lamports uint64) (JsonRpcResponse[string], error) {
	return c.processRequestAirdrop(c.Call(ctx, "requestAirdrop", base58Addr, lamports))
}

// RequestAirdropWithConfig requests an airdrop of lamports to a Pubkey
func (c *RpcClient) RequestAirdropWithConfig(ctx context.Context, base58Addr string, lamports uint64, cfg RequestAirdropConfig) (JsonRpcResponse[string], error) {
	return c.processRequestAirdrop(c.Call(ctx, "requestAirdrop", base58Addr, lamports, cfg))
}

func (c *RpcClient) processRequestAirdrop(body []byte, rpcErr error) (res JsonRpcResponse[string], err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
