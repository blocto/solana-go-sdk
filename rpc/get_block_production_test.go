package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetBlockProduction(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockProduction"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":6850},"value":{"byIdentity":{"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn":[6851,6851]},"range":{"firstSlot":0,"lastSlot":6850}}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockProduction(
						context.TODO(),
					)
				},
				ExpectedValue: JsonRpcResponse[GetBlockProduction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetBlockProduction{
						Context: Context{
							Slot: 6850,
						},
						Value: GetBlockProductionResponseResultValue{
							ByIdentity: map[string][]uint64{
								"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn": {6851, 6851},
							},
							Range: GetBlockProductionRange{
								FirstSlot: 0,
								LastSlot:  6850,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockProduction", "params":[{"identity": "8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":6850},"value":{"byIdentity":{"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn":[6851,6851]},"range":{"firstSlot":0,"lastSlot":6850}}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockProductionWithConfig(
						context.TODO(),
						GetBlockProductionConfig{
							Identity: "8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn",
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetBlockProduction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetBlockProduction{
						Context: Context{
							Slot: 6850,
						},
						Value: GetBlockProductionResponseResultValue{
							ByIdentity: map[string][]uint64{
								"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn": {6851, 6851},
							},
							Range: GetBlockProductionRange{
								FirstSlot: 0,
								LastSlot:  6850,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockProduction", "params":[{"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":6850},"value":{"byIdentity":{"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn":[6851,6851]},"range":{"firstSlot":0,"lastSlot":6850}}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockProductionWithConfig(
						context.TODO(),
						GetBlockProductionConfig{
							Commitment: CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetBlockProduction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetBlockProduction{
						Context: Context{
							Slot: 6850,
						},
						Value: GetBlockProductionResponseResultValue{
							ByIdentity: map[string][]uint64{
								"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn": {6851, 6851},
							},
							Range: GetBlockProductionRange{
								FirstSlot: 0,
								LastSlot:  6850,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockProduction", "params":[{"range": {"firstSlot": 6000}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":6850},"value":{"byIdentity":{"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn":[6851,6851]},"range":{"firstSlot":0,"lastSlot":6850}}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockProductionWithConfig(
						context.TODO(),
						GetBlockProductionConfig{
							Range: &GetBlockProductionRange{
								FirstSlot: 6000,
							},
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetBlockProduction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetBlockProduction{
						Context: Context{
							Slot: 6850,
						},
						Value: GetBlockProductionResponseResultValue{
							ByIdentity: map[string][]uint64{
								"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn": {6851, 6851},
							},
							Range: GetBlockProductionRange{
								FirstSlot: 0,
								LastSlot:  6850,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockProduction", "params":[{"range": {"firstSlot": 6000, "lastSlot": 6100}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":6850},"value":{"byIdentity":{"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn":[6851,6851]},"range":{"firstSlot":0,"lastSlot":6850}}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockProductionWithConfig(
						context.TODO(),
						GetBlockProductionConfig{
							Range: &GetBlockProductionRange{
								FirstSlot: 6000,
								LastSlot:  6100,
							},
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetBlockProduction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetBlockProduction{
						Context: Context{
							Slot: 6850,
						},
						Value: GetBlockProductionResponseResultValue{
							ByIdentity: map[string][]uint64{
								"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn": {6851, 6851},
							},
							Range: GetBlockProductionRange{
								FirstSlot: 0,
								LastSlot:  6850,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockProduction", "params":[{"identity": "8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn", "range": {"firstSlot": 6000, "lastSlot": 6100}}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":6851},"value":{"byIdentity":{"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn":[101,101]},"range":{"firstSlot":6000,"lastSlot":6100}}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockProductionWithConfig(
						context.TODO(),
						GetBlockProductionConfig{
							Identity: "8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn",
							Range: &GetBlockProductionRange{
								FirstSlot: 6000,
								LastSlot:  6100,
							},
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetBlockProduction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetBlockProduction{
						Context: Context{
							Slot: 6851,
						},
						Value: GetBlockProductionResponseResultValue{
							ByIdentity: map[string][]uint64{
								"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn": {101, 101},
							},
							Range: GetBlockProductionRange{
								FirstSlot: 6000,
								LastSlot:  6100,
							},
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
