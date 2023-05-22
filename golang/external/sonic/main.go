package main

import (
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

type LogTxnEvent struct {
	ChainAddress       string `dynamodbav:"chain_addr" json:"chainaddr"` // chainname#address
	TimestampWithIndex string `dynamodbav:"ts_idx"     json:"tsidx"`     // timestamp#block_number#tx_index#log_index
	EventType          string `dynamodbav:"event_type" json:"eventtype"`
	TransactionHash    string `dynamodbav:"txn_hash"   json:"txnhash"`
	TransactionFee     string `dynamodbav:"txn_fee"    json:"txnfee"`
	ContractAddress    string `dynamodbav:"ca_addr"    json:"caaddr"`
	ContractType       string `dynamodbav:"ca_type"    json:"catype"`
	MethodName         string `dynamodbav:"method"     json:"method"`

	// optional attributes
	From  string `dynamodbav:"from"  json:"from"`
	To    string `dynamodbav:"to"    json:"to"`
	Value string `dynamodbav:"value" json:"value"`

	chainName        string
	address          string
	timestamp        string
	blockNumber      string
	transactionIndex string
	logIndex         string
}

type logData struct {
	TxHash   string `json:"tx_hash"`
	TxFee    string `json:"tx_fee"`
	Action   string `json:"action"`
	Method   string `json:"method"`
	Datetime string `json:"datetime"`
	From     string `json:"from"`
	To       string `json:"to"`
	Value    string `json:"value"`
}

type logEventsResData struct {
	Contract string    `json:"contract"`
	Name     string    `json:"name"`
	Symbol   string    `json:"symbol"`
	Logs     []logData `json:"logs"`
}

func main() {
	alice := LogTxnEvent{}

	// Marshal
	output, err := sonic.Marshal(&alice)
	fmt.Println(string(output))

	data := LogTxnEvent{}

	// Unmarshal
	err = sonic.Unmarshal(output, &data)
	if err != nil {
		panic(err)
	}
	sData, err := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(sData))
}
