package model

type TransactionLog struct {
	TxId      string `json:"txId" keytype:"partition" attr:"S"`
	Timestamp int64  `json:"timestamp" keytype:"sort" attr:"N"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Data      TransactionLogData
}

func (_ TransactionLog) GetTableName() string {
	return "TRANSACTION_LOG_DEV"
}

// TODO
// func (_ *TransactionLog) GetTableName() string {
// 	return "TRANSACTION_LOG_DEV"
// }

type TransactionLogData struct {
	Amount   uint   `json:"amount"`
	Currency string `json:"currency"`
}
