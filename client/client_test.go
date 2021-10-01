package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestGetTransaction(t *testing.T) {
	type args struct {
		ctx    context.Context
		txhash string
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         GetTransactionResponse
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF", {"encoding":"base64"}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"blockTime":1631380624,"meta":{"err":null,"fee":5000,"innerInstructions":[{"index":0,"instructions":[{"accounts":[0,1],"data":"3Bxs4h24hBtQy9rw","programIdIndex":3},{"accounts":[1],"data":"9krTDU2LzCSUJuVZ","programIdIndex":3},{"accounts":[1],"data":"SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg","programIdIndex":3},{"accounts":[1,2,0,5],"data":"2","programIdIndex":4}]}],"logMessages":["Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]","Program log: Transfer 2039280 lamports to the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Allocate space for the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Assign the associated token account to the SPL Token program","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Initialize the associated token account","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]","Program log: Instruction: InitializeAccount","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3412 of 177045 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 27016 of 200000 compute units","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success"],"postBalances":[38024615601,2039280,1461600,1,1089991680,1,898174080],"postTokenBalances":[{"accountIndex":1,"mint":"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","uiTokenAmount":{"amount":"0","decimals":9,"uiAmount":null,"uiAmountString":"0"}}],"preBalances":[38026659881,0,1461600,1,1089991680,1,898174080],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":80218681,"transaction":["AaEGlsrjwHOjXODEvEGb5Zade8QelkWx2l9VvseP/g1olewFxKkJEwRDJyZ2wel8p2Dilp3wnBu6AEbRB4LthwABAAUHEJZZF158ZDMhpe1GQqAnsKvZe43ZetG8xtxkcThszdyUJGGIseU8n4crN7gTTkkjZvTPQVkY2NPZnO+5BTpTqzO9mOFbcsDwmqTwyIZje2Ppd9PY6hWpndBzwVYYhseQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqQan1RcZLFxRIYzJTD1K8X9Y2u4Im6H9ROPb2YoAAAAAjJclj04kifG7PRApFI4NgwtaE5na/xCEBI572Nvp+FnrFE6iq1ZbCKVJ+UiBaEkoE9dTFWqba+nWyTsH21qhygEGBwABAAIDBAUA","base64"]},"id":1}`,
			args: args{
				context.Background(),
				"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
			},
			want: GetTransactionResponse{
				Slot:      80218681,
				BlockTime: pointer.Int64(1631380624),
				Meta: &TransactionMeta{
					Fee: 5000,
					InnerInstructions: []TransactionMetaInnerInstruction{
						{
							Index: 0,
							Instructions: []types.CompiledInstruction{
								{
									ProgramIDIndex: 3,
									Data:           []byte{0x2, 0x0, 0x0, 0x0, 0xf0, 0x1d, 0x1f, 0x0, 0x0, 0x0, 0x0, 0x0},
									Accounts:       []int{0, 1},
								},
								{
									ProgramIDIndex: 3,
									Data:           []byte{0x8, 0x0, 0x0, 0x0, 0xa5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},

									Accounts: []int{1},
								},
								{
									ProgramIDIndex: 3,
									Data:           []byte{0x1, 0x0, 0x0, 0x0, 0x6, 0xdd, 0xf6, 0xe1, 0xd7, 0x65, 0xa1, 0x93, 0xd9, 0xcb, 0xe1, 0x46, 0xce, 0xeb, 0x79, 0xac, 0x1c, 0xb4, 0x85, 0xed, 0x5f, 0x5b, 0x37, 0x91, 0x3a, 0x8c, 0xf5, 0x85, 0x7e, 0xff, 0x0, 0xa9},
									Accounts:       []int{1},
								},
								{
									ProgramIDIndex: 4,
									Data:           []byte{0x1},
									Accounts:       []int{1, 2, 0, 5},
								},
							},
						},
					},
					LogMessages: []string{
						"Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]",
						"Program log: Transfer 2039280 lamports to the associated token account",
						"Program 11111111111111111111111111111111 invoke [2]",
						"Program 11111111111111111111111111111111 success",
						"Program log: Allocate space for the associated token account",
						"Program 11111111111111111111111111111111 invoke [2]",
						"Program 11111111111111111111111111111111 success",
						"Program log: Assign the associated token account to the SPL Token program",
						"Program 11111111111111111111111111111111 invoke [2]",
						"Program 11111111111111111111111111111111 success",
						"Program log: Initialize the associated token account",
						"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]",
						"Program log: Instruction: InitializeAccount",
						"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3412 of 177045 compute units",
						"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success",
						"Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 27016 of 200000 compute units",
						"Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success",
					},
					PreBalances: []int64{
						38026659881,
						0,
						1461600,
						1,
						1089991680,
						1,
						898174080,
					},
					PostBalances: []int64{
						38024615601,
						2039280,
						1461600,
						1,
						1089991680,
						1,
						898174080,
					},
					PreTokenBalances: []rpc.TransactionMetaTokenBalance{},
					PostTokenBalances: []rpc.TransactionMetaTokenBalance{
						{
							AccountIndex: 1,
							Mint:         "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
							UITokenAmount: rpc.GetTokenAccountBalanceResultValue{
								Amount:         "0",
								Decimals:       9,
								UIAmountString: "0",
							},
						},
					},
				},
				Transaction: types.Transaction{
					Signatures: []types.Signature{[]byte{0xa1, 0x6, 0x96, 0xca, 0xe3, 0xc0, 0x73, 0xa3, 0x5c, 0xe0, 0xc4, 0xbc, 0x41, 0x9b, 0xe5, 0x96, 0x9d, 0x7b, 0xc4, 0x1e, 0x96, 0x45, 0xb1, 0xda, 0x5f, 0x55, 0xbe, 0xc7, 0x8f, 0xfe, 0xd, 0x68, 0x95, 0xec, 0x5, 0xc4, 0xa9, 0x9, 0x13, 0x4, 0x43, 0x27, 0x26, 0x76, 0xc1, 0xe9, 0x7c, 0xa7, 0x60, 0xe2, 0x96, 0x9d, 0xf0, 0x9c, 0x1b, 0xba, 0x0, 0x46, 0xd1, 0x7, 0x82, 0xed, 0x87, 0x0}},
					Message: types.Message{
						Header: types.MessageHeader{
							NumRequireSignatures:        1,
							NumReadonlySignedAccounts:   0,
							NumReadonlyUnsignedAccounts: 5,
						},
						Accounts: []common.PublicKey{
							common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ"),
							common.PublicKeyFromString("AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"),
							common.PublicKeyFromString("4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"),
							common.PublicKeyFromString("11111111111111111111111111111111"),
							common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
							common.PublicKeyFromString("SysvarRent111111111111111111111111111111111"),
							common.PublicKeyFromString("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"),
						},
						Instructions: []types.CompiledInstruction{
							{
								ProgramIDIndex: 6,
								Accounts:       []int{0, 1, 0, 2, 3, 4, 5},
								Data:           []byte{},
							},
						},
						RecentBlockHash: "Gpemb2whtMogoSGVe5KMjuoueeqNNkQ1kKnw7fsYKZHj",
					},
				},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				body, err := ioutil.ReadAll(req.Body)
				assert.Nil(t, err)
				assert.JSONEq(t, tt.requestBody, string(body))
				n, err := rw.Write([]byte(tt.responseBody))
				assert.Nil(t, err)
				assert.Equal(t, len([]byte(tt.responseBody)), n)
			}))
			c := NewClient(server.URL)
			got, err := c.GetTransaction(tt.args.ctx, tt.args.txhash)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
			server.Close()
		})
	}
}

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
				ParentSLot:        32,
				BlockHeight:       pointer.Int64(33),
				BlockTime:         pointer.Int64(1631803928),
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
							InnerInstructions: []TransactionMetaInnerInstruction{},
						},
						Transaction: types.Transaction{
							Signatures: []types.Signature{
								[]byte{0x75, 0xd4, 0xf0, 0x96, 0x2, 0x22, 0xb7, 0x3b, 0xdc, 0x9c, 0x31, 0x2b, 0xdb, 0x6b, 0x69, 0x34, 0xa9, 0xdc, 0x46, 0x5d, 0x9a, 0x72, 0x69, 0xd0, 0x1e, 0x65, 0xb9, 0x6f, 0x2, 0xb, 0x3a, 0xfb, 0x1a, 0x71, 0x5d, 0xfd, 0x28, 0xc8, 0x31, 0x78, 0x86, 0xe2, 0xa6, 0xcd, 0xe4, 0x2c, 0x32, 0x11, 0x58, 0x5b, 0xa4, 0xe4, 0x73, 0x5, 0x4a, 0x41, 0xff, 0x6e, 0x77, 0x70, 0xdb, 0x52, 0xd4, 0xf},
								[]byte{0xb4, 0xad, 0x61, 0x3e, 0x9e, 0xb0, 0xea, 0x59, 0x14, 0x36, 0x39, 0x55, 0x3b, 0xdb, 0xb3, 0xd4, 0xe0, 0xdb, 0x51, 0xeb, 0x5d, 0x63, 0x63, 0xed, 0x74, 0x43, 0xe1, 0x7a, 0x94, 0x26, 0x43, 0xb0, 0x9b, 0x78, 0xe9, 0xbe, 0xa8, 0x5f, 0x3d, 0xcf, 0x37, 0x95, 0x86, 0xc5, 0x34, 0x24, 0xee, 0x60, 0xb8, 0x62, 0x10, 0x9b, 0x62, 0xc4, 0xdc, 0x6a, 0x55, 0x81, 0x92, 0x5b, 0x40, 0xf0, 0xa8, 0xf},
							},
							Message: types.Message{
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
				body, err := ioutil.ReadAll(req.Body)
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
				body, err := ioutil.ReadAll(req.Body)
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
