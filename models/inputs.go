package models

type CreateBroadcast struct {
	Symbol    string
	Price     uint64
	Timestamp uint64
}

type GetTransaction struct {
	Tx_hash string
}
