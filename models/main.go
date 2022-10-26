package models

// Block data structure
type Block struct {
	BlockNumber       int64         `json:"blockNumber"`
	Timestamp         uint64        `json:"timestamp"`
	Difficulty        uint64        `json:"difficulty"`
	Hash              string        `json:"hash"`
	TransactionsCount int           `json:"transactionsCount"`
	Transactions      []Transaction `json:"transactions"`
}

// Transaction data structure
type Transaction struct {
	Hash      string `json:"transactionHash"`
	Type 	  string `json:"type"`
	Address   string `json:"address"`
	BlockHash string `json:"blockHash"`
	BlockNumber int64 `json:"blockNumber"`
	Gas       uint64 `json:"gas"`
	GasPrice  uint64 `json:"price"`
	Nonce     uint64 `json:"nonce"`
	To        string `json:"to"`
	From 	  string `json:"from"`
	Pending   bool   `json:"pending"`
	Date 	  string `json:"date"`

	CollectionName string `json:"collection_name"`
	CollectionAddress string `json:"collection_address"`
}

// TransferEthRequest data str	ucture
type TransferEthRequest struct {
	PrivKey string `json:"privKey"`
	To      string `json:"to"`
	Amount  int64  `json:"amount"`
}

// HashResponse data structure
type HashResponse struct {
	Hash string `json:"hash"`
}
// Error data structure
type Error struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}
