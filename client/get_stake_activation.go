package client

import "errors"

type StakeActivationState string

const (
	StakeActivationStateActive       StakeActivationState = "active"
	StakeActivationStateInactive     StakeActivationState = "inactive"
	StakeActivationStateActivating   StakeActivationState = "activating"
	StakeActivationStateDeactivating StakeActivationState = "deactivating"
)

type GetStakeActivationConfig struct {
	Commitment Commitment `json:"commitment,omitempty"`
	Epoch      uint64     `json:"epoch,omitempty"`
}

type GetStakeActivationResponse struct {
	State    StakeActivationState `json:"state"`
	Active   uint64               `json:"active"`
	Inactive uint64               `json:"inactive"`
}

func (s *Client) GetStakeActivation(address string, cfg GetStakeActivationConfig) (GetStakeActivationResponse, error) {
	res := struct {
		GeneralResponse
		Result GetStakeActivationResponse `json:"result"`
	}{}

	err := s.request("getStakeActivation", []interface{}{address, cfg}, &res)
	if err != nil {
		return GetStakeActivationResponse{}, err
	}
	if res.Error != (ErrorResponse{}) {
		return GetStakeActivationResponse{}, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
