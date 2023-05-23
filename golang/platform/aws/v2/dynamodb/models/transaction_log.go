package models

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TransactionLog struct {
	TxId      string `dynamodbav:"txId"      json:"tx_id"`
	Timestamp int64  `dynamodbav:"timestamp" json:"timestamp"`
	Type      string `dynamodbav:"type"      json:"type"`
	Message   string `dynamodbav:"message"   json:"message"`
	Data      Data   `dynamodbav:"data"      json:"-"`
	Amount    uint   `dynamodbav:"-"         json:"amount"`
	Currency  string `dynamodbav:"-"         json:"currency"`
}

type Data struct {
	Amount   uint   `dynamodbav:"amount"     json:"-"`
	Currency string `dynamodbav:"currency"   json:"-"`
}

func (_ TransactionLog) GetTableName() string {
	return "TRANSACTION_LOG_DEV_BY_OWEN"
}

func (model *TransactionLog) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	if m, ok := av.(*types.AttributeValueMemberM); ok {
		typ := reflect.TypeOf(model).Elem()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			avTagName := field.Tag.Get("dynamodbav")
			if avTagName != "-" {
				fmt.Println(avTagName, m.Value[avTagName])
			}

			// fmt.Println(typ.Field(i).Tag.Get("dynamodbav"))
			// fmt.Println(m.Value[typ.Field(i).Tag.Get("dynamodbav")])
		}
	}

	return nil
}

// func ConvertAttributeValueMemberMToAttributValueS(m *types.AttributeValueMemberM, attrName string) (*types.AttributeValueMemberS, error) {

// 	target = "ca_type"
// 	sAttr, err = utils.ConvertAttributeValueMemberMToAttributValueS(m, target)
// 	if err != nil {
// 		return err
// 	}
// 	l.ContractType = sAttr.Value

// 	v, ok := m.Value[attrName]
// 	if !ok {
// 		return nil, fmt.Errorf("expected `%s` map key", attrName)
// 	}

// 	vv, kk := v.(*types.AttributeValueMemberS)
// 	if !kk || vv == nil {
// 		return nil, fmt.Errorf("expected `%s` map value string", attrName)
// 	}

// 	return vv, nil
// }
