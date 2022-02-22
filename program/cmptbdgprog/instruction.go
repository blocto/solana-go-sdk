package cmptbdgprog

import (
	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction borsh.Enum

const (
	InstructionRequestUnits Instruction = iota
	InstructionRequestHeapFrame
)

type RequestUnitsParam struct {
	Units         uint32
	AdditionalFee uint32
}

// RequestUnits ...
func RequestUnits(param RequestUnitsParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction   Instruction
		Units         uint32
		AdditionalFee uint32
	}{
		Instruction:   InstructionRequestUnits,
		Units:         param.Units,
		AdditionalFee: param.AdditionalFee,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.ComputeBudgetProgramID,
		Accounts:  []types.AccountMeta{},
		Data:      data,
	}
}

type RequestHeapFrameParam struct {
	Bytes uint32
}

// RequestHeapFrame ...
func RequestHeapFrame(param RequestHeapFrameParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Bytes       uint32
	}{
		Instruction: InstructionRequestHeapFrame,
		Bytes:       param.Bytes,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.ComputeBudgetProgramID,
		Accounts:  []types.AccountMeta{},
		Data:      data,
	}
}
