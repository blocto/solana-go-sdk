package address_lookup_table

import "errors"

var (
	ErrInvalidAccountOwner    = errors.New("invalid account owner")
	ErrInvalidAccountDataSize = errors.New("invalid account data size")
	ErrInvalidAccountData     = errors.New("invalid account data")
)
