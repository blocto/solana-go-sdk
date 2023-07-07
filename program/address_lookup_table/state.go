package address_lookup_table

import (
	"encoding/binary"

	"github.com/blocto/solana-go-sdk/common"
)

const LOOKUP_TABLE_MAX_ADDRESSES uint = 256

const LOOKUP_TABLE_META_SIZE uint = 56

func DeriveLookupTableAddress(authorityAddr common.PublicKey, recentBlockSlot uint64) (common.PublicKey, uint8) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, recentBlockSlot)
	pubkey, bump, _ := common.FindProgramAddress(
		[][]byte{
			authorityAddr.Bytes(),
			b[:],
		},
		common.AddressLookupTableProgramID,
	)
	return pubkey, uint8(bump)
}

type ProgramStateEnum uint32

const (
	ProgramStateUninitialized ProgramStateEnum = iota
	ProgramStateLookupTable
)

type AddressLookupTable struct {
	ProgramState               ProgramStateEnum
	DeactivationSlot           uint64
	LastExtendedSlot           uint64
	LastExtendedSlotStartIndex uint8
	Authority                  *common.PublicKey
	padding                    uint16
	Addresses                  []common.PublicKey
}

func DeserializeLookupTable(data []byte, accountOwner common.PublicKey) (AddressLookupTable, error) {
	if accountOwner != common.AddressLookupTableProgramID {
		return AddressLookupTable{}, ErrInvalidAccountOwner
	}
	if len(data) < 4 {
		return AddressLookupTable{}, ErrInvalidAccountDataSize
	}

	current := 0
	programState := ProgramStateEnum(binary.LittleEndian.Uint32(data[current : current+4]))
	current += 4

	switch programState {
	case ProgramStateUninitialized:
		return AddressLookupTable{}, nil
	case ProgramStateLookupTable:
		if uint(len(data)) < LOOKUP_TABLE_META_SIZE {
			return AddressLookupTable{}, ErrInvalidAccountDataSize
		}
		addressLookupTable := AddressLookupTable{
			ProgramState: programState,
		}

		addressLookupTable.DeactivationSlot = binary.LittleEndian.Uint64(data[current : current+8])
		current += 8

		addressLookupTable.LastExtendedSlot = binary.LittleEndian.Uint64(data[current : current+8])
		current += 8

		addressLookupTable.LastExtendedSlotStartIndex = data[current]
		current += 1

		some := bool(data[current] == 1)
		current += 1
		if some {
			pubkey := common.PublicKeyFromBytes(data[current : current+32])
			current += 32
			addressLookupTable.Authority = &pubkey
		}

		addressLookupTable.padding = binary.LittleEndian.Uint16(data[current : current+2])
		current += 2

		l := (len(data) - current) / 32
		addresses := make([]common.PublicKey, 0, l)
		for i := 0; i < l; i++ {
			addresses = append(addresses, common.PublicKeyFromBytes(data[current:current+32]))
			current += 32
		}
		addressLookupTable.Addresses = addresses

		return addressLookupTable, nil
	}

	return AddressLookupTable{}, ErrInvalidAccountData
}
