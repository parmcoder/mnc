package models

type CreateBroadcastResponse struct {
	TxHash string `json:"tx_hash"`
}

type GetTransactionResponse struct {
	TxStatus string `json:"tx_status"`
}
