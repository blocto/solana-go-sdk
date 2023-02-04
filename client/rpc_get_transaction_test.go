package client

import (
	"context"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

func TestClient_GetTransaction(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp", {"encoding":"base64", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":null,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTransaction(
						context.Background(),
						"25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp",
					)
				},
				ExpectedValue: (*Transaction)(nil),
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF", {"encoding":"base64", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1631380624,"meta":{"err":null,"fee":5000,"innerInstructions":[{"index":0,"instructions":[{"accounts":[0,1],"data":"3Bxs4h24hBtQy9rw","programIdIndex":3},{"accounts":[1],"data":"9krTDU2LzCSUJuVZ","programIdIndex":3},{"accounts":[1],"data":"SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg","programIdIndex":3},{"accounts":[1,2,0,5],"data":"2","programIdIndex":4}]}],"logMessages":["Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]","Program log: Transfer 2039280 lamports to the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Allocate space for the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Assign the associated token account to the SPL Token program","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Initialize the associated token account","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]","Program log: Instruction: InitializeAccount","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3412 of 177045 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 27016 of 200000 compute units","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success"],"postBalances":[38024615601,2039280,1461600,1,1089991680,1,898174080],"postTokenBalances":[{"accountIndex":1,"mint":"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","uiTokenAmount":{"amount":"0","decimals":9,"uiAmount":null,"uiAmountString":"0"}}],"preBalances":[38026659881,0,1461600,1,1089991680,1,898174080],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":80218681,"transaction":["AaEGlsrjwHOjXODEvEGb5Zade8QelkWx2l9VvseP/g1olewFxKkJEwRDJyZ2wel8p2Dilp3wnBu6AEbRB4LthwABAAUHEJZZF158ZDMhpe1GQqAnsKvZe43ZetG8xtxkcThszdyUJGGIseU8n4crN7gTTkkjZvTPQVkY2NPZnO+5BTpTqzO9mOFbcsDwmqTwyIZje2Ppd9PY6hWpndBzwVYYhseQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqQan1RcZLFxRIYzJTD1K8X9Y2u4Im6H9ROPb2YoAAAAAjJclj04kifG7PRApFI4NgwtaE5na/xCEBI572Nvp+FnrFE6iq1ZbCKVJ+UiBaEkoE9dTFWqba+nWyTsH21qhygEGBwABAAIDBAUA","base64"]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTransaction(
						context.Background(),
						"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
					)
				},
				ExpectedValue: &Transaction{
					Slot:      80218681,
					BlockTime: pointer.Get[int64](1631380624),
					Meta: &TransactionMeta{
						Fee: 5000,
						InnerInstructions: []InnerInstruction{
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
							RecentBlockHash:     "Gpemb2whtMogoSGVe5KMjuoueeqNNkQ1kKnw7fsYKZHj",
							AddressLookupTables: []types.CompiledAddressLookupTable{},
						},
					},
					AccountKeys: []common.PublicKey{
						common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ"),
						common.PublicKeyFromString("AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ"),
						common.PublicKeyFromString("4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3"),
						common.PublicKeyFromString("11111111111111111111111111111111"),
						common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						common.PublicKeyFromString("SysvarRent111111111111111111111111111111111"),
						common.PublicKeyFromString("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"),
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp", {"encoding":"base64", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1631744159,"meta":{"err":null,"fee":5000,"innerInstructions":[],"logMessages":["Program H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm invoke [1]","Program consumption: 199622 units remaining","Program consumption: 199621 units remaining","Program consumption: 199620 units remaining","Program log: update here","Program log: update here","Program log: program id H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm","Program log: accounts [AccountInfo { key: 11111111111111111111111111111111 owner: NativeLoader1111111111111111111111111111111 is_signer: false is_writable: false executable: true rent_epoch: 55 lamports: 1 data.len: 14  data: 73797374656d5f70726f6772616d ... }]","Program log: data []","Program H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm consumed 32192 of 200000 compute units","Program H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm success"],"postBalances":[109107166519,1,1141440],"postTokenBalances":[],"preBalances":[109107171519,1,1141440],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":81103164,"transaction":["ATWlpjPdm+8muj2Gw5etBJABHggGthzIiQxcFO+Tizs4krrFB2rWui2DBN+Zz/N0x8tKp6731l5ZWnigQDuMQQ0BAAEDBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAO9leZEN4av+0/cP0pb3UfT4YZeMVMzaq+GAwcjoYx/Y4ueeH6yFx+7mz1QHKS/wM0DumafPn5kBqjpYmzd0eeABAgEBAA==","base64"]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTransaction(
						context.Background(),
						"25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp",
					)
				},
				ExpectedValue: &Transaction{
					Slot:      81103164,
					BlockTime: pointer.Get[int64](1631744159),
					Meta: &TransactionMeta{
						Fee:               5000,
						InnerInstructions: []InnerInstruction{},
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
							RecentBlockHash:     "GGjz3cjABNTaCA9w1pP3y5FtpsZtKLR5taBk5MF8ijQj",
							AddressLookupTables: []types.CompiledAddressLookupTable{},
						},
					},
					AccountKeys: []common.PublicKey{
						common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
						common.PublicKeyFromString("11111111111111111111111111111111"),
						common.PublicKeyFromString("H7WBiBDaZpWwGfhPLmXrdD3r86d6eQfzb184a2arM7Bm"),
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params": ["4fSTSDTTuYa1XXAFxFenuY3SoZWUwCzpMq7kUiya1zW6uqqh6C76GFqTQ3wvegEbZhbPJyr33iDAbieQVWCtVXmf", {"encoding": "base64", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1675511254,"meta":{"computeUnitsConsumed":12344,"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":["F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw"],"writable":["3Yvq7e9UXLoFK4PKyxrpEA3y3TKmFK2Wb1f5tVFUgwPu","5McxjaxNKYLHtv9DqbMfoi6GNs7ZEMHGkJDrouPib4sW","GAXzq8BWdAWaS1kWFiL5tzV2h3AbRBtYGP5psNTWrM9g"]},"logMessages":["Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 400000 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 393828 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success"],"postBalances":[112595188235,2039280,934087680,2039280,2039280,2039280,1461600,1461600],"postTokenBalances":[{"accountIndex":1,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"101","decimals":0,"uiAmount":101.0,"uiAmountString":"101"}},{"accountIndex":3,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"99","decimals":0,"uiAmount":99.0,"uiAmountString":"99"}},{"accountIndex":4,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"99","decimals":0,"uiAmount":99.0,"uiAmountString":"99"}},{"accountIndex":5,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"101","decimals":0,"uiAmount":101.0,"uiAmountString":"101"}}],"preBalances":[112595193235,2039280,934087680,2039280,2039280,2039280,1461600,1461600],"preTokenBalances":[{"accountIndex":1,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}},{"accountIndex":3,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}},{"accountIndex":4,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}},{"accountIndex":5,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}}],"rewards":[],"status":{"Ok":null}},"slot":193487858,"transaction":["AbczATLXANCJ0Y2NoK0du6pwKuLSbYyG7YaFgJgQVtvjd7oKxHCE11YBK9DlyS2t2Fslh+oDT02oSJNGpJuCsQaAAQABAwY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWze+il9VuGydqFkeFhh/iremTB8Ngd13K3Xt+TOOJY8/QQG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqUTB7DdvVxpi/fsG318JDpL57X6sICK5kJnx/HugOWK7AgIEAwYFAAoMAQAAAAAAAAAAAgQEBwEACgwBAAAAAAAAAAACWt1BI7yRb9qO/G87o+tplZPL5F1W7UbkIFKWOJjtmUECAQIBAKMGCIabnF0TqEjGtz+67okLc/n3dwUqej+EGtkfc+eaAQIBAA==","base64"],"version":0},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTransaction(
						context.Background(),
						"4fSTSDTTuYa1XXAFxFenuY3SoZWUwCzpMq7kUiya1zW6uqqh6C76GFqTQ3wvegEbZhbPJyr33iDAbieQVWCtVXmf",
					)
				},
				ExpectedValue: &Transaction{
					Slot:      193487858,
					BlockTime: pointer.Get[int64](1675511254),
					Meta: &TransactionMeta{
						Err: nil,
						Fee: 5000,
						PreBalances: []int64{
							112595193235,
							2039280,
							934087680,
							2039280,
							2039280,
							2039280,
							1461600,
							1461600,
						},
						PostBalances: []int64{
							112595188235,
							2039280,
							934087680,
							2039280,
							2039280,
							2039280,
							1461600,
							1461600,
						},
						LogMessages: []string{
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]",
							"Program log: Instruction: TransferChecked",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 400000 compute units",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]",
							"Program log: Instruction: TransferChecked",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 393828 compute units",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success",
						},
						LoadedAddresses: rpc.TransactionLoadedAddresses{
							Readonly: []string{
								"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
							},
							Writable: []string{
								"3Yvq7e9UXLoFK4PKyxrpEA3y3TKmFK2Wb1f5tVFUgwPu",
								"5McxjaxNKYLHtv9DqbMfoi6GNs7ZEMHGkJDrouPib4sW",
								"GAXzq8BWdAWaS1kWFiL5tzV2h3AbRBtYGP5psNTWrM9g",
							},
						},
						PreTokenBalances: []rpc.TransactionMetaTokenBalance{
							{
								AccountIndex: 1,
								Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "100",
									Decimals:       0,
									UIAmountString: "100",
								},
							},
							{
								AccountIndex: 3,
								Mint:         "F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "100",
									Decimals:       0,
									UIAmountString: "100",
								},
							},
							{
								AccountIndex: 4,
								Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "100",
									Decimals:       0,
									UIAmountString: "100",
								},
							},
							{
								AccountIndex: 5,
								Mint:         "F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "100",
									Decimals:       0,
									UIAmountString: "100",
								},
							},
						},
						PostTokenBalances: []rpc.TransactionMetaTokenBalance{
							{
								AccountIndex: 1,
								Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "101",
									Decimals:       0,
									UIAmountString: "101",
								},
							},
							{
								AccountIndex: 3,
								Mint:         "F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "99",
									Decimals:       0,
									UIAmountString: "99",
								},
							},
							{
								AccountIndex: 4,
								Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "99",
									Decimals:       0,
									UIAmountString: "99",
								},
							},
							{
								AccountIndex: 5,
								Mint:         "F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "101",
									Decimals:       0,
									UIAmountString: "101",
								},
							},
						},
						InnerInstructions:    []InnerInstruction{},
						ComputeUnitsConsumed: pointer.Get[uint64](12344),
					},
					Transaction: types.Transaction{
						Signatures: []types.Signature{
							mustBase58Decode(t, "4fSTSDTTuYa1XXAFxFenuY3SoZWUwCzpMq7kUiya1zW6uqqh6C76GFqTQ3wvegEbZhbPJyr33iDAbieQVWCtVXmf"),
						},
						Message: types.Message{
							Version: types.MessageVersionV0,
							Header: types.MessageHeader{
								NumRequireSignatures:        1,
								NumReadonlySignedAccounts:   0,
								NumReadonlyUnsignedAccounts: 1,
							},
							Accounts: []common.PublicKey{
								common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
								common.PublicKeyFromString("HqXcr9ja8jTZAfWN4YSSL8PPWFN3BFJsoxrCvSLaqww1"),
								common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
							},
							RecentBlockHash: "5dQEKfLJt77vfrw2UxWrPrDFwFmxRui6Rk6FBjGnuZBg",
							Instructions: []types.CompiledInstruction{
								{
									ProgramIDIndex: 2,
									Accounts:       []int{3, 6, 5, 0},
									Data:           []byte{12, 1, 0, 0, 0, 0, 0, 0, 0, 0},
								},
								{
									ProgramIDIndex: 2,
									Accounts:       []int{4, 7, 1, 0},
									Data:           []byte{12, 1, 0, 0, 0, 0, 0, 0, 0, 0},
								},
							},
							AddressLookupTables: []types.CompiledAddressLookupTable{
								{
									AccountKey:      common.PublicKeyFromString("77hNYFDx74WFBD1jfM1gHFYk3naH8CxLzLG4KRJAHcRv"),
									ReadonlyIndexes: []uint8{0},
									WritableIndexes: []uint8{1, 2},
								},
								{
									AccountKey:      common.PublicKeyFromString("ByNnrePVpmJTXGiU3Nm9UxTN36tsbaahQcvUNFWmX2Do"),
									ReadonlyIndexes: []uint8{0},
									WritableIndexes: []uint8{2},
								},
							},
						},
					},
					AccountKeys: []common.PublicKey{
						common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
						common.PublicKeyFromString("HqXcr9ja8jTZAfWN4YSSL8PPWFN3BFJsoxrCvSLaqww1"),
						common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						common.PublicKeyFromString("3Yvq7e9UXLoFK4PKyxrpEA3y3TKmFK2Wb1f5tVFUgwPu"),
						common.PublicKeyFromString("5McxjaxNKYLHtv9DqbMfoi6GNs7ZEMHGkJDrouPib4sW"),
						common.PublicKeyFromString("GAXzq8BWdAWaS1kWFiL5tzV2h3AbRBtYGP5psNTWrM9g"),
						common.PublicKeyFromString("F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8"),
						common.PublicKeyFromString("5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw"),
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func mustBase58Decode(t *testing.T, s string) []byte {
	b, err := base58.Decode(s)
	if err != nil {
		t.Fatalf("failed to base58 decode %v", s)
	}
	return b
}
