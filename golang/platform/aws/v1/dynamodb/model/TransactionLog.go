package model

import "github.com/aws/aws-sdk-go/service/dynamodb"

type TransactionLog struct {
	TxId      string `json:"txId"`
	Timestamp int    `json:"timestamp"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Data      map[string]*dynamodb.AttributeValue
}

func (_ *TransactionLog) GetTableName() string {
	return "TRANSACTION_LOG_DEV"
}
