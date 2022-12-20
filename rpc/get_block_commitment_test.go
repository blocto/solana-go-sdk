package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetBlockCommitment(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockCommitment", "params":[86708800]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"commitment":null,"totalStake":156502861915805458},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockCommitment(
						context.TODO(),
						86708800,
					)
				},
				ExpectedValue: JsonRpcResponse[GetBlockCommitment]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetBlockCommitment{
						Commitment: nil,
						TotalStake: 156502861915805458,
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockCommitment", "params":[86708895]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"commitment":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,140000814436952564],"totalStake":156502861915805458},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockCommitment(
						context.TODO(),
						86708895,
					)
				},
				ExpectedValue: JsonRpcResponse[GetBlockCommitment]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetBlockCommitment{
						Commitment: &[]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 140000814436952564},
						TotalStake: 156502861915805458,
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
