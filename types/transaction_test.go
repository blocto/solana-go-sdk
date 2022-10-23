package types

import (
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
	"github.com/stretchr/testify/assert"
)

func TestTransaction_Serialize(t *testing.T) {
	type fields struct {
		Signatures []Signature
		Message    Message
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			fields: fields{
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
			},
			want:    []byte{1, 189, 98, 67, 19, 102, 99, 124, 234, 70, 209, 28, 10, 33, 66, 167, 162, 222, 122, 16, 68, 248, 129, 46, 111, 221, 255, 40, 40, 236, 84, 233, 213, 234, 185, 235, 222, 155, 204, 139, 164, 184, 155, 32, 54, 151, 73, 235, 65, 200, 76, 127, 111, 244, 72, 183, 208, 21, 247, 114, 176, 181, 21, 77, 8, 1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Transaction{
				Signatures: tt.fields.Signatures,
				Message:    tt.fields.Message,
			}
			got, err := tx.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Transaction.Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transaction.Serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionDeserialize(t *testing.T) {
	type args struct {
		tx []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Transaction
		wantErr bool
	}{
		{
			args: args{
				tx: []byte{1, 189, 98, 67, 19, 102, 99, 124, 234, 70, 209, 28, 10, 33, 66, 167, 162, 222, 122, 16, 68, 248, 129, 46, 111, 221, 255, 40, 40, 236, 84, 233, 213, 234, 185, 235, 222, 155, 204, 139, 164, 184, 155, 32, 54, 151, 73, 235, 65, 200, 76, 127, 111, 244, 72, 183, 208, 21, 247, 114, 176, 181, 21, 77, 8, 1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			},
			want: Transaction{
				Signatures: []Signature{[]byte{189, 98, 67, 19, 102, 99, 124, 234, 70, 209, 28, 10, 33, 66, 167, 162, 222, 122, 16, 68, 248, 129, 46, 111, 221, 255, 40, 40, 236, 84, 233, 213, 234, 185, 235, 222, 155, 204, 139, 164, 184, 155, 32, 54, 151, 73, 235, 65, 200, 76, 127, 111, 244, 72, 183, 208, 21, 247, 114, 176, 181, 21, 77, 8}},
				Message: Message{
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
					AddressLookupTables: []CompiledAddressLookupTable{},
				},
			},
			wantErr: false,
		},
		{
			args: args{
				tx: []byte{2, 74, 231, 188, 191, 144, 39, 14, 161, 169, 155, 174, 83, 136, 177, 49, 105, 154, 137, 23, 153, 145, 47, 130, 208, 246, 195, 244, 141, 52, 228, 21, 190, 130, 99, 162, 145, 30, 133, 140, 2, 103, 40, 95, 141, 116, 111, 249, 205, 59, 137, 56, 204, 67, 132, 148, 152, 74, 69, 48, 200, 227, 0, 156, 8, 33, 150, 49, 151, 221, 70, 119, 149, 120, 244, 227, 186, 179, 109, 146, 176, 20, 58, 224, 180, 254, 64, 210, 181, 208, 226, 151, 52, 192, 198, 242, 20, 184, 23, 238, 214, 165, 140, 56, 190, 100, 122, 29, 216, 79, 196, 144, 239, 203, 64, 106, 255, 216, 27, 153, 242, 78, 154, 235, 204, 72, 58, 227, 3, 2, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 131, 56, 140, 99, 57, 71, 67, 79, 102, 217, 86, 239, 231, 34, 85, 48, 147, 87, 18, 236, 176, 227, 54, 6, 201, 50, 117, 164, 3, 220, 147, 222, 2, 2, 2, 0, 1, 12, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 2, 2, 1, 0, 12, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			},
			want: Transaction{
				Signatures: []Signature{
					[]byte{74, 231, 188, 191, 144, 39, 14, 161, 169, 155, 174, 83, 136, 177, 49, 105, 154, 137, 23, 153, 145, 47, 130, 208, 246, 195, 244, 141, 52, 228, 21, 190, 130, 99, 162, 145, 30, 133, 140, 2, 103, 40, 95, 141, 116, 111, 249, 205, 59, 137, 56, 204, 67, 132, 148, 152, 74, 69, 48, 200, 227, 0, 156, 8},
					[]byte{33, 150, 49, 151, 221, 70, 119, 149, 120, 244, 227, 186, 179, 109, 146, 176, 20, 58, 224, 180, 254, 64, 210, 181, 208, 226, 151, 52, 192, 198, 242, 20, 184, 23, 238, 214, 165, 140, 56, 190, 100, 122, 29, 216, 79, 196, 144, 239, 203, 64, 106, 255, 216, 27, 153, 242, 78, 154, 235, 204, 72, 58, 227, 3},
				},
				Message: Message{
					Version: MessageVersionLegacy,
					Header: MessageHeader{
						NumRequireSignatures:        2,
						NumReadonlySignedAccounts:   0,
						NumReadonlyUnsignedAccounts: 1,
					},
					Accounts: []common.PublicKey{
						common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"),
						common.SystemProgramID,
					},
					RecentBlockHash: "9qERNBLXzCqchyfquh2DjUT21xsLym6ynZPRh9TZbEiq",
					Instructions: []CompiledInstruction{
						{
							ProgramIDIndex: 2,
							Accounts:       []int{0, 1},
							Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
						},
						{
							ProgramIDIndex: 2,
							Accounts:       []int{1, 0},
							Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
						},
					},
					AddressLookupTables: []CompiledAddressLookupTable{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransactionDeserialize(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionDeserialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionDeserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTransaction(t *testing.T) {
	testAccount1 := NewAccount()
	testAccount2 := NewAccount()
	testAccount3 := NewAccount()

	emptySig := make([]byte, 64)

	msg := []Message{
		NewMessage(NewMessageParam{
			FeePayer: testAccount1.PublicKey,
			Instructions: []Instruction{
				{
					ProgramID: common.PublicKeyFromString("CustomProgram111111111111111111111111111111"),
					Accounts: []AccountMeta{
						{
							PubKey:     testAccount2.PublicKey,
							IsSigner:   false,
							IsWritable: false,
						},
						{
							PubKey:     testAccount3.PublicKey,
							IsSigner:   false,
							IsWritable: false,
						},
					},
					Data: []byte{},
				},
			},
			RecentBlockhash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
		}),
		NewMessage(NewMessageParam{
			FeePayer: testAccount1.PublicKey,
			Instructions: []Instruction{
				{
					ProgramID: common.PublicKeyFromString("CustomProgram111111111111111111111111111111"),
					Accounts: []AccountMeta{
						{
							PubKey:     testAccount2.PublicKey,
							IsSigner:   true,
							IsWritable: false,
						},
						{
							PubKey:     testAccount3.PublicKey,
							IsSigner:   false,
							IsWritable: false,
						},
					},
					Data: []byte{},
				},
			},
			RecentBlockhash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
		}),
		NewMessage(NewMessageParam{
			FeePayer: testAccount1.PublicKey,
			Instructions: []Instruction{
				{
					ProgramID: common.PublicKeyFromString("CustomProgram111111111111111111111111111111"),
					Accounts: []AccountMeta{
						{
							PubKey:     testAccount2.PublicKey,
							IsSigner:   true,
							IsWritable: true,
						},
						{
							PubKey:     testAccount3.PublicKey,
							IsSigner:   true,
							IsWritable: false,
						},
					},
					Data: []byte{},
				},
			},
			RecentBlockhash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
		}),
	}
	serMsg := make([][]byte, 0, len(msg))
	for _, m := range msg {
		sm, _ := m.Serialize()
		serMsg = append(serMsg, sm)
	}

	type args struct {
		message Message
		signers []Account
	}
	tests := []struct {
		name string
		args args
		want Transaction
		err  error
	}{
		{
			args: args{
				message: msg[0],
				signers: []Account{},
			},
			want: Transaction{
				Signatures: []Signature{
					emptySig,
				},
				Message: msg[0],
			},
		},
		{
			args: args{
				message: msg[0],
				signers: []Account{testAccount2},
			},
			want: Transaction{},
			err:  ErrTransactionAddNotNecessarySignatures,
		},
		{
			args: args{
				message: msg[0],
				signers: []Account{
					testAccount1,
				},
			},
			want: Transaction{
				Signatures: []Signature{
					testAccount1.Sign(serMsg[0]),
				},
				Message: msg[0],
			},
		},
		{
			args: args{
				message: msg[1],
				signers: []Account{},
			},
			want: Transaction{
				Signatures: []Signature{
					emptySig,
					emptySig,
				},
				Message: msg[1],
			},
		},
		{
			args: args{
				message: msg[1],
				signers: []Account{testAccount1},
			},
			want: Transaction{
				Signatures: []Signature{
					testAccount1.Sign(serMsg[1]),
					emptySig,
				},
				Message: msg[1],
			},
		},
		{
			args: args{
				message: msg[1],
				signers: []Account{testAccount2},
			},
			want: Transaction{
				Signatures: []Signature{
					emptySig,
					testAccount2.Sign(serMsg[1]),
				},
				Message: msg[1],
			},
		},
		{
			args: args{
				message: msg[1],
				signers: []Account{testAccount2, testAccount1},
			},
			want: Transaction{
				Signatures: []Signature{
					testAccount1.Sign(serMsg[1]),
					testAccount2.Sign(serMsg[1]),
				},
				Message: msg[1],
			},
		},
		{
			args: args{
				message: msg[1],
				signers: []Account{testAccount1, testAccount2},
			},
			want: Transaction{
				Signatures: []Signature{
					testAccount1.Sign(serMsg[1]),
					testAccount2.Sign(serMsg[1]),
				},
				Message: msg[1],
			},
		},
		{
			args: args{
				message: msg[2],
				signers: []Account{testAccount2},
			},
			want: Transaction{
				Signatures: []Signature{
					emptySig,
					testAccount2.Sign(serMsg[2]),
					emptySig,
				},
				Message: msg[2],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTransaction(NewTransactionParam{
				Message: tt.args.message,
				Signers: tt.args.signers,
			})
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTransaction_AddSignature(t *testing.T) {
	testAccount1 := NewAccount()
	testAccount2 := NewAccount()
	testAccount3 := NewAccount()
	testAccount4 := NewAccount()

	emptySig := make([]byte, 64)
	msg := NewMessage(NewMessageParam{
		FeePayer: testAccount1.PublicKey,
		Instructions: []Instruction{
			{
				ProgramID: common.PublicKeyFromString("CustomProgram111111111111111111111111111111"),
				Accounts: []AccountMeta{
					{
						PubKey:     testAccount2.PublicKey,
						IsSigner:   true,
						IsWritable: true,
					},
					{
						PubKey:     testAccount3.PublicKey,
						IsSigner:   true,
						IsWritable: false,
					},
				},
				Data: []byte{},
			},
		},
		RecentBlockhash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
	})
	serMsg, _ := msg.Serialize()

	type args struct {
		sig []byte
	}
	tests := []struct {
		name string
		args args
		tx   Transaction
		want Transaction
		err  error
	}{
		{
			name: "add",
			tx: Transaction{
				Signatures: []Signature{
					emptySig,
					emptySig,
					emptySig,
				},
				Message: msg,
			},
			args: args{
				sig: testAccount1.Sign(serMsg),
			},
			want: Transaction{
				Signatures: []Signature{
					testAccount1.Sign(serMsg),
					emptySig,
					emptySig,
				},
				Message: msg,
			},
		},
		{
			name: "add duplicate",
			tx: Transaction{
				Signatures: []Signature{
					testAccount1.Sign(serMsg),
					emptySig,
					emptySig,
				},
				Message: msg,
			},
			args: args{
				sig: testAccount1.Sign(serMsg),
			},
			want: Transaction{
				Signatures: []Signature{
					testAccount1.Sign(serMsg),
					emptySig,
					emptySig,
				},
				Message: msg,
			},
		},
		{
			name: "add no match",
			tx: Transaction{
				Signatures: []Signature{
					testAccount1.Sign(serMsg),
					emptySig,
					emptySig,
				},
				Message: msg,
			},
			args: args{
				sig: testAccount4.Sign(serMsg),
			},
			want: Transaction{
				Signatures: []Signature{
					testAccount1.Sign(serMsg),
					emptySig,
					emptySig,
				},
				Message: msg,
			},
			err: ErrTransactionAddNotNecessarySignatures,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := tt.tx
			err := tx.AddSignature(tt.args.sig)
			assert.Equal(t, tt.want, tx)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}
