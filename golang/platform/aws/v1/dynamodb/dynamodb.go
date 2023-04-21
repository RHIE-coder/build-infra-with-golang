package dynamodb

import (
	"fmt"
	table "golang/platform/aws/v1/dynamodb/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type dynamoDB struct {
	client *dynamodb.DynamoDB
}

func New(sess *session.Session) *dynamoDB {
	dynamodbClient := &dynamoDB{}
	dynamodbClient.client = dynamodb.New(sess)
	return dynamodbClient
}

func (db *dynamoDB) QueryOne(accessToken string) (model DynamoDBModel, err error) {

	// GetItemInput 객체 생성
	input := &dynamodb.GetItemInput{
		TableName: aws.String("BC_AUTH"), // 테이블 이름
		Key: map[string]*dynamodb.AttributeValue{
			"accessToken": {
				S: aws.String(accessToken), // 파티션 키 값
			},
		},
	}

	// 항목 가져오기 요청 전송
	result, err := db.client.GetItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 결과를 BC_AUTH 객체로 언마샬링
	var auth table.BC_AUTH
	err = dynamodbattribute.UnmarshalMap(result.Item, &auth)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 결과 출력
	fmt.Printf("AccessToken: %s, RefreshToken: %s, Active: %t, CreatedAt: %s, UpdatedAt: %s\n",
		auth.AccessToken, auth.RefreshToken, auth.Active, auth.CreatedAt, auth.UpdatedAt)
	return
}

func (db *dynamoDB) Query(accessToken string, refreshToken string) (model DynamoDBModel, err error) {

	// GetItemInput 객체 생성
	input := &dynamodb.GetItemInput{
		TableName: aws.String("BC_AUTH"), // 테이블 이름
		Key: map[string]*dynamodb.AttributeValue{
			"accessToken": {
				S: aws.String(accessToken), // 파티션 키 값
			},
			"refreshToken": {
				S: aws.String(refreshToken), // 정렬 키 값
			},
		},
	}

	// 항목 가져오기 요청 전송
	result, err := db.client.GetItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 결과를 BC_AUTH 객체로 언마샬링
	var auth table.BC_AUTH
	err = dynamodbattribute.UnmarshalMap(result.Item, &auth)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 결과 출력
	fmt.Printf("AccessToken: %s, RefreshToken: %s, Active: %t, CreatedAt: %s, UpdatedAt: %s\n",
		auth.AccessToken, auth.RefreshToken, auth.Active, auth.CreatedAt, auth.UpdatedAt)
	return
}
