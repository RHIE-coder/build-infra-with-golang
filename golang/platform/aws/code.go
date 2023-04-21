package main

import (
	"golang/platform/aws/v1/client"
	"golang/platform/aws/v1/dynamodb"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func READ_ENV_FILE() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	READ_ENV_FILE()
	ACCESS_KEY := os.Getenv("ACCESS_KEY")
	SECRET_KEY := os.Getenv("SECRET_KEY")
	REGION := "ap-northeast-2"
	session := client.NewSession(ACCESS_KEY, SECRET_KEY, REGION)
	dynamodbClient := dynamodb.New(session)
	dynamodbClient.QueryOne(
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoiSldUIiwic2VjIjoiQ0hPU1VOX0FQSV9LRVkiLCJpYXQiOjE2NjM2NTIwMTIsImV4cCI6MTY2NjI0NDAxMiwiaXNzIjoiVG9tbXkifQ.9Iig6q9xdd8woninhlXuvU-0bvQa_5Cey_3qGl8d9Fs",
		// "c9f1e0cb-f989-4717-8971-b0b6669b3d2f",
	)

}
