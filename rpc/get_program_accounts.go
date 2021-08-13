package client

import (
	"context"
	"encoding/json"
	"errors"
)

type GetProgramAccountsConfig struct {
	Encoding   Encoding                          `json:"encoding"`
	Commitment Commitment                        `json:"commitment"`
	DataSlice  GetProgramAccountsConfigDataSlice `json:"dataSlice"`
	// filter should be either GetProgramAccountsConfigFilterMemCmp or GetProgramAccountsConfigFilterDataSize
	Filters []interface{} `json:"filters"`
}

type getProgramAccountsConfig struct {
	Encoding    *Encoding                          `json:"encoding,omitempty"`
	Commitment  *Commitment                        `json:"commitment,omitempty"`
	DataSlice   *GetProgramAccountsConfigDataSlice `json:"dataSlice,omitempty"`
	Filters     *[]interface{}                     `json:"filters,omitempty"`
	WithContext bool                               `json:"withContext"`
}

func (cfg GetProgramAccountsConfig) MarshalJSON() ([]byte, error) {
	config := getProgramAccountsConfig{}

	if cfg.Encoding != "" {
		config.Encoding = &cfg.Encoding
	}

	if cfg.Commitment != "" {
		config.Commitment = &cfg.Commitment
	}

	if cfg.DataSlice != (GetProgramAccountsConfigDataSlice{}) {
		config.DataSlice = &cfg.DataSlice
	}

	if len(cfg.Filters) != 0 {
		config.Filters = &cfg.Filters
	}

	config.WithContext = true

	return json.Marshal(config)
}

type GetProgramAccountsConfigFilterMemCmp struct {
	Offset uint64 `json:"offset"`
	Bytes  string `json:"bytes"`
}

type GetProgramAccountsConfigFilterDataSize struct {
	DataSize uint64 `json:"dataSize"`
}

type GetProgramAccountsConfigDataSlice struct {
	Offset uint64 `json:"offset"`
	Length uint64 `json:"length"`
}

type GetProgramAccountsAccount struct {
	Data       interface{} `json:"data"`
	Executable bool        `json:"executable"`
	Lamports   uint64      `json:"lamports"`
	Owner      string      `json:"owner"`
	Rentepoch  uint64      `json:"rentEpoch"`
}

type GetProgramAccounts struct {
	Account GetProgramAccountsAccount `json:"account"`
	Pubkey  string                    `json:"pubkey"`
}

// GetProgramAccounts returns all accounts owned by the provided program pubkey
func (s *RpcClient) GetProgramAccounts(ctx context.Context, base58Addr string, cfg GetProgramAccountsConfig) ([]GetProgramAccounts, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context              `json:"context"`
			Value   []GetProgramAccounts `json:"value"`
		} `json:"result"`
	}{}

	err := s.request(ctx, "getProgramAccounts", []interface{}{base58Addr, cfg}, &res)
	if err != nil {
		return []GetProgramAccounts{}, err
	}
	if res.Error != (ErrorResponse{}) {
		return []GetProgramAccounts{}, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
