package rpc

import (
	"context"
	"testing"

	"github.com/portto/solana-go-sdk/pkg/pointer"
)

func TestGetBlock(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}],"transactions":[{"meta":{"err":null,"fee":10000,"innerInstructions":[],"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499999835001,1000000000000000,143487360,1169280,1],"postTokenBalances":[],"preBalances":[499999845001,1000000000000000,143487360,1169280,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"transaction":{"message":{"accountKeys":["9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","3UbyTNpi3omt7hfEqQRB5844VANQFWiq8uEDNCrSwAVG","SysvarS1otHashes111111111111111111111111111","SysvarC1ock11111111111111111111111111111111","Vote111111111111111111111111111111111111111"],"header":{"numReadonlySignedAccounts":0,"numReadonlyUnsignedAccounts":3,"numRequiredSignatures":2},"instructions":[{"accounts":[1,2,3,1],"data":"2ZjTR1vUs2pHXyTLuZA9zjpNqav47YU1uqenSEcYn6xkrdmMkUJK8JDHd5TcEU7K5R9pbB2UxbY95zDzHio","programIdIndex":4}],"recentBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q"},"signatures":["3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A","4cWqSVUcxTujZ6eHtNWESwCrBUfidbZ1J124VU2jY9TQpXxyHSDku1NiZhw95SzXe1mGihiP9AdQNEkLMAdvBYPQ"]}}]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlock(
					context.TODO(),
					33,
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlock]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlock{
					ParentSLot:        32,
					BlockHeight:       pointer.Get[int64](33),
					BlockTime:         pointer.Get[int64](1631803928),
					PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
					Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
					Rewards: []GetBlockReward{
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
								PreTokenBalances:  []TransactionMetaTokenBalance{},
								PostTokenBalances: []TransactionMetaTokenBalance{},
								LogMessages: []string{
									"Program Vote111111111111111111111111111111111111111 invoke [1]",
									"Program Vote111111111111111111111111111111111111111 success",
								},
								InnerInstructions: []TransactionMetaInnerInstruction{},
							},
							Transaction: map[string]any{
								"signatures": []any{
									"3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A",
									"4cWqSVUcxTujZ6eHtNWESwCrBUfidbZ1J124VU2jY9TQpXxyHSDku1NiZhw95SzXe1mGihiP9AdQNEkLMAdvBYPQ",
								},
								"message": map[string]any{
									"accountKeys": []any{
										"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX",
										"3UbyTNpi3omt7hfEqQRB5844VANQFWiq8uEDNCrSwAVG",
										"SysvarS1otHashes111111111111111111111111111",
										"SysvarC1ock11111111111111111111111111111111",
										"Vote111111111111111111111111111111111111111",
									},
									"header": map[string]any{
										"numReadonlySignedAccounts":   0.,
										"numReadonlyUnsignedAccounts": 3.,
										"numRequiredSignatures":       2.,
									},
									"instructions": []any{
										map[string]any{
											"accounts":       []any{1., 2., 3., 1.},
											"data":           "2ZjTR1vUs2pHXyTLuZA9zjpNqav47YU1uqenSEcYn6xkrdmMkUJK8JDHd5TcEU7K5R9pbB2UxbY95zDzHio",
											"programIdIndex": 4.,
										},
									},
									"recentBlockhash": "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
								},
							},
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64", "rewards": false}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","transactions":[{"meta":{"err":null,"fee":10000,"innerInstructions":[],"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499999835001,1000000000000000,143487360,1169280,1],"postTokenBalances":[],"preBalances":[499999845001,1000000000000000,143487360,1169280,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"transaction":["AnXU8JYCIrc73JwxK9traTSp3EZdmnJp0B5luW8CCzr7GnFd/SjIMXiG4qbN5CwyEVhbpORzBUpB/253cNtS1A+0rWE+nrDqWRQ2OVU727PU4NtR611jY+10Q+F6lCZDsJt46b6oXz3PN5WGxTQk7mC4YhCbYsTcalWBkltA8KgPAgADBXszyT4GLb26BFuAAUXtW0B75zurDhXE7UOYKHFkpIlKJMmZpq+FRXTx8jzBMy1YsdkCo0kyLDdF2Q3NhXRdEosGp9UXGS8Kr8byZeP7d8x62oLFKdC+OxNuLQBVIAAAAAan1RcYx3TJKFZjmGkdXraLXrijm0ttXHNVWyEAAAAAB2FIHTV0dLt8TXYk69O9s9g1XnPREEP8DaNTgAAAAACrUBylgzc0SSCUPSfMJC3TI6KJEzs834KdMIMJci+UYAEEBAECAwE9AgAAAAEAAAAAAAAAIAAAAAAAAAAGCHSVIc5Betdf+NkRi4YR2D3abNLvpbI83qnB7EvNsAEZWkNhAAAAAA==","base64"]}]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockWithConfig(
					context.TODO(),
					33,
					GetBlockConfig{
						Encoding: GetBlockConfigEncodingBase64,
						Rewards:  pointer.Get[bool](false),
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlock]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlock{
					ParentSLot:        32,
					BlockHeight:       pointer.Get[int64](33),
					BlockTime:         pointer.Get[int64](1631803928),
					PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
					Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
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
								PreTokenBalances:  []TransactionMetaTokenBalance{},
								PostTokenBalances: []TransactionMetaTokenBalance{},
								LogMessages: []string{
									"Program Vote111111111111111111111111111111111111111 invoke [1]",
									"Program Vote111111111111111111111111111111111111111 success",
								},
								InnerInstructions: []TransactionMetaInnerInstruction{},
							},
							Transaction: []any{
								"AnXU8JYCIrc73JwxK9traTSp3EZdmnJp0B5luW8CCzr7GnFd/SjIMXiG4qbN5CwyEVhbpORzBUpB/253cNtS1A+0rWE+nrDqWRQ2OVU727PU4NtR611jY+10Q+F6lCZDsJt46b6oXz3PN5WGxTQk7mC4YhCbYsTcalWBkltA8KgPAgADBXszyT4GLb26BFuAAUXtW0B75zurDhXE7UOYKHFkpIlKJMmZpq+FRXTx8jzBMy1YsdkCo0kyLDdF2Q3NhXRdEosGp9UXGS8Kr8byZeP7d8x62oLFKdC+OxNuLQBVIAAAAAan1RcYx3TJKFZjmGkdXraLXrijm0ttXHNVWyEAAAAAB2FIHTV0dLt8TXYk69O9s9g1XnPREEP8DaNTgAAAAACrUBylgzc0SSCUPSfMJC3TI6KJEzs834KdMIMJci+UYAEEBAECAwE9AgAAAAEAAAAAAAAAIAAAAAAAAAAGCHSVIc5Betdf+NkRi4YR2D3abNLvpbI83qnB7EvNsAEZWkNhAAAAAA==",
								"base64",
							},
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}],"transactions":[{"meta":{"err":null,"fee":10000,"innerInstructions":[],"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499999835001,1000000000000000,143487360,1169280,1],"postTokenBalances":[],"preBalances":[499999845001,1000000000000000,143487360,1169280,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"transaction":["AnXU8JYCIrc73JwxK9traTSp3EZdmnJp0B5luW8CCzr7GnFd/SjIMXiG4qbN5CwyEVhbpORzBUpB/253cNtS1A+0rWE+nrDqWRQ2OVU727PU4NtR611jY+10Q+F6lCZDsJt46b6oXz3PN5WGxTQk7mC4YhCbYsTcalWBkltA8KgPAgADBXszyT4GLb26BFuAAUXtW0B75zurDhXE7UOYKHFkpIlKJMmZpq+FRXTx8jzBMy1YsdkCo0kyLDdF2Q3NhXRdEosGp9UXGS8Kr8byZeP7d8x62oLFKdC+OxNuLQBVIAAAAAan1RcYx3TJKFZjmGkdXraLXrijm0ttXHNVWyEAAAAAB2FIHTV0dLt8TXYk69O9s9g1XnPREEP8DaNTgAAAAACrUBylgzc0SSCUPSfMJC3TI6KJEzs834KdMIMJci+UYAEEBAECAwE9AgAAAAEAAAAAAAAAIAAAAAAAAAAGCHSVIc5Betdf+NkRi4YR2D3abNLvpbI83qnB7EvNsAEZWkNhAAAAAA==","base64"]}]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockWithConfig(
					context.TODO(),
					33,
					GetBlockConfig{
						Encoding: GetBlockConfigEncodingBase64,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlock]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlock{
					ParentSLot:        32,
					BlockHeight:       pointer.Get[int64](33),
					BlockTime:         pointer.Get[int64](1631803928),
					PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
					Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
					Rewards: []GetBlockReward{
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
								PreTokenBalances:  []TransactionMetaTokenBalance{},
								PostTokenBalances: []TransactionMetaTokenBalance{},
								LogMessages: []string{
									"Program Vote111111111111111111111111111111111111111 invoke [1]",
									"Program Vote111111111111111111111111111111111111111 success",
								},
								InnerInstructions: []TransactionMetaInnerInstruction{},
							},
							Transaction: []any{
								"AnXU8JYCIrc73JwxK9traTSp3EZdmnJp0B5luW8CCzr7GnFd/SjIMXiG4qbN5CwyEVhbpORzBUpB/253cNtS1A+0rWE+nrDqWRQ2OVU727PU4NtR611jY+10Q+F6lCZDsJt46b6oXz3PN5WGxTQk7mC4YhCbYsTcalWBkltA8KgPAgADBXszyT4GLb26BFuAAUXtW0B75zurDhXE7UOYKHFkpIlKJMmZpq+FRXTx8jzBMy1YsdkCo0kyLDdF2Q3NhXRdEosGp9UXGS8Kr8byZeP7d8x62oLFKdC+OxNuLQBVIAAAAAan1RcYx3TJKFZjmGkdXraLXrijm0ttXHNVWyEAAAAAB2FIHTV0dLt8TXYk69O9s9g1XnPREEP8DaNTgAAAAACrUBylgzc0SSCUPSfMJC3TI6KJEzs834KdMIMJci+UYAEEBAECAwE9AgAAAAEAAAAAAAAAIAAAAAAAAAAGCHSVIc5Betdf+NkRi4YR2D3abNLvpbI83qnB7EvNsAEZWkNhAAAAAA==",
								"base64",
							},
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64", "transactionDetails":"signatures"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}],"signatures":["3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A"]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockWithConfig(
					context.TODO(),
					33,
					GetBlockConfig{
						Encoding:           GetBlockConfigEncodingBase64,
						TransactionDetails: GetBlockConfigTransactionDetailsSignatures,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlock]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlock{
					ParentSLot:        32,
					BlockHeight:       pointer.Get[int64](33),
					BlockTime:         pointer.Get[int64](1631803928),
					PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
					Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
					Rewards: []GetBlockReward{
						{
							Pubkey:       "9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX",
							Lamports:     5000,
							PostBalances: 499999840001,
							RewardType:   "Fee",
							Commission:   nil,
						},
					},
					Signatures: []string{
						"3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64", "transactionDetails":"none"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockWithConfig(
					context.TODO(),
					33,
					GetBlockConfig{
						Encoding:           GetBlockConfigEncodingBase64,
						TransactionDetails: GetBlockConfigTransactionDetailsNone,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlock]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlock{
					ParentSLot:        32,
					BlockHeight:       pointer.Get[int64](33),
					BlockTime:         pointer.Get[int64](1631803928),
					PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
					Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
					Rewards: []GetBlockReward{
						{
							Pubkey:       "9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX",
							Lamports:     5000,
							PostBalances: 499999840001,
							RewardType:   "Fee",
							Commission:   nil,
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64", "transactionDetails":"none", "commitment": "confirmed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockWithConfig(
					context.TODO(),
					33,
					GetBlockConfig{
						Encoding:           GetBlockConfigEncodingBase64,
						TransactionDetails: GetBlockConfigTransactionDetailsNone,
						Commitment:         CommitmentConfirmed,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlock]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlock{
					ParentSLot:        32,
					BlockHeight:       pointer.Get[int64](33),
					BlockTime:         pointer.Get[int64](1631803928),
					PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
					Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
					Rewards: []GetBlockReward{
						{
							Pubkey:       "9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX",
							Lamports:     5000,
							PostBalances: 499999840001,
							RewardType:   "Fee",
							Commission:   nil,
						},
					},
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
