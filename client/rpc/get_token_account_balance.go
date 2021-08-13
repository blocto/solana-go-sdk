package rpc

import "context"

type GetTokenAccountBalance struct {
	Amount         string `json:"amount"`
	Decimals       int64  `json:"decimals"`
	UIAmountString string `json:"uiAmountString"`
}

// GetTokenAccountBalance returns the token balance of an SPL Token account.
func (s *RpcClient) GetTokenAccountBalance(ctx context.Context, base58Addr string, commitment Commitment) (GetTokenAccountBalance, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context                `json:"context"`
			Value   GetTokenAccountBalance `json:"value"`
		} `json:"result"`
	}{}
	err := s.request(ctx, "getTokenAccountBalance", []interface{}{base58Addr, map[string]interface{}{"commitment": commitment}}, &res)
	if err != nil {
		return GetTokenAccountBalance{}, err
	}
	return res.Result.Value, nil
}
