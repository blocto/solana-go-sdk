package client

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/rpc"
)

func TestClient_SimulateTransaction(t *testing.T) {
	tx := mustDeserializeBase64Tx(t, "Ab/yMEK7qNgGxaPMg2XaVnwwLMqnY8FTeJrA9qJ1nOBFX08BHycnp3/9WOxOY53+eZnbkT2/+6Mx7w+DsuVN8ggBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYEsGZsdWhvaw9ENEPFBEi4eBna4CphPQWWcgU4yARSnVAQEAAA==")
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"simulateTransaction","params":["Ab/yMEK7qNgGxaPMg2XaVnwwLMqnY8FTeJrA9qJ1nOBFX08BHycnp3/9WOxOY53+eZnbkT2/+6Mx7w+DsuVN8ggBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYEsGZsdWhvaw9ENEPFBEi4eBna4CphPQWWcgU4yARSnVAQEAAA==", {"encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.5","slot":159776096},"value":{"accounts":null,"err":null,"logs":["Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units","Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success"],"returnData":{"data":["AQIDBAU=","base64"],"programId":"35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP"},"unitsConsumed":185}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.SimulateTransaction(
						context.Background(),
						tx,
					)
				},
				ExpectedValue: SimulateTransaction{
					Accounts: nil,
					Logs: []string{
						"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]",
						"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units",
						"Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=",
						"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success",
					},
					ReturnData: &ReturnData{
						ProgramId: common.PublicKeyFromString("35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP"),
						Data:      []byte{1, 2, 3, 4, 5},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_SimulateTransactionAndContext(t *testing.T) {
	tx := mustDeserializeBase64Tx(t, "Ab/yMEK7qNgGxaPMg2XaVnwwLMqnY8FTeJrA9qJ1nOBFX08BHycnp3/9WOxOY53+eZnbkT2/+6Mx7w+DsuVN8ggBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYEsGZsdWhvaw9ENEPFBEi4eBna4CphPQWWcgU4yARSnVAQEAAA==")
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"simulateTransaction","params":["Ab/yMEK7qNgGxaPMg2XaVnwwLMqnY8FTeJrA9qJ1nOBFX08BHycnp3/9WOxOY53+eZnbkT2/+6Mx7w+DsuVN8ggBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYEsGZsdWhvaw9ENEPFBEi4eBna4CphPQWWcgU4yARSnVAQEAAA==", {"encoding": "base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.5","slot":159776096},"value":{"accounts":null,"err":null,"logs":["Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units","Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success"],"returnData":{"data":["AQIDBAU=","base64"],"programId":"35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP"},"unitsConsumed":185}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.SimulateTransactionAndContext(
						context.Background(),
						tx,
					)
				},
				ExpectedValue: rpc.ValueWithContext[SimulateTransaction]{
					Context: rpc.Context{
						Slot:       159776096,
						ApiVersion: "1.14.5",
					},
					Value: SimulateTransaction{
						Accounts: nil,
						Logs: []string{
							"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]",
							"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units",
							"Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=",
							"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success",
						},
						ReturnData: &ReturnData{
							ProgramId: common.PublicKeyFromString("35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP"),
							Data:      []byte{1, 2, 3, 4, 5},
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
