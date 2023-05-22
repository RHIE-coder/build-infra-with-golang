package dynamodb_lib

type DynamoDBModel interface {
	GetTableName() string
}
