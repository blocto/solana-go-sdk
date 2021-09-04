package types

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/portto/solana-go-sdk/common"
)

type Account struct {
	PublicKey  common.PublicKey
	PrivateKey ed25519.PrivateKey
}

func NewAccount() Account {
	_, X, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(fmt.Sprintf("ed25519 genkey error: %s", err))
	}
	return AccountFromPrivateKeyBytes(X)
}

func AccountFromPrivateKeyBytes(privateKey []byte) Account {
	sk := ed25519.PrivateKey(privateKey)
	return Account{
		PublicKey:  common.PublicKeyFromBytes(sk.Public().(ed25519.PublicKey)),
		PrivateKey: sk,
	}
}

func (a *Account) UnmarshalText(b []byte) error {
	// hex.Decode(dst []byte, src []byte)
	key, err := hex.DecodeString(string(b))
	if err != nil {
		return fmt.Errorf("decode private key: %w", err)
	}

	if len(key) != ed25519.PrivateKeySize {
		return errors.New("invalid private key size")
	}

	*a = AccountFromPrivateKeyBytes(key)

	return nil
}
