package system

import (
	"encoding/binary"
	"fmt"

	"github.com/blocto/solana-go-sdk/common"
)

const FeeCalculatorSize = 8

type FeeCalculator struct {
	LamportsPerSignature uint64
}

func FeeCalculatorDeserialize(data []byte) (FeeCalculator, error) {
	if len(data) < FeeCalculatorSize {
		return FeeCalculator{}, fmt.Errorf("fee calculator data size is not enough")
	}
	lamportsPerSignature := binary.LittleEndian.Uint64(data[:8])
	return FeeCalculator{
		LamportsPerSignature: lamportsPerSignature,
	}, nil
}

const NonceAccountSize = 80

type NonceAccount struct {
	Version          uint32
	State            uint32
	AuthorizedPubkey common.PublicKey
	Nonce            common.PublicKey
	FeeCalculator    FeeCalculator
}

func NonceAccountDeserialize(data []byte) (NonceAccount, error) {
	if len(data) < NonceAccountSize {
		return NonceAccount{}, fmt.Errorf("nonce account data size is not enough")
	}
	version := binary.LittleEndian.Uint32(data[:4])
	state := binary.LittleEndian.Uint32(data[4:8])
	authorizedPubkey := common.PublicKeyFromBytes(data[8:40])
	nonce := common.PublicKeyFromBytes(data[40:72])
	feeCalculator, err := FeeCalculatorDeserialize(data[72:])
	if err != nil {
		return NonceAccount{}, err
	}
	return NonceAccount{
		Version:          version,
		State:            state,
		AuthorizedPubkey: authorizedPubkey,
		Nonce:            nonce,
		FeeCalculator:    feeCalculator,
	}, nil
}
