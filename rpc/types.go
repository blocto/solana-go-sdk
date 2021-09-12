package rpc

type MessageHeader struct {
	NumRequiredSignatures       uint8 `json:"numRequiredSignatures"`
	NumReadonlySignedAccounts   uint8 `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}

type Message struct {
	Header          MessageHeader `json:"header"`
	AccountKeys     []string      `json:"accountKeys"`
	RecentBlockhash string        `json:"recentBlockhash"`
	Instructions    []Instruction `json:"instructions"`
}

type Transaction struct {
	Signatures []string `json:"signatures"`
	Message    Message  `json:"message"`
}

type Encoding string

const (
	EncodingBase58     Encoding = "base58" // limited to Account data of less than 128 bytes
	EncodingBase64     Encoding = "base64"
	EncodingBase64Zstd Encoding = "base64+zstd"
)
