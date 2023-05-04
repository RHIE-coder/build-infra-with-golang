package main

import (
	"golang/platform/aws/v1/client"
	"golang/platform/aws/v1/dynamodb"
	"golang/platform/aws/v1/dynamodb/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("fail to load .env file")
	}
}

func main() {
	LoadEnvFile()
	ACCESS_KEY := os.Getenv("ACCESS_KEY")
	SECRET_KEY := os.Getenv("SECRET_KEY")
	REGION := os.Getenv("REGION")
	session := client.NewSession(ACCESS_KEY, SECRET_KEY, REGION)
	dataSource := dynamodb.GetInstance(session)

	inputItem := &model.TransactionLog{
		TxId:      "456daad5-0cb6-443f-96d0-79c45491a300",
		Timestamp: 1683165917415,
	}
	// dataSource.Put(&model.TransactionLog{
	// 	TxId:      utils.GetUUID4(),
	// 	Timestamp: nowTimestamp,
	// 	Type:      "coin",
	// 	Message:   "ethereum coin balance",
	// 	Data: model.TransactionLogData{
	// 		Amount:   80,
	// 		Currency: "ETH",
	// 	},
	// })

	dataSource.Query(inputItem.GetTableName(), inputItem.TxId)
	// item := dataSource.Get(&model.TransactionLog{})
	// fmt.Println(utils.StringifyJSON(item))
	/////////////////////////////////////////////////

}

func generateDump() {

}
