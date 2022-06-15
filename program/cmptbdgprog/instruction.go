package cmptbdgprog

import (
	"github.com/OldSmokeGun/solana-go-sdk/common"
	"github.com/OldSmokeGun/solana-go-sdk/types"
	"github.com/near/borsh-go"
)

type Instruction borsh.Enum

const (
	InstructionRequestUnits Instruction = iota
	InstructionRequestHeapFrame
	InstructionSetComputeUnitLimit
	InstructionSetComputeUnitPrice
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

type SetComputeUnitLimitParam struct {
	Units uint32
}

// SetComputeUnitLimit set a specific compute unit limit that the transaction is allowed to consume.
func SetComputeUnitLimit(param SetComputeUnitLimitParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Units       uint32
	}{
		Instruction: InstructionSetComputeUnitLimit,
		Units:       param.Units,
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

type SetComputeUnitPriceParam struct {
	MicroLamports uint64
}

// SetComputeUnitPrice set a compute unit price in "micro-lamports" to pay a higher transaction
// fee for higher transaction prioritization.
func SetComputeUnitPrice(param SetComputeUnitPriceParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction   Instruction
		MicroLamports uint64
	}{
		Instruction:   InstructionSetComputeUnitPrice,
		MicroLamports: param.MicroLamports,
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
