package tokenprog

import "github.com/portto/solana-go-sdk/common"

const MintAccountSize = 82

type MintAccount struct {
	MintAuthorityOption   uint32
	MintAuthority         common.PublicKey
	Supply                uint64
	Decimals              uint8
	IsInitialized         bool
	FreezeAuthorityOption uint32
	FreezeAuthority       common.PublicKey
}
