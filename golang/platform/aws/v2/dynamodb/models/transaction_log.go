package models

import (
	"fmt"
	"reflect"
	"strconv"

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

// `dynamodbav`를 중심으로 구성된 데이터를 Unmarshal하는 작업
func (model *TransactionLog) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	err := convertAttributeValueToObject(model, av)
	if err != nil {
		return err
	}
	model.Amount = model.Data.Amount
	model.Currency = model.Data.Currency
	return nil
}

func convertAttributeValueToObject(model interface{}, av types.AttributeValue) error {
	modelTyp := reflect.TypeOf(model).Elem()
	modelVal := reflect.ValueOf(model)

	if avM, ok := av.(*types.AttributeValueMemberM); ok {
		for i := 0; i < modelTyp.NumField(); i++ {
			modelField := modelTyp.Field(i)
			avTagName := modelField.Tag.Get("dynamodbav")

			if avTagName == "" || avTagName == "-" {
				continue
			}

			modelKeyName := modelField.Name

			var elemOfModelVal reflect.Value
			if modelVal.Kind() == reflect.Ptr {
				elemOfModelVal = modelVal.Elem().FieldByName(modelKeyName)
			} else {
				elemOfModelVal = modelVal.FieldByName(modelKeyName)
			}

			elemOfavM := avM.Value[avTagName]

			switch elemOfModelVal.Type().Kind() {
			case reflect.String:
				avS, ok := elemOfavM.(*types.AttributeValueMemberS)
				if !ok {
					return fmt.Errorf("fail to assert type to member S")
				}
				elemOfModelVal.SetString(avS.Value)
			case reflect.Int, reflect.Int64:
				avN, ok := elemOfavM.(*types.AttributeValueMemberN)
				if !ok {
					return fmt.Errorf("fail to assert type to member N")
				}
				num, err := strconv.Atoi(avN.Value)
				if err != nil {
					return err
				}
				elemOfModelVal.SetInt(int64(num))
			case reflect.Uint:
				avN, ok := elemOfavM.(*types.AttributeValueMemberN)
				if !ok {
					return fmt.Errorf("fail to assert type to member N")
				}
				num, err := strconv.Atoi(avN.Value)
				if err != nil {
					return err
				}
				elemOfModelVal.SetUint(uint64(num))
			case reflect.Struct:
				address := elemOfModelVal.Addr()
				err := convertAttributeValueToObject(address.Interface(), elemOfavM)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
