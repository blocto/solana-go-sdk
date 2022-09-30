package sysvar

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/pkg/bytes_decoder"
)

type SlotHash struct {
	Slot uint64
	Hash [32]byte
}

type SlotHashes []SlotHash

func DeserializeSlotHashes(data []byte, owner common.PublicKey) (SlotHashes, error) {
	if owner != common.SysVarPubkey {
		return SlotHashes{}, ErrInvalidAccountOwner
	}

	current := 0
	len, err := bytes_decoder.GetUint64(&current, data)
	if err != nil {
		return SlotHashes{}, err
	}

	v := make([]SlotHash, 0, len)
	for i := uint64(0); i < len; i++ {
		slot, err := bytes_decoder.GetUint64(&current, data)
		if err != nil {
			return SlotHashes{}, err
		}
		hash, err := bytes_decoder.GetBytes32(&current, data)
		if err != nil {
			return SlotHashes{}, err
		}

		v = append(v, SlotHash{Slot: slot, Hash: hash})
	}
	return v, nil
}
