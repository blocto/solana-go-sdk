package secp256k1

import (
	"fmt"

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

func zeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}

func NewSecp256k1Instruction(msgs [][]byte, sigs [][]byte, addrs [][]byte, thisInstrIndex uint8) (types.Instruction, error) {
	if len(msgs) != len(sigs) || len(sigs) != len(addrs) {
		return types.Instruction{}, fmt.Errorf("Provided a different number of keys, messages, or signatures")
	}
	n := len(msgs)
	instrData := []byte{}
	instrData = append(instrData, uint8(n)) // Count of Offsets structs
	// Append empty offsets structs to be filled it once we add our message data
	preData := [OffsetsSerializedSize]byte{}
	for i := 0; i < n; i++ {
		instrData = append(instrData, preData[:]...)
	}

	for i, msg := range msgs {
		sig := sigs[i]
		addr := addrs[i]

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
