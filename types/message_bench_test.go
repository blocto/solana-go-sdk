package types

import (
	"testing"

	"github.com/portto/solana-go-sdk/common"
)

func BenchmarkSerializeLegacyMessage(b *testing.B) {
	message := Message{
		Version: MessageVersionLegacy,
		Header: MessageHeader{
			NumRequireSignatures:        1,
			NumReadonlySignedAccounts:   0,
			NumReadonlyUnsignedAccounts: 1,
		},
		Accounts: []common.PublicKey{
			common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
			common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"),
			common.SystemProgramID,
		},
		RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
		Instructions: []CompiledInstruction{
			{
				ProgramIDIndex: 2,
				Accounts:       []int{0, 1},
				Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := message.Serialize()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDeserializeLegacyMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MessageDeserialize([]byte{1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0})
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSerializeV0Message(b *testing.B) {
	message := Message{
		Version: MessageVersionV0,
		Header: MessageHeader{
			NumRequireSignatures:        1,
			NumReadonlySignedAccounts:   0,
			NumReadonlyUnsignedAccounts: 1,
		},
		Accounts: []common.PublicKey{
			common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
			common.SystemProgramID,
		},
		RecentBlockHash: "5EvWPqKeYfN2P7SAQZ2TLnXhV3Ltjn6qEhK1F279dUUW",
		Instructions: []CompiledInstruction{
			{
				ProgramIDIndex: 1,
				Accounts:       []int{0, 2},
				Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		AddressLookupTable: &CompiledAddressLookupTable{
			AccountKey:      common.PublicKeyFromString("HEhDGuxaxGr9LuNtBdvbX2uggyAKoxYgHFaAiqxVu8UY"),
			WritableIndexes: []uint8{1},
			ReadonlyIndexes: []uint8{},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := message.Serialize()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDeserializeV0Message(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MessageDeserialize([]byte{128, 1, 0, 1, 2, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 62, 255, 204, 109, 44, 223, 1, 225, 41, 92, 205, 204, 199, 90, 32, 104, 6, 123, 211, 72, 233, 131, 88, 65, 115, 38, 138, 217, 189, 202, 86, 39, 1, 1, 2, 0, 2, 12, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 241, 61, 2, 62, 211, 181, 33, 219, 74, 147, 127, 38, 231, 159, 99, 194, 103, 129, 201, 15, 51, 106, 114, 199, 122, 142, 121, 87, 112, 78, 138, 249, 1, 1, 0})
		if err != nil {
			b.Error(err)
		}
	}
}
