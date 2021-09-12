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
