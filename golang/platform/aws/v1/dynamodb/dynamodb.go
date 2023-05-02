package dynamodb

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var instance *dynamodb.DynamoDB
var singleton sync.Once

func Initialize(sess *session.Session) *dynamodb.DynamoDB {
	singleton.Do(func() {
		instance = dynamodb.New(sess)
	})
	return instance
}

func GetClient() *dynamodb.DynamoDB {
	if instance == nil {
		panic("not initialized")
	}
	return instance
}
