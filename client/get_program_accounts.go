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
	MemCmp struct {
		Offset int    `json:"offset"`
		Bytes  string `json:"bytes"`
	} `json:"memcmp"`
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
	getProgramAccountDataSizeAllTokenAccounts = 165 // The dataSize 165 filter selects all Token Accounts
	getProgramAccountOffsetMintAddress        = 0
	getProgramAccountOffsetWalletAddress      = 32
)

// GetProgramAccounts get all token sub accounts of specified program address
// for a public token, the programAddress is generally common.TokenProgramID
// mintAddress and walletAddress can be empty
// if you fill them, they will be used as filters
func (s *Client) GetProgramAccounts(
	ctx context.Context, programAddress, mintAddress, walletAddress string) ([]GetProgramAccountsResponse, error) {

	res := struct {
		GeneralResponse
		Result []GetProgramAccountsResponse `json:"result"`
	}{}

	cfg := GetProgramAccountsConfig{}
	cfg.Filters = append(cfg.Filters,
		GetProgramAccountsConfigDataSize{DataSize: getProgramAccountDataSizeAllTokenAccounts})
	if mintAddress != "" { // this filter is based on the mint address within each token account
		cfg.Filters = append(cfg.Filters, GetProgramAccountsConfigMemCmpFilter{
			struct {
				Offset int    `json:"offset"`
				Bytes  string `json:"bytes"`
			}{
				Offset: getProgramAccountOffsetMintAddress,
				Bytes:  mintAddress,
			},
		})
	}
	if walletAddress != "" { // this filter is based on the owner address within each token account
		cfg.Filters = append(cfg.Filters, GetProgramAccountsConfigMemCmpFilter{
			struct {
				Offset int    `json:"offset"`
				Bytes  string `json:"bytes"`
			}{
				Offset: getProgramAccountOffsetWalletAddress,
				Bytes:  walletAddress,
			},
		})
	}

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
