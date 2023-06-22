package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/pkg/pointer"
)

func TestGetSignatureStatuses(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				Name:         "get 1 signature",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignatureStatuses(
						context.TODO(),
						[]string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[SignatureStatuses]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[SignatureStatuses]{
						Context: Context{
							Slot: 86136583,
						},
						Value: []*SignatureStatus{
							{
								Slot:               86136524,
								Confirmations:      nil,
								ConfirmationStatus: (*Commitment)(pointer.Get(string(CommitmentFinalized))),
								Err:                nil,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "get 2 signature",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg", "3wPgLJ1e34v41VKzq4AN7jiS2m5VWZU98NT2AWSTtobLwXGdUbazV8cXrw9ooi1LiP37imfQyfjZat4rufFmc2VK"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}},{"confirmationStatus":"confirmed","confirmations":25,"err":null,"slot":86136551,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignatureStatuses(
						context.TODO(),
						[]string{
							"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
							"3wPgLJ1e34v41VKzq4AN7jiS2m5VWZU98NT2AWSTtobLwXGdUbazV8cXrw9ooi1LiP37imfQyfjZat4rufFmc2VK",
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[SignatureStatuses]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[SignatureStatuses]{
						Context: Context{
							Slot: 86136583,
						},
						Value: []*SignatureStatus{
							{
								Slot:               86136524,
								Confirmations:      nil,
								ConfirmationStatus: (*Commitment)(pointer.Get(string(CommitmentFinalized))),
								Err:                nil,
							},
							{
								Slot:               86136551,
								Confirmations:      pointer.Get[uint64](25),
								ConfirmationStatus: (*Commitment)(pointer.Get(string(CommitmentConfirmed))),
								Err:                nil,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "get 1 signature with config",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignatureStatusesWithConfig(
						context.TODO(),
						[]string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
						GetSignatureStatusesConfig{
							SearchTransactionHistory: true,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[SignatureStatuses]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[SignatureStatuses]{
						Context: Context{
							Slot: 86136583,
						},
						Value: []*SignatureStatus{
							{
								Slot:               86136524,
								Confirmations:      nil,
								ConfirmationStatus: (*Commitment)(pointer.Get(string(CommitmentFinalized))),
								Err:                nil,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "get 1 failed signature with config",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["5K4QuDxZjFkAhv8Kfj6pocDf59Cuo8EmxbRcM1d4PVaSWekF6RdC2krJHucF1FhkuSFVHCMcZ4GXnSx7zzygXMdi"], {"searchTransactionHistory": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86143308},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":{"InstructionError":[0,{"Custom":1}]},"slot":85500622,"status":{"Err":{"InstructionError":[0,{"Custom":1}]}}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignatureStatusesWithConfig(
						context.TODO(),
						[]string{"5K4QuDxZjFkAhv8Kfj6pocDf59Cuo8EmxbRcM1d4PVaSWekF6RdC2krJHucF1FhkuSFVHCMcZ4GXnSx7zzygXMdi"},
						GetSignatureStatusesConfig{
							SearchTransactionHistory: true,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[SignatureStatuses]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[SignatureStatuses]{
						Context: Context{
							Slot: 86143308,
						},
						Value: []*SignatureStatus{
							{
								Slot:               85500622,
								Confirmations:      nil,
								ConfirmationStatus: (*Commitment)(pointer.Get(string(CommitmentFinalized))),
								Err: map[string]any{
									"InstructionError": []any{
										0.,
										map[string]any{
											"Custom": 1.,
										},
									},
								},
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "get not found signature",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86138946},"value":[null]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignatureStatusesWithConfig(
						context.TODO(),
						[]string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
						GetSignatureStatusesConfig{
							SearchTransactionHistory: true,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[SignatureStatuses]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[SignatureStatuses]{
						Context: Context{
							Slot: 86138946,
						},
						Value: []*SignatureStatus{
							nil,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "get not found signature 2",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg","4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86142683},"value":[null,{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86142617,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignatureStatuses(
						context.TODO(),
						[]string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg", "4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ"},
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[SignatureStatuses]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[SignatureStatuses]{
						Context: Context{
							Slot: 86142683,
						},
						Value: []*SignatureStatus{
							nil,
							{
								Slot:               86142617,
								Confirmations:      nil,
								ConfirmationStatus: (*Commitment)(pointer.Get(string(CommitmentFinalized))),
								Err:                nil,
							},
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
