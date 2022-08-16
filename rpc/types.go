package rpc

// Commitment describes how finalized a block is at that point in time
type Commitment string

const (
	CommitmentFinalized Commitment = "finalized"
	CommitmentConfirmed Commitment = "confirmed"
	CommitmentProcessed Commitment = "processed"
)

type Context struct {
	Slot       uint64 `json:"slot"`
	ApiVersion string `json:"apiVersion,omitempty"`
}

type AccountEncoding string

const (
	AccountEncodingBase58     AccountEncoding = "base58"
	AccountEncodingJsonParsed AccountEncoding = "jsonParsed"
	AccountEncodingBase64     AccountEncoding = "base64"
	AccountEncodingBase64Zstd AccountEncoding = "base64+zstd"
)

type TransactionEncoding string

const (
	TransactionEncodingBinary     TransactionEncoding = "binary"
	TransactionEncodingBase64     TransactionEncoding = "base64"
	TransactionEncodingBase58     TransactionEncoding = "base58"
	TransactionEncodingJson       TransactionEncoding = "json"
	TransactionEncodingJsonParsed TransactionEncoding = "jsonParsed"
)

type DataSlice struct {
	Offset uint64 `json:"offset,omitempty"`
	Length uint64 `json:"length,omitempty"`
}

type AccountInfo struct {
	Lamports   uint64 `json:"lamports"`
	Owner      string `json:"owner"`
	RentEpoch  uint64 `json:"rentEpoch"`
	Data       any    `json:"data"`
	Executable bool   `json:"executable"`
}
