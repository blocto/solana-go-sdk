package pointer

import "github.com/portto/solana-go-sdk/common"

func Bool(v bool) *bool {
	return &v
}

func Uint8(v uint8) *uint8 {
	return &v
}

func Uint64(v uint64) *uint64 {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Pubkey(v common.PublicKey) *common.PublicKey {
	return &v
}
