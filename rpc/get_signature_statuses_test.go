package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/pkg/pointer"
)

func TestGetSignatureStatuses(t *testing.T) {
	tests := []testRpcCallParam{
		{
			Name:         "get 1 signature",
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}]},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetSignatureStatuses(
					context.TODO(),
					[]string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
				)
			},
			ExpectedResponse: GetSignatureStatusesResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetSignatureStatusesResult{
					Context: Context{
						Slot: 86136583,
					},
					Value: []*GetSignatureStatusesResultValue{
						{
							Slot:               86136524,
							Confirmations:      nil,
							ConfirmationStatus: (*Commitment)(pointer.String(string(CommitmentFinalized))),
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
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetSignatureStatuses(
					context.TODO(),
					[]string{
						"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
						"3wPgLJ1e34v41VKzq4AN7jiS2m5VWZU98NT2AWSTtobLwXGdUbazV8cXrw9ooi1LiP37imfQyfjZat4rufFmc2VK",
					},
				)
			},
			ExpectedResponse: GetSignatureStatusesResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetSignatureStatusesResult{
					Context: Context{
						Slot: 86136583,
					},
					Value: []*GetSignatureStatusesResultValue{
						{
							Slot:               86136524,
							Confirmations:      nil,
							ConfirmationStatus: (*Commitment)(pointer.String(string(CommitmentFinalized))),
							Err:                nil,
						},
						{
							Slot:               86136551,
							Confirmations:      pointer.Uint64(25),
							ConfirmationStatus: (*Commitment)(pointer.String(string(CommitmentConfirmed))),
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
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetSignatureStatusesWithConfig(
					context.TODO(),
					[]string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
					GetSignatureStatusesConfig{
						SearchTransactionHistory: true,
					},
				)
			},
			ExpectedResponse: GetSignatureStatusesResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetSignatureStatusesResult{
					Context: Context{
						Slot: 86136583,
					},
					Value: []*GetSignatureStatusesResultValue{
						{
							Slot:               86136524,
							Confirmations:      nil,
							ConfirmationStatus: (*Commitment)(pointer.String(string(CommitmentFinalized))),
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
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetSignatureStatusesWithConfig(
					context.TODO(),
					[]string{"5K4QuDxZjFkAhv8Kfj6pocDf59Cuo8EmxbRcM1d4PVaSWekF6RdC2krJHucF1FhkuSFVHCMcZ4GXnSx7zzygXMdi"},
					GetSignatureStatusesConfig{
						SearchTransactionHistory: true,
					},
				)
			},
			ExpectedResponse: GetSignatureStatusesResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetSignatureStatusesResult{
					Context: Context{
						Slot: 86143308,
					},
					Value: []*GetSignatureStatusesResultValue{
						{
							Slot:               85500622,
							Confirmations:      nil,
							ConfirmationStatus: (*Commitment)(pointer.String(string(CommitmentFinalized))),
							Err: map[string]interface{}{
								"InstructionError": []interface{}{
									0.,
									map[string]interface{}{
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
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetSignatureStatusesWithConfig(
					context.TODO(),
					[]string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
					GetSignatureStatusesConfig{
						SearchTransactionHistory: true,
					},
				)
			},
			ExpectedResponse: GetSignatureStatusesResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetSignatureStatusesResult{
					Context: Context{
						Slot: 86138946,
					},
					Value: []*GetSignatureStatusesResultValue{
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
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetSignatureStatuses(
					context.TODO(),
					[]string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg", "4jSRBMUh8HvfPkjZt8enXBFvKykhGdDW5uqtXy8ys52rqeEQuG8Y9hyRSjsSwjEkYYQht7aqEoLwJnuxv9YD99EQ"},
				)
			},
			ExpectedResponse: GetSignatureStatusesResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetSignatureStatusesResult{
					Context: Context{
						Slot: 86142683,
					},
					Value: []*GetSignatureStatusesResultValue{
						nil,
						{
							Slot:               86142617,
							Confirmations:      nil,
							ConfirmationStatus: (*Commitment)(pointer.String(string(CommitmentFinalized))),
							Err:                nil,
						},
					},
				},
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			testRpcCall(t, tt)
		})
	}
}
