package secp256k1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"fmt"

	secp256k1 "github.com/ipsn/go-secp256k1"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/bincode"
	"github.com/portto/solana-go-sdk/types"
	"golang.org/x/crypto/sha3"
)

const (
	OffsetsSerializedSize = 11
	DataStart             = OffsetsSerializedSize + 1
)

type SecpSignatureOffsets struct {
	SignatureOffsets           uint16
	SignatureInstructionIndex  uint8
	EthAddressOffset           uint16
	EthAddressInstructionIndex uint8
	MessageDataOffset          uint16
	MessageDataSize            uint16
	MessageInstructionIndex    uint8
}

type Secp256k1InstructionParam struct {
	Offsets []SecpSignatureOffsets
}

// Sign calculates an ECDSA signature.
//
// This function is susceptible to chosen plaintext attacks that can leak
// information about the private key that is used for signing. Callers must
// be aware that the given digest cannot be chosen by an adversery. Common
// solution is to hash any input before calculating the signature.
//
// The produced signature is in the [R || S || V] format where V is 0 or 1.
func sign(digestHash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error) {
	if len(digestHash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly %d bytes (%d)", 32, len(digestHash))
	}
	seckey := [32]byte{}
	prv.D.FillBytes(seckey[:])
	defer zeroBytes(seckey[:])
	return secp256k1.Sign(digestHash, seckey[:])
}

func zeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}

func pubkeyToAddress(pub *ecdsa.PublicKey) ([]byte, error) {
	pubBytes := elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)
	hasher := sha3.NewLegacyKeccak256()
	if _, err := hasher.Write(pubBytes[1:]); err != nil {
		return nil, err
	}
	return hasher.Sum([]byte{})[12:], nil
}

func NewSecp256k1Instruction(priv *ecdsa.PrivateKey, msg []byte, thisInstrIndex uint8) (types.Instruction, error) {
	return NewSecp256k1InstructionMultipleSigs([]*ecdsa.PrivateKey{priv}, [][]byte{msg}, thisInstrIndex)
}

func NewSecp256k1InstructionMultipleSigs(privs []*ecdsa.PrivateKey, msgs [][]byte, thisInstrIndex uint8) (types.Instruction, error) {
	if len(privs) != len(msgs) {
		return types.Instruction{}, errors.New(fmt.Sprintf("Provided a different number of keys and messages: %d private keys and %d messages.", len(privs), len(msgs)))
	}

	n := len(privs)

	instrData := []byte{}
	instrData = append(instrData, uint8(n)) // Count of Offsets structs
	// Append empty offsets structs to be filled it once we add our message data
	preData := [OffsetsSerializedSize]byte{}
	for i := 0; i < n; i++ {
		instrData = append(instrData, preData[:]...)
	}

	for i, msg := range msgs {
		priv := privs[i]
		hasher := sha3.NewLegacyKeccak256()
		_, err := hasher.Write(msg)
		if err != nil {
			return types.Instruction{}, err
		}
		hash := hasher.Sum([]byte{})
		sig, err := sign(hash[:], priv)
		if err != nil {
			return types.Instruction{}, err
		}
		addr, err := pubkeyToAddress(&priv.PublicKey)
		if err != nil {
			return types.Instruction{}, err
		}

		ethOffset := len(instrData)
		instrData = append(instrData, addr[:]...)
		sigOffset := len(instrData)
		instrData = append(instrData, sig...)
		msgOffset := len(instrData)
		instrData = append(instrData, msg[:]...)

		osets := SecpSignatureOffsets{
			SignatureOffsets:           uint16(sigOffset),
			SignatureInstructionIndex:  thisInstrIndex,
			EthAddressOffset:           uint16(ethOffset),
			EthAddressInstructionIndex: thisInstrIndex,
			MessageDataOffset:          uint16(msgOffset),
			MessageInstructionIndex:    thisInstrIndex,
			MessageDataSize:            uint16(len(msg)),
		}

		osetsBytes, err := bincode.SerializeData(osets)
		if err != nil {
			return types.Instruction{}, err
		}

		osetsStart := 1 + i*OffsetsSerializedSize
		osetsEnd := osetsStart + OffsetsSerializedSize
		copy(instrData[osetsStart:osetsEnd], osetsBytes)
	}

	return types.Instruction{
		ProgramID: common.Secp256k1ProgramID,
		Data:      instrData,
	}, nil
}
