package types

import (
	"reflect"
	"testing"

	"github.com/olegfomenko/solana-go-sdk/common"
)

func TestMessage_Serialize(t *testing.T) {
	type fields struct {
		Header          MessageHeader
		Accounts        []common.PublicKey
		RecentBlockHash string
		Instructions    []CompiledInstruction
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
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
			want:    []byte{1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Header:          tt.fields.Header,
				Accounts:        tt.fields.Accounts,
				RecentBlockHash: tt.fields.RecentBlockHash,
				Instructions:    tt.fields.Instructions,
			}
			got, err := m.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_DecompileInstructions(t *testing.T) {
	type fields struct {
		Header          MessageHeader
		Accounts        []common.PublicKey
		RecentBlockHash string
		Instructions    []CompiledInstruction
	}
	tests := []struct {
		name   string
		fields fields
		want   []Instruction
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Header:          tt.fields.Header,
				Accounts:        tt.fields.Accounts,
				RecentBlockHash: tt.fields.RecentBlockHash,
				Instructions:    tt.fields.Instructions,
			}
			if got := m.DecompileInstructions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.DecompileInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}
