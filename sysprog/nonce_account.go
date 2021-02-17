package sysprog

import (
	"encoding/binary"
	"fmt"

	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

const NonceAccountSize = 80

type NonceAccount struct {
	Version          uint32
	State            uint32
	AuthorizedPubkey common.PublicKey
	Nonce            common.PublicKey
	FeeCalculator    FeeCalculator
}

func NonceAccountDeserialize(data []byte) (NonceAccount, error) {
	if len(data) < NonceAccountSize {
		return NonceAccount{}, fmt.Errorf("nonce account data size is not enough")
	}
	version := binary.LittleEndian.Uint32(data[:4])
	state := binary.LittleEndian.Uint32(data[4:8])
	authorizedPubkey := common.PublicKeyFromBytes(data[8:40])
	nonce := common.PublicKeyFromBytes(data[40:72])
	feeCalculator, err := FeeCalculatorDeserialize(data[72:])
	if err != nil {
		return NonceAccount{}, err
	}
	return NonceAccount{
		Version:          version,
		State:            state,
		AuthorizedPubkey: authorizedPubkey,
		Nonce:            nonce,
		FeeCalculator:    feeCalculator,
	}, nil
}

type CreateNonceAccountParam struct {
	FromPubkey       common.PublicKey
	NoncePubkey      common.PublicKey
	AuthorizedPubkey common.PublicKey
	Lamports         uint64
}

func CreateNonceAccountInstruction(param CreateNonceAccountParam) []types.Instruction {
	return []types.Instruction{
		CreateAccount(param.FromPubkey, param.NoncePubkey, common.SystemProgramID, param.Lamports, NonceAccountSize),
		InitializeNonce(param.NoncePubkey, param.AuthorizedPubkey),
	}
}

func InitializeNonce(noncePubkey, authPubkey common.PublicKey) types.Instruction {
	instruction := make([]byte, 4)
	binary.LittleEndian.PutUint32(instruction, 6)

	data := make([]byte, 0, 36)
	data = append(data, instruction...)
	data = append(data, authPubkey.Bytes()...)

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: noncePubkey, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRecentBlockhashsPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		ProgramID: common.SystemProgramID,
		Data:      data,
	}
}
