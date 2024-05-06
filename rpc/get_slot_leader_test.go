package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
)

func TestGetSlotLeader(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlotLeader"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot","id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSlotLeader(context.TODO())
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  "q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot",
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestGetSlotLeaderWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlotLeader","params":[{"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot","id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSlotLeaderWithConfig(
						context.TODO(),
						GetSlotLeaderConfig{
							Commitment: pointer.Get(CommitmentProcessed),
						},
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  "q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot",
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlotLeader","params":[{"minContextSlot":0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot","id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSlotLeaderWithConfig(
						context.TODO(),
						GetSlotLeaderConfig{
							MinContextSlot: pointer.Get[uint64](0),
						},
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  "q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot",
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlotLeader","params":[{"commitment": "confirmed","minContextSlot":10}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot","id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSlotLeaderWithConfig(
						context.TODO(),
						GetSlotLeaderConfig{
							Commitment:     pointer.Get(CommitmentConfirmed),
							MinContextSlot: pointer.Get[uint64](10),
						},
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  "q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot",
				},
				ExpectedError: nil,
			},
		},
	)
}
