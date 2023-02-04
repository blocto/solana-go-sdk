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

func TestGetBlock(t *testing.T) {
	type args struct {
		ctx  context.Context
		slot uint64
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         GetBlockResponse
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64"}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}],"transactions":[{"meta":{"err":null,"fee":10000,"innerInstructions":[],"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499999835001,1000000000000000,143487360,1169280,1],"postTokenBalances":[],"preBalances":[499999845001,1000000000000000,143487360,1169280,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"transaction":["AnXU8JYCIrc73JwxK9traTSp3EZdmnJp0B5luW8CCzr7GnFd/SjIMXiG4qbN5CwyEVhbpORzBUpB/253cNtS1A+0rWE+nrDqWRQ2OVU727PU4NtR611jY+10Q+F6lCZDsJt46b6oXz3PN5WGxTQk7mC4YhCbYsTcalWBkltA8KgPAgADBXszyT4GLb26BFuAAUXtW0B75zurDhXE7UOYKHFkpIlKJMmZpq+FRXTx8jzBMy1YsdkCo0kyLDdF2Q3NhXRdEosGp9UXGS8Kr8byZeP7d8x62oLFKdC+OxNuLQBVIAAAAAan1RcYx3TJKFZjmGkdXraLXrijm0ttXHNVWyEAAAAAB2FIHTV0dLt8TXYk69O9s9g1XnPREEP8DaNTgAAAAACrUBylgzc0SSCUPSfMJC3TI6KJEzs834KdMIMJci+UYAEEBAECAwE9AgAAAAEAAAAAAAAAIAAAAAAAAAAGCHSVIc5Betdf+NkRi4YR2D3abNLvpbI83qnB7EvNsAEZWkNhAAAAAA==","base64"]}]},"id":1}`,
			args: args{
				context.Background(),
				33,
			},
			want: GetBlockResponse{
				ParentSlot:        32,
				BlockHeight:       pointer.Get[int64](33),
				BlockTime:         pointer.Get[int64](1631803928),
				PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
				Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
				Rewards: []rpc.GetBlockReward{
					{
						Pubkey:       "9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX",
						Lamports:     5000,
						PostBalances: 499999840001,
						RewardType:   "Fee",
						Commission:   nil,
					},
				},
				Transactions: []GetBlockTransaction{
					{
						Meta: &TransactionMeta{
							Err: nil,
							Fee: 10000,
							PreBalances: []int64{
								499999845001,
								1000000000000000,
								143487360,
								1169280,
								1,
							},
							PostBalances: []int64{
								499999835001,
								1000000000000000,
								143487360,
								1169280,
								1,
							},
							PreTokenBalances:  []rpc.TransactionMetaTokenBalance{},
							PostTokenBalances: []rpc.TransactionMetaTokenBalance{},
							LogMessages: []string{
								"Program Vote111111111111111111111111111111111111111 invoke [1]",
								"Program Vote111111111111111111111111111111111111111 success",
							},
							InnerInstructions: []InnerInstruction{},
						},
						Transaction: types.Transaction{
							Signatures: []types.Signature{
								[]byte{0x75, 0xd4, 0xf0, 0x96, 0x2, 0x22, 0xb7, 0x3b, 0xdc, 0x9c, 0x31, 0x2b, 0xdb, 0x6b, 0x69, 0x34, 0xa9, 0xdc, 0x46, 0x5d, 0x9a, 0x72, 0x69, 0xd0, 0x1e, 0x65, 0xb9, 0x6f, 0x2, 0xb, 0x3a, 0xfb, 0x1a, 0x71, 0x5d, 0xfd, 0x28, 0xc8, 0x31, 0x78, 0x86, 0xe2, 0xa6, 0xcd, 0xe4, 0x2c, 0x32, 0x11, 0x58, 0x5b, 0xa4, 0xe4, 0x73, 0x5, 0x4a, 0x41, 0xff, 0x6e, 0x77, 0x70, 0xdb, 0x52, 0xd4, 0xf},
								[]byte{0xb4, 0xad, 0x61, 0x3e, 0x9e, 0xb0, 0xea, 0x59, 0x14, 0x36, 0x39, 0x55, 0x3b, 0xdb, 0xb3, 0xd4, 0xe0, 0xdb, 0x51, 0xeb, 0x5d, 0x63, 0x63, 0xed, 0x74, 0x43, 0xe1, 0x7a, 0x94, 0x26, 0x43, 0xb0, 0x9b, 0x78, 0xe9, 0xbe, 0xa8, 0x5f, 0x3d, 0xcf, 0x37, 0x95, 0x86, 0xc5, 0x34, 0x24, 0xee, 0x60, 0xb8, 0x62, 0x10, 0x9b, 0x62, 0xc4, 0xdc, 0x6a, 0x55, 0x81, 0x92, 0x5b, 0x40, 0xf0, 0xa8, 0xf},
							},
							Message: types.Message{
								Version: types.MessageVersionLegacy,
								Header: types.MessageHeader{
									NumRequireSignatures:        2,
									NumReadonlySignedAccounts:   0,
									NumReadonlyUnsignedAccounts: 3,
								},
								Accounts: []common.PublicKey{
									common.PublicKeyFromString("9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX"),
									common.PublicKeyFromString("3UbyTNpi3omt7hfEqQRB5844VANQFWiq8uEDNCrSwAVG"),
									common.PublicKeyFromString("SysvarS1otHashes111111111111111111111111111"),
									common.PublicKeyFromString("SysvarC1ock11111111111111111111111111111111"),
									common.PublicKeyFromString("Vote111111111111111111111111111111111111111"),
								},
								RecentBlockHash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
								Instructions: []types.CompiledInstruction{
									{
										ProgramIDIndex: 4,
										Accounts:       []int{1, 2, 3, 1},
										Data:           []byte{0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x8, 0x74, 0x95, 0x21, 0xce, 0x41, 0x7a, 0xd7, 0x5f, 0xf8, 0xd9, 0x11, 0x8b, 0x86, 0x11, 0xd8, 0x3d, 0xda, 0x6c, 0xd2, 0xef, 0xa5, 0xb2, 0x3c, 0xde, 0xa9, 0xc1, 0xec, 0x4b, 0xcd, 0xb0, 0x1, 0x19, 0x5a, 0x43, 0x61, 0x0, 0x0, 0x0, 0x0},
									},
								},
								AddressLookupTables: []types.CompiledAddressLookupTable{},
							},
						},
					},
				},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
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
			got, err := c.GetBlock(tt.args.ctx, tt.args.slot)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
			server.Close()
		})
	}
}

func TestGetBlockWithConfig(t *testing.T) {
	type args struct {
		ctx  context.Context
		slot uint64
		cfg  rpc.GetBlockConfig
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         GetBlockResponse
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64"}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}],"transactions":[{"meta":{"err":null,"fee":10000,"innerInstructions":[],"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499999835001,1000000000000000,143487360,1169280,1],"postTokenBalances":[],"preBalances":[499999845001,1000000000000000,143487360,1169280,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"transaction":["AnXU8JYCIrc73JwxK9traTSp3EZdmnJp0B5luW8CCzr7GnFd/SjIMXiG4qbN5CwyEVhbpORzBUpB/253cNtS1A+0rWE+nrDqWRQ2OVU727PU4NtR611jY+10Q+F6lCZDsJt46b6oXz3PN5WGxTQk7mC4YhCbYsTcalWBkltA8KgPAgADBXszyT4GLb26BFuAAUXtW0B75zurDhXE7UOYKHFkpIlKJMmZpq+FRXTx8jzBMy1YsdkCo0kyLDdF2Q3NhXRdEosGp9UXGS8Kr8byZeP7d8x62oLFKdC+OxNuLQBVIAAAAAan1RcYx3TJKFZjmGkdXraLXrijm0ttXHNVWyEAAAAAB2FIHTV0dLt8TXYk69O9s9g1XnPREEP8DaNTgAAAAACrUBylgzc0SSCUPSfMJC3TI6KJEzs834KdMIMJci+UYAEEBAECAwE9AgAAAAEAAAAAAAAAIAAAAAAAAAAGCHSVIc5Betdf+NkRi4YR2D3abNLvpbI83qnB7EvNsAEZWkNhAAAAAA==","base64"]}]},"id":1}`,
			args: args{
				context.Background(),
				33,
				rpc.GetBlockConfig{
					Encoding: rpc.GetBlockConfigEncodingBase64,
				},
			},
			want: GetBlockResponse{
				ParentSlot:        32,
				BlockHeight:       pointer.Get[int64](33),
				BlockTime:         pointer.Get[int64](1631803928),
				PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
				Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
				Rewards: []rpc.GetBlockReward{
					{
						Pubkey:       "9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX",
						Lamports:     5000,
						PostBalances: 499999840001,
						RewardType:   "Fee",
						Commission:   nil,
					},
				},
				Transactions: []GetBlockTransaction{
					{
						Meta: &TransactionMeta{
							Err: nil,
							Fee: 10000,
							PreBalances: []int64{
								499999845001,
								1000000000000000,
								143487360,
								1169280,
								1,
							},
							PostBalances: []int64{
								499999835001,
								1000000000000000,
								143487360,
								1169280,
								1,
							},
							PreTokenBalances:  []rpc.TransactionMetaTokenBalance{},
							PostTokenBalances: []rpc.TransactionMetaTokenBalance{},
							LogMessages: []string{
								"Program Vote111111111111111111111111111111111111111 invoke [1]",
								"Program Vote111111111111111111111111111111111111111 success",
							},
							InnerInstructions: []InnerInstruction{},
						},
						Transaction: types.Transaction{
							Signatures: []types.Signature{
								[]byte{0x75, 0xd4, 0xf0, 0x96, 0x2, 0x22, 0xb7, 0x3b, 0xdc, 0x9c, 0x31, 0x2b, 0xdb, 0x6b, 0x69, 0x34, 0xa9, 0xdc, 0x46, 0x5d, 0x9a, 0x72, 0x69, 0xd0, 0x1e, 0x65, 0xb9, 0x6f, 0x2, 0xb, 0x3a, 0xfb, 0x1a, 0x71, 0x5d, 0xfd, 0x28, 0xc8, 0x31, 0x78, 0x86, 0xe2, 0xa6, 0xcd, 0xe4, 0x2c, 0x32, 0x11, 0x58, 0x5b, 0xa4, 0xe4, 0x73, 0x5, 0x4a, 0x41, 0xff, 0x6e, 0x77, 0x70, 0xdb, 0x52, 0xd4, 0xf},
								[]byte{0xb4, 0xad, 0x61, 0x3e, 0x9e, 0xb0, 0xea, 0x59, 0x14, 0x36, 0x39, 0x55, 0x3b, 0xdb, 0xb3, 0xd4, 0xe0, 0xdb, 0x51, 0xeb, 0x5d, 0x63, 0x63, 0xed, 0x74, 0x43, 0xe1, 0x7a, 0x94, 0x26, 0x43, 0xb0, 0x9b, 0x78, 0xe9, 0xbe, 0xa8, 0x5f, 0x3d, 0xcf, 0x37, 0x95, 0x86, 0xc5, 0x34, 0x24, 0xee, 0x60, 0xb8, 0x62, 0x10, 0x9b, 0x62, 0xc4, 0xdc, 0x6a, 0x55, 0x81, 0x92, 0x5b, 0x40, 0xf0, 0xa8, 0xf},
							},
							Message: types.Message{
								Version: types.MessageVersionLegacy,
								Header: types.MessageHeader{
									NumRequireSignatures:        2,
									NumReadonlySignedAccounts:   0,
									NumReadonlyUnsignedAccounts: 3,
								},
								Accounts: []common.PublicKey{
									common.PublicKeyFromString("9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX"),
									common.PublicKeyFromString("3UbyTNpi3omt7hfEqQRB5844VANQFWiq8uEDNCrSwAVG"),
									common.PublicKeyFromString("SysvarS1otHashes111111111111111111111111111"),
									common.PublicKeyFromString("SysvarC1ock11111111111111111111111111111111"),
									common.PublicKeyFromString("Vote111111111111111111111111111111111111111"),
								},
								RecentBlockHash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
								Instructions: []types.CompiledInstruction{
									{
										ProgramIDIndex: 4,
										Accounts:       []int{1, 2, 3, 1},
										Data:           []byte{0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x8, 0x74, 0x95, 0x21, 0xce, 0x41, 0x7a, 0xd7, 0x5f, 0xf8, 0xd9, 0x11, 0x8b, 0x86, 0x11, 0xd8, 0x3d, 0xda, 0x6c, 0xd2, 0xef, 0xa5, 0xb2, 0x3c, 0xde, 0xa9, 0xc1, 0xec, 0x4b, 0xcd, 0xb0, 0x1, 0x19, 0x5a, 0x43, 0x61, 0x0, 0x0, 0x0, 0x0},
									},
								},
								AddressLookupTables: []types.CompiledAddressLookupTable{},
							},
						},
					},
				},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
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
			got, err := c.GetBlockWithConfig(tt.args.ctx, tt.args.slot, tt.args.cfg)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
			server.Close()
		})
	}
}

func TestGetMinimumBalanceForRentExemption(t *testing.T) {
	type args struct {
		ctx     context.Context
		dataLen uint64
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         uint64
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getMinimumBalanceForRentExemption", "params":[100]}`,
			responseBody: `{"jsonrpc":"2.0","result":1586880,"id":1}`,
			args: args{
				context.Background(),
				100,
			},
			want: 1586880,
			err:  nil,
		},
	}
	for _, tt := range tests {
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
			got, err := c.GetMinimumBalanceForRentExemption(tt.args.ctx, tt.args.dataLen)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
			server.Close()
		})
	}
}

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
