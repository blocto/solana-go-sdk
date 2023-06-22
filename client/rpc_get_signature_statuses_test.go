package client

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/rpc"
)

func TestClient_GetSignatureStatus(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"confirmed","confirmations":25,"err":null,"slot":86136551,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatus(context.Background(), "3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg")
				},
				ExpectedValue: &rpc.SignatureStatus{
					ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentConfirmed))),
					Confirmations:      pointer.Get[uint64](25),
					Err:                nil,
					Slot:               86136551,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatus(context.Background(), "3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg")
				},
				ExpectedValue: &rpc.SignatureStatus{
					ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentFinalized))),
					Confirmations:      nil,
					Err:                nil,
					Slot:               86136524,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[null]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatus(context.Background(), "3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg")
				},
				ExpectedValue: (*rpc.SignatureStatus)(nil),
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetSignatureStatusWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"confirmed","confirmations":25,"err":null,"slot":86136551,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatusWithConfig(
						context.Background(),
						"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
						GetSignatureStatusesConfig{
							SearchTransactionHistory: true,
						},
					)
				},
				ExpectedValue: &rpc.SignatureStatus{
					ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentConfirmed))),
					Confirmations:      pointer.Get[uint64](25),
					Err:                nil,
					Slot:               86136551,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatusWithConfig(
						context.Background(),
						"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
						GetSignatureStatusesConfig{
							SearchTransactionHistory: true,
						},
					)
				},
				ExpectedValue: &rpc.SignatureStatus{
					ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentFinalized))),
					Confirmations:      nil,
					Err:                nil,
					Slot:               86136524,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[null]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatusWithConfig(
						context.Background(),
						"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
						GetSignatureStatusesConfig{
							SearchTransactionHistory: true,
						},
					)
				},
				ExpectedValue: (*rpc.SignatureStatus)(nil),
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetSignatureStatuses(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"confirmed","confirmations":25,"err":null,"slot":86136551,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatuses(
						context.Background(),
						[]string{
							"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
						},
					)
				},
				ExpectedValue: rpc.SignatureStatuses{
					{
						ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentConfirmed))),
						Confirmations:      pointer.Get[uint64](25),
						Err:                nil,
						Slot:               86136551,
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg", "4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}, null]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatuses(
						context.Background(),
						[]string{
							"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
							"4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ",
						},
					)
				},
				ExpectedValue: rpc.SignatureStatuses{
					{
						ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentFinalized))),
						Confirmations:      nil,
						Err:                nil,
						Slot:               86136524,
					},
					nil,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg", "4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ"]]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[null,null]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatuses(
						context.Background(),
						[]string{
							"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
							"4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ",
						},
					)
				},
				ExpectedValue: rpc.SignatureStatuses{nil, nil},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetSignatureStatusesWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"confirmed","confirmations":25,"err":null,"slot":86136551,"status":{"Ok":null}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatusesWithConfig(
						context.Background(),
						[]string{
							"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
						},
						GetSignatureStatusesConfig{
							SearchTransactionHistory: true,
						},
					)
				},
				ExpectedValue: rpc.SignatureStatuses{
					{
						ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentConfirmed))),
						Confirmations:      pointer.Get[uint64](25),
						Err:                nil,
						Slot:               86136551,
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg", "4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ"], {"searchTransactionHistory": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}, null]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatusesWithConfig(
						context.Background(),
						[]string{
							"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
							"4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ",
						},
						GetSignatureStatusesConfig{
							SearchTransactionHistory: true,
						},
					)
				},
				ExpectedValue: rpc.SignatureStatuses{
					{
						ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentFinalized))),
						Confirmations:      nil,
						Err:                nil,
						Slot:               86136524,
					},
					nil,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg", "4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ"], {"searchTransactionHistory": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[null,null]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSignatureStatusesWithConfig(
						context.Background(),
						[]string{
							"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
							"4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ",
						},
						GetSignatureStatusesConfig{
							SearchTransactionHistory: true,
						},
					)
				},
				ExpectedValue: rpc.SignatureStatuses{nil, nil},
				ExpectedError: nil,
			},
		},
	)
}
