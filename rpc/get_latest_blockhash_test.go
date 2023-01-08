package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
)

func TestGetLatestBlockhash(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112872139},"value":{"blockhash":"9K9GnvWXn9zYitQdHUSYzvjLjebnviwEFaWgWqHDU3ve","lastValidBlockHeight":92248597}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetLatestBlockhash(context.TODO())
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetLatestBlockhashValue]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetLatestBlockhashValue]{
						Context: Context{
							Slot: 112872139,
						},
						Value: GetLatestBlockhashValue{
							Blockhash:              "9K9GnvWXn9zYitQdHUSYzvjLjebnviwEFaWgWqHDU3ve",
							LatestValidBlockHeight: 92248597,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash", "params":[{"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112871314},"value":{"blockhash":"3H2pwJD6pTrEveh5xcwHXToLn7txt5uTW6CPzCan4ZKL","lastValidBlockHeight":92247902}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetLatestBlockhashWithConfig(context.TODO(), GetLatestBlockhashConfig{Commitment: CommitmentProcessed})
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetLatestBlockhashValue]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetLatestBlockhashValue]{
						Context: Context{
							Slot: 112871314,
						},
						Value: GetLatestBlockhashValue{
							Blockhash:              "3H2pwJD6pTrEveh5xcwHXToLn7txt5uTW6CPzCan4ZKL",
							LatestValidBlockHeight: 92247902,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash", "params":[{"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112871311},"value":{"blockhash":"FXuaK93DmxWt98bv3wYMdE3TMnY2o8e3h85KrGWEUAzv","lastValidBlockHeight":92247899}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetLatestBlockhashWithConfig(context.TODO(), GetLatestBlockhashConfig{Commitment: CommitmentConfirmed})
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetLatestBlockhashValue]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetLatestBlockhashValue]{
						Context: Context{
							Slot: 112871311,
						},
						Value: GetLatestBlockhashValue{
							Blockhash:              "FXuaK93DmxWt98bv3wYMdE3TMnY2o8e3h85KrGWEUAzv",
							LatestValidBlockHeight: 92247899,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash", "params":[{"commitment": "finalized"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112871221},"value":{"blockhash":"21f41sJRvMV8Tc3R5bTTA3n3yBLuoocSkgb8zj1vmEJa","lastValidBlockHeight":92247838}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetLatestBlockhashWithConfig(context.TODO(), GetLatestBlockhashConfig{Commitment: CommitmentFinalized})
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[GetLatestBlockhashValue]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[GetLatestBlockhashValue]{
						Context: Context{
							Slot: 112871221,
						},
						Value: GetLatestBlockhashValue{
							Blockhash:              "21f41sJRvMV8Tc3R5bTTA3n3yBLuoocSkgb8zj1vmEJa",
							LatestValidBlockHeight: 92247838,
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
