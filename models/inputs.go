package models

type CreateBroadcast struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type GetTransaction struct {
	TxHash string `json:"tx_hash"`
}
