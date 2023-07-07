package rpc

import (
	"context"
	"testing"

	"github.com/blocto/solana-go-sdk/internal/client_test"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
)

func TestGetSignaturesForAddress(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignaturesForAddress", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"blockTime":1633580920,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"26UNKcerqcmHhRgFqtvtWUJZod91dGpkkAtRzKewvsZSoah33VZiFgKPmnHmMhBwsHT4bQjKdCgH88Faex5WkChh","slot":63372},{"blockTime":1633580919,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"vB73C8nWXv2ZwLjCRhTQXnkqqKDafE9uWyycQqajFzQcNFzjRUYq8ZqmtCm4qnwpGxXJWbkkRuyhiQ26zEpfk28","slot":63370},{"blockTime":1633580918,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"4ZAtdQ7wF8EPXVf43ZpyqMiKqpaiGppysSxQ54X31Q92ht1pKgQymiJwRnUc3h4cUjRJfFrVprNJhW6UfybJMAQP","slot":63369},{"blockTime":1633580916,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"63BWysCcY6CcVT9G4FrfD1XuZ2DUswSTuJM5jeULUWwgp3BJhMnAStH3gTkUPXeUEjYoQhVNfd61RxxxWwsezi8y","slot":63367},{"blockTime":1633580915,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"pxS5UZhzvk8p5qiKonkAeVBjP1ipujERPMMNZ2ZexK1PyU8RvZxQwYjiv9YJYP4CpRrHAxTET2rNZ6LNf4aYvDN","slot":63365}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignaturesForAddress(
						context.Background(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
					)
				},
				ExpectedValue: JsonRpcResponse[GetSignaturesForAddress]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetSignaturesForAddress{
						{
							Signature: "26UNKcerqcmHhRgFqtvtWUJZod91dGpkkAtRzKewvsZSoah33VZiFgKPmnHmMhBwsHT4bQjKdCgH88Faex5WkChh",
							Slot:      63372,
							BlockTime: pointer.Get[int64](1633580920),
						},
						{
							Signature: "vB73C8nWXv2ZwLjCRhTQXnkqqKDafE9uWyycQqajFzQcNFzjRUYq8ZqmtCm4qnwpGxXJWbkkRuyhiQ26zEpfk28",
							Slot:      63370,
							BlockTime: pointer.Get[int64](1633580919),
						},
						{
							Signature: "4ZAtdQ7wF8EPXVf43ZpyqMiKqpaiGppysSxQ54X31Q92ht1pKgQymiJwRnUc3h4cUjRJfFrVprNJhW6UfybJMAQP",
							Slot:      63369,
							BlockTime: pointer.Get[int64](1633580918),
						},
						{
							Signature: "63BWysCcY6CcVT9G4FrfD1XuZ2DUswSTuJM5jeULUWwgp3BJhMnAStH3gTkUPXeUEjYoQhVNfd61RxxxWwsezi8y",
							Slot:      63367,
							BlockTime: pointer.Get[int64](1633580916),
						},
						{
							Signature: "pxS5UZhzvk8p5qiKonkAeVBjP1ipujERPMMNZ2ZexK1PyU8RvZxQwYjiv9YJYP4CpRrHAxTET2rNZ6LNf4aYvDN",
							Slot:      63365,
							BlockTime: pointer.Get[int64](1633580915),
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignaturesForAddress", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"limit": 1, "commitment": "confirmed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"blockTime":1633581563,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"66z7UgyzEozBQ1moxc2ThzGtzbwETZ9bR5ExSUubWhQzqGuX1hQgCGSV1n22o96yuDCFY2dHeMNNLDnf6zjewy7C","slot":64265}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignaturesForAddressWithConfig(
						context.Background(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						GetSignaturesForAddressConfig{
							Limit:      1,
							Commitment: CommitmentConfirmed,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetSignaturesForAddress]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetSignaturesForAddress{
						{
							Signature: "66z7UgyzEozBQ1moxc2ThzGtzbwETZ9bR5ExSUubWhQzqGuX1hQgCGSV1n22o96yuDCFY2dHeMNNLDnf6zjewy7C",
							Slot:      64265,
							BlockTime: pointer.Get[int64](1633581563),
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignaturesForAddress", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"until": "vB73C8nWXv2ZwLjCRhTQXnkqqKDafE9uWyycQqajFzQcNFzjRUYq8ZqmtCm4qnwpGxXJWbkkRuyhiQ26zEpfk28"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"blockTime":1633581563,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"66z7UgyzEozBQ1moxc2ThzGtzbwETZ9bR5ExSUubWhQzqGuX1hQgCGSV1n22o96yuDCFY2dHeMNNLDnf6zjewy7C","slot":64265},{"blockTime":1633580920,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"26UNKcerqcmHhRgFqtvtWUJZod91dGpkkAtRzKewvsZSoah33VZiFgKPmnHmMhBwsHT4bQjKdCgH88Faex5WkChh","slot":63372}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignaturesForAddressWithConfig(
						context.Background(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						GetSignaturesForAddressConfig{
							Until: "vB73C8nWXv2ZwLjCRhTQXnkqqKDafE9uWyycQqajFzQcNFzjRUYq8ZqmtCm4qnwpGxXJWbkkRuyhiQ26zEpfk28",
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetSignaturesForAddress]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetSignaturesForAddress{
						{
							Signature: "66z7UgyzEozBQ1moxc2ThzGtzbwETZ9bR5ExSUubWhQzqGuX1hQgCGSV1n22o96yuDCFY2dHeMNNLDnf6zjewy7C",
							Slot:      64265,
							BlockTime: pointer.Get[int64](1633581563),
						},
						{
							Signature: "26UNKcerqcmHhRgFqtvtWUJZod91dGpkkAtRzKewvsZSoah33VZiFgKPmnHmMhBwsHT4bQjKdCgH88Faex5WkChh",
							Slot:      63372,
							BlockTime: pointer.Get[int64](1633580920),
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignaturesForAddress", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"before": "4ZAtdQ7wF8EPXVf43ZpyqMiKqpaiGppysSxQ54X31Q92ht1pKgQymiJwRnUc3h4cUjRJfFrVprNJhW6UfybJMAQP"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"blockTime":1633580916,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"63BWysCcY6CcVT9G4FrfD1XuZ2DUswSTuJM5jeULUWwgp3BJhMnAStH3gTkUPXeUEjYoQhVNfd61RxxxWwsezi8y","slot":63367},{"blockTime":1633580915,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"pxS5UZhzvk8p5qiKonkAeVBjP1ipujERPMMNZ2ZexK1PyU8RvZxQwYjiv9YJYP4CpRrHAxTET2rNZ6LNf4aYvDN","slot":63365}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignaturesForAddressWithConfig(
						context.Background(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						GetSignaturesForAddressConfig{
							Before: "4ZAtdQ7wF8EPXVf43ZpyqMiKqpaiGppysSxQ54X31Q92ht1pKgQymiJwRnUc3h4cUjRJfFrVprNJhW6UfybJMAQP",
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetSignaturesForAddress]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetSignaturesForAddress{
						{
							Signature: "63BWysCcY6CcVT9G4FrfD1XuZ2DUswSTuJM5jeULUWwgp3BJhMnAStH3gTkUPXeUEjYoQhVNfd61RxxxWwsezi8y",
							Slot:      63367,
							BlockTime: pointer.Get[int64](1633580916),
						},
						{
							Signature: "pxS5UZhzvk8p5qiKonkAeVBjP1ipujERPMMNZ2ZexK1PyU8RvZxQwYjiv9YJYP4CpRrHAxTET2rNZ6LNf4aYvDN",
							Slot:      63365,
							BlockTime: pointer.Get[int64](1633580915),
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSignaturesForAddress", "params":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ", {"limit": 1, "before": "4ZAtdQ7wF8EPXVf43ZpyqMiKqpaiGppysSxQ54X31Q92ht1pKgQymiJwRnUc3h4cUjRJfFrVprNJhW6UfybJMAQP"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"blockTime":1633580916,"confirmationStatus":"finalized","err":null,"memo":null,"signature":"63BWysCcY6CcVT9G4FrfD1XuZ2DUswSTuJM5jeULUWwgp3BJhMnAStH3gTkUPXeUEjYoQhVNfd61RxxxWwsezi8y","slot":63367}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetSignaturesForAddressWithConfig(
						context.Background(),
						"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
						GetSignaturesForAddressConfig{
							Limit:  1,
							Before: "4ZAtdQ7wF8EPXVf43ZpyqMiKqpaiGppysSxQ54X31Q92ht1pKgQymiJwRnUc3h4cUjRJfFrVprNJhW6UfybJMAQP",
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetSignaturesForAddress]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetSignaturesForAddress{
						{
							Signature: "63BWysCcY6CcVT9G4FrfD1XuZ2DUswSTuJM5jeULUWwgp3BJhMnAStH3gTkUPXeUEjYoQhVNfd61RxxxWwsezi8y",
							Slot:      63367,
							BlockTime: pointer.Get[int64](1633580916),
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
