package client

import (
	"context"
	"errors"
)

type GetProgramAccountsConfig struct {
	Filters []interface{} `json:"filters,omitempty"`
}

type GetProgramAccountsConfigDataSize struct {
	DataSize int `json:"dataSize"`
}

type GetProgramAccountsConfigMemCmpFilter struct {
	MemCmp GetProgramAccountsConfigMemCmp `json:"memcmp"`
}

type GetProgramAccountsConfigMemCmp struct {
	Offset int    `json:"offset"`
	Bytes  string `json:"bytes"`
}

type GetProgramAccountsResponse struct {
	Account struct {
		Data       string `json:"data"`
		Executable bool   `json:"executable"`
		Lamports   int    `json:"lamports"`
		Owner      string `json:"owner"`
		Rentepoch  int    `json:"rentEpoch"`
	} `json:"account"`
	Pubkey string `json:"pubkey"`
}

const (
	GetProgramAccountDataSizeAllTokenAccounts = 165 // The dataSize 165 filter selects all Token Accounts
	GetProgramAccountOffsetMintAddress        = 0
	GetProgramAccountOffsetWalletAddress      = 32
)

// GetProgramAccounts get token sub accounts of specified program address
// for a public token, the programAddress is generally common.TokenProgramID
// filters can be used to filter mint address, wallet address or program address
// simply use GetProgramAccountsConfigMemCmpFilter to filter them
// if you don't pass any filter, a dataSize 165 filter will be added to get all token accounts
func (s *Client) GetProgramAccounts(
	ctx context.Context, programAddress string, filters ...interface{}) ([]GetProgramAccountsResponse, error) {

	res := struct {
		GeneralResponse
		Result []GetProgramAccountsResponse `json:"result"`
	}{}

	cfg := GetProgramAccountsConfig{}
	if len(filters) == 0 {
		cfg.Filters = append(cfg.Filters,
			GetProgramAccountsConfigDataSize{DataSize: GetProgramAccountDataSizeAllTokenAccounts})
	}
	cfg.Filters = append(cfg.Filters, filters...)

	err := s.request(ctx, "getProgramAccounts",
		[]interface{}{
			programAddress,
			cfg,
		}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != (ErrorResponse{}) {
		return nil, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
