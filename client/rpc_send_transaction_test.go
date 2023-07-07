package client

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

func TestClient_SendTransaction(t *testing.T) {
	b, _ := base64.StdEncoding.DecodeString("AS0vVCOi6XOkuufPHS3HyoJPInhwLzT11XpPBYBC9gp/bK9yC94aoeyiuZHZBF7MdddUJ2TPhKZiVyuJuaKp1QQBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4FSlNamSkhBk0k6HFg2jh8fDW13bySu4HkH6hAQQVEjYoKT69yJM5QhVyj/TbbwW+0VbubU5Ssg4cY/m97ik7YAQEABPCfkbs=")
	tx, _ := types.TransactionDeserialize(b)
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"sendTransaction", "params":["AS0vVCOi6XOkuufPHS3HyoJPInhwLzT11XpPBYBC9gp/bK9yC94aoeyiuZHZBF7MdddUJ2TPhKZiVyuJuaKp1QQBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4FSlNamSkhBk0k6HFg2jh8fDW13bySu4HkH6hAQQVEjYoKT69yJM5QhVyj/TbbwW+0VbubU5Ssg4cY/m97ik7YAQEABPCfkbs=", {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.SendTransaction(
						context.Background(),
						tx,
					)
				},
				ExpectedValue: "uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV",
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_SendTransactionWithConfig(t *testing.T) {
	b, _ := base64.StdEncoding.DecodeString("AS0vVCOi6XOkuufPHS3HyoJPInhwLzT11XpPBYBC9gp/bK9yC94aoeyiuZHZBF7MdddUJ2TPhKZiVyuJuaKp1QQBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4FSlNamSkhBk0k6HFg2jh8fDW13bySu4HkH6hAQQVEjYoKT69yJM5QhVyj/TbbwW+0VbubU5Ssg4cY/m97ik7YAQEABPCfkbs=")
	tx, _ := types.TransactionDeserialize(b)
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				Name:         "empty",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"sendTransaction", "params":["AS0vVCOi6XOkuufPHS3HyoJPInhwLzT11XpPBYBC9gp/bK9yC94aoeyiuZHZBF7MdddUJ2TPhKZiVyuJuaKp1QQBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4FSlNamSkhBk0k6HFg2jh8fDW13bySu4HkH6hAQQVEjYoKT69yJM5QhVyj/TbbwW+0VbubU5Ssg4cY/m97ik7YAQEABPCfkbs=", {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.SendTransaction(
						context.Background(),
						tx,
					)
				},
				ExpectedValue: "uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV",
				ExpectedError: nil,
			},
			{
				Name:         "with skip preflight",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"sendTransaction", "params":["AS0vVCOi6XOkuufPHS3HyoJPInhwLzT11XpPBYBC9gp/bK9yC94aoeyiuZHZBF7MdddUJ2TPhKZiVyuJuaKp1QQBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4FSlNamSkhBk0k6HFg2jh8fDW13bySu4HkH6hAQQVEjYoKT69yJM5QhVyj/TbbwW+0VbubU5Ssg4cY/m97ik7YAQEABPCfkbs=", {"encoding":"base64", "skipPreflight": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.SendTransactionWithConfig(
						context.Background(),
						tx,
						SendTransactionConfig{
							SkipPreflight: true,
						},
					)
				},
				ExpectedValue: "uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV",
				ExpectedError: nil,
			},
			{
				Name:         "with max reties",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"sendTransaction", "params":["AS0vVCOi6XOkuufPHS3HyoJPInhwLzT11XpPBYBC9gp/bK9yC94aoeyiuZHZBF7MdddUJ2TPhKZiVyuJuaKp1QQBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4FSlNamSkhBk0k6HFg2jh8fDW13bySu4HkH6hAQQVEjYoKT69yJM5QhVyj/TbbwW+0VbubU5Ssg4cY/m97ik7YAQEABPCfkbs=", {"encoding":"base64", "maxRetries": 5}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.SendTransactionWithConfig(
						context.Background(),
						tx,
						SendTransactionConfig{
							MaxRetries: 5,
						},
					)
				},
				ExpectedValue: "uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV",
				ExpectedError: nil,
			},
			{
				Name:         "with preflight commitment",
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"sendTransaction", "params":["AS0vVCOi6XOkuufPHS3HyoJPInhwLzT11XpPBYBC9gp/bK9yC94aoeyiuZHZBF7MdddUJ2TPhKZiVyuJuaKp1QQBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4FSlNamSkhBk0k6HFg2jh8fDW13bySu4HkH6hAQQVEjYoKT69yJM5QhVyj/TbbwW+0VbubU5Ssg4cY/m97ik7YAQEABPCfkbs=", {"encoding":"base64", "preflightCommitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.SendTransactionWithConfig(
						context.Background(),
						tx,
						SendTransactionConfig{
							PreflightCommitment: rpc.CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: "uQ1KB2ZS7WDN5Jf4nFxDCC75reGMdUW8S7mybWfZPzMPo4TULPE8NCkJAaQ5ifCoDmreCnzdPmFjLrDTRJ6QLbV",
				ExpectedError: nil,
			},
		},
	)
}
