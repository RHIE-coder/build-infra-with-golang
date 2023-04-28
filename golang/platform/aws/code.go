package main

import (
	"fmt"
	"golang/platform/aws/v1/client"
	"golang/platform/aws/v1/dynamodb"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func update(dataSource *dynamodb.DynamoDB) {
	// UpdateItemInput 객체 생성
	expr, err := expression.NewBuilder().
		WithUpdate(expression.Set(
			expression.Name("Active"), expression.Value(true),
		)).
		Build()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	input := &dataSource.client.UpdateItemInput{
		TableName: aws.String("BC_AUTH"), // 테이블 이름
		Key: map[string]*dynamodb.AttributeValue{
			"accessToken": {
				S: aws.String("myAccessToken"), // 파티션 키 값
			},
			"refreshToken": {
				S: aws.String("myRefreshToken"), // 정렬 키 값
			},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	// 항목 업데이트 요청 전송
	_, err = svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("UpdateItem operation successful")
}

func main() {
	LoadEnvFile()
	ACCESS_KEY := os.Getenv("ACCESS_KEY")
	SECRET_KEY := os.Getenv("SECRET_KEY")
	REGION := os.Getenv("REGION")
	session := client.NewSession(ACCESS_KEY, SECRET_KEY, REGION)
	dataSource := dynamodb.GetInstance(session)

}
