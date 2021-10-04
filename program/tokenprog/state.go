package tokenprog

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/portto/solana-go-sdk/common"
)

var (
	Some = []byte{1, 0, 0, 0}
	None = []byte{0, 0, 0, 0}
)

const MultisigAccountSize uint64 = 355

const MintAccountSize = 82

type MintAccount struct {
	MintAuthority   *common.PublicKey
	Supply          uint64
	Decimals        uint8
	IsInitialized   bool
	FreezeAuthority *common.PublicKey
}

func MintAccountFromData(data []byte) (MintAccount, error) {
	if len(data) != MintAccountSize {
		return MintAccount{}, fmt.Errorf("data length not match")
	}

	var mint *common.PublicKey
	if bytes.Equal(data[:4], Some) {
		key := common.PublicKeyFromBytes(data[4:36])
		mint = &key
	}

	supply := binary.LittleEndian.Uint64(data[36:44])

	decimals := uint8(data[44])

	isInitialized := data[45] == 1

	var freezeAuthority *common.PublicKey
	if bytes.Equal(data[46:50], Some) {
		key := common.PublicKeyFromBytes(data[50:82])
		freezeAuthority = &key
	}

	return MintAccount{
		MintAuthority:   mint,
		Supply:          supply,
		Decimals:        decimals,
		IsInitialized:   isInitialized,
		FreezeAuthority: freezeAuthority,
	}, nil
}

const TokenAccountSize = 165

type TokenAccountState uint8

const (
	TokenAccountStateUninitialized TokenAccountState = iota
	TokenAccountStateInitialized
	TokenAccountFrozen
)

// TokenAccount is token program account
type TokenAccount struct {
	Mint     common.PublicKey
	Owner    common.PublicKey
	Amount   uint64
	Delegate *common.PublicKey
	State    TokenAccountState
	// if is wrapped SOL, IsNative is the rent-exempt value
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
