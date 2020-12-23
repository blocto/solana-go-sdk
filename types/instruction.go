package types

import "github.com/portto/solana-go-sdk/common"

type compiledInstruction struct {
	ProgramIDIndex int
	Accounts       []int
	Data           []byte
}

type Instruction struct {
	ProgramID common.PublicKey
	Accounts  []AccountMeta
	Data      []byte
}
