package dynamodb

import (
	"fmt"
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

// TODO
func (ds *DynamoDBDataSource) Query(tableName string, txId string) {
	builder := expression.NewBuilder()

	keyCondition := expression.Key("txId").
		Equal(expression.Value(txId)).
		And(
			expression.Key("timestamp").
				Between( //1683165917415
					expression.Value(1683165917410),
					expression.Value(1683165917420),
				),
		)

	expr, err := builder.WithKeyCondition(keyCondition).Build()
	if err != nil {
		panic(err)
	}
	input := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	fmt.Println(expr.KeyCondition())
	fmt.Println(expr.Names())
	fmt.Println(expr.Values())

	result, err := ds.sess.Query(input)
	if err != nil {
		panic(err)
	}

	items := len(result.Items)

	for 

}

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
