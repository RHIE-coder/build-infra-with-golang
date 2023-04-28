package dynamodb

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDB struct {
	client *dynamodb.DynamoDB
}

var instance *DynamoDB
var singleton sync.Once

func GetInstance(sess *session.Session) *DynamoDB {
	singleton.Do(func() {
		dynamo := &DynamoDB{}
		dynamo.client = dynamodb.New(sess)
		instance = dynamo
	})
	transaction.NewDynamoDBTransaction(instance.client)
	return instance
}
