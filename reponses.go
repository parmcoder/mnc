package mnc

// CreateBroadcastResponse a basic GoLang struct which includes the following fields: TxHash.
// It is a struct that is used as a response from broadcasting the transaction.
type CreateBroadcastResponse struct {
	TxHash string `json:"tx_hash"`
}

// GetTransactionResponse a basic GoLang struct which includes the following fields: TxStatus.
// It is a struct that is used as a response from gettting the transaction's status.
type GetTransactionResponse struct {
	TxStatus string `json:"tx_status"`
}
