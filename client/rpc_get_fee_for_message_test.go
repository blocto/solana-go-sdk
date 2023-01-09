package client

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

func TestClient_GetFeeForMessage(t *testing.T) {
	b, err := base64.StdEncoding.DecodeString("AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA")
	if err != nil {
		t.Errorf("failed to decode message, err: %v", err)
	}
	message, err := types.MessageDeserialize(b)
	if err != nil {
		t.Errorf("failed to deserialize message, err: %v", err)
	}
	var nilValue *uint64

	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187830352},"value":null},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetFeeForMessage(
						context.TODO(),
						message,
					)
				},
				ExpectedValue: nilValue,
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187830352},"value":5000},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetFeeForMessage(
						context.TODO(),
						message,
					)
				},
				ExpectedValue: pointer.Get[uint64](5000),
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetFeeForMessageWithConfig(t *testing.T) {
	b, err := base64.StdEncoding.DecodeString("AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA")
	if err != nil {
		t.Errorf("failed to decode message, err: %v", err)
	}
	message, err := types.MessageDeserialize(b)
	if err != nil {
		t.Errorf("failed to deserialize message, err: %v", err)
	}
	var nilValue *uint64

	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187830352},"value":null},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetFeeForMessageWithConfig(
						context.TODO(),
						message,
						GetFeeForMessageConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: nilValue,
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187830352},"value":5000},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetFeeForMessageWithConfig(
						context.TODO(),
						message,
						GetFeeForMessageConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: pointer.Get[uint64](5000),
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetFeeForMessageAndContext(t *testing.T) {
	b, err := base64.StdEncoding.DecodeString("AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA")
	if err != nil {
		t.Errorf("failed to decode message, err: %v", err)
	}
	message, err := types.MessageDeserialize(b)
	if err != nil {
		t.Errorf("failed to deserialize message, err: %v", err)
	}
	var nilValue *uint64

	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187830352},"value":null},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetFeeForMessageAndContext(
						context.TODO(),
						message,
					)
				},
				ExpectedValue: rpc.ValueWithContext[*uint64]{
					Context: rpc.Context{
						Slot:       187830352,
						ApiVersion: "1.14.10",
					},
					Value: nilValue,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187830352},"value":5000},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetFeeForMessageAndContext(
						context.TODO(),
						message,
					)
				},
				ExpectedValue: rpc.ValueWithContext[*uint64]{
					Context: rpc.Context{
						Slot:       187830352,
						ApiVersion: "1.14.10",
					},
					Value: pointer.Get[uint64](5000),
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetFeeForMessageAndContextWithConfig(t *testing.T) {
	b, err := base64.StdEncoding.DecodeString("AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA")
	if err != nil {
		t.Errorf("failed to decode message, err: %v", err)
	}
	message, err := types.MessageDeserialize(b)
	if err != nil {
		t.Errorf("failed to deserialize message, err: %v", err)
	}
	var nilValue *uint64

	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187830352},"value":null},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetFeeForMessageAndContextWithConfig(
						context.TODO(),
						message,
						GetFeeForMessageConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[*uint64]{
					Context: rpc.Context{
						Slot:       187830352,
						ApiVersion: "1.14.10",
					},
					Value: nilValue,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA", {"commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187830352},"value":5000},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetFeeForMessageAndContextWithConfig(
						context.TODO(),
						message,
						GetFeeForMessageConfig{
							Commitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: rpc.ValueWithContext[*uint64]{
					Context: rpc.Context{
						Slot:       187830352,
						ApiVersion: "1.14.10",
					},
					Value: pointer.Get[uint64](5000),
				},
				ExpectedError: nil,
			},
		},
	)
}
