package types

import (
	"github.com/portto/solana-go-sdk/common"
)

type AccountMeta struct {
	PubKey     common.PublicKey
	IsSigner   bool
	IsWritable bool
}
