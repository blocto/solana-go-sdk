package tokenprog

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/portto/solana-go-sdk/common"
)

const MultisigAccountSize uint64 = 355

const MintAccountSize = 82

type MintAccount struct {
	// MintAuthorityOption   uint32
	MintAuthority *common.PublicKey
	Supply        uint64
	Decimals      uint8
	IsInitialized bool
	// FreezeAuthorityOption uint32
	FreezeAuthority *common.PublicKey
}

const (
	mintMintAuthorityOptionOffset   = 0
	mintMintAuthorityOffset         = mintMintAuthorityOptionOffset + 4
	mintSupplyOffset                = mintMintAuthorityOffset + 32
	mintDecimalsOffset              = mintSupplyOffset + 8
	mintIsInitializedOffset         = mintDecimalsOffset + 1
	mintFreezeAuthorityOptionOffset = mintIsInitializedOffset + 1
	mintFreezeAuthorityOffset       = mintFreezeAuthorityOptionOffset + 4
)

func isSome(option []byte) bool {
	return bytes.Equal(option, Some)
}

func MintAccountFromData(data []byte) (*MintAccount, error) {
	if len(data) != MintAccountSize {
		return nil, fmt.Errorf("mint account data length mismatch")
	}

	var mint MintAccount

	mintAuthorityOption := data[0:4]
	if isSome(mintAuthorityOption) {
		key := common.PublicKeyFromBytes(data[mintMintAuthorityOffset : mintSupplyOffset+32])
		mint.MintAuthority = &key
	}

	mint.Supply = binary.LittleEndian.Uint64(data[mintSupplyOffset : mintSupplyOffset+8])
	mint.Decimals = uint8(data[mintDecimalsOffset])
	mint.IsInitialized = data[mintIsInitializedOffset] == 1

	if isSome(data[mintFreezeAuthorityOptionOffset:mintFreezeAuthorityOptionOffset+4]) {
		key := common.PublicKeyFromBytes(data[mintFreezeAuthorityOffset : mintFreezeAuthorityOffset+32])
		mint.FreezeAuthority = &key
	}

	return &mint, nil
}

const TokenAccountSize = 165

type TokenAccountState uint8

const (
	TokenAccountStateUninitialized TokenAccountState = iota
	TokenAccountStateInitialized
	TokenAccountFrozen
)

var (
	Some = []byte{1, 0, 0, 0}
	None = []byte{0, 0, 0, 0}
)

// TokenAccount is token program account
type TokenAccount struct {
	Mint            common.PublicKey
	Owner           common.PublicKey
	Amount          uint64
	Delegate        *common.PublicKey
	State           TokenAccountState
	IsNative        *uint64
	DelegatedAmount uint64
	CloseAuthority  *common.PublicKey
}

func TokenAccountFromData(data []byte) (*TokenAccount, error) {
	if len(data) != TokenAccountSize {
		return nil, fmt.Errorf("data length not match")
	}

	mint := common.PublicKeyFromBytes(data[:32])

	owner := common.PublicKeyFromBytes(data[32:64])

	amount := binary.LittleEndian.Uint64(data[64:72])

	var delegate *common.PublicKey
	if bytes.Equal(data[72:76], Some) {
		key := common.PublicKeyFromBytes(data[76:108])
		delegate = &key
	}

	state := TokenAccountState(data[108])

	var isNative *uint64
	if bytes.Equal(data[109:113], Some) {
		num := binary.LittleEndian.Uint64(data[113:121])
		isNative = &num
	}

	delegatedAmount := binary.LittleEndian.Uint64(data[121:129])

	var closeAuthority *common.PublicKey
	if bytes.Equal(data[129:133], Some) {
		key := common.PublicKeyFromBytes(data[133:165])
		closeAuthority = &key
	}

	return &TokenAccount{
		Mint:            mint,
		Owner:           owner,
		Amount:          amount,
		Delegate:        delegate,
		State:           state,
		IsNative:        isNative,
		DelegatedAmount: delegatedAmount,
		CloseAuthority:  closeAuthority,
	}, nil
}
