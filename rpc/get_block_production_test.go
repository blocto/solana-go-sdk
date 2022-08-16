package rpc

import (
	"context"
	"testing"
)

func TestGetBlockProduction(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlockProduction"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":6850},"value":{"byIdentity":{"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn":[6851,6851]},"range":{"firstSlot":0,"lastSlot":6850}}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockProduction(
					context.TODO(),
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlockProduction]{
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
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockProductionWithConfig(
					context.TODO(),
					GetBlockProductionConfig{
						Identity: "8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn",
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlockProduction]{
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
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockProductionWithConfig(
					context.TODO(),
					GetBlockProductionConfig{
						Commitment: CommitmentConfirmed,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlockProduction]{
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
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockProductionWithConfig(
					context.TODO(),
					GetBlockProductionConfig{
						Range: &GetBlockProductionRange{
							FirstSlot: 6000,
						},
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlockProduction]{
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
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockProductionWithConfig(
					context.TODO(),
					GetBlockProductionConfig{
						Range: &GetBlockProductionRange{
							FirstSlot: 6000,
							LastSlot:  6100,
						},
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlockProduction]{
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
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockProductionWithConfig(
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
			ExpectedResponse: JsonRpcResponse[GetBlockProduction]{
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
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			testRpcCall(t, tt)
		})
	}
}
