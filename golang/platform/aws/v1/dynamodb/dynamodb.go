package dynamodb

import (
	"fmt"
	"golang/platform/aws/v1/dynamodb/model"
	"golang/platform/aws/v1/utils"
	"reflect"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var instance *DynamoDBDataSource
var singleton sync.Once

type AttributeValue dynamodb.AttributeValue

type DynamoDBDataSource struct {
	sess *dynamodb.DynamoDB
}

func GetInstance(sess *session.Session) *DynamoDBDataSource {
	singleton.Do(func() {
		dataSource := &DynamoDBDataSource{
			sess: dynamodb.New(sess),
		}
		instance = dataSource
	})
	return instance
}

func (ds *DynamoDBDataSource) Put(model DynamoDBModel) *dynamodb.PutItemOutput {
	av, err := dynamodbattribute.MarshalMap(model)
	if err != nil {
		panic(err)
	}

	params := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(model.GetTableName()),
	}

	output, err := ds.sess.PutItem(params)

	if err != nil {
		panic(err)
	}

	return output
}

func (ds *DynamoDBDataSource) Scan() {

	tablename := "TRANSACTION_LOG_DEV"
	// address := "0x5b62c110b69dc4c9cee8d54603d503679af7678e"

	// Set up the scan input parameters
	// params := &dynamodb.ScanInput{
	// 	TableName:        aws.String(tablename),
	// 	FilterExpression: aws.String("address = :address"),
	// 	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	// 		":address": {
	// 			N: aws.String(address),
	// 		},
	// 	},
	// }

	typeName := "coin"

	params := &dynamodb.ScanInput{
		TableName:        aws.String(tablename),
		FilterExpression: aws.String("#type = :type"),
		ExpressionAttributeNames: map[string]*string{
			"#type": aws.String("type"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":type": {
				S: aws.String(typeName),
			},
		},
	}

	// Perform the scan operation
	result, err := ds.sess.Scan(params)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println("---------------")

	itemsLen := len(result.Items)

	modelList := make([]model.TransactionLog, itemsLen)
	for i := 0; i < itemsLen; i++ {
		err := dynamodbattribute.UnmarshalMap(result.Items[i], &modelList[i])

		if err != nil {
			panic(err)
		}
	}
	fmt.Println(utils.StringifyJSON(modelList))

}

// TODO
func (ds *DynamoDBDataSource) Query(tableName string, txId string) []model.TransactionLog {
	builder := expression.NewBuilder()

	keyCondition := expression.Key("txId").
		Equal(expression.Value(txId))
		// Equal(expression.Value(txId)).
		// And(
		// 	expression.Key("timestamp").
		// 		Between( //1683165917415
		// 			expression.Value(0),
		// 			expression.Value(1683165917415),
		// 		),
		// )

	expr, err := builder.WithKeyCondition(keyCondition).Build()
	if err != nil {
		panic(err)
	}
	input := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Limit:                     aws.Int64(1),
		ExclusiveStartKey: map[string]*dynamodb.AttributeValue{
			"txId": {
				S: aws.String("456daad5-0cb6-443f-96d0-79c45491a300"),
			},
			"timestamp": {
				N: aws.String("1683510603697"),
			},
		},
	}

	result, err := ds.sess.Query(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println("---------------")

	itemsLen := len(result.Items)

	modelList := make([]model.TransactionLog, itemsLen)
	for i := 0; i < itemsLen; i++ {
		err := dynamodbattribute.UnmarshalMap(result.Items[i], &modelList[i])

		if err != nil {
			panic(err)
		}
	}

	return modelList
}

// TODO
func ParseModelKeys(model DynamoDBModel) map[string][]string {
	var keyMap map[string][]string = make(map[string][]string)
	fields := reflect.TypeOf(model)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		key := field.Tag.Get("keytype")

		if key == "partition" {
			keyMap["partition"] = []string{
				field.Tag.Get("attr"),
				field.Tag.Get("json"),
			}
			continue
		}

		if key == "sort" {
			keyMap["sort"] = []string{
				field.Tag.Get("attr"),
				field.Tag.Get("json"),
			}
			continue
		}
	}

	if keyMap["partition"] == nil {
		panic("DynamoDB model has to have partition key")
	}

	return keyMap
}
