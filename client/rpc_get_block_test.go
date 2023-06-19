package client

import (
	"context"
	"testing"
	"time"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/internal/client_test"
	"github.com/portto/solana-go-sdk/pkg/pointer"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

func TestClient_BlockWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				Name:         "transaction detail: signatures",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64", "transactionDetails":"signatures", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}],"signatures":["3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A"]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetBlockWithConfig(context.Background(), 33, GetBlockConfig{TransactionDetails: rpc.GetBlockConfigTransactionDetailsSignatures})
				},
				ExpectedValue: &Block{
					Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
					BlockTime:         pointer.Get[time.Time](time.Unix(1631803928, 0)),
					BlockHeight:       pointer.Get[int64](33),
					PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
					ParentSlot:        32,
					Transactions:      nil,
					Signatures:        []string{"3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A"},
					Rewards: []Reward{
						{
							Pubkey:       common.PublicKeyFromString("9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX"),
							Lamports:     5000,
							PostBalances: 499999840001,
							RewardType:   pointer.Get[rpc.RewardType](rpc.RewardTypeFee),
							Commission:   nil,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "transaciton detail: none",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64", "transactionDetails":"none", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetBlockWithConfig(context.Background(), 33, GetBlockConfig{TransactionDetails: rpc.GetBlockConfigTransactionDetailsNone})
				},
				ExpectedValue: &Block{
					Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
					BlockTime:         pointer.Get[time.Time](time.Unix(1631803928, 0)),
					BlockHeight:       pointer.Get[int64](33),
					PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
					ParentSlot:        32,
					Transactions:      nil,
					Signatures:        nil,
					Rewards: []Reward{
						{
							Pubkey:       common.PublicKeyFromString("9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX"),
							Lamports:     5000,
							PostBalances: 499999840001,
							RewardType:   pointer.Get[rpc.RewardType](rpc.RewardTypeFee),
							Commission:   nil,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				Name:         "no reward",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64", "rewards": false, "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","transactions":[{"meta":{"err":null,"fee":10000,"innerInstructions":[],"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499999835001,1000000000000000,143487360,1169280,1],"postTokenBalances":[],"preBalances":[499999845001,1000000000000000,143487360,1169280,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"transaction":["AnXU8JYCIrc73JwxK9traTSp3EZdmnJp0B5luW8CCzr7GnFd/SjIMXiG4qbN5CwyEVhbpORzBUpB/253cNtS1A+0rWE+nrDqWRQ2OVU727PU4NtR611jY+10Q+F6lCZDsJt46b6oXz3PN5WGxTQk7mC4YhCbYsTcalWBkltA8KgPAgADBXszyT4GLb26BFuAAUXtW0B75zurDhXE7UOYKHFkpIlKJMmZpq+FRXTx8jzBMy1YsdkCo0kyLDdF2Q3NhXRdEosGp9UXGS8Kr8byZeP7d8x62oLFKdC+OxNuLQBVIAAAAAan1RcYx3TJKFZjmGkdXraLXrijm0ttXHNVWyEAAAAAB2FIHTV0dLt8TXYk69O9s9g1XnPREEP8DaNTgAAAAACrUBylgzc0SSCUPSfMJC3TI6KJEzs834KdMIMJci+UYAEEBAECAwE9AgAAAAEAAAAAAAAAIAAAAAAAAAAGCHSVIc5Betdf+NkRi4YR2D3abNLvpbI83qnB7EvNsAEZWkNhAAAAAA==","base64"]}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetBlockWithConfig(context.Background(), 33, GetBlockConfig{Rewards: pointer.Get(false)})
				},
				ExpectedValue: &Block{
					Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
					BlockTime:         pointer.Get[time.Time](time.Unix(1631803928, 0)),
					BlockHeight:       pointer.Get[int64](33),
					PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
					ParentSlot:        32,
					Transactions: []BlockTransaction{
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
									mustBase58Decode(t, "3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A"),
									mustBase58Decode(t, "4cWqSVUcxTujZ6eHtNWESwCrBUfidbZ1J124VU2jY9TQpXxyHSDku1NiZhw95SzXe1mGihiP9AdQNEkLMAdvBYPQ"),
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
											Data:           mustBase58Decode(t, "2ZjTR1vUs2pHXyTLuZA9zjpNqav47YU1uqenSEcYn6xkrdmMkUJK8JDHd5TcEU7K5R9pbB2UxbY95zDzHio"),
										},
									},
									AddressLookupTables: []types.CompiledAddressLookupTable{},
								},
							},
							AccountKeys: []common.PublicKey{
								common.PublicKeyFromString("9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX"),
								common.PublicKeyFromString("3UbyTNpi3omt7hfEqQRB5844VANQFWiq8uEDNCrSwAVG"),
								common.PublicKeyFromString("SysvarS1otHashes111111111111111111111111111"),
								common.PublicKeyFromString("SysvarC1ock11111111111111111111111111111111"),
								common.PublicKeyFromString("Vote111111111111111111111111111111111111111"),
							},
						},
					},
					Signatures: nil,
					Rewards:    nil,
				},
				ExpectedError: nil,
			},
		},
	)
}
