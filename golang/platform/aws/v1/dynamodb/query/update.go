package query

import (
	"golang/platform/aws/v1/dynamodb"
	"golang/platform/aws/v1/dynamodb/interfaces"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func Update(model *interfaces.DynamoModel) {
	client := dynamodb.GetClient()

	expr, err := expression.NewBuilder().Build()
}
