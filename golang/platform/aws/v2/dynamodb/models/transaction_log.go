package models

type TransactionLog struct {
	TxId      string `dynamodbav:"txId"`
	Timestamp int64  `dynamodbav:"timestamp"`
	Type      string `dynamodbav:"type"`
	Message   string `dynamodbav:"message"`
	Data      Data   `dynamodbav:"data" json:"data"`
}

type Data struct {
	Amount   uint   `dynamodbav:"amount"`
	Currency string `dynamodbav:"currency"`
}

func (_ TransactionLog) GetTableName() string {
	return "TRANSACTION_LOG_DEV_BY_OWEN"
}
