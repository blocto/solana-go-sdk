package rpc

import (
	"context"
	"testing"
)

func TestGetIdentity(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getIdentity"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"identity":"BjHeMczor9oycGJHLepRTCU2LpkZNtpy2mdQKianx1EJ"},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetIdentity(
					context.TODO(),
				)
			},
			ExpectedResponse: GetIdentityResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetIdentityResult{
					Identity: "BjHeMczor9oycGJHLepRTCU2LpkZNtpy2mdQKianx1EJ",
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
