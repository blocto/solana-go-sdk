package pointer

import (
	"github.com/portto/solana-go-sdk/common"
)

func Get[T any](v T) *T {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Pubkey(v common.PublicKey) *common.PublicKey {
	return &v
}

func String(v string) *string {
	return &v
}
