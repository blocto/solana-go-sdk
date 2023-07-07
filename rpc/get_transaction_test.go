package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
)

func TestGetTransaction(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1631380624,"meta":{"err":null,"fee":5000,"innerInstructions":[{"index":0,"instructions":[{"accounts":[0,1],"data":"3Bxs4h24hBtQy9rw","programIdIndex":3},{"accounts":[1],"data":"9krTDU2LzCSUJuVZ","programIdIndex":3},{"accounts":[1],"data":"SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg","programIdIndex":3},{"accounts":[1,2,0,5],"data":"2","programIdIndex":4}]}],"logMessages":["Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]","Program log: Transfer 2039280 lamports to the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Allocate space for the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Assign the associated token account to the SPL Token program","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Initialize the associated token account","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]","Program log: Instruction: InitializeAccount","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3412 of 177045 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 27016 of 200000 compute units","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success"],"postBalances":[38024615601,2039280,1461600,1,1089991680,1,898174080],"postTokenBalances":[{"accountIndex":1,"mint":"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","uiTokenAmount":{"amount":"0","decimals":9,"uiAmount":null,"uiAmountString":"0"}}],"preBalances":[38026659881,0,1461600,1,1089991680,1,898174080],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":80218681,"transaction":{"message":{"accountKeys":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ","AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ","4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","11111111111111111111111111111111","TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","SysvarRent111111111111111111111111111111111","ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"],"header":{"numReadonlySignedAccounts":0,"numReadonlyUnsignedAccounts":5,"numRequiredSignatures":1},"instructions":[{"accounts":[0,1,0,2,3,4,5],"data":"","programIdIndex":6}],"recentBlockhash":"Gpemb2whtMogoSGVe5KMjuoueeqNNkQ1kKnw7fsYKZHj"},"signatures":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF"]}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransaction(
						context.TODO(),
						"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
					)
				},
				ExpectedValue: JsonRpcResponse[*GetTransaction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: &GetTransaction{
						Slot:      80218681,
						BlockTime: pointer.Get[int64](1631380624),
						Meta: &TransactionMeta{
							Fee: 5000,
							InnerInstructions: []TransactionMetaInnerInstruction{
								{
									Index: 0,
									Instructions: []any{
										map[string]any{
											"programIdIndex": 3.,
											"data":           "3Bxs4h24hBtQy9rw",
											"accounts":       []any{0., 1.},
										},
										map[string]any{
											"programIdIndex": 3.,
											"data":           "9krTDU2LzCSUJuVZ",
											"accounts":       []any{1.},
										},
										map[string]any{
											"programIdIndex": 3.,
											"data":           "SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg",
											"accounts":       []any{1.},
										},
										map[string]any{
											"programIdIndex": 4.,
											"data":           "2",
											"accounts":       []any{1., 2., 0., 5.},
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
							Rewards: []Reward{},
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
							PreTokenBalances: []TransactionMetaTokenBalance{},
							PostTokenBalances: []TransactionMetaTokenBalance{
								{
									AccountIndex: 1,
									Mint:         "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
									UITokenAmount: TokenAccountBalance{
										Amount:         "0",
										Decimals:       9,
										UIAmountString: "0",
									},
								},
							},
						},
						Transaction: map[string]any{
							"signatures": []any{
								"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
							},
							"message": map[string]any{
								"header": map[string]any{
									"numReadonlySignedAccounts":   0.,
									"numReadonlyUnsignedAccounts": 5.,
									"numRequiredSignatures":       1.,
								},
								"accountKeys": []any{
									"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
									"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
									"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
									"11111111111111111111111111111111",
									"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									"SysvarRent111111111111111111111111111111111",
									"ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL",
								},
								"instructions": []any{
									map[string]any{
										"accounts":       []any{0., 1., 0., 2., 3., 4., 5.},
										"data":           "",
										"programIdIndex": 6.,
									},
								},
								"recentBlockhash": "Gpemb2whtMogoSGVe5KMjuoueeqNNkQ1kKnw7fsYKZHj",
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF", {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1631380624,"meta":{"err":null,"fee":5000,"innerInstructions":[{"index":0,"instructions":[{"accounts":[0,1],"data":"3Bxs4h24hBtQy9rw","programIdIndex":3},{"accounts":[1],"data":"9krTDU2LzCSUJuVZ","programIdIndex":3},{"accounts":[1],"data":"SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg","programIdIndex":3},{"accounts":[1,2,0,5],"data":"2","programIdIndex":4}]}],"logMessages":["Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]","Program log: Transfer 2039280 lamports to the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Allocate space for the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Assign the associated token account to the SPL Token program","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Initialize the associated token account","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]","Program log: Instruction: InitializeAccount","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3412 of 177045 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 27016 of 200000 compute units","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success"],"postBalances":[38024615601,2039280,1461600,1,1089991680,1,898174080],"postTokenBalances":[{"accountIndex":1,"mint":"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","uiTokenAmount":{"amount":"0","decimals":9,"uiAmount":null,"uiAmountString":"0"}}],"preBalances":[38026659881,0,1461600,1,1089991680,1,898174080],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":80218681,"transaction":["AaEGlsrjwHOjXODEvEGb5Zade8QelkWx2l9VvseP/g1olewFxKkJEwRDJyZ2wel8p2Dilp3wnBu6AEbRB4LthwABAAUHEJZZF158ZDMhpe1GQqAnsKvZe43ZetG8xtxkcThszdyUJGGIseU8n4crN7gTTkkjZvTPQVkY2NPZnO+5BTpTqzO9mOFbcsDwmqTwyIZje2Ppd9PY6hWpndBzwVYYhseQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqQan1RcZLFxRIYzJTD1K8X9Y2u4Im6H9ROPb2YoAAAAAjJclj04kifG7PRApFI4NgwtaE5na/xCEBI572Nvp+FnrFE6iq1ZbCKVJ+UiBaEkoE9dTFWqba+nWyTsH21qhygEGBwABAAIDBAUA","base64"]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransactionWithConfig(
						context.TODO(),
						"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
						GetTransactionConfig{
							Encoding: TransactionEncodingBase64,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[*GetTransaction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: &GetTransaction{
						Slot:      80218681,
						BlockTime: pointer.Get[int64](1631380624),
						Meta: &TransactionMeta{
							Fee: 5000,
							InnerInstructions: []TransactionMetaInnerInstruction{
								{
									Index: 0,
									Instructions: []any{
										map[string]any{
											"programIdIndex": 3.,
											"data":           "3Bxs4h24hBtQy9rw",
											"accounts":       []any{0., 1.},
										},
										map[string]any{
											"programIdIndex": 3.,
											"data":           "9krTDU2LzCSUJuVZ",
											"accounts":       []any{1.},
										},
										map[string]any{
											"programIdIndex": 3.,
											"data":           "SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg",
											"accounts":       []any{1.},
										},
										map[string]any{
											"programIdIndex": 4.,
											"data":           "2",
											"accounts":       []any{1., 2., 0., 5.},
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
							Rewards: []Reward{},
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
							PreTokenBalances: []TransactionMetaTokenBalance{},
							PostTokenBalances: []TransactionMetaTokenBalance{
								{
									AccountIndex: 1,
									Mint:         "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
									UITokenAmount: TokenAccountBalance{
										Amount:         "0",
										Decimals:       9,
										UIAmountString: "0",
									},
								},
							},
						},
						Transaction: []any{"AaEGlsrjwHOjXODEvEGb5Zade8QelkWx2l9VvseP/g1olewFxKkJEwRDJyZ2wel8p2Dilp3wnBu6AEbRB4LthwABAAUHEJZZF158ZDMhpe1GQqAnsKvZe43ZetG8xtxkcThszdyUJGGIseU8n4crN7gTTkkjZvTPQVkY2NPZnO+5BTpTqzO9mOFbcsDwmqTwyIZje2Ppd9PY6hWpndBzwVYYhseQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqQan1RcZLFxRIYzJTD1K8X9Y2u4Im6H9ROPb2YoAAAAAjJclj04kifG7PRApFI4NgwtaE5na/xCEBI572Nvp+FnrFE6iq1ZbCKVJ+UiBaEkoE9dTFWqba+nWyTsH21qhygEGBwABAAIDBAUA", "base64"},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF", {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":null,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransactionWithConfig(
						context.TODO(),
						"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
						GetTransactionConfig{
							Encoding: TransactionEncodingBase64,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[*GetTransaction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  nil,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"getTransaction","params":["2vFRWHPwbxFkZMtgcyRPx8Hx1v4TTyvmhYZgYejfwWXQehiaNijP18FEQMpMFZHyJFmTXrWxncZsS6yKYESDBucQ",{"encoding":"base64", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1666897943,"meta":{"computeUnitsConsumed":185,"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units","Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success"],"postBalances":[34358789287,1141440],"postTokenBalances":[],"preBalances":[34358794287,1141440],"preTokenBalances":[],"returnData":{"data":["AQIDBAU=","base64"],"programId":"35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP"},"rewards":[],"status":{"Ok":null}},"slot":159780566,"transaction":["AV/vxqOrdrGio45xsX7l9jdCcQDy3VuY/wHlleHEuBuwDUIYO2ce/YpjkRfZHCq7tYSNEwFCRdolqNg2oibR5wUBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYGK9azWH/heD6vSj5deOv9pmPQoZfCmIFJqrW8ixgJLtAQEAAA==","base64"],"version":"legacy"},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransactionWithConfig(
						context.TODO(),
						"2vFRWHPwbxFkZMtgcyRPx8Hx1v4TTyvmhYZgYejfwWXQehiaNijP18FEQMpMFZHyJFmTXrWxncZsS6yKYESDBucQ",
						GetTransactionConfig{
							Encoding:                       TransactionEncodingBase64,
							MaxSupportedTransactionVersion: pointer.Get[uint8](0),
						},
					)
				},
				ExpectedValue: JsonRpcResponse[*GetTransaction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: &GetTransaction{
						Slot:      159780566,
						BlockTime: pointer.Get[int64](1666897943),
						Version:   "legacy",
						Meta: &TransactionMeta{
							Err: nil,
							Fee: 5000,
							PreBalances: []int64{
								34358794287,
								1141440,
							},
							PostBalances: []int64{
								34358789287,
								1141440,
							},
							LogMessages: []string{
								"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]",
								"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units",
								"Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=",
								"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success",
							},
							Rewards: []Reward{},
							LoadedAddresses: TransactionLoadedAddresses{
								Writable: []string{},
								Readonly: []string{},
							},
							PreTokenBalances:  []TransactionMetaTokenBalance{},
							PostTokenBalances: []TransactionMetaTokenBalance{},
							InnerInstructions: []TransactionMetaInnerInstruction{},
							ReturnData: &ReturnData{
								ProgramId: "35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP",
								Data:      []any{"AQIDBAU=", "base64"},
							},
							ComputeUnitsConsumed: pointer.Get[uint64](185),
						},
						Transaction: []any{
							"AV/vxqOrdrGio45xsX7l9jdCcQDy3VuY/wHlleHEuBuwDUIYO2ce/YpjkRfZHCq7tYSNEwFCRdolqNg2oibR5wUBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYGK9azWH/heD6vSj5deOv9pmPQoZfCmIFJqrW8ixgJLtAQEAAA==",
							"base64",
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params": ["4fSTSDTTuYa1XXAFxFenuY3SoZWUwCzpMq7kUiya1zW6uqqh6C76GFqTQ3wvegEbZhbPJyr33iDAbieQVWCtVXmf", {"encoding": "base64", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1675511254,"meta":{"computeUnitsConsumed":12344,"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":["F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw"],"writable":["3Yvq7e9UXLoFK4PKyxrpEA3y3TKmFK2Wb1f5tVFUgwPu","5McxjaxNKYLHtv9DqbMfoi6GNs7ZEMHGkJDrouPib4sW","GAXzq8BWdAWaS1kWFiL5tzV2h3AbRBtYGP5psNTWrM9g"]},"logMessages":["Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 400000 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 393828 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success"],"postBalances":[112595188235,2039280,934087680,2039280,2039280,2039280,1461600,1461600],"postTokenBalances":[{"accountIndex":1,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"101","decimals":0,"uiAmount":101.0,"uiAmountString":"101"}},{"accountIndex":3,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"99","decimals":0,"uiAmount":99.0,"uiAmountString":"99"}},{"accountIndex":4,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"99","decimals":0,"uiAmount":99.0,"uiAmountString":"99"}},{"accountIndex":5,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"101","decimals":0,"uiAmount":101.0,"uiAmountString":"101"}}],"preBalances":[112595193235,2039280,934087680,2039280,2039280,2039280,1461600,1461600],"preTokenBalances":[{"accountIndex":1,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}},{"accountIndex":3,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}},{"accountIndex":4,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}},{"accountIndex":5,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}}],"rewards":[],"status":{"Ok":null}},"slot":193487858,"transaction":["AbczATLXANCJ0Y2NoK0du6pwKuLSbYyG7YaFgJgQVtvjd7oKxHCE11YBK9DlyS2t2Fslh+oDT02oSJNGpJuCsQaAAQABAwY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWze+il9VuGydqFkeFhh/iremTB8Ngd13K3Xt+TOOJY8/QQG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqUTB7DdvVxpi/fsG318JDpL57X6sICK5kJnx/HugOWK7AgIEAwYFAAoMAQAAAAAAAAAAAgQEBwEACgwBAAAAAAAAAAACWt1BI7yRb9qO/G87o+tplZPL5F1W7UbkIFKWOJjtmUECAQIBAKMGCIabnF0TqEjGtz+67okLc/n3dwUqej+EGtkfc+eaAQIBAA==","base64"],"version":0},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransactionWithConfig(
						context.TODO(),
						"4fSTSDTTuYa1XXAFxFenuY3SoZWUwCzpMq7kUiya1zW6uqqh6C76GFqTQ3wvegEbZhbPJyr33iDAbieQVWCtVXmf",
						GetTransactionConfig{
							Encoding:                       TransactionEncodingBase64,
							MaxSupportedTransactionVersion: pointer.Get[uint8](0),
						},
					)
				},
				ExpectedValue: JsonRpcResponse[*GetTransaction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: &GetTransaction{
						Slot:      193487858,
						BlockTime: pointer.Get[int64](1675511254),
						Version:   0.,
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
							Rewards: []Reward{},
							LoadedAddresses: TransactionLoadedAddresses{
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
							PreTokenBalances: []TransactionMetaTokenBalance{
								{
									AccountIndex: 1,
									Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
									Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
									ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									UITokenAmount: TokenAccountBalance{
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
									UITokenAmount: TokenAccountBalance{
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
									UITokenAmount: TokenAccountBalance{
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
									UITokenAmount: TokenAccountBalance{
										Amount:         "100",
										Decimals:       0,
										UIAmountString: "100",
									},
								},
							},
							PostTokenBalances: []TransactionMetaTokenBalance{
								{
									AccountIndex: 1,
									Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
									Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
									ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									UITokenAmount: TokenAccountBalance{
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
									UITokenAmount: TokenAccountBalance{
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
									UITokenAmount: TokenAccountBalance{
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
									UITokenAmount: TokenAccountBalance{
										Amount:         "101",
										Decimals:       0,
										UIAmountString: "101",
									},
								},
							},
							InnerInstructions:    []TransactionMetaInnerInstruction{},
							ComputeUnitsConsumed: pointer.Get[uint64](12344),
						},
						Transaction: []any{
							"AbczATLXANCJ0Y2NoK0du6pwKuLSbYyG7YaFgJgQVtvjd7oKxHCE11YBK9DlyS2t2Fslh+oDT02oSJNGpJuCsQaAAQABAwY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWze+il9VuGydqFkeFhh/iremTB8Ngd13K3Xt+TOOJY8/QQG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqUTB7DdvVxpi/fsG318JDpL57X6sICK5kJnx/HugOWK7AgIEAwYFAAoMAQAAAAAAAAAAAgQEBwEACgwBAAAAAAAAAAACWt1BI7yRb9qO/G87o+tplZPL5F1W7UbkIFKWOJjtmUECAQIBAKMGCIabnF0TqEjGtz+67okLc/n3dwUqej+EGtkfc+eaAQIBAA==",
							"base64",
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
