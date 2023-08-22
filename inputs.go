package mnc

// CreateBroadcast a basic GoLang struct which includes the following fields: Symbol, Price, Timestamp.
// It is a struct that is used as an input for broadcasting a transaction.
type CreateBroadcast struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

// GetTransaction a basic GoLang struct which includes the following fields: TxHash.
// It is a struct that is used as an input for getting the transaction status.
type GetTransaction struct {
	TxHash string `json:"tx_hash"`
}
