package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/pkg/pointer"
)

func TestGetTransaction(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1631380624,"meta":{"err":null,"fee":5000,"innerInstructions":[{"index":0,"instructions":[{"accounts":[0,1],"data":"3Bxs4h24hBtQy9rw","programIdIndex":3},{"accounts":[1],"data":"9krTDU2LzCSUJuVZ","programIdIndex":3},{"accounts":[1],"data":"SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg","programIdIndex":3},{"accounts":[1,2,0,5],"data":"2","programIdIndex":4}]}],"logMessages":["Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]","Program log: Transfer 2039280 lamports to the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Allocate space for the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Assign the associated token account to the SPL Token program","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Initialize the associated token account","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]","Program log: Instruction: InitializeAccount","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3412 of 177045 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 27016 of 200000 compute units","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success"],"postBalances":[38024615601,2039280,1461600,1,1089991680,1,898174080],"postTokenBalances":[{"accountIndex":1,"mint":"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","uiTokenAmount":{"amount":"0","decimals":9,"uiAmount":null,"uiAmountString":"0"}}],"preBalances":[38026659881,0,1461600,1,1089991680,1,898174080],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":80218681,"transaction":{"message":{"accountKeys":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ","AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ","4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","11111111111111111111111111111111","TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","SysvarRent111111111111111111111111111111111","ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"],"header":{"numReadonlySignedAccounts":0,"numReadonlyUnsignedAccounts":5,"numRequiredSignatures":1},"instructions":[{"accounts":[0,1,0,2,3,4,5],"data":"","programIdIndex":6}],"recentBlockhash":"Gpemb2whtMogoSGVe5KMjuoueeqNNkQ1kKnw7fsYKZHj"},"signatures":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF"]}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetTransaction(
					context.TODO(),
					"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
				)
			},
			ExpectedResponse: GetTransactionResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetTransactionResult{
					Slot:      80218681,
					BlockTime: pointer.Int64(1631380624),
					Meta: &TransactionMeta{
						Fee: 5000,
						InnerInstructions: []TransactionMetaInnerInstruction{
							{
								Index: 0,
								Instructions: []Instruction{
									{
										ProgramIDIndex: 3,
										Data:           "3Bxs4h24hBtQy9rw",
										Accounts:       []int{0, 1},
									},
									{
										ProgramIDIndex: 3,
										Data:           "9krTDU2LzCSUJuVZ",
										Accounts:       []int{1},
									},
									{
										ProgramIDIndex: 3,
										Data:           "SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg",
										Accounts:       []int{1},
									},
									{
										ProgramIDIndex: 4,
										Data:           "2",
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
						PreTokenBalances: []TransactionMetaTokenBalance{},
						PostTokenBalances: []TransactionMetaTokenBalance{
							{
								AccountIndex: 1,
								Mint:         "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
								UITokenAmount: GetTokenAccountBalanceResultValue{
									Amount:         "0",
									Decimals:       9,
									UIAmountString: "0",
								},
							},
						},
					},
					Transaction: map[string]interface{}{
						"signatures": []interface{}{
							"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
						},
						"message": map[string]interface{}{
							"header": map[string]interface{}{
								"numReadonlySignedAccounts":   0.,
								"numReadonlyUnsignedAccounts": 5.,
								"numRequiredSignatures":       1.,
							},
							"accountKeys": []interface{}{
								"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
								"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
								"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
								"11111111111111111111111111111111",
								"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								"SysvarRent111111111111111111111111111111111",
								"ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL",
							},
							"instructions": []interface{}{
								map[string]interface{}{
									"accounts":       []interface{}{0., 1., 0., 2., 3., 4., 5.},
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
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetTransactionWithConfig(
					context.TODO(),
					"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
					GetTransactionConfig{
						Encoding: GetTransactionConfigEncodingBase64,
					},
				)
			},
			ExpectedResponse: GetTransactionResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetTransactionResult{
					Slot:      80218681,
					BlockTime: pointer.Int64(1631380624),
					Meta: &TransactionMeta{
						Fee: 5000,
						InnerInstructions: []TransactionMetaInnerInstruction{
							{
								Index: 0,
								Instructions: []Instruction{
									{
										ProgramIDIndex: 3,
										Data:           "3Bxs4h24hBtQy9rw",
										Accounts:       []int{0, 1},
									},
									{
										ProgramIDIndex: 3,
										Data:           "9krTDU2LzCSUJuVZ",
										Accounts:       []int{1},
									},
									{
										ProgramIDIndex: 3,
										Data:           "SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg",
										Accounts:       []int{1},
									},
									{
										ProgramIDIndex: 4,
										Data:           "2",
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
						PreTokenBalances: []TransactionMetaTokenBalance{},
						PostTokenBalances: []TransactionMetaTokenBalance{
							{
								AccountIndex: 1,
								Mint:         "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
								UITokenAmount: GetTokenAccountBalanceResultValue{
									Amount:         "0",
									Decimals:       9,
									UIAmountString: "0",
								},
							},
						},
					},
					Transaction: []interface{}{"AaEGlsrjwHOjXODEvEGb5Zade8QelkWx2l9VvseP/g1olewFxKkJEwRDJyZ2wel8p2Dilp3wnBu6AEbRB4LthwABAAUHEJZZF158ZDMhpe1GQqAnsKvZe43ZetG8xtxkcThszdyUJGGIseU8n4crN7gTTkkjZvTPQVkY2NPZnO+5BTpTqzO9mOFbcsDwmqTwyIZje2Ppd9PY6hWpndBzwVYYhseQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqQan1RcZLFxRIYzJTD1K8X9Y2u4Im6H9ROPb2YoAAAAAjJclj04kifG7PRApFI4NgwtaE5na/xCEBI572Nvp+FnrFE6iq1ZbCKVJ+UiBaEkoE9dTFWqba+nWyTsH21qhygEGBwABAAIDBAUA", "base64"},
				},
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			testRpcCall(t, tt)
		})
	}
}
