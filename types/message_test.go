package types

import (
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/stretchr/testify/assert"
)

func TestMessage_Serialize(t *testing.T) {
	type fields struct {
		Version            MessageVersion
		Header             MessageHeader
		Accounts           []common.PublicKey
		RecentBlockHash    string
		Instructions       []CompiledInstruction
		AddressLookupTable *CompiledAddressLookupTable
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
		err    error
	}{
		{
			fields: fields{
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
			want: []byte{1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			err:  nil,
		},
		{
			fields: fields{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					common.SystemProgramID,
				},
				RecentBlockHash: "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			want: []byte{128, 1, 0, 1, 3, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185, 29, 11, 113, 4, 101, 239, 39, 167, 201, 112, 156, 239, 236, 36, 251, 140, 76, 199, 150, 228, 218, 214, 20, 123, 180, 181, 103, 160, 71, 251, 237, 123, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 131, 118, 36, 248, 169, 123, 97, 98, 215, 133, 18, 92, 220, 162, 163, 79, 201, 66, 96, 112, 57, 224, 101, 105, 255, 83, 217, 144, 233, 242, 195, 102, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			err:  nil,
		},
		{
			fields: fields{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					common.SystemProgramID,
				},
				RecentBlockHash: "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				AddressLookupTable: &CompiledAddressLookupTable{
					AccountKey:      common.PublicKeyFromString("HEhDGuxaxGr9LuNtBdvbX2uggyAKoxYgHFaAiqxVu8UY"),
					WritableIndexes: []uint8{},
					ReadonlyIndexes: []uint8{},
				},
			},
			want: []byte{128, 1, 0, 1, 3, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185, 29, 11, 113, 4, 101, 239, 39, 167, 201, 112, 156, 239, 236, 36, 251, 140, 76, 199, 150, 228, 218, 214, 20, 123, 180, 181, 103, 160, 71, 251, 237, 123, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 131, 118, 36, 248, 169, 123, 97, 98, 215, 133, 18, 92, 220, 162, 163, 79, 201, 66, 96, 112, 57, 224, 101, 105, 255, 83, 217, 144, 233, 242, 195, 102, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			err:  nil,
		},
		{
			fields: fields{
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
			},
			want: []byte{128, 1, 0, 1, 2, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 62, 255, 204, 109, 44, 223, 1, 225, 41, 92, 205, 204, 199, 90, 32, 104, 6, 123, 211, 72, 233, 131, 88, 65, 115, 38, 138, 217, 189, 202, 86, 39, 1, 1, 2, 0, 2, 12, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 241, 61, 2, 62, 211, 181, 33, 219, 74, 147, 127, 38, 231, 159, 99, 194, 103, 129, 201, 15, 51, 106, 114, 199, 122, 142, 121, 87, 112, 78, 138, 249, 1, 1, 0},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Version:            tt.fields.Version,
				Header:             tt.fields.Header,
				Accounts:           tt.fields.Accounts,
				RecentBlockHash:    tt.fields.RecentBlockHash,
				Instructions:       tt.fields.Instructions,
				AddressLookupTable: tt.fields.AddressLookupTable,
			}
			got, err := m.Serialize()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestMessage_DecompileInstructions(t *testing.T) {
	type fields struct {
		Version            MessageVersion
		Header             MessageHeader
		Accounts           []common.PublicKey
		RecentBlockHash    string
		Instructions       []CompiledInstruction
		AddressLookupTable *CompiledAddressLookupTable
	}
	tests := []struct {
		name   string
		fields fields
		want   []Instruction
		panic  string
	}{
		{
			fields: fields{
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
			want: []Instruction{
				{
					Accounts: []AccountMeta{
						{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: true},
						{PubKey: common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"), IsSigner: false, IsWritable: true},
					},
					ProgramID: common.SystemProgramID,
					Data:      []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				},
			},
		},
		{
			fields: fields{
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
						Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			want: []Instruction{
				{
					Accounts: []AccountMeta{
						{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: true},
						{PubKey: common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"), IsSigner: false, IsWritable: true},
					},
					ProgramID: common.SystemProgramID,
					Data:      []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				},
			},
		},
		{
			fields: fields{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					common.SystemProgramID,
				},
				RecentBlockHash: "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			panic: "hasn't supported",
		},
		{
			fields: fields{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					common.SystemProgramID,
				},
				RecentBlockHash: "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				AddressLookupTable: &CompiledAddressLookupTable{
					AccountKey:      common.PublicKeyFromString("HEhDGuxaxGr9LuNtBdvbX2uggyAKoxYgHFaAiqxVu8UY"),
					WritableIndexes: []uint8{},
					ReadonlyIndexes: []uint8{},
				},
			},
			panic: "hasn't supported",
		},
		{
			fields: fields{
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
			},
			panic: "hasn't supported",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Version:            tt.fields.Version,
				Header:             tt.fields.Header,
				Accounts:           tt.fields.Accounts,
				RecentBlockHash:    tt.fields.RecentBlockHash,
				Instructions:       tt.fields.Instructions,
				AddressLookupTable: tt.fields.AddressLookupTable,
			}
			if len(tt.panic) == 0 {
				assert.Equal(t, tt.want, m.DecompileInstructions())
			} else {
				assert.PanicsWithValue(t, tt.panic, func() {
					m.DecompileInstructions()
				})
			}
		})
	}
}

func TestNewMessage(t *testing.T) {
	type args struct {
		param NewMessageParam
	}
	tests := []struct {
		name string
		args args
		want Message
	}{
		{
			args: args{
				NewMessageParam{
					FeePayer: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Instructions: []Instruction{
						{
							ProgramID: common.SystemProgramID,
							Accounts: []AccountMeta{
								{
									PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
									IsSigner:   true,
									IsWritable: true,
								},
								{
									PubKey:     common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"),
									IsSigner:   false,
									IsWritable: true,
								},
							},
							Data: []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
						},
					},
					RecentBlockhash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
				},
			},
			want: Message{
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
						Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					},
				},
			},
		},
		{
			args: args{
				NewMessageParam{
					FeePayer: common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					Instructions: []Instruction{
						{
							ProgramID: common.SystemProgramID,
							Accounts: []AccountMeta{
								{
									PubKey:     common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
									IsSigner:   true,
									IsWritable: true,
								},
								{
									PubKey:     common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
									IsSigner:   false,
									IsWritable: true,
								},
							},
							Data: []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
						},
					},
					RecentBlockhash: "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
				},
			},
			want: Message{
				Version: MessageVersionLegacy,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					common.SystemProgramID,
				},
				RecentBlockHash: "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
		},
		{
			args: args{
				NewMessageParam{
					FeePayer: common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					Instructions: []Instruction{
						{
							ProgramID: common.SystemProgramID,
							Accounts: []AccountMeta{
								{
									PubKey:     common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
									IsSigner:   true,
									IsWritable: true,
								},
								{
									PubKey:     common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
									IsSigner:   false,
									IsWritable: true,
								},
							},
							Data: []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
						},
					},
					RecentBlockhash:    "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
					AddressLookupTable: common.PublicKeyFromString("HEhDGuxaxGr9LuNtBdvbX2uggyAKoxYgHFaAiqxVu8UY"),
				},
			},
			want: Message{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					common.SystemProgramID,
				},
				RecentBlockHash: "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
		},
		{
			args: args{
				NewMessageParam{
					FeePayer: common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					Instructions: []Instruction{
						{
							ProgramID: common.SystemProgramID,
							Accounts: []AccountMeta{
								{
									PubKey:     common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
									IsSigner:   true,
									IsWritable: true,
								},
								{
									PubKey:     common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
									IsSigner:   false,
									IsWritable: true,
								},
							},
							Data: []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
						},
					},
					RecentBlockhash:    "5EvWPqKeYfN2P7SAQZ2TLnXhV3Ltjn6qEhK1F279dUUW",
					AddressLookupTable: common.PublicKeyFromString("HEhDGuxaxGr9LuNtBdvbX2uggyAKoxYgHFaAiqxVu8UY"),
					AddressLookupTableAddresses: []common.PublicKey{
						common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
						common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					},
				},
			},
			want: Message{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewMessage(tt.args.param), tt.want)
		})
	}
}

func TestMessageDeserialize(t *testing.T) {
	type args struct {
		messageData []byte
	}
	tests := []struct {
		name string
		args args
		want Message
		err  error
	}{
		{
			args: args{messageData: []byte{1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
			want: Message{
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
						Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			err: nil,
		},
		{
			args: args{messageData: []byte{128, 1, 0, 1, 3, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185, 29, 11, 113, 4, 101, 239, 39, 167, 201, 112, 156, 239, 236, 36, 251, 140, 76, 199, 150, 228, 218, 214, 20, 123, 180, 181, 103, 160, 71, 251, 237, 123, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 131, 118, 36, 248, 169, 123, 97, 98, 215, 133, 18, 92, 220, 162, 163, 79, 201, 66, 96, 112, 57, 224, 101, 105, 255, 83, 217, 144, 233, 242, 195, 102, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
			want: Message{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					common.SystemProgramID,
				},
				RecentBlockHash: "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			err: nil,
		},
		{
			args: args{messageData: []byte{128, 1, 0, 1, 3, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185, 29, 11, 113, 4, 101, 239, 39, 167, 201, 112, 156, 239, 236, 36, 251, 140, 76, 199, 150, 228, 218, 214, 20, 123, 180, 181, 103, 160, 71, 251, 237, 123, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 131, 118, 36, 248, 169, 123, 97, 98, 215, 133, 18, 92, 220, 162, 163, 79, 201, 66, 96, 112, 57, 224, 101, 105, 255, 83, 217, 144, 233, 242, 195, 102, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
			want: Message{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					common.SystemProgramID,
				},
				RecentBlockHash: "9rAtxuhtKn8qagc3UtZFyhLrw5zkh6etv43TibaXuSKo",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			err: nil,
		},
		{
			args: args{messageData: []byte{128, 1, 0, 1, 2, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 62, 255, 204, 109, 44, 223, 1, 225, 41, 92, 205, 204, 199, 90, 32, 104, 6, 123, 211, 72, 233, 131, 88, 65, 115, 38, 138, 217, 189, 202, 86, 39, 1, 1, 2, 0, 2, 12, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 241, 61, 2, 62, 211, 181, 33, 219, 74, 147, 127, 38, 231, 159, 99, 194, 103, 129, 201, 15, 51, 106, 114, 199, 122, 142, 121, 87, 112, 78, 138, 249, 1, 1, 0}},
			want: Message{
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
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MessageDeserialize(tt.args.messageData)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, err, tt.err)
		})
	}
}
