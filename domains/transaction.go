package domains

struct CreateBroadcast {
	Symbol string           
    Price uint64          
    Timestamp uint64 
}

struct CreateBroadcastResponse {     
    Tx_hash string 
}

struct GetTransactionResponse {     
    Tx_status string 
}