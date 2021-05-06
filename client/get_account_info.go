package client

import (
	"context"
	"errors"
)

// GetAccountInfoConfig s
// encoding:
//		- "base58": is limited to Account data of less than 128 bytes
//		- "base64": will return base64 encoded data for Account data of any size
// 		- "base64+zstd": compresses the Account data using Zstandard and base64-encodes the result
// 		- "jsonParsed": encoding attempts to use program-specific state parsers to return more human-readable and explicit account state data. If "jsonParsed" is requested but a parser cannot be found, the field falls back to "base64" encoding, detectable when the data field is type <string>
// dataSlice: limit the returned account data using the provided offset: <usize> and length: <usize> fields;
// 			  only available for "base58", "base64" or "base64+zstd" encodings.
type GetAccountInfoConfig struct {
	Encoding  string                        `json:"encoding"`
	DataSlice GetAccountInfoConfigDataSlice `json:"dataSlice"`
}

// GetAccountInfoConfigDataSlice limit the returned account data using the provided offset: <usize> and length: <usize> fields; only available for "base58", "base64" or "base64+zstd" encodings.
type GetAccountInfoConfigDataSlice struct {
	Offset uint64 `json:"offset"`
	Length uint64 `json:"length"`
}

type GetAccountInfoResponse struct {
	Lamports  uint64      `json:"lamports"`
	Owner     string      `json:"owner"`
	Excutable bool        `json:"excutable"`
	RentEpoch uint64      `json:"rentEpoch"`
	Data      interface{} `json:"data"`
}

func (s *Client) GetAccountInfo(ctx context.Context, account string, cfg GetAccountInfoConfig) (GetAccountInfoResponse, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context                `json:"context"`
			Value   GetAccountInfoResponse `json:"value"`
		} `json:"result"`
	}{}
	err := s.request(ctx, "getAccountInfo", []interface{}{account, cfg}, &res)
	if err != nil {
		return GetAccountInfoResponse{}, err
	}
	if res.Error != (ErrorResponse{}) {
		return GetAccountInfoResponse{}, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
