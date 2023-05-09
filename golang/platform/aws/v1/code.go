package main

import (
	"golang/platform/aws/v1/client"
	"golang/platform/aws/v1/dynamodb"
	"golang/platform/aws/v1/dynamodb/model"
	"golang/platform/aws/v1/utils"
	"log"
	"math/rand"
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

	/* PutItem */
	typeName := []string{"coin", "token"}
	messages := []string{
		"ethereum coin balance",
		"klaytn coin balance",
		"news token balance",
	}
	currency := []string{"ETH", "KLAYTN", "NEWS"}
	for i := 0; i < 1; i++ {
		dataSource.Put(&model.TransactionLog{
			// TxId:      utils.GetUUID4(),
			TxId:      "456daad5-0cb6-443f-96d0-79c45491a300",
			Timestamp: utils.GetNowTimestamp(),
			Type:      typeName[rand.Intn(2)],
			Message:   messages[rand.Intn(3)],
			Data: model.TransactionLogData{
				Amount:   uint(rand.Intn(100-10) + 10),
				Currency: currency[rand.Intn(3)],
			},
		})
	}

	// inputItem := &model.TransactionLog{
	// 	TxId:      "456daad5-0cb6-443f-96d0-79c45491a300",
	// 	Timestamp: 1683165917415,
	// }
	// /* Query */
	// modellist := dataSource.Query(inputItem.GetTableName(), inputItem.TxId)
	// fmt.Println(utils.StringifyJSON(modellist))
	/////////////////////////////////////////////////

}

func generateDump() {

}
