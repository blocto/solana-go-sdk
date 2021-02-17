package sysprog

import (
	"encoding/binary"
	"fmt"
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
