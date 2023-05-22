package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"golang/platform/aws/v2/common/utils"
	dynamodb_module "golang/platform/aws/v2/dynamodb"
	"golang/platform/aws/v2/dynamodb/models"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	// _, _, execPath, _ := runtime.Caller(0)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid args")
	}
	cmd := os.Args[1]

	LoadEnvFile()
	// environment.LoadEnvFile()
	ACCESS_KEY := os.Getenv("ACCESS_KEY")
	SECRET_KEY := os.Getenv("SECRET_KEY")
	REGION := os.Getenv("REGION")

	ddbClient, err := dynamodb_module.GetClient(ACCESS_KEY, SECRET_KEY, REGION)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch cmd {
	case "put":
		typeName := []string{"coin", "token"}
		messages := []string{
			"ethereum coin balance",
			"klaytn coin balance",
			"news token balance",
		}
		currency := []string{"ETH", "KLAYTN", "NEWS"}

		txLogItem := models.TransactionLog{
			TxId:      "a120f758-88cb-4ae2-9dc4-006159f05624",
			Timestamp: utils.GetNowTimestamp(),
			Type:      typeName[rand.Intn(2)],
			Message:   messages[rand.Intn(3)],
			Data: models.Data{
				Amount:   uint(rand.Intn(100-10) + 10),
				Currency: currency[rand.Intn(3)],
			},
		}
		err = ddbClient.Put(txLogItem)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "query":
		items, err := ddbClient.Query()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(utils.StringifyJSON(items))
	}

	/* Query */
	// Get(svc)

}

// func Get(svc *dynamodb.Client) {
// 	// 아래 코드도 동작함
// 	// filter := expression.Name("txId").Equal(expression.Value("a120f758-88cb-4ae2-9dc4-006159f05624")).
// 	// 	And(expression.Name("timestamp").Between(
// 	// 		expression.Value(0),
// 	// 		expression.Value(1683597820540),
// 	// 	))
// 	// expr, err := expression.NewBuilder().WithFilter(filter).Build()

// 	keyFilter := expression.Key("txId").Equal(expression.Value("a120f758-88cb-4ae2-9dc4-006159f05624")).
// 		And(expression.Key("timestamp").Between(
// 			expression.Value(0),
// 			expression.Value(1683599326844),
// 		))

// 	filter := expression.Name("type").Equal(expression.Value("token"))
// 	expr, err := expression.NewBuilder().WithKeyCondition(keyFilter).WithFilter(filter).Build()
// 	if err != nil {
// 		panic(err)
// 	}

// 	params := &dynamodb.QueryInput{
// 		TableName:                 aws.String("TRANSACTION_LOG_DEV_BY_OWEN"),
// 		ExpressionAttributeNames:  expr.Names(),
// 		ExpressionAttributeValues: expr.Values(),
// 		KeyConditionExpression:    expr.KeyCondition(),
// 		FilterExpression:          expr.Filter(),
// 	}

// 	output, err := svc.Query(context.TODO(), params)

// 	if err != nil {
// 		panic(err)
// 	}

// 	var items []TransactionLog
// 	err = attributevalue.UnmarshalListOfMaps(output.Items, &items)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// fmt.Println(utils.StringifyJSON(output))
// 	fmt.Println(utils.StringifyJSON(items))
// }

// func Put(svc *dynamodb.Client, model TransactionLog) {
// 	av, err := attributevalue.MarshalMap(model)
// 	if err != nil {
// 		panic(err)
// 	}

// 	params := &dynamodb.PutItemInput{
// 		Item:      av,
// 		TableName: aws.String(model.GetTableName()),
// 	}

// 	output, err := svc.PutItem(context.TODO(), params)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(utils.StringifyJSON(output))
// }
