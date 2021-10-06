package rpc

import (
	"context"
)

// RequestAirdropResponse is a full raw rpc response of `requestAirdrop`
type RequestAirdropResponse struct {
	GeneralResponse
	Result string `json:"result"`
}

// RequestAirdropConfig is a option config for `requestAirdrop`
type RequestAirdropConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
}

// RequestAirdrop requests an airdrop of lamports to a Pubkey
func (c *RpcClient) RequestAirdrop(ctx context.Context, base58Addr string, lamports uint64) (RequestAirdropResponse, error) {
	return c.processRequestAirdrop(c.Call(ctx, "requestAirdrop", base58Addr, lamports))
}

// RequestAirdropWithConfig requests an airdrop of lamports to a Pubkey
func (c *RpcClient) RequestAirdropWithConfig(ctx context.Context, base58Addr string, lamports uint64, cfg RequestAirdropConfig) (RequestAirdropResponse, error) {
	return c.processRequestAirdrop(c.Call(ctx, "requestAirdrop", base58Addr, lamports, cfg))
}

func (c *RpcClient) processRequestAirdrop(body []byte, rpcErr error) (res RequestAirdropResponse, err error) {
	err = c.processRpcCall(body, rpcErr, &res)
	return
}
