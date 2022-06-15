package pointer

import (
	"github.com/OldSmokeGun/solana-go-sdk/common"
)

func Bool(v bool) *bool {
	return &v
}

func Uint8(v uint8) *uint8 {
	return &v
}

func Uint16(v uint16) *uint16 {
	return &v
}

func Uint32(v uint32) *uint32 {
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

func String(v string) *string {
	return &v
}
