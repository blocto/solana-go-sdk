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
					ParentSlot:        32,
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
					ParentSlot:        32,
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
					ParentSlot:        32,
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
					ParentSlot:        32,
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
					ParentSlot:        32,
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
					ParentSlot:        32,
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
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64", "rewards": false, "maxSupportedTransactionVersion": 0}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","transactions":[{"meta":{"err":null,"fee":10000,"innerInstructions":[],"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499999835001,1000000000000000,143487360,1169280,1],"postTokenBalances":[],"preBalances":[499999845001,1000000000000000,143487360,1169280,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"transaction":["AQiClQkvASAMI63iTE4VCNKpvDttDlM70bXlosqCRJ4kPeiPcPmIwW4AFNFTjmil/X1BSQJV6yUnXdQ+1+KSlAKAAQACBNcUkx66ahmo9NxsAZr/Jk9fv2jFoo7gs7mHVc451knTB0TvGBo/9u4tBgYjoLxl6Y29BFXLmb1J7Q+3GgPKqJ8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJ3pq+XM5t6yI0YkskERvUAAjCSZvYuU8EYJEmCAAAAc4JUQCBBCFja0HaW7x24Mm7k1W45VWHvtEvczqYmigABAwQBAAACDQAAAAC4J9QJAAAAAP4A","base64"], "version": "legacy"}]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockWithConfig(
					context.TODO(),
					33,
					GetBlockConfig{
						Encoding:                       GetBlockConfigEncodingBase64,
						Rewards:                        pointer.Get[bool](false),
						MaxSupportedTransactionVersion: pointer.Get[uint8](0),
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlock]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlock{
					ParentSlot:        32,
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
								"AQiClQkvASAMI63iTE4VCNKpvDttDlM70bXlosqCRJ4kPeiPcPmIwW4AFNFTjmil/X1BSQJV6yUnXdQ+1+KSlAKAAQACBNcUkx66ahmo9NxsAZr/Jk9fv2jFoo7gs7mHVc451knTB0TvGBo/9u4tBgYjoLxl6Y29BFXLmb1J7Q+3GgPKqJ8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJ3pq+XM5t6yI0YkskERvUAAjCSZvYuU8EYJEmCAAAAc4JUQCBBCFja0HaW7x24Mm7k1W45VWHvtEvczqYmigABAwQBAAACDQAAAAC4J9QJAAAAAP4A",
								"base64",
							},
							Version: "legacy",
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[46986, {"encoding": "base64"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32015,"message":"Transaction version (0) is not supported"},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockWithConfig(
					context.TODO(),
					46986,
					GetBlockConfig{
						Encoding: GetBlockConfigEncodingBase64,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlock]{
				JsonRpc: "2.0",
				Id:      1,
				Error: &JsonRpcError{
					Code:    -32015,
					Message: "Transaction version (0) is not supported",
				},
				Result: GetBlock{},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[46986, {"encoding": "base64", "rewards": false, "maxSupportedTransactionVersion": 0}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":46984,"blockTime":1663315273,"blockhash":"5GLYVkvcF7ygCQpFVbeZsaCcSBaw9qofJjdtzwVXYPu1","parentSlot":46985,"previousBlockhash":"2ChS54SSzba9zFqsbQJviVkmXXN7zdpMRXPjvzY93cfX","transactions":[{"meta":{"computeUnitsConsumed":0,"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":["test111111111111111111111111111111111111111"]},"logMessages":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"],"postBalances":[9998999995000,1,1000000000],"postTokenBalances":[],"preBalances":[10000000000000,1,0],"preTokenBalances":[],"rewards":null,"status":{"Ok":null}},"transaction":["Ad8N7teP2bRE2xzLJRc+5zw2OOJWF4cXtfSxbMTwijdxQMkGB4KMAAzG/svcdNGzaA5QiH4tb1ly0e7oFyJ53gCAAQABAn9ga/qYhdDgSftxl4CLVlBlRooyjZnabjgnerV4N1a5AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADzHkIXiJ+eq12wXs5ZseAwuBZu+EOsM5tQogx7iqiY+AEBAgACDAIAAAAAypo7AAAAAAGWXqBzSZja4R/6KEg0G0P5Y1bvQBhRNhSSdA0KHQ84BAECAA==","base64"],"version":0},{"meta":{"computeUnitsConsumed":0,"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"],"postBalances":[9997999990000,2000000000,1],"postTokenBalances":[],"preBalances":[9998999995000,1000000000,1],"preTokenBalances":[],"rewards":null,"status":{"Ok":null}},"transaction":["Achxxn2FVKasGFueEHTMBFmON/5ztN7JdYvc+Ss5CtLwfj6TDPgqGJvkBbWoIsZn6jM7j4WaV1eRMzlu9jz6cgUBAAEDf2Br+piF0OBJ+3GXgItWUGVGijKNmdpuOCd6tXg3VrkNO3L9o/rpfyK75/JQr7NPRByyLQ5llixat3eAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA8x5CF4ifnqtdsF7OWbHgMLgWbvhDrDObUKIMe4qomPgBAgIAAQwCAAAAAMqaOwAAAAA=","base64"],"version":"legacy"},{"meta":{"computeUnitsConsumed":0,"err":null,"fee":10000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499765100000,1000000000000000,1],"postTokenBalances":[],"preBalances":[499765110000,1000000000000000,1],"preTokenBalances":[],"rewards":null,"status":{"Ok":null}},"transaction":["AkC0jBUH+KTvqrw+XF/h4aA2u09VmQDbI8dEqPiGay0clJI52fhNEJ41IPmaBBjaweVM1Mh1B7J3VFNr7nlAxgDyt+2B+en7Q0PKoWT1ONDfEElmzIWVLhXuFX5KXZRpHQSrvKmP4x8j1dKJ2fbHxK0mMZLescxNMQsszjaXCT0GAgABA0DSMuP/AXO4osA2Lche6ae8CYVL66c/suPx4F4or2rxK8rwBthv3Sd+0di7ooNIjpqXz5HxZ1BkFltiD/91+EcHYUgdNXR0u3xNdiTr072z2DVec9EQQ/wNo1OAAAAAABHayu7YzW3oqcYF3ZBq9Na6SY+f1bF5L/g44JhaYX6WAQICAQF0DAAAAGq3AAAAAAAAHwEfAR4BHQEcARsBGgEZARgBFwEWARUBFAETARIBEQEQAQ8BDgENAQwBCwEKAQkBCAEHAQYBBQEEAQMBAgEBUqiGzWLvHkwvxR/wG+1DSHfhah9IsndPja3vZj6mCxAB83w2YwAAAAA=","base64"],"version":"legacy"}]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockWithConfig(
					context.TODO(),
					46986,
					GetBlockConfig{
						Encoding:                       GetBlockConfigEncodingBase64,
						Rewards:                        pointer.Get(false),
						MaxSupportedTransactionVersion: pointer.Get[uint8](0),
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetBlock]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetBlock{
					ParentSlot:        46985,
					BlockHeight:       pointer.Get[int64](46984),
					BlockTime:         pointer.Get[int64](1663315273),
					PreviousBlockhash: "2ChS54SSzba9zFqsbQJviVkmXXN7zdpMRXPjvzY93cfX",
					Blockhash:         "5GLYVkvcF7ygCQpFVbeZsaCcSBaw9qofJjdtzwVXYPu1",
					Transactions: []GetBlockTransaction{
						{
							Meta: &TransactionMeta{
								Err: nil,
								Fee: 5000,
								PreBalances: []int64{
									10000000000000,
									1,
									0,
								},
								PostBalances: []int64{
									9998999995000,
									1,
									1000000000,
								},
								LogMessages: []string{
									"Program 11111111111111111111111111111111 invoke [1]",
									"Program 11111111111111111111111111111111 success",
								},
								LoadedAddresses: TransactionLoadedAddresses{
									Writable: []string{
										"test111111111111111111111111111111111111111",
									},
									Readonly: []string{},
								},
								PreTokenBalances:  []TransactionMetaTokenBalance{},
								PostTokenBalances: []TransactionMetaTokenBalance{},
								InnerInstructions: []TransactionMetaInnerInstruction{},
							},
							Transaction: []any{
								"Ad8N7teP2bRE2xzLJRc+5zw2OOJWF4cXtfSxbMTwijdxQMkGB4KMAAzG/svcdNGzaA5QiH4tb1ly0e7oFyJ53gCAAQABAn9ga/qYhdDgSftxl4CLVlBlRooyjZnabjgnerV4N1a5AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADzHkIXiJ+eq12wXs5ZseAwuBZu+EOsM5tQogx7iqiY+AEBAgACDAIAAAAAypo7AAAAAAGWXqBzSZja4R/6KEg0G0P5Y1bvQBhRNhSSdA0KHQ84BAECAA==",
								"base64",
							},
							Version: float64(0),
						},
						{
							Meta: &TransactionMeta{
								Err: nil,
								Fee: 5000,
								PreBalances: []int64{
									9998999995000,
									1000000000,
									1,
								},
								PostBalances: []int64{
									9997999990000,
									2000000000,
									1,
								},
								LogMessages: []string{
									"Program 11111111111111111111111111111111 invoke [1]",
									"Program 11111111111111111111111111111111 success",
								},
								LoadedAddresses: TransactionLoadedAddresses{
									Writable: []string{},
									Readonly: []string{},
								},
								PreTokenBalances:  []TransactionMetaTokenBalance{},
								PostTokenBalances: []TransactionMetaTokenBalance{},
								InnerInstructions: []TransactionMetaInnerInstruction{},
							},
							Transaction: []any{
								"Achxxn2FVKasGFueEHTMBFmON/5ztN7JdYvc+Ss5CtLwfj6TDPgqGJvkBbWoIsZn6jM7j4WaV1eRMzlu9jz6cgUBAAEDf2Br+piF0OBJ+3GXgItWUGVGijKNmdpuOCd6tXg3VrkNO3L9o/rpfyK75/JQr7NPRByyLQ5llixat3eAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA8x5CF4ifnqtdsF7OWbHgMLgWbvhDrDObUKIMe4qomPgBAgIAAQwCAAAAAMqaOwAAAAA=",
								"base64",
							},
							Version: "legacy",
						},
						{
							Meta: &TransactionMeta{
								Err: nil,
								Fee: 10000,
								PreBalances: []int64{
									499765110000,
									1000000000000000,
									1,
								},
								PostBalances: []int64{
									499765100000,
									1000000000000000,
									1,
								},
								LogMessages: []string{
									"Program Vote111111111111111111111111111111111111111 invoke [1]",
									"Program Vote111111111111111111111111111111111111111 success",
								},
								LoadedAddresses: TransactionLoadedAddresses{
									Writable: []string{},
									Readonly: []string{},
								},
								PreTokenBalances:  []TransactionMetaTokenBalance{},
								PostTokenBalances: []TransactionMetaTokenBalance{},
								InnerInstructions: []TransactionMetaInnerInstruction{},
							},
							Transaction: []any{
								"AkC0jBUH+KTvqrw+XF/h4aA2u09VmQDbI8dEqPiGay0clJI52fhNEJ41IPmaBBjaweVM1Mh1B7J3VFNr7nlAxgDyt+2B+en7Q0PKoWT1ONDfEElmzIWVLhXuFX5KXZRpHQSrvKmP4x8j1dKJ2fbHxK0mMZLescxNMQsszjaXCT0GAgABA0DSMuP/AXO4osA2Lche6ae8CYVL66c/suPx4F4or2rxK8rwBthv3Sd+0di7ooNIjpqXz5HxZ1BkFltiD/91+EcHYUgdNXR0u3xNdiTr072z2DVec9EQQ/wNo1OAAAAAABHayu7YzW3oqcYF3ZBq9Na6SY+f1bF5L/g44JhaYX6WAQICAQF0DAAAAGq3AAAAAAAAHwEfAR4BHQEcARsBGgEZARgBFwEWARUBFAETARIBEQEQAQ8BDgENAQwBCwEKAQkBCAEHAQYBBQEEAQMBAgEBUqiGzWLvHkwvxR/wG+1DSHfhah9IsndPja3vZj6mCxAB83w2YwAAAAA=",
								"base64",
							},
							Version: "legacy",
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"getBlock","params":[237,{"encoding":"base64"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":237,"blockTime":1666899961,"blockhash":"EdSZ7r12ZipGivMX4NDpFUT9pCmDkCPFHQyWwLSB2Wmv","parentSlot":236,"previousBlockhash":"DhaQz2FuLkredqU2VGBkoUVZzk4QTCskDezYZS7ua9rE","rewards":[{"commission":null,"lamports":7500,"postBalance":499998890000,"pubkey":"D7SRyWHVRfZngyzUciWJXNg6qPRV1Ga5ickasA3RKWbs","rewardType":"Fee"}],"transactions":[{"meta":{"computeUnitsConsumed":185,"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units","Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success"],"postBalances":[19999995000,1141440],"postTokenBalances":[],"preBalances":[20000000000,1141440],"preTokenBalances":[],"returnData":{"data":["AQIDBAU=","base64"],"programId":"35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP"},"rewards":[],"status":{"Ok":null}},"transaction":["ATzOQjvgBIzjOQc6WnMjWrjVwswSbhuFctuEe07qdy8aQUKCN2wnhexRJOtnuRaO2Ej+0ZVHHUbUuK4dlEXYDQMBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYBIFo+itoZLn+XpB0rgSk0J0TAnzTRp1+XLycCCIW6nbAQEAAA==","base64"]},{"meta":{"computeUnitsConsumed":0,"err":null,"fee":10000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499998882500,1000000000000000,1],"postTokenBalances":[],"preBalances":[499998892500,1000000000000000,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"transaction":["AvTyTNWKN/M/HMqK7PHNq92HzS+KzQgUOv+2/9gM5eCIblqBhtNggpRwS8QdEdTGZNPy9sGV4QnQdbwgd29ceAOEIqk0ngBBrb2L3/X1cnxOi3r9HeH5qsMVSF3mrWPQQY0HQMJcJ4P9h6/ZoSWPQ5EoBY6pTITUTYGnfOh4QiMCAgABA7Pye4CQ39EMJf/39tiZcvcBNSICJuXthpHIWtsgthjKpNP8p02KRZAs+x1AyHhXECmbuluryPHNRdTp7sPN4Z4HYUgdNXR0u3xNdiTr072z2DVec9EQQ/wNo1OAAAAAALyxQZliVJ+XOv1JCOY6vzkCK+ooOH3hT8ccNVIAeRWvAQICAQF0DAAAAM0AAAAAAAAAHwEfAR4BHQEcARsBGgEZARgBFwEWARUBFAETARIBEQEQAQ8BDgENAQwBCwEKAQkBCAEHAQYBBQEEAQMBAgEBgvsUbTEPD4Zz4/0zYUTxwrHil3ufJ2oDtVhxgydDp9cB+t9aYwAAAAA=","base64"]}]},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetBlockWithConfig(
					context.TODO(),
					237,
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
					ParentSlot:        236,
					BlockHeight:       pointer.Get[int64](237),
					BlockTime:         pointer.Get[int64](1666899961),
					PreviousBlockhash: "DhaQz2FuLkredqU2VGBkoUVZzk4QTCskDezYZS7ua9rE",
					Blockhash:         "EdSZ7r12ZipGivMX4NDpFUT9pCmDkCPFHQyWwLSB2Wmv",
					Rewards: []GetBlockReward{
						{
							Pubkey:       "D7SRyWHVRfZngyzUciWJXNg6qPRV1Ga5ickasA3RKWbs",
							Lamports:     7500,
							PostBalances: 499998890000,
							RewardType:   "Fee",
							Commission:   nil,
						},
					},
					Transactions: []GetBlockTransaction{
						{
							Meta: &TransactionMeta{
								Err: nil,
								Fee: 5000,
								PreBalances: []int64{
									20000000000,
									1141440,
								},
								PostBalances: []int64{
									19999995000,
									1141440,
								},
								LogMessages: []string{
									"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]",
									"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units",
									"Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=",
									"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success",
								},
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
							},
							Transaction: []any{
								"ATzOQjvgBIzjOQc6WnMjWrjVwswSbhuFctuEe07qdy8aQUKCN2wnhexRJOtnuRaO2Ej+0ZVHHUbUuK4dlEXYDQMBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYBIFo+itoZLn+XpB0rgSk0J0TAnzTRp1+XLycCCIW6nbAQEAAA==",
								"base64",
							},
						},
						{
							Meta: &TransactionMeta{
								Err: nil,
								Fee: 10000,
								PreBalances: []int64{
									499998892500,
									1000000000000000,
									1,
								},
								PostBalances: []int64{
									499998882500,
									1000000000000000,
									1,
								},
								LogMessages: []string{
									"Program Vote111111111111111111111111111111111111111 invoke [1]",
									"Program Vote111111111111111111111111111111111111111 success",
								},
								LoadedAddresses: TransactionLoadedAddresses{
									Writable: []string{},
									Readonly: []string{},
								},
								PreTokenBalances:  []TransactionMetaTokenBalance{},
								PostTokenBalances: []TransactionMetaTokenBalance{},
								InnerInstructions: []TransactionMetaInnerInstruction{},
							},
							Transaction: []any{
								"AvTyTNWKN/M/HMqK7PHNq92HzS+KzQgUOv+2/9gM5eCIblqBhtNggpRwS8QdEdTGZNPy9sGV4QnQdbwgd29ceAOEIqk0ngBBrb2L3/X1cnxOi3r9HeH5qsMVSF3mrWPQQY0HQMJcJ4P9h6/ZoSWPQ5EoBY6pTITUTYGnfOh4QiMCAgABA7Pye4CQ39EMJf/39tiZcvcBNSICJuXthpHIWtsgthjKpNP8p02KRZAs+x1AyHhXECmbuluryPHNRdTp7sPN4Z4HYUgdNXR0u3xNdiTr072z2DVec9EQQ/wNo1OAAAAAALyxQZliVJ+XOv1JCOY6vzkCK+ooOH3hT8ccNVIAeRWvAQICAQF0DAAAAM0AAAAAAAAAHwEfAR4BHQEcARsBGgEZARgBFwEWARUBFAETARIBEQEQAQ8BDgENAQwBCwEKAQkBCAEHAQYBBQEEAQMBAgEBgvsUbTEPD4Zz4/0zYUTxwrHil3ufJ2oDtVhxgydDp9cB+t9aYwAAAAA=",
								"base64",
							},
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
