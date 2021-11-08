package secp256k1

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/bincode"
	"github.com/portto/solana-go-sdk/types"
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
		hash := crypto.Keccak256(msg)
		sig, err := crypto.Sign(hash, priv)
		if err != nil {
			return types.Instruction{}, err
		}
		addr := crypto.PubkeyToAddress(priv.PublicKey)

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
