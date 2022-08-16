package rpc

// GetAccountInfoConfigEncoding is account's data encode format
type GetAccountInfoConfigEncoding string

const (
	// GetAccountInfoConfigEncodingBase58 limited to Account data of less than 128 bytes
	GetAccountInfoConfigEncodingBase58     GetAccountInfoConfigEncoding = "base58"
	GetAccountInfoConfigEncodingJsonParsed GetAccountInfoConfigEncoding = "jsonParsed"
	GetAccountInfoConfigEncodingBase64     GetAccountInfoConfigEncoding = "base64"
	GetAccountInfoConfigEncodingBase64Zstd GetAccountInfoConfigEncoding = "base64+zstd"
)

// GetAccountInfoConfigDataSlice is a part of GetAccountInfoConfig
type GetAccountInfoConfigDataSlice struct {
	Offset uint64 `json:"offset,omitempty"`
	Length uint64 `json:"length,omitempty"`
}
