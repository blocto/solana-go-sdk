package rpc

type Instruction struct {
	ProgramIDIndex uint64   `json:"programIdIndex"`
	Accounts       []uint64 `json:"accounts"`
	Data           string   `json:"data"`
}

type TransactionMetaTokenBalance struct {
	AccountIndex  int    `json:"accountIndex"`
	Mint          string `json:"mint"`
	UITokenAmount struct {
		Amount         string `json:"amount"`
		Decimals       int64  `json:"decimals"`
		UIAmountString string `json:"uiAmountString"`
	} `json:"uiTokenAmount"`
}

type TransactionMeta struct {
	Fee               uint64                        `json:"fee"`
	PreBalances       []int64                       `json:"preBalances"`
	PostBalances      []int64                       `json:"postBalances"`
	PreTokenBalances  []TransactionMetaTokenBalance `json:"preTokenBalances"`
	PostTokenBalances []TransactionMetaTokenBalance `json:"postTokenBalances"`
	LogMessages       []string                      `json:"logMessages"`
	InnerInstructions []struct {
		Index        uint64        `json:"index"`
		Instructions []Instruction `json:"instructions"`
	} `json:"innerInstructions"`
	Err    interface{}            `json:"err"`
	Status map[string]interface{} `json:"status"`
}

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
