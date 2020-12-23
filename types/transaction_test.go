package types

import (
	"crypto/ed25519"
	"reflect"
	"testing"

	"github.com/portto/solana-go-sdk/common"
)

func TestTransaction_Sign(t *testing.T) {
	type fields struct {
		Signatures []Signature
		Message    Message
	}
	type args struct {
		privateKeys []ed25519.PrivateKey
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Transaction
		wantErr bool
	}{
		{
			name: "sign a tx",
			fields: fields{
				Signatures: []Signature{},
				Message: Message{
					Header: MessageHeader{
						NumRequireSignatures:        1,
						NumReadonlySignedAccounts:   0,
						NumReadonlyUnsignedAccounts: 1,
					},
					Accounts: []common.PublicKey{
						common.PublicKeyFromHex("ced387e6c36f57fe93ef8f516e9f318c6d89e0c51831df3d7b084e6d6e88e4f0"),
						common.PublicKeyFromHex("86acd1d5e3893d6c74abcd7c36443d6e501ff0756c8961de26f2449c1b411d8e"),
						common.PublicKeyFromHex("0"),
					},
					RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
					Instructions: []compiledInstruction{
						{
							ProgramIDIndex: 2,
							Accounts:       []int{0, 1},
							Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
						},
					},
				},
			},
			args: args{[]ed25519.PrivateKey{[]byte{220, 190, 97, 243, 86, 180, 6, 192, 121, 120, 30, 246, 134, 81, 46, 27, 181, 181, 148, 200, 182, 184, 13, 124, 51, 186, 141, 11, 125, 116, 9, 203, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240}}},
			want: &Transaction{
				Signatures: []Signature{[]byte{189, 98, 67, 19, 102, 99, 124, 234, 70, 209, 28, 10, 33, 66, 167, 162, 222, 122, 16, 68, 248, 129, 46, 111, 221, 255, 40, 40, 236, 84, 233, 213, 234, 185, 235, 222, 155, 204, 139, 164, 184, 155, 32, 54, 151, 73, 235, 65, 200, 76, 127, 111, 244, 72, 183, 208, 21, 247, 114, 176, 181, 21, 77, 8}},
				Message: Message{
					Header: MessageHeader{
						NumRequireSignatures:        1,
						NumReadonlySignedAccounts:   0,
						NumReadonlyUnsignedAccounts: 1,
					},
					Accounts: []common.PublicKey{
						common.PublicKeyFromHex("ced387e6c36f57fe93ef8f516e9f318c6d89e0c51831df3d7b084e6d6e88e4f0"),
						common.PublicKeyFromHex("86acd1d5e3893d6c74abcd7c36443d6e501ff0756c8961de26f2449c1b411d8e"),
						common.PublicKeyFromHex("0"),
					},
					RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
					Instructions: []compiledInstruction{
						{
							ProgramIDIndex: 2,
							Accounts:       []int{0, 1},
							Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "sign a tx which has signed before",
			fields: fields{
				Signatures: []Signature{[]byte("Signatures1"), []byte("Signatures2")},
				Message: Message{
					Header: MessageHeader{
						NumRequireSignatures:        1,
						NumReadonlySignedAccounts:   0,
						NumReadonlyUnsignedAccounts: 1,
					},
					Accounts: []common.PublicKey{
						common.PublicKeyFromHex("ced387e6c36f57fe93ef8f516e9f318c6d89e0c51831df3d7b084e6d6e88e4f0"),
						common.PublicKeyFromHex("86acd1d5e3893d6c74abcd7c36443d6e501ff0756c8961de26f2449c1b411d8e"),
						common.PublicKeyFromHex("0"),
					},
					RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
					Instructions: []compiledInstruction{
						{
							ProgramIDIndex: 2,
							Accounts:       []int{0, 1},
							Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
						},
					},
				},
			},
			args: args{[]ed25519.PrivateKey{[]byte{220, 190, 97, 243, 86, 180, 6, 192, 121, 120, 30, 246, 134, 81, 46, 27, 181, 181, 148, 200, 182, 184, 13, 124, 51, 186, 141, 11, 125, 116, 9, 203, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240}}},
			want: &Transaction{
				Signatures: []Signature{[]byte{189, 98, 67, 19, 102, 99, 124, 234, 70, 209, 28, 10, 33, 66, 167, 162, 222, 122, 16, 68, 248, 129, 46, 111, 221, 255, 40, 40, 236, 84, 233, 213, 234, 185, 235, 222, 155, 204, 139, 164, 184, 155, 32, 54, 151, 73, 235, 65, 200, 76, 127, 111, 244, 72, 183, 208, 21, 247, 114, 176, 181, 21, 77, 8}},
				Message: Message{
					Header: MessageHeader{
						NumRequireSignatures:        1,
						NumReadonlySignedAccounts:   0,
						NumReadonlyUnsignedAccounts: 1,
					},
					Accounts: []common.PublicKey{
						common.PublicKeyFromHex("ced387e6c36f57fe93ef8f516e9f318c6d89e0c51831df3d7b084e6d6e88e4f0"),
						common.PublicKeyFromHex("86acd1d5e3893d6c74abcd7c36443d6e501ff0756c8961de26f2449c1b411d8e"),
						common.PublicKeyFromHex("0"),
					},
					RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
					Instructions: []compiledInstruction{
						{
							ProgramIDIndex: 2,
							Accounts:       []int{0, 1},
							Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Transaction{
				Signatures: tt.fields.Signatures,
				Message:    tt.fields.Message,
			}
			got, err := tx.sign(tt.args.privateKeys)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transaction.Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transaction.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
						common.PublicKeyFromHex("ced387e6c36f57fe93ef8f516e9f318c6d89e0c51831df3d7b084e6d6e88e4f0"),
						common.PublicKeyFromHex("86acd1d5e3893d6c74abcd7c36443d6e501ff0756c8961de26f2449c1b411d8e"),
						common.PublicKeyFromHex("0"),
					},
					RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
					Instructions: []compiledInstruction{
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

func TestCreateRawTransaction(t *testing.T) {
	type args struct {
		param CreateRawTransactionParam
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{
				param: CreateRawTransactionParam{
					Instructions: []Instruction{
						{
							ProgramID: common.PublicKeyFromString("11111111111111111111111111111111"),
							Accounts: []AccountMeta{
								{
									PubKey:     common.PublicKeyFromHex("ced387e6c36f57fe93ef8f516e9f318c6d89e0c51831df3d7b084e6d6e88e4f0"),
									IsSigner:   true,
									IsWritable: true,
								},
								{
									PubKey:     common.PublicKeyFromHex("86acd1d5e3893d6c74abcd7c36443d6e501ff0756c8961de26f2449c1b411d8e"),
									IsSigner:   false,
									IsWritable: true,
								},
							},
							Data: []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
						},
					},
					PrivateKeys:     []ed25519.PrivateKey{[]byte{220, 190, 97, 243, 86, 180, 6, 192, 121, 120, 30, 246, 134, 81, 46, 27, 181, 181, 148, 200, 182, 184, 13, 124, 51, 186, 141, 11, 125, 116, 9, 203, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240}},
					FeePayer:        common.PublicKeyFromHex("0"),
					RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
				},
			},
			want:    []byte{1, 189, 98, 67, 19, 102, 99, 124, 234, 70, 209, 28, 10, 33, 66, 167, 162, 222, 122, 16, 68, 248, 129, 46, 111, 221, 255, 40, 40, 236, 84, 233, 213, 234, 185, 235, 222, 155, 204, 139, 164, 184, 155, 32, 54, 151, 73, 235, 65, 200, 76, 127, 111, 244, 72, 183, 208, 21, 247, 114, 176, 181, 21, 77, 8, 1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTransaction(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
