package main

import (
	"context"
	"fmt"
	"golang/platform/aws/v2/utils"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err.Error())
	}
}

type TransactionLog struct {
	TxId      string `dynamodbav:"txId" keytype:"partition" attr:"S"`
	Timestamp int64  `dynamodbav:"timestamp" keytype:"sort" attr:"N"`
	Type      string `dynamodbav:"type"`
	Message   string `dynamodbav:"message"`
	Data      TransactionLogData
}

func (_ TransactionLog) GetTableName() string {
	return "TRANSACTION_LOG_DEV_BY_OWEN"
}

type TransactionLogData struct {
	Amount   uint   `dynamodbav:"amount"`
	Currency string `dynamodbav:"currency"`
}

func main() {
	LoadEnvFile()
	ACCESS_KEY := os.Getenv("ACCESS_KEY")
	SECRET_KEY := os.Getenv("SECRET_KEY")

	creds := credentials.NewStaticCredentialsProvider(ACCESS_KEY, SECRET_KEY, "")

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(creds),
		config.WithRegion("ap-northeast-2"),
	)
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	/* PutItem */
	// typeName := []string{"coin", "token"}
	// messages := []string{
	// 	"ethereum coin balance",
	// 	"klaytn coin balance",
	// 	"news token balance",
	// }
	// currency := []string{"ETH", "KLAYTN", "NEWS"}
	// for i := 0; i < 10; i++ {
	// 	Put(svc, TransactionLog{
	// 		// TxId: utils.GetUUID4(),
	// 		TxId:      "a120f758-88cb-4ae2-9dc4-006159f05624",
	// 		Timestamp: utils.GetNowTimestamp(),
	// 		Type:      typeName[rand.Intn(2)],
	// 		Message:   messages[rand.Intn(3)],
	// 		Data: TransactionLogData{
	// 			Amount:   uint(rand.Intn(100-10) + 10),
	// 			Currency: currency[rand.Intn(3)],
	// 		},
	// 	})
	// }

	/* Query */
	Get(svc)

}

func Get(svc *dynamodb.Client) {
	// 아래 코드도 동작함
	// filter := expression.Name("txId").Equal(expression.Value("a120f758-88cb-4ae2-9dc4-006159f05624")).
	// 	And(expression.Name("timestamp").Between(
	// 		expression.Value(0),
	// 		expression.Value(1683597820540),
	// 	))
	// expr, err := expression.NewBuilder().WithFilter(filter).Build()

	keyFilter := expression.Key("txId").Equal(expression.Value("a120f758-88cb-4ae2-9dc4-006159f05624")).
		And(expression.Key("timestamp").Between(
			expression.Value(0),
			expression.Value(1683599326844),
		))

	filter := expression.Name("type").Equal(expression.Value("token"))
	expr, err := expression.NewBuilder().WithKeyCondition(keyFilter).WithFilter(filter).Build()
	if err != nil {
		panic(err)
	}

	params := &dynamodb.QueryInput{
		TableName:                 aws.String("TRANSACTION_LOG_DEV_BY_OWEN"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
	}

	output, err := svc.Query(context.TODO(), params)

	if err != nil {
		panic(err)
	}

	var items []TransactionLog
	err = attributevalue.UnmarshalListOfMaps(output.Items, &items)
	if err != nil {
		panic(err)
	}

	// fmt.Println(utils.StringifyJSON(output))
	fmt.Println(utils.StringifyJSON(items))
}

func Put(svc *dynamodb.Client, model TransactionLog) {
	av, err := attributevalue.MarshalMap(model)
	if err != nil {
		panic(err)
	}

	params := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(model.GetTableName()),
	}

	output, err := svc.PutItem(context.TODO(), params)
	if err != nil {
		panic(err)
	}

	fmt.Println(utils.StringifyJSON(output))
}
