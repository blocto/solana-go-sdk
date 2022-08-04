package rpc

type AccountInfo struct {
	Lamports   uint64      `json:"lamports"`
	Owner      string      `json:"owner"`
	RentEpoch  uint64      `json:"rentEpoch"`
	Data       interface{} `json:"data"`
	Executable bool        `json:"executable"`
}
