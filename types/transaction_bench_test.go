package types

import (
	"testing"

	"github.com/blocto/solana-go-sdk/common"
)

func BenchmarkSerializeTransaction(b *testing.B) {
	tx := Transaction{
		Signatures: []Signature{[]byte{189, 98, 67, 19, 102, 99, 124, 234, 70, 209, 28, 10, 33, 66, 167, 162, 222, 122, 16, 68, 248, 129, 46, 111, 221, 255, 40, 40, 236, 84, 233, 213, 234, 185, 235, 222, 155, 204, 139, 164, 184, 155, 32, 54, 151, 73, 235, 65, 200, 76, 127, 111, 244, 72, 183, 208, 21, 247, 114, 176, 181, 21, 77, 8}},
		Message: Message{
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
					Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				},
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tx.Serialize()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDeserializeTransaction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := TransactionDeserialize([]byte{1, 189, 98, 67, 19, 102, 99, 124, 234, 70, 209, 28, 10, 33, 66, 167, 162, 222, 122, 16, 68, 248, 129, 46, 111, 221, 255, 40, 40, 236, 84, 233, 213, 234, 185, 235, 222, 155, 204, 139, 164, 184, 155, 32, 54, 151, 73, 235, 65, 200, 76, 127, 111, 244, 72, 183, 208, 21, 247, 114, 176, 181, 21, 77, 8, 1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0})
		if err != nil {
			b.Error(err)
		}
	}
}
