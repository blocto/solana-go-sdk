package rpc

import (
	"context"
	"testing"
)

func TestSendTransaction(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"sendTransaction","params":["37u9WtQpcm6ULa3Vmu7ySnANv"]}`,
			ResponseBody: `{"error":{"code":-32602,"message":"io error: failed to fill whole buffer"},"id":1,"jsonrpc":"2.0"}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.SendTransaction(context.Background(), "37u9WtQpcm6ULa3Vmu7ySnANv")
			},
			ExpectedResponse: SendTransactionResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error: &ErrorResponse{
						Code:    -32602,
						Message: `io error: failed to fill whole buffer`,
					},
				},
				Result: "",
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"sendTransaction","params":["582FZW9Ed2W4mLJVXPbHjDNw2LPZHL3UB3FRovi2qwzN3DZyRY26mhifhs6E9cpgaBqMeoPwE8vjU2LT63EfjxmH8BqkpJd2huKNBbZiV9DsBrTQwez8gfuiv5d9nt9af7pYdRhj1LH1PGF3cYVz1pkGwuyRnmnPqRFHFerXnNtDR5LncKLxoTi2So8LqJcdrQWfMRvqUJNB7EJBGrsvXtYSFdkXa3wqmsyY9a6PVwoGh2tcAY9o9DzTS16Cp1JFceBq4dQkNHxxnpUpPw6Rgey1SLveY7XUv61j5"]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32002,"message":"Transaction simulation failed: Blockhash not found","data":{"accounts":null,"err":"BlockhashNotFound","logs":[]}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.SendTransaction(context.Background(), "582FZW9Ed2W4mLJVXPbHjDNw2LPZHL3UB3FRovi2qwzN3DZyRY26mhifhs6E9cpgaBqMeoPwE8vjU2LT63EfjxmH8BqkpJd2huKNBbZiV9DsBrTQwez8gfuiv5d9nt9af7pYdRhj1LH1PGF3cYVz1pkGwuyRnmnPqRFHFerXnNtDR5LncKLxoTi2So8LqJcdrQWfMRvqUJNB7EJBGrsvXtYSFdkXa3wqmsyY9a6PVwoGh2tcAY9o9DzTS16Cp1JFceBq4dQkNHxxnpUpPw6Rgey1SLveY7XUv61j5")
			},
			ExpectedResponse: SendTransactionResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error: &ErrorResponse{
						Code:    -32002,
						Message: `Transaction simulation failed: Blockhash not found`,
						Data: map[string]interface{}{
							"accounts": nil,
							"err":      "BlockhashNotFound",
							"logs":     []interface{}{},
						},
					},
				},
				Result: "",
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"sendTransaction","params":["6RU3epX3frKoMKHLGjcH1pJQKt2ngYU2sTMB1siJfZFZFucUorBNpUT4JF3VCLjWRh1FMhadsbLnbUwoRRfWuVXrpiHbUa42pz9HpSvwospNeMzZLNTV9PBvS3CZQtLwrEzdv93kPo1uKquNWPdzBRSfLo13aCRaNfxvZNwg8mZb1cibSLRTGqvYKJhFqrVRaxCovZyQ9jVPBNG71hVXhzUnNHSKMTDRZAGqodZLodjZU4k4Eo4RZQ5g56R7cAMcPd4Pet8g5WGmSZwUAHauZjFY7jKxAB4eYRr3R"]}`,
			ResponseBody: `{"jsonrpc":"2.0","error":{"code":-32005,"message":"Node is behind by 42 slots","data":{"numSlotsBehind":42}},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.SendTransaction(context.Background(), "6RU3epX3frKoMKHLGjcH1pJQKt2ngYU2sTMB1siJfZFZFucUorBNpUT4JF3VCLjWRh1FMhadsbLnbUwoRRfWuVXrpiHbUa42pz9HpSvwospNeMzZLNTV9PBvS3CZQtLwrEzdv93kPo1uKquNWPdzBRSfLo13aCRaNfxvZNwg8mZb1cibSLRTGqvYKJhFqrVRaxCovZyQ9jVPBNG71hVXhzUnNHSKMTDRZAGqodZLodjZU4k4Eo4RZQ5g56R7cAMcPd4Pet8g5WGmSZwUAHauZjFY7jKxAB4eYRr3R")
			},
			ExpectedResponse: SendTransactionResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error: &ErrorResponse{
						Code:    -32005,
						Message: `Node is behind by 42 slots`,
						Data: map[string]interface{}{
							"numSlotsBehind": float64(42),
						},
					},
				},
				Result: "",
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"sendTransaction","params":["KAmmQBoPnTWMnMCsoa8chJJN6R2EdP2X9i48wDayjdg6WBpaj88yYQ655LNpXiSa1pPVsp6ynCBaaQeSi1ZKQKuziYss4LgmfjuWhQyRyzZfkfM5W7AsQCG1XTNWbKmXzXpNPLXLLx3tTN49n3Q7eq7g5kJnsJZNJJKnYgHSbwEQVcTfqecC6P3V6CtMEeCjM9PzRKnaUXc5seW6gcMDhgTfgPbacVjkZsrUh6mwC84amEJrwioYcYqyh"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":"5rgpegm86vwXotD2Z7WWW1DxpSxmWGQ9g4RMoBJvxJ2xiVF6TNCvGseZ3A1uisew9tGrdKirkkHUGjQW8uNqz9BW","id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.SendTransaction(context.Background(), "KAmmQBoPnTWMnMCsoa8chJJN6R2EdP2X9i48wDayjdg6WBpaj88yYQ655LNpXiSa1pPVsp6ynCBaaQeSi1ZKQKuziYss4LgmfjuWhQyRyzZfkfM5W7AsQCG1XTNWbKmXzXpNPLXLLx3tTN49n3Q7eq7g5kJnsJZNJJKnYgHSbwEQVcTfqecC6P3V6CtMEeCjM9PzRKnaUXc5seW6gcMDhgTfgPbacVjkZsrUh6mwC84amEJrwioYcYqyh")
			},
			ExpectedResponse: SendTransactionResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
				},
				Result: "5rgpegm86vwXotD2Z7WWW1DxpSxmWGQ9g4RMoBJvxJ2xiVF6TNCvGseZ3A1uisew9tGrdKirkkHUGjQW8uNqz9BW",
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"sendTransaction","params":["ATRYBqVekVOUHarhI0hXlMCqlcAxn7dNuH/TYXZFa9aPAGDT6tMC7NIqExjRV2EXy3HJB7kPjiNwT4fLdxBFQw4BAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGESezVpevNbRbWDSu3ezgl24yUCjGiZzveFPdbyKi4TAQECAAAMAgAAAAEAAAAAAAAA",{"encoding":"base64"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":"23hVrUsx17XuRbGndEPhShvaMT7HnxEs4dppq2NqvFJTDbEFm11a16f6W4Abs7RfXpzKQRKRoCiyHSNvBmvhVwR7","id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.SendTransactionWithConfig(
					context.Background(),
					"ATRYBqVekVOUHarhI0hXlMCqlcAxn7dNuH/TYXZFa9aPAGDT6tMC7NIqExjRV2EXy3HJB7kPjiNwT4fLdxBFQw4BAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGESezVpevNbRbWDSu3ezgl24yUCjGiZzveFPdbyKi4TAQECAAAMAgAAAAEAAAAAAAAA",
					SendTransactionConfig{
						Encoding: SendTransactionConfigEncodingBase64,
					},
				)
			},
			ExpectedResponse: SendTransactionResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
				},
				Result: "23hVrUsx17XuRbGndEPhShvaMT7HnxEs4dppq2NqvFJTDbEFm11a16f6W4Abs7RfXpzKQRKRoCiyHSNvBmvhVwR7",
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"sendTransaction","params":["CayPuTQXGJTdD36zymjYnGPDovnMr9ZaTvqN4W4JDe3bNs6WySjg4qut7AQYHbMxf38f95qd7cQ3GtzHM1CWjoui6qPPkaMMAu9fyCvfsGXFkVjeTczjSrBCWz6t74m3voiTaLpVEG8WHosKfSVUUC1UMHdgHKp63ZZeA1k9ZH2hgwfByfnEftgkMTEGyQ8mMx1q8MZVbbQGs2eNeTKxbUupCp8WotHZ9YrtqwJfLXF8HMHqGHZ8VdpMV",{"skipPreflight":true}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":"2F53DggXYWLzczigoMr7smSEZtWSKmsWr7HMJQiNbTBdjjcN54LUMWdvTLj46MH7rAnJVPjJEjRjjXKeG7mssmZb","id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.SendTransactionWithConfig(
					context.Background(),
					"CayPuTQXGJTdD36zymjYnGPDovnMr9ZaTvqN4W4JDe3bNs6WySjg4qut7AQYHbMxf38f95qd7cQ3GtzHM1CWjoui6qPPkaMMAu9fyCvfsGXFkVjeTczjSrBCWz6t74m3voiTaLpVEG8WHosKfSVUUC1UMHdgHKp63ZZeA1k9ZH2hgwfByfnEftgkMTEGyQ8mMx1q8MZVbbQGs2eNeTKxbUupCp8WotHZ9YrtqwJfLXF8HMHqGHZ8VdpMV",
					SendTransactionConfig{
						SkipPreflight: true,
					},
				)
			},
			ExpectedResponse: SendTransactionResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
				},
				Result: "2F53DggXYWLzczigoMr7smSEZtWSKmsWr7HMJQiNbTBdjjcN54LUMWdvTLj46MH7rAnJVPjJEjRjjXKeG7mssmZb",
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"sendTransaction","params":["HvPMZonNNzD9M2VY3DBJUHVw8fXuym23SB193SX7qMgHu2BhTwaanTDmaCg4XiTFqHnLAx5Tirim87BqYuvEdZsEcEaTRjPBnFhMR8cXBbKGkZnhNNoU6F8GcZ2gjYfFV8WkABQa2gimsyiTLzifHroVYuB7qpH8VFUGkbvDuqsJPykmhWx1dk94LUsic2e1PRLJkeKTPojSvRZomjXHDQV2d4izfNNZVTViKRfhwvdqiauX7niFBraes",{"preflightCommitment":"finalized","maxRetries":5}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":"2F53DggXYWLzczigoMr7smSEZtWSKmsWr7HMJQiNbTBdjjcN54LUMWdvTLj46MH7rAnJVPjJEjRjjXKeG7mssmZb","id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.SendTransactionWithConfig(
					context.Background(),
					"HvPMZonNNzD9M2VY3DBJUHVw8fXuym23SB193SX7qMgHu2BhTwaanTDmaCg4XiTFqHnLAx5Tirim87BqYuvEdZsEcEaTRjPBnFhMR8cXBbKGkZnhNNoU6F8GcZ2gjYfFV8WkABQa2gimsyiTLzifHroVYuB7qpH8VFUGkbvDuqsJPykmhWx1dk94LUsic2e1PRLJkeKTPojSvRZomjXHDQV2d4izfNNZVTViKRfhwvdqiauX7niFBraes",
					SendTransactionConfig{
						PreflightCommitment: CommitmentFinalized,
						MaxRetries:          5,
					},
				)
			},
			ExpectedResponse: SendTransactionResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
				},
				Result: "2F53DggXYWLzczigoMr7smSEZtWSKmsWr7HMJQiNbTBdjjcN54LUMWdvTLj46MH7rAnJVPjJEjRjjXKeG7mssmZb",
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
