package dynamodb

type DynamoDBModel interface {
	GetTableName() string
}
