package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/program/system"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetSignatureStatus(t *testing.T) {
	type args struct {
		ctx context.Context
		sig string
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         *rpc.SignatureStatus
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"confirmed","confirmations":25,"err":null,"slot":86136551,"status":{"Ok":null}}]},"id":1}`,
			args: args{
				ctx: context.Background(),
				sig: "3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
			},
			want: &rpc.SignatureStatus{
				ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentConfirmed))),
				Confirmations:      pointer.Get[uint64](25),
				Err:                nil,
				Slot:               86136551,
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}]},"id":1}`,
			args: args{
				ctx: context.Background(),
				sig: "3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
			},
			want: &rpc.SignatureStatus{
				ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentFinalized))),
				Confirmations:      nil,
				Err:                nil,
				Slot:               86136524,
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[null]},"id":1}`,
			args: args{
				ctx: context.Background(),
				sig: "3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
			},
			want: nil,
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := io.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.requestBody, string(body))
					n, err := rw.Write([]byte(tt.responseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.responseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.GetSignatureStatus(tt.args.ctx, tt.args.sig)
				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
				server.Close()
			})
		})
	}
}

func TestClient_GetSignatureStatusWithConfig(t *testing.T) {
	type args struct {
		ctx context.Context
		sig string
		cfg rpc.GetSignatureStatusesConfig
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         *rpc.SignatureStatus
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"confirmed","confirmations":25,"err":null,"slot":86136551,"status":{"Ok":null}}]},"id":1}`,
			args: args{
				ctx: context.Background(),
				sig: "3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
				cfg: rpc.GetSignatureStatusesConfig{
					SearchTransactionHistory: true,
				},
			},
			want: &rpc.SignatureStatus{
				ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentConfirmed))),
				Confirmations:      pointer.Get[uint64](25),
				Err:                nil,
				Slot:               86136551,
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}]},"id":1}`,
			args: args{
				ctx: context.Background(),
				sig: "3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
				cfg: rpc.GetSignatureStatusesConfig{
					SearchTransactionHistory: true,
				},
			},
			want: &rpc.SignatureStatus{
				ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentFinalized))),
				Confirmations:      nil,
				Err:                nil,
				Slot:               86136524,
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[null]},"id":1}`,
			args: args{
				ctx: context.Background(),
				sig: "3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg",
				cfg: rpc.GetSignatureStatusesConfig{
					SearchTransactionHistory: true,
				},
			},
			want: nil,
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := io.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.requestBody, string(body))
					n, err := rw.Write([]byte(tt.responseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.responseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.GetSignatureStatusWithConfig(tt.args.ctx, tt.args.sig, tt.args.cfg)
				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
				server.Close()
			})
		})
	}
}

func TestClient_GetSignatureStatuses(t *testing.T) {
	type args struct {
		ctx  context.Context
		sigs []string
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         rpc.SignatureStatuses
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"confirmed","confirmations":25,"err":null,"slot":86136551,"status":{"Ok":null}}]},"id":1}`,
			args: args{
				ctx:  context.Background(),
				sigs: []string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
			},
			want: rpc.SignatureStatuses{
				{
					ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentConfirmed))),
					Confirmations:      pointer.Get[uint64](25),
					Err:                nil,
					Slot:               86136551,
				},
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}]},"id":1}`,
			args: args{
				ctx:  context.Background(),
				sigs: []string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
			},
			want: rpc.SignatureStatuses{
				{
					ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentFinalized))),
					Confirmations:      nil,
					Err:                nil,
					Slot:               86136524,
				},
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"]]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[null]},"id":1}`,
			args: args{
				ctx:  context.Background(),
				sigs: []string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
			},
			want: rpc.SignatureStatuses{nil},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := io.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.requestBody, string(body))
					n, err := rw.Write([]byte(tt.responseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.responseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.GetSignatureStatuses(tt.args.ctx, tt.args.sigs)
				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
				server.Close()
			})
		})
	}
}

func TestClient_GetSignatureStatusesWithConfig(t *testing.T) {
	type args struct {
		ctx  context.Context
		sigs []string
		cfg  rpc.GetSignatureStatusesConfig
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         rpc.SignatureStatuses
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"confirmed","confirmations":25,"err":null,"slot":86136551,"status":{"Ok":null}}]},"id":1}`,
			args: args{
				ctx:  context.Background(),
				sigs: []string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
				cfg: rpc.GetSignatureStatusesConfig{
					SearchTransactionHistory: true,
				},
			},
			want: rpc.SignatureStatuses{
				{
					ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentConfirmed))),
					Confirmations:      pointer.Get[uint64](25),
					Err:                nil,
					Slot:               86136551,
				},
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[{"confirmationStatus":"finalized","confirmations":null,"err":null,"slot":86136524,"status":{"Ok":null}}]},"id":1}`,
			args: args{
				ctx:  context.Background(),
				sigs: []string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
				cfg: rpc.GetSignatureStatusesConfig{
					SearchTransactionHistory: true,
				},
			},
			want: rpc.SignatureStatuses{
				{
					ConfirmationStatus: (*rpc.Commitment)(pointer.Get(string(rpc.CommitmentFinalized))),
					Confirmations:      nil,
					Err:                nil,
					Slot:               86136524,
				},
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignatureStatuses", "params":[["3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"], {"searchTransactionHistory": true}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":86136583},"value":[null]},"id":1}`,
			args: args{
				ctx:  context.Background(),
				sigs: []string{"3E6jD48LnMeNDs1QTXXunXGaqYybZKHXYdriDwqXGJbCXzVkMZNexuiGnTtUSba7PcmbKcsxKsAcBKLSmqjUKDRg"},
				cfg: rpc.GetSignatureStatusesConfig{
					SearchTransactionHistory: true,
				},
			},
			want: rpc.SignatureStatuses{nil},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := io.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.requestBody, string(body))
					n, err := rw.Write([]byte(tt.responseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.responseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.GetSignatureStatusesWithConfig(tt.args.ctx, tt.args.sigs, tt.args.cfg)
				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
				server.Close()
			})
		})
	}
}

func TestClient_SimulateTransactionWithConfig(t *testing.T) {
	type args struct {
		ctx context.Context
		tx  types.Transaction
		cfg SimulateTransactionConfig
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         SimulateTransaction
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"simulateTransaction", "params":["ATGaPuMIw9CsfadV9Nb4MMzJYklmBijmw264ueG2dnwVYaozlV/a3s2rFgznOaJPTdO4vuWJDFNceOk+6XSulA8BAAEDJGfwe2rYvkoBCI3IHIybkz+rysOkG7upqFJ/ZsFJmrgkaRt6Oq89b8zO7fpC3krN2PFyLB32negX+tFWkhOEaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAq6APAPRzMLN9vl6W+qjRKQbb1TnwZvWUH8lNXaW4SxkBAgIAAQwCAAAAAQAAAAAAAAA=", {"encoding": "base64", "accounts": {"encoding": "base64", "addresses": ["7amubLCYTZMnR5fBo9fi8r5NfhZXFFQWej2UXi5Q71nq", "3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"]}}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":95948750},"value":{"accounts":[null,{"data":["","base64"],"executable":false,"lamports":999994999,"owner":"11111111111111111111111111111111","rentEpoch":222}],"err":null,"logs":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"]}},"id":1}`,
			args: args{
				ctx: context.Background(),
				cfg: SimulateTransactionConfig{
					Addresses: []string{"7amubLCYTZMnR5fBo9fi8r5NfhZXFFQWej2UXi5Q71nq", "3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"},
				},
				tx: types.Transaction{
					Signatures: []types.Signature{
						{0x31, 0x9a, 0x3e, 0xe3, 0x8, 0xc3, 0xd0, 0xac, 0x7d, 0xa7, 0x55, 0xf4, 0xd6, 0xf8, 0x30, 0xcc, 0xc9, 0x62, 0x49, 0x66, 0x6, 0x28, 0xe6, 0xc3, 0x6e, 0xb8, 0xb9, 0xe1, 0xb6, 0x76, 0x7c, 0x15, 0x61, 0xaa, 0x33, 0x95, 0x5f, 0xda, 0xde, 0xcd, 0xab, 0x16, 0xc, 0xe7, 0x39, 0xa2, 0x4f, 0x4d, 0xd3, 0xb8, 0xbe, 0xe5, 0x89, 0xc, 0x53, 0x5c, 0x78, 0xe9, 0x3e, 0xe9, 0x74, 0xae, 0x94, 0xf},
					},
					Message: types.NewMessage(types.NewMessageParam{
						FeePayer: common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
						Instructions: []types.Instruction{
							system.Transfer(system.TransferParam{
								From:   common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
								To:     common.PublicKeyFromString("3T8doiqPi8XDvmPVMA3UyEwvssqA3QrAJYjg92jMT7D9"),
								Amount: 1,
							}),
						},
						RecentBlockhash: "CYxGqF7n3vZtvbjdBZf35e1ejTebaPoPATn8VvKqg83J",
					}),
				},
			},
			want: SimulateTransaction{
				Logs: []string{
					"Program 11111111111111111111111111111111 invoke [1]",
					"Program 11111111111111111111111111111111 success",
				},
				Accounts: []*AccountInfo{
					nil,
					{
						Owner:      common.SystemProgramID,
						RentEpoch:  222,
						Lamports:   999994999,
						Executable: false,
						Data:       []byte{},
					},
				},
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"simulateTransaction", "params":["ARNAYOrxLauRNXP42vVIm9L+hhrhksvKtnSdCc2o/XyUpJI4ajj1v+XhDu+zPiupczOSv8WYS63voKqgbJJtYwIBAAEDJGfwe2rYvkoBCI3IHIybkz+rysOkG7upqFJ/ZsFJmrgkaRt6Oq89b8zO7fpC3krN2PFyLB32negX+tFWkhOEaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAvl9qvCSc7MmYX+8JkCg6fi/H72HezoUKWLJjwmx2y64BAgIAAQwCAAAAAQAAAAAAAAA=", {"encoding": "base64", "replaceRecentBlockhash": true}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":95949958},"value":{"accounts":null,"err":null,"logs":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"]}},"id":1}`,
			args: args{
				ctx: context.Background(),
				cfg: SimulateTransactionConfig{
					ReplaceRecentBlockhash: true,
				},
				tx: types.Transaction{
					Signatures: []types.Signature{
						{0x13, 0x40, 0x60, 0xea, 0xf1, 0x2d, 0xab, 0x91, 0x35, 0x73, 0xf8, 0xda, 0xf5, 0x48, 0x9b, 0xd2, 0xfe, 0x86, 0x1a, 0xe1, 0x92, 0xcb, 0xca, 0xb6, 0x74, 0x9d, 0x9, 0xcd, 0xa8, 0xfd, 0x7c, 0x94, 0xa4, 0x92, 0x38, 0x6a, 0x38, 0xf5, 0xbf, 0xe5, 0xe1, 0xe, 0xef, 0xb3, 0x3e, 0x2b, 0xa9, 0x73, 0x33, 0x92, 0xbf, 0xc5, 0x98, 0x4b, 0xad, 0xef, 0xa0, 0xaa, 0xa0, 0x6c, 0x92, 0x6d, 0x63, 0x2},
					},
					Message: types.NewMessage(types.NewMessageParam{
						FeePayer: common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
						Instructions: []types.Instruction{
							system.Transfer(system.TransferParam{
								From:   common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
								To:     common.PublicKeyFromString("3T8doiqPi8XDvmPVMA3UyEwvssqA3QrAJYjg92jMT7D9"),
								Amount: 1,
							}),
						},
						RecentBlockhash: "Dp8rHi3URdtT3VEEyJj4w9Vv4797XM7CArdGJVxLaxb7",
					}),
				},
			},
			want: SimulateTransaction{
				Logs: []string{
					"Program 11111111111111111111111111111111 invoke [1]",
					"Program 11111111111111111111111111111111 success",
				},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := io.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.requestBody, string(body))
					n, err := rw.Write([]byte(tt.responseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.responseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.SimulateTransactionWithConfig(tt.args.ctx, tt.args.tx, tt.args.cfg)
				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
				server.Close()
			})
		})
	}
}

func Test_checkJsonRpcResponse(t *testing.T) {
	type args struct {
		res rpc.JsonRpcResponse[rpc.GetBlock]
		err error
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			args: args{
				res: rpc.JsonRpcResponse[rpc.GetBlock]{
					JsonRpc: "2.0",
					Id:      1,
					Result:  rpc.GetBlock{},
				},
				err: nil,
			},
			expectedErr: nil,
		},
		{
			args: args{
				res: rpc.JsonRpcResponse[rpc.GetBlock]{
					JsonRpc: "2.0",
					Id:      1,
					Result:  rpc.GetBlock{},
				},
				err: errors.New("rpc error"),
			},
			expectedErr: errors.New("rpc error"),
		},
		{
			args: args{
				res: rpc.JsonRpcResponse[rpc.GetBlock]{
					JsonRpc: "2.0",
					Id:      1,
					Result:  rpc.GetBlock{},
					Error: &rpc.JsonRpcError{
						Code:    -1,
						Message: "error",
						Data:    nil,
					},
				},
				err: nil,
			},
			expectedErr: &rpc.JsonRpcError{
				Code:    -1,
				Message: "error",
				Data:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkJsonRpcResponse(tt.args.res, tt.args.err)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
