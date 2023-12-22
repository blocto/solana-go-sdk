package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
)

func TestGetRecentPrioritizationFees(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentPrioritizationFees", "params":[["CxELquR1gPP8wHe33gZ4QxqGB3sZ9RSwsJ2KshVewkFY"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"slot":348125,"prioritizationFee":0},{"slot":348126,"prioritizationFee":1000},{"slot":348127,"prioritizationFee":500},{"slot":348128,"prioritizationFee":0},{"slot":348129,"prioritizationFee":1234}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetRecentPrioritizationFees(
						context.TODO(),
						[]string{"CxELquR1gPP8wHe33gZ4QxqGB3sZ9RSwsJ2KshVewkFY"},
					)
				},
				ExpectedValue: JsonRpcResponse[PrioritizationFees]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: PrioritizationFees{
						{
							Slot:              348125,
							PrioritizationFee: 0,
						},
						{
							Slot:              348126,
							PrioritizationFee: 1000,
						},
						{
							Slot:              348127,
							PrioritizationFee: 500,
						},
						{
							Slot:              348128,
							PrioritizationFee: 0,
						},
						{
							Slot:              348129,
							PrioritizationFee: 1234,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentPrioritizationFees", "params":[[]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"slot":348125,"prioritizationFee":0},{"slot":348126,"prioritizationFee":1000},{"slot":348127,"prioritizationFee":500},{"slot":348128,"prioritizationFee":0},{"slot":348129,"prioritizationFee":1234}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetRecentPrioritizationFees(
						context.TODO(),
						[]string{},
					)
				},
				ExpectedValue: JsonRpcResponse[PrioritizationFees]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: PrioritizationFees{
						{
							Slot:              348125,
							PrioritizationFee: 0,
						},
						{
							Slot:              348126,
							PrioritizationFee: 1000,
						},
						{
							Slot:              348127,
							PrioritizationFee: 500,
						},
						{
							Slot:              348128,
							PrioritizationFee: 0,
						},
						{
							Slot:              348129,
							PrioritizationFee: 1234,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getRecentPrioritizationFees", "params":[[]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"slot":348125,"prioritizationFee":0},{"slot":348126,"prioritizationFee":1000},{"slot":348127,"prioritizationFee":500},{"slot":348128,"prioritizationFee":0},{"slot":348129,"prioritizationFee":1234}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetRecentPrioritizationFees(
						context.TODO(),
						[]string{},
					)
				},
				ExpectedValue: JsonRpcResponse[PrioritizationFees]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: PrioritizationFees{
						{
							Slot:              348125,
							PrioritizationFee: 0,
						},
						{
							Slot:              348126,
							PrioritizationFee: 1000,
						},
						{
							Slot:              348127,
							PrioritizationFee: 500,
						},
						{
							Slot:              348128,
							PrioritizationFee: 0,
						},
						{
							Slot:              348129,
							PrioritizationFee: 1234,
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
