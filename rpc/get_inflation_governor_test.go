package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetInflationGovernor(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getInflationGovernor"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"foundation":0.05,"foundationTerm":7.0,"initial":0.08,"taper":0.15,"terminal":0.015},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetInflationGovernor(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[GetInflationGovernor]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetInflationGovernor{
						Foundation:     0.05,
						FoundationTerm: 7.0,
						Initial:        0.08,
						Taper:          0.15,
						Terminal:       0.015,
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getInflationGovernor", "params":[{"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"foundation":0.05,"foundationTerm":7.0,"initial":0.08,"taper":0.15,"terminal":0.015},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetInflationGovernorWithConfig(
						context.TODO(),
						GetInflationGovernorConfig{
							Commitment: CommitmentProcessed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetInflationGovernor]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetInflationGovernor{
						Foundation:     0.05,
						FoundationTerm: 7.0,
						Initial:        0.08,
						Taper:          0.15,
						Terminal:       0.015,
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
