package bytes_decoder

import (
	"encoding/binary"
	"fmt"
)

func GetUint64(curr *int, data []byte) (uint64, error) {
	if curr == nil {
		return 0, fmt.Errorf("index is nil")
	}
	if data == nil {
		return 0, fmt.Errorf("data is nil")
	}
	if len(data[*curr:]) < 8 {
		return 0, fmt.Errorf("insufficient data length")
	}

	v := binary.LittleEndian.Uint64(data[*curr : *curr+8])
	*curr += 8

	return v, nil
}

func GetBytes32(curr *int, data []byte) ([32]byte, error) {
	var v [32]byte

	if curr == nil {
		return v, fmt.Errorf("index is nil")
	}
	if data == nil {
		return v, fmt.Errorf("data is nil")
	}
	if len(data[*curr:]) < 32 {
		return v, fmt.Errorf("insufficient data length")
	}

	copy(v[:], data[*curr:*curr+32])
	*curr += 32

	return v, nil
}
