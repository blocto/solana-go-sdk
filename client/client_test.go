package client

import (
	"context"
	"io/ioutil"
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
		want         *GetTransactionResponse
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF", {"encoding":"base64"}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"blockTime":1631380624,"meta":{"err":null,"fee":5000,"innerInstructions":[{"index":0,"instructions":[{"accounts":[0,1],"data":"3Bxs4h24hBtQy9rw","programIdIndex":3},{"accounts":[1],"data":"9krTDU2LzCSUJuVZ","programIdIndex":3},{"accounts":[1],"data":"SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg","programIdIndex":3},{"accounts":[1,2,0,5],"data":"2","programIdIndex":4}]}],"logMessages":["Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]","Program log: Transfer 2039280 lamports to the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Allocate space for the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Assign the associated token account to the SPL Token program","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Initialize the associated token account","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]","Program log: Instruction: InitializeAccount","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3412 of 177045 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 27016 of 200000 compute units","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success"],"postBalances":[38024615601,2039280,1461600,1,1089991680,1,898174080],"postTokenBalances":[{"accountIndex":1,"mint":"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","uiTokenAmount":{"amount":"0","decimals":9,"uiAmount":null,"uiAmountString":"0"}}],"preBalances":[38026659881,0,1461600,1,1089991680,1,898174080],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":80218681,"transaction":["AaEGlsrjwHOjXODEvEGb5Zade8QelkWx2l9VvseP/g1olewFxKkJEwRDJyZ2wel8p2Dilp3wnBu6AEbRB4LthwABAAUHEJZZF158ZDMhpe1GQqAnsKvZe43ZetG8xtxkcThszdyUJGGIseU8n4crN7gTTkkjZvTPQVkY2NPZnO+5BTpTqzO9mOFbcsDwmqTwyIZje2Ppd9PY6hWpndBzwVYYhseQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqQan1RcZLFxRIYzJTD1K8X9Y2u4Im6H9ROPb2YoAAAAAjJclj04kifG7PRApFI4NgwtaE5na/xCEBI572Nvp+FnrFE6iq1ZbCKVJ+UiBaEkoE9dTFWqba+nWyTsH21qhygEGBwABAAIDBAUA","base64"]},"id":1}`,
			args: args{
				context.Background(),
				"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
			},
			want: &GetTransactionResponse{
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
									Accounts:       []int{1},
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
							UITokenAmount: rpc.TokenAccountBalance{
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
						Version: types.MessageVersionLegacy,
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
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp", {"encoding":"base64"}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"blockTime":1631744159,"meta":{"err":null,"fee":5000,"innerInstructions":[],"logMessages":["Program H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm invoke [1]","Program consumption: 199622 units remaining","Program consumption: 199621 units remaining","Program consumption: 199620 units remaining","Program log: update here","Program log: update here","Program log: program id H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm","Program log: accounts [AccountInfo { key: 11111111111111111111111111111111 owner: NativeLoader1111111111111111111111111111111 is_signer: false is_writable: false executable: true rent_epoch: 55 lamports: 1 data.len: 14  data: 73797374656d5f70726f6772616d ... }]","Program log: data []","Program H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm consumed 32192 of 200000 compute units","Program H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm success"],"postBalances":[109107166519,1,1141440],"postTokenBalances":[],"preBalances":[109107171519,1,1141440],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":81103164,"transaction":["ATWlpjPdm+8muj2Gw5etBJABHggGthzIiQxcFO+Tizs4krrFB2rWui2DBN+Zz/N0x8tKp6731l5ZWnigQDuMQQ0BAAEDBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAO9leZEN4av+0/cP0pb3UfT4YZeMVMzaq+GAwcjoYx/Y4ueeH6yFx+7mz1QHKS/wM0DumafPn5kBqjpYmzd0eeABAgEBAA==","base64"]},"id":1}`,
			args: args{
				context.Background(),
				"25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp",
			},
			want: &GetTransactionResponse{
				Slot:      81103164,
				BlockTime: pointer.Int64(1631744159),
				Meta: &TransactionMeta{
					Fee:               5000,
					InnerInstructions: []TransactionMetaInnerInstruction{},
					LogMessages: []string{
						"Program H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm invoke [1]",
						"Program consumption: 199622 units remaining",
						"Program consumption: 199621 units remaining",
						"Program consumption: 199620 units remaining",
						"Program log: update here",
						"Program log: update here",
						"Program log: program id H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm",
						"Program log: accounts [AccountInfo { key: 11111111111111111111111111111111 owner: NativeLoader1111111111111111111111111111111 is_signer: false is_writable: false executable: true rent_epoch: 55 lamports: 1 data.len: 14  data: 73797374656d5f70726f6772616d ... }]",
						"Program log: data []",
						"Program H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm consumed 32192 of 200000 compute units",
						"Program H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm success",
					},
					PreBalances: []int64{
						109107171519,
						1,
						1141440,
					},
					PostBalances: []int64{
						109107166519,
						1,
						1141440,
					},
					PreTokenBalances:  []rpc.TransactionMetaTokenBalance{},
					PostTokenBalances: []rpc.TransactionMetaTokenBalance{},
				},
				Transaction: types.Transaction{
					Signatures: []types.Signature{[]byte{0x35, 0xa5, 0xa6, 0x33, 0xdd, 0x9b, 0xef, 0x26, 0xba, 0x3d, 0x86, 0xc3, 0x97, 0xad, 0x4, 0x90, 0x1, 0x1e, 0x8, 0x6, 0xb6, 0x1c, 0xc8, 0x89, 0xc, 0x5c, 0x14, 0xef, 0x93, 0x8b, 0x3b, 0x38, 0x92, 0xba, 0xc5, 0x7, 0x6a, 0xd6, 0xba, 0x2d, 0x83, 0x4, 0xdf, 0x99, 0xcf, 0xf3, 0x74, 0xc7, 0xcb, 0x4a, 0xa7, 0xae, 0xf7, 0xd6, 0x5e, 0x59, 0x5a, 0x78, 0xa0, 0x40, 0x3b, 0x8c, 0x41, 0xd}},
					Message: types.Message{
						Version: types.MessageVersionLegacy,
						Header: types.MessageHeader{
							NumRequireSignatures:        1,
							NumReadonlySignedAccounts:   0,
							NumReadonlyUnsignedAccounts: 1,
						},
						Accounts: []common.PublicKey{
							common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
							common.PublicKeyFromString("11111111111111111111111111111111"),
							common.PublicKeyFromString("H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm"),
						},
						Instructions: []types.CompiledInstruction{
							{
								ProgramIDIndex: 2,
								Accounts:       []int{1},
								Data:           []byte{},
							},
						},
						RecentBlockHash: "GGjz3cjABNTaCA9w1pP3y5FtpsZtKLR5taBk5MF8ijQj",
					},
				},
			},
			err: nil,
		},
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp", {"encoding":"base64"}]}`,
			responseBody: `{"jsonrpc":"2.0","result":null,"id":1}`,
			args: args{
				context.Background(),
				"25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp",
			},
			want: nil,
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

func TestClient_GetClusterNodes(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         []ClusterNode
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getClusterNodes"}`,
			responseBody: `{"jsonrpc":"2.0","result":[{"featureSet":1797267350,"gossip":"127.0.0.1:1024","pubkey":"8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn","rpc":"127.0.0.1:8899","shredVersion":23492,"tpu":"127.0.0.1:1027","version":"1.8.1"}],"id":1}`,
			args: args{
				context.Background(),
			},
			want: []ClusterNode{
				{
					Pubkey:       common.PublicKeyFromString("8gNdbr9dG6oj8bhaQ44icyMYsfG3t1dhXKUJLGVav4tn"),
					Gossip:       pointer.String("127.0.0.1:1024"),
					Tpu:          pointer.String("127.0.0.1:1027"),
					Rpc:          pointer.String("127.0.0.1:8899"),
					Version:      pointer.String("1.8.1"),
					FeatureSet:   pointer.Uint32(1797267350),
					ShredVersion: pointer.Get[uint16](23492),
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
			got, err := c.GetClusterNodes(tt.args.ctx)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
			server.Close()
		})
	}
}

func TestClient_GetAccountInfo(t *testing.T) {
	type args struct {
		ctx        context.Context
		base58Addr string
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         AccountInfo
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64"}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317717},"value":{"data":["AQAAAAY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWzeAAAAAAAAAAAJAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
			args: args{
				ctx:        context.Background(),
				base58Addr: "F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
			},
			want: AccountInfo{
				RentEpoch:  178,
				Lamports:   1461600,
				Owner:      common.TokenProgramID,
				Executable: false,
				Data:       []byte{0x1, 0x0, 0x0, 0x0, 0x6, 0x3e, 0x70, 0xd9, 0x91, 0x57, 0x98, 0xdc, 0xa3, 0xee, 0xdb, 0x91, 0x37, 0xcf, 0x64, 0xc7, 0xf, 0xfa, 0xfb, 0x62, 0xcd, 0xd3, 0x82, 0x81, 0x5, 0x25, 0xb, 0xd7, 0x6b, 0x1d, 0x6c, 0xde, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
				got, err := c.GetAccountInfo(tt.args.ctx, tt.args.base58Addr)
				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
				server.Close()
			})
		})
	}
}

func TestClient_GetAccountInfoWithConfig(t *testing.T) {
	type args struct {
		ctx        context.Context
		base58Addr string
		cfg        GetAccountInfoConfig
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         AccountInfo
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb", {"encoding": "base64", "dataSlice": {"offset": 4, "length": 32}}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":95887894},"value":{"data":["Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4=","base64"],"executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":221}},"id":1}`,
			args: args{
				ctx:        context.Background(),
				base58Addr: "F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
				cfg: GetAccountInfoConfig{
					DataSlice: &rpc.DataSlice{
						Offset: 4,
						Length: 32,
					},
				},
			},
			want: AccountInfo{
				RentEpoch:  221,
				Lamports:   1461600,
				Owner:      common.TokenProgramID,
				Executable: false,
				Data:       []byte{0x6, 0x3e, 0x70, 0xd9, 0x91, 0x57, 0x98, 0xdc, 0xa3, 0xee, 0xdb, 0x91, 0x37, 0xcf, 0x64, 0xc7, 0xf, 0xfa, 0xfb, 0x62, 0xcd, 0xd3, 0x82, 0x81, 0x5, 0x25, 0xb, 0xd7, 0x6b, 0x1d, 0x6c, 0xde},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
				got, err := c.GetAccountInfoWithConfig(tt.args.ctx, tt.args.base58Addr, tt.args.cfg)
				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
				server.Close()
			})
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
				ConfirmationStatus: (*rpc.Commitment)(pointer.String(string(rpc.CommitmentConfirmed))),
				Confirmations:      pointer.Uint64(25),
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
				ConfirmationStatus: (*rpc.Commitment)(pointer.String(string(rpc.CommitmentFinalized))),
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
					body, err := ioutil.ReadAll(req.Body)
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
				ConfirmationStatus: (*rpc.Commitment)(pointer.String(string(rpc.CommitmentConfirmed))),
				Confirmations:      pointer.Uint64(25),
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
				ConfirmationStatus: (*rpc.Commitment)(pointer.String(string(rpc.CommitmentFinalized))),
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
					body, err := ioutil.ReadAll(req.Body)
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
					ConfirmationStatus: (*rpc.Commitment)(pointer.String(string(rpc.CommitmentConfirmed))),
					Confirmations:      pointer.Uint64(25),
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
					ConfirmationStatus: (*rpc.Commitment)(pointer.String(string(rpc.CommitmentFinalized))),
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
					body, err := ioutil.ReadAll(req.Body)
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
					ConfirmationStatus: (*rpc.Commitment)(pointer.String(string(rpc.CommitmentConfirmed))),
					Confirmations:      pointer.Uint64(25),
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
					ConfirmationStatus: (*rpc.Commitment)(pointer.String(string(rpc.CommitmentFinalized))),
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
					body, err := ioutil.ReadAll(req.Body)
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

func TestClient_SimulateTransaction(t *testing.T) {
	type args struct {
		ctx context.Context
		tx  types.Transaction
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
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"simulateTransaction", "params":["AYiC5ORGn1mWvOc0GvqzBeXCW8t+pGUkXLue8SQ3LEr/ByvfYmkCQEDEAQqfTNN+yA1DEibKeBPxMvs087fuxAABAAEDJGfwe2rYvkoBCI3IHIybkz+rysOkG7upqFJ/ZsFJmrgkaRt6Oq89b8zO7fpC3krN2PFyLB32negX+tFWkhOEaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAaHN3i7POgmrBKRwjsBGZVnodEydldAs+WyXzyAnMowIBAgIAAQwCAAAAAQAAAAAAAAA=", {"encoding": "base64"}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":95946904},"value":{"accounts":null,"err":null,"logs":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"]}},"id":1}`,
			args: args{
				ctx: context.Background(),
				tx: types.Transaction{
					Signatures: []types.Signature{
						{0x88, 0x82, 0xe4, 0xe4, 0x46, 0x9f, 0x59, 0x96, 0xbc, 0xe7, 0x34, 0x1a, 0xfa, 0xb3, 0x5, 0xe5, 0xc2, 0x5b, 0xcb, 0x7e, 0xa4, 0x65, 0x24, 0x5c, 0xbb, 0x9e, 0xf1, 0x24, 0x37, 0x2c, 0x4a, 0xff, 0x7, 0x2b, 0xdf, 0x62, 0x69, 0x2, 0x40, 0x40, 0xc4, 0x1, 0xa, 0x9f, 0x4c, 0xd3, 0x7e, 0xc8, 0xd, 0x43, 0x12, 0x26, 0xca, 0x78, 0x13, 0xf1, 0x32, 0xfb, 0x34, 0xf3, 0xb7, 0xee, 0xc4, 0x0},
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
						RecentBlockhash: "82jXAmeXyz6ZJxuw1etGweaDEBzPnKc1XbV2KgYWAstM",
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
					body, err := ioutil.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.requestBody, string(body))
					n, err := rw.Write([]byte(tt.responseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.responseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.SimulateTransaction(tt.args.ctx, tt.args.tx)
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
					body, err := ioutil.ReadAll(req.Body)
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

func TestClient_GetLatestBlockhash(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         rpc.GetLatestBlockhashValue
		err          error
	}{
		{
			requestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash"}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112872139},"value":{"blockhash":"9K9GnvWXn9zYitQdHUSYzvjLjebnviwEFaWgWqHDU3ve","lastValidBlockHeight":92248597}},"id":1}`,
			args: args{
				ctx: context.Background(),
			},
			want: rpc.GetLatestBlockhashValue{
				Blockhash:              "9K9GnvWXn9zYitQdHUSYzvjLjebnviwEFaWgWqHDU3ve",
				LatestValidBlockHeight: 92248597,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
				got, err := c.GetLatestBlockhash(tt.args.ctx)
				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
				server.Close()
			})
		})
	}
}

func TestClient_GetLatestBlockhashWithConfig(t *testing.T) {
	type Args struct {
		ctx context.Context
		cfg GetLatestBlockhashConfig
	}
	tests := []struct {
		Name         string
		RequestBody  string
		ResponseBody string
		Args         Args
		Want         rpc.GetLatestBlockhashValue
		Err          error
	}{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getLatestBlockhash", "params":[{"commitment": "processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112871314},"value":{"blockhash":"3H2pwJD6pTrEveh5xcwHXToLn7txt5uTW6CPzCan4ZKL","lastValidBlockHeight":92247902}},"id":1}`,
			Args: Args{
				ctx: context.Background(),
				cfg: GetLatestBlockhashConfig{
					Commitment: rpc.CommitmentProcessed,
				},
			},
			Want: rpc.GetLatestBlockhashValue{
				Blockhash:              "3H2pwJD6pTrEveh5xcwHXToLn7txt5uTW6CPzCan4ZKL",
				LatestValidBlockHeight: 92247902,
			},
			Err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Run(tt.Name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := ioutil.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.RequestBody, string(body))
					n, err := rw.Write([]byte(tt.ResponseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.ResponseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.GetLatestBlockhashWithConfig(tt.Args.ctx, tt.Args.cfg)
				assert.Equal(t, tt.Err, err)
				assert.Equal(t, tt.Want, got)
				server.Close()
			})
		})
	}
}

func TestClient_IsBlockhashValid(t *testing.T) {
	type Args struct {
		ctx       context.Context
		blockhash string
	}
	tests := []struct {
		Name         string
		RequestBody  string
		ResponseBody string
		Args         Args
		Want         bool
		Err          error
	}{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params":["14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112890169},"value":false},"id":1}`,
			Args: Args{
				ctx:       context.Background(),
				blockhash: "14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N",
			},
			Want: false,
			Err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Run(tt.Name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := ioutil.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.RequestBody, string(body))
					n, err := rw.Write([]byte(tt.ResponseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.ResponseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.IsBlockhashValid(tt.Args.ctx, tt.Args.blockhash)
				assert.Equal(t, tt.Err, err)
				assert.Equal(t, tt.Want, got)
				server.Close()
			})
		})
	}
}

func TestClient_IsBlockhashValidWithConfig(t *testing.T) {
	type Args struct {
		ctx       context.Context
		blockhash string
		cfg       IsBlockhashConfig
	}
	tests := []struct {
		Name         string
		RequestBody  string
		ResponseBody string
		Args         Args
		Want         bool
		Err          error
	}{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"isBlockhashValid", "params":["14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N", {"commitment": "processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112890231},"value":true},"id":1}`,
			Args: Args{
				ctx:       context.Background(),
				blockhash: "14PVzxGGU4WQ7qbQffn3XJV1pasafs4wApFUs5sps89N",
				cfg: IsBlockhashConfig{
					Commitment: rpc.CommitmentProcessed,
				},
			},
			Want: true,
			Err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Run(tt.Name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := ioutil.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.RequestBody, string(body))
					n, err := rw.Write([]byte(tt.ResponseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.ResponseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.IsBlockhashValidWithConfig(tt.Args.ctx, tt.Args.blockhash, tt.Args.cfg)
				assert.Equal(t, tt.Err, err)
				assert.Equal(t, tt.Want, got)
				server.Close()
			})
		})
	}
}

func TestClient_GetFeeForMessage(t *testing.T) {
	type Args struct {
		ctx     context.Context
		message types.Message
	}
	tests := []struct {
		Name         string
		RequestBody  string
		ResponseBody string
		Args         Args
		Want         *uint64
		Err          error
	}{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHtjfNL0642B6qTQXHXH7TCzYl7SJhtOdw7hTYnd/mMpAQICAAEMAgAAAAEAAAAAAAAA"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112884237},"value":null},"id":1}`,
			Args: Args{
				ctx: context.Background(),
				message: types.NewMessage(types.NewMessageParam{
					FeePayer: common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
					Instructions: []types.Instruction{
						system.Transfer(system.TransferParam{
							From:   common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
							To:     common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
							Amount: 1,
						}),
					},
					RecentBlockhash: "9Jf8nJep3oubyeGVernU2kVVmrmnADJFyHu8Bmq23S2C",
				}),
			},
			Want: nil,
			Err:  nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHtjfNL0642B6qTQXHXH7TCzYl7SJhtOdw7hTYnd/mMpAQICAAEMAgAAAAEAAAAAAAAA"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112884237},"value":5000},"id":1}`,
			Args: Args{
				ctx: context.Background(),
				message: types.NewMessage(types.NewMessageParam{
					FeePayer: common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
					Instructions: []types.Instruction{
						system.Transfer(system.TransferParam{
							From:   common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
							To:     common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
							Amount: 1,
						}),
					},
					RecentBlockhash: "9Jf8nJep3oubyeGVernU2kVVmrmnADJFyHu8Bmq23S2C",
				}),
			},
			Want: pointer.Uint64(5000),
			Err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Run(tt.Name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := ioutil.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.RequestBody, string(body))
					n, err := rw.Write([]byte(tt.ResponseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.ResponseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.GetFeeForMessage(tt.Args.ctx, tt.Args.message)
				assert.Equal(t, tt.Err, err)
				assert.Equal(t, tt.Want, got)
				server.Close()
			})
		})
	}
}

func TestClient_GetFeeForMessageWithConfig(t *testing.T) {
	type Args struct {
		ctx     context.Context
		message types.Message
		cfg     GetFeeForMessageConfig
	}
	tests := []struct {
		Name         string
		RequestBody  string
		ResponseBody string
		Args         Args
		Want         *uint64
		Err          error
	}{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHtjfNL0642B6qTQXHXH7TCzYl7SJhtOdw7hTYnd/mMpAQICAAEMAgAAAAEAAAAAAAAA", {"commitment":"processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112884237},"value":null},"id":1}`,
			Args: Args{
				ctx: context.Background(),
				message: types.NewMessage(types.NewMessageParam{
					FeePayer: common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
					Instructions: []types.Instruction{
						system.Transfer(system.TransferParam{
							From:   common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
							To:     common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
							Amount: 1,
						}),
					},
					RecentBlockhash: "9Jf8nJep3oubyeGVernU2kVVmrmnADJFyHu8Bmq23S2C",
				}),
				cfg: GetFeeForMessageConfig{
					Commitment: rpc.CommitmentProcessed,
				},
			},
			Want: nil,
			Err:  nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHtjfNL0642B6qTQXHXH7TCzYl7SJhtOdw7hTYnd/mMpAQICAAEMAgAAAAEAAAAAAAAA", {"commitment":"processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112884237},"value":5000},"id":1}`,
			Args: Args{
				ctx: context.Background(),
				message: types.NewMessage(types.NewMessageParam{
					FeePayer: common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
					Instructions: []types.Instruction{
						system.Transfer(system.TransferParam{
							From:   common.PublicKeyFromString("3T7btuZcLDHxRqKJ7YzxH22toGhGedaJnecn5h4mBeL7"),
							To:     common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
							Amount: 1,
						}),
					},
					RecentBlockhash: "9Jf8nJep3oubyeGVernU2kVVmrmnADJFyHu8Bmq23S2C",
				}),
				cfg: GetFeeForMessageConfig{
					Commitment: rpc.CommitmentProcessed,
				},
			},
			Want: pointer.Uint64(5000),
			Err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Run(tt.Name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := ioutil.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.RequestBody, string(body))
					n, err := rw.Write([]byte(tt.ResponseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.ResponseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.GetFeeForMessageWithConfig(tt.Args.ctx, tt.Args.message, tt.Args.cfg)
				assert.Equal(t, tt.Err, err)
				assert.Equal(t, tt.Want, got)
				server.Close()
			})
		})
	}
}
