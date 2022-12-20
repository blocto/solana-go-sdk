package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/pkg/pointer"
)

func TestGetInflationReward(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getInflationReward", "params":[["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", "BJhtevCiNKrWsc2pkJP1TFhxAhheZ9FNJ7F567FayhSD"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[null,{"amount":154995,"commission":0,"effectiveSlot":1120,"epoch":34,"postBalance":10003564885}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetInflationReward(
						context.TODO(),
						[]string{"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", "BJhtevCiNKrWsc2pkJP1TFhxAhheZ9FNJ7F567FayhSD"},
					)
				},
				ExpectedValue: JsonRpcResponse[[]*GetInflationReward]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: []*GetInflationReward{
						nil,
						{
							Epoch:         34,
							EffectiveSlot: 1120,
							Commission:    pointer.Get[uint8](0),
							Amount:        154995,
							PostBalance:   10003564885,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getInflationReward", "params":[["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", "BJhtevCiNKrWsc2pkJP1TFhxAhheZ9FNJ7F567FayhSD"], {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[null,{"amount":154995,"commission":0,"effectiveSlot":1152,"epoch":35,"postBalance":10003719880}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetInflationRewardWithConfig(
						context.TODO(),
						[]string{"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", "BJhtevCiNKrWsc2pkJP1TFhxAhheZ9FNJ7F567FayhSD"},
						GetInflationRewardConfig{
							Commitment: CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[[]*GetInflationReward]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: []*GetInflationReward{
						nil,
						{
							Epoch:         35,
							EffectiveSlot: 1152,
							Commission:    pointer.Get[uint8](0),
							Amount:        154995,
							PostBalance:   10003719880,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getInflationReward", "params":[["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", "BJhtevCiNKrWsc2pkJP1TFhxAhheZ9FNJ7F567FayhSD"], {"epoch": 31}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[null,{"amount":154995,"commission":0,"effectiveSlot":1024,"epoch":31,"postBalance":10003099900}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetInflationRewardWithConfig(
						context.TODO(),
						[]string{"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", "BJhtevCiNKrWsc2pkJP1TFhxAhheZ9FNJ7F567FayhSD"},
						GetInflationRewardConfig{
							Epoch: 31,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[[]*GetInflationReward]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: []*GetInflationReward{
						nil,
						{
							Epoch:         31,
							EffectiveSlot: 1024,
							Commission:    pointer.Get[uint8](0),
							Amount:        154995,
							PostBalance:   10003099900,
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
