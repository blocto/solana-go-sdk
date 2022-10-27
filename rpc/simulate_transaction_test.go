package rpc

import (
	"context"
	"testing"
)

func TestSimulateTransaction(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"simulateTransaction", "params":["5nxpoKAc5anKyiJuwj5f2SLxnruNHDjpFFz1TAw5VvpL1L4TbF4mVUMwiH36uBMnJEhpxqtKvjPFMaBms2vNe2LYQjydRs2niy5pBsBjjxif5mxkEa3S27pc5epeYATPA9Xhgagz2TDzniZEYQgQ6uEGyKGJRQ2AX9qpTY7LtHxN8sUqn5SuZAMnM27iZ9bwwyjjBGepRmz1mfQfFvSV92exnJRjCrzcR5VPuViSAxDtwZFVzB8CVcA3M9ZFaUn8mhTe9U8wKFYEm1mH9cPjWHpwm5h4S2yvMVSw1"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":80208054},"value":{"accounts":null,"err":null,"logs":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"]}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.SimulateTransaction(
					context.TODO(),
					"5nxpoKAc5anKyiJuwj5f2SLxnruNHDjpFFz1TAw5VvpL1L4TbF4mVUMwiH36uBMnJEhpxqtKvjPFMaBms2vNe2LYQjydRs2niy5pBsBjjxif5mxkEa3S27pc5epeYATPA9Xhgagz2TDzniZEYQgQ6uEGyKGJRQ2AX9qpTY7LtHxN8sUqn5SuZAMnM27iZ9bwwyjjBGepRmz1mfQfFvSV92exnJRjCrzcR5VPuViSAxDtwZFVzB8CVcA3M9ZFaUn8mhTe9U8wKFYEm1mH9cPjWHpwm5h4S2yvMVSw1",
				)
			},
			ExpectedResponse: JsonRpcResponse[SimulateTransaction]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: SimulateTransaction{
					Context: Context{
						Slot: 80208054,
					},
					Value: SimulateTransactionValue{
						Logs: []string{"Program 11111111111111111111111111111111 invoke [1]", "Program 11111111111111111111111111111111 success"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"simulateTransaction", "params":["3vDU6xomZYLVZDefJLUEKXdzQLxRvJ8m1u31a4m6ynNDrPmFUUC9ogWH1yTnqaKm5SRcYbrp1xXExzhVKdCiv1KTkZdmZ7oNgTMSq4SN1nu1nL4hkZPSKGNGxXk6fViefXGiaHvzmC6mR2coVhvjs75eayuGyhomCfEUnKfUuQK99UC8pYJNenHTQQ4DX92sJmuiPoQHGDBVDwtVAkvLEfav89uSUxS1jbpfMPBs7fTNGiEhgMgYq5p4rsvPAYG6EZpDnK3VFjbBGBeUAdkCbjxLQGRJs3UEUNwjD"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":80207873},"value":{"accounts":null,"err":{"InstructionError":[0,{"Custom":1}]},"logs":["Program 11111111111111111111111111111111 invoke [1]","Transfer: insufficient lamports 109112817160, need 10000000000000","Program 11111111111111111111111111111111 failed: custom program error: 0x1"]}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.SimulateTransaction(
					context.TODO(),
					"3vDU6xomZYLVZDefJLUEKXdzQLxRvJ8m1u31a4m6ynNDrPmFUUC9ogWH1yTnqaKm5SRcYbrp1xXExzhVKdCiv1KTkZdmZ7oNgTMSq4SN1nu1nL4hkZPSKGNGxXk6fViefXGiaHvzmC6mR2coVhvjs75eayuGyhomCfEUnKfUuQK99UC8pYJNenHTQQ4DX92sJmuiPoQHGDBVDwtVAkvLEfav89uSUxS1jbpfMPBs7fTNGiEhgMgYq5p4rsvPAYG6EZpDnK3VFjbBGBeUAdkCbjxLQGRJs3UEUNwjD",
				)
			},
			ExpectedResponse: JsonRpcResponse[SimulateTransaction]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: SimulateTransaction{
					Context: Context{
						Slot: 80207873,
					},
					Value: SimulateTransactionValue{
						Err: map[string]any{
							"InstructionError": []any{
								0.,
								map[string]any{
									"Custom": 1.,
								},
							},
						},
						Logs: []string{"Program 11111111111111111111111111111111 invoke [1]", "Transfer: insufficient lamports 109112817160, need 10000000000000", "Program 11111111111111111111111111111111 failed: custom program error: 0x1"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"simulateTransaction", "params":["3vDU6xomZYLVZDefJLUEKXdzQLxRvJ8m1u31a4m6ynNDrPmFUUC9ogWH1yTnqaKm5SRcYbrp1xXExzhVKdCiv1KTkZdmZ7oNgTMSq4SN1nu1nL4hkZPSKGNGxXk6fViefXGiaHvzmC6mR2coVhvjs75eayuGyhomCfEUnKfUuQK99UC8pYJNenHTQQ4DX92sJmuiPoQHGDBVDwtVAkvLEfav89uSUxS1jbpfMPBs7fTNGiEhgMgYq5p4rsvPAYG6EZpDnK3VFjbBGBeUAdkCbjxLQGRJs3UEUNwjD"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":80208056},"value":{"accounts":null,"err":"BlockhashNotFound","logs":[]}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.SimulateTransaction(
					context.TODO(),
					"3vDU6xomZYLVZDefJLUEKXdzQLxRvJ8m1u31a4m6ynNDrPmFUUC9ogWH1yTnqaKm5SRcYbrp1xXExzhVKdCiv1KTkZdmZ7oNgTMSq4SN1nu1nL4hkZPSKGNGxXk6fViefXGiaHvzmC6mR2coVhvjs75eayuGyhomCfEUnKfUuQK99UC8pYJNenHTQQ4DX92sJmuiPoQHGDBVDwtVAkvLEfav89uSUxS1jbpfMPBs7fTNGiEhgMgYq5p4rsvPAYG6EZpDnK3VFjbBGBeUAdkCbjxLQGRJs3UEUNwjD",
				)
			},
			ExpectedResponse: JsonRpcResponse[SimulateTransaction]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: SimulateTransaction{
					Context: Context{
						Slot: 80208056,
					},
					Value: SimulateTransactionValue{
						Err:  "BlockhashNotFound",
						Logs: []string{},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"simulateTransaction", "params":["3vDU6xomZYLVZDefJLUEKXdzQLxRvJ8m1u31a4m6ynNDrPmFUUC9ogWH1yTnqaKm5SRcYbrp1xXExzhVKdCiv1KTkZdmZ7oNgTMSq4SN1nu1nL4hkZPSKGNGxXk6fViefXGiaHvzmC6mR2coVhvjs75eayuGyhomCfEUnKfUuQK99UC8pYJNenHTQQ4DX92sJmuiPoQHGDBVDwtVAkvLEfav89uSUxS1jbpfMPBs7fTNGiEhgMgYq5p4rsvPAYG6EZpDnK3VFjbBGBeUAdkCbjxLQGRJs3UEUNwjD", {"replaceRecentBlockhash": true}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":80208226},"value":{"accounts":null,"err":{"InstructionError":[0,{"Custom":1}]},"logs":["Program 11111111111111111111111111111111 invoke [1]","Transfer: insufficient lamports 109112817160, need 10000000000000","Program 11111111111111111111111111111111 failed: custom program error: 0x1"]}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.SimulateTransactionWithConfig(
					context.TODO(),
					"3vDU6xomZYLVZDefJLUEKXdzQLxRvJ8m1u31a4m6ynNDrPmFUUC9ogWH1yTnqaKm5SRcYbrp1xXExzhVKdCiv1KTkZdmZ7oNgTMSq4SN1nu1nL4hkZPSKGNGxXk6fViefXGiaHvzmC6mR2coVhvjs75eayuGyhomCfEUnKfUuQK99UC8pYJNenHTQQ4DX92sJmuiPoQHGDBVDwtVAkvLEfav89uSUxS1jbpfMPBs7fTNGiEhgMgYq5p4rsvPAYG6EZpDnK3VFjbBGBeUAdkCbjxLQGRJs3UEUNwjD",
					SimulateTransactionConfig{
						ReplaceRecentBlockhash: true,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[SimulateTransaction]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: SimulateTransaction{
					Context: Context{
						Slot: 80208226,
					},
					Value: SimulateTransactionValue{
						Err: map[string]any{
							"InstructionError": []any{
								0.,
								map[string]any{
									"Custom": 1.,
								},
							},
						},
						Logs: []string{"Program 11111111111111111111111111111111 invoke [1]", "Transfer: insufficient lamports 109112817160, need 10000000000000", "Program 11111111111111111111111111111111 failed: custom program error: 0x1"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"simulateTransaction", "params":["AbsEK+X7n9gAZ6giVYXDsyjtpHWaz8DA8IzsvBZSBGrtusZXpDaRD90P5HkZfc/JSzJTKfdmwG57tk6BvUkp4AYBAAEDBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4QllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA06LaMHqdmx+mBhPvnVRx5R2BobhuEZaAxJvCwM6btcUBAgIAAQwCAAAAAMqaOwAAAAA", {"encoding": "base64", "replaceRecentBlockhash": true, "accounts": {"encoding": "base64+zstd", "addresses": ["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"]}}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":80208978},"value":{"accounts":[{"data":["KLUv/SAAAQAA","base64+zstd"],"executable":false,"lamports":108112817160,"owner":"11111111111111111111111111111111","rentEpoch":185}],"err":null,"logs":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"]}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.SimulateTransactionWithConfig(
					context.TODO(),
					"AbsEK+X7n9gAZ6giVYXDsyjtpHWaz8DA8IzsvBZSBGrtusZXpDaRD90P5HkZfc/JSzJTKfdmwG57tk6BvUkp4AYBAAEDBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4QllkXXnxkMyGl7UZCoCewq9l7jdl60bzG3GRxOGzN3AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA06LaMHqdmx+mBhPvnVRx5R2BobhuEZaAxJvCwM6btcUBAgIAAQwCAAAAAMqaOwAAAAA",
					SimulateTransactionConfig{
						Encoding:               SimulateTransactionEncodingBase64,
						ReplaceRecentBlockhash: true,
						Accounts: &SimulateTransactionConfigAccounts{
							Encoding:  AccountEncodingBase64Zstd,
							Addresses: []string{"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"},
						},
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[SimulateTransaction]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: SimulateTransaction{
					Context: Context{
						Slot: 80208978,
					},
					Value: SimulateTransactionValue{
						Accounts: []*AccountInfo{
							{
								Owner:      "11111111111111111111111111111111",
								Lamports:   108112817160,
								Data:       []any{"KLUv/SAAAQAA", "base64+zstd"},
								Executable: false,
								RentEpoch:  185,
							},
						},
						Logs: []string{"Program 11111111111111111111111111111111 invoke [1]", "Program 11111111111111111111111111111111 success"},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"simulateTransaction","params":["Ab/yMEK7qNgGxaPMg2XaVnwwLMqnY8FTeJrA9qJ1nOBFX08BHycnp3/9WOxOY53+eZnbkT2/+6Mx7w+DsuVN8ggBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYEsGZsdWhvaw9ENEPFBEi4eBna4CphPQWWcgU4yARSnVAQEAAA==",{"encoding":"base64"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.5","slot":159776096},"value":{"accounts":null,"err":null,"logs":["Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units","Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success"],"returnData":{"data":["AQIDBAU=","base64"],"programId":"35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP"},"unitsConsumed":185}},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.SimulateTransactionWithConfig(
					context.TODO(),
					"Ab/yMEK7qNgGxaPMg2XaVnwwLMqnY8FTeJrA9qJ1nOBFX08BHycnp3/9WOxOY53+eZnbkT2/+6Mx7w+DsuVN8ggBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYEsGZsdWhvaw9ENEPFBEi4eBna4CphPQWWcgU4yARSnVAQEAAA==",
					SimulateTransactionConfig{
						Encoding: SimulateTransactionEncodingBase64,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[SimulateTransaction]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: SimulateTransaction{
					Context: Context{
						ApiVersion: "1.14.5",
						Slot:       159776096,
					},
					Value: SimulateTransactionValue{
						Accounts: nil,
						Logs: []string{
							"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]",
							"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units",
							"Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=",
							"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success",
						},
						ReturnData: &ReturnData{
							ProgramId: "35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP",
							Data:      []any{"AQIDBAU=", "base64"},
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
