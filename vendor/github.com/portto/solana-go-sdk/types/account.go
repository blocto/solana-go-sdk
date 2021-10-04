package types

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/mr-tron/base58"
	"github.com/portto/solana-go-sdk/common"
)

var (
	ErrAccountFailedToBase58Decode     = errors.New("failed to base58 decode")
	ErrAccountFailedToHexDecode        = errors.New("failed to hex decode")
	ErrAccountPrivateKeyLengthMismatch = errors.New("key length mismatch")
)

type Account struct {
	PublicKey  common.PublicKey
	PrivateKey ed25519.PrivateKey
}

func NewAccount() Account {
	_, X, _ := ed25519.GenerateKey(nil)
	return AccountFromPrivateKeyBytes(X)
}

// AccountFromPrivateKeyBytes DEPRECATED, use AccountFromBytes instead. it will be removed in v2.0.0
func AccountFromPrivateKeyBytes(privateKey []byte) Account {
	sk := ed25519.PrivateKey(privateKey)
	return Account{
		PublicKey:  common.PublicKeyFromBytes(sk.Public().(ed25519.PublicKey)),
		PrivateKey: sk,
	}
}

// AccountFromBytes generate a account by bytes private key
func AccountFromBytes(key []byte) (Account, error) {
	if len(key) != ed25519.PrivateKeySize {
		return Account{}, fmt.Errorf("%w, expected: %v, got: %v", ErrAccountPrivateKeyLengthMismatch, ed25519.PrivateKeySize, len(key))
	}
	// TODO check is an ed25519 key
	priKey := ed25519.PrivateKey(key)
	return Account{
		PublicKey:  common.PublicKeyFromBytes(priKey.Public().(ed25519.PublicKey)),
		PrivateKey: priKey,
	}, nil
}

// AccountFromBase58 generate a account by base58 private key
func AccountFromBase58(key string) (Account, error) {
	b, err := base58.Decode(key)
	if err != nil {
		return Account{}, fmt.Errorf("%w, err: %v", ErrAccountFailedToBase58Decode, err)
	}
	return AccountFromBytes(b)
}

// AccountFromHex generate a account by hex private key
func AccountFromHex(key string) (Account, error) {
	b, err := hex.DecodeString(key)
	if err != nil {
		return Account{}, fmt.Errorf("%w, err: %v", ErrAccountFailedToHexDecode, err)
	}
	return AccountFromBytes(b)
}

func (a Account) Sign(message []byte) []byte {
	return ed25519.Sign(a.PrivateKey, message)
}
