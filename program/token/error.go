package token

import "errors"

var (
	ErrInvalidAccountOwner    = errors.New("invalid account owner")
	ErrInvalidAccountDataSize = errors.New("invalid account data size")
)
