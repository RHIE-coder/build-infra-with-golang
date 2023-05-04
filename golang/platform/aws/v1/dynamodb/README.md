# AWS-GO-SDK V1: DynamoDB


## [ Structure ]

### - `dynamodb`
 - github.com/aws/aws-sdk-go/service/`dynamodb`

DynamoDB를 다루는데 기본적으로 필요한 API 재공

### - `dynamodb/dynamodbattribute`
 - github.com/aws/aws-sdk-go/service/`dynamodb/dynamodbattribute`

DynamoDB 데이터와 Go 객체 간의 Marshal, Unmarshal을 제공함

### - `dynamodb/expression`
 - github.com/aws/aws-sdk-go/service/`dynamodb/expression`

DynamoDB 자료형

### - `dynamodb/dynamodbiface`
 - github.com/aws/aws-sdk-go/service/`dynamodb/dynamodbiface`

Testing



## [ Research ]

### - DynamoDB PutItem

```go
// {"amount":{"S":"100"},"currency":{"S":"ETH"}}
type TransactionLog struct {
	TxId      string `json:"txId"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Data      map[string]interface{}
}

// {"amount":{"S":"100"},"currency":{"S":"ETH"}}
type TransactionLog struct {
	TxId      string `json:"txId"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Data      map[string]string
}

// {"amount":{ "N" : "100" },"currency":{"S":"ETH"}}
type TransactionLog struct {
	TxId      string `json:"txId"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Data      TransactionLogData
}
```

아래 경우는 사용하지 말 것.

```go
// {"amount":{"M":{"SS":{"NULL":true},"BS":{"NULL":true},"B":{"NULL":true},"NULL":{"NULL":true},"S":{"NULL":true},"BOOL":{"NULL":true},"NS":{"NULL":true},"L":{"NULL":true},"M":{"NULL":true},"N":{"S":"100"}}},"currency":{"M":{"SS":{"NULL":true},"BS":{"NULL":true},"B":{"NULL":true},"NULL":{"NULL":true},"S":{"S":"USD"},"BOOL":{"NULL":true},"NS":{"NULL":true},"L":{"NULL":true},"M":{"NULL":true},"N":{"NULL":true}}}}
type TransactionLog struct {
	TxId      string `json:"txId"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Data      *dynamodb.AttributeValue
}
```


























 - 조회

```go
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
```

 - 업데이트

```go
import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type BC_AUTH struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Active       bool   `json:"active"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

func main() {
	// AWS 세션 생성
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("ap-northeast-2"), // 리전 설정
		},
	}))

	// DynamoDB 클라이언트 생성
	svc := dynamodb.New(sess)

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

	input := &dynamodb.UpdateItemInput{
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
```

 - 추가
```go
import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type BC_AUTH struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Active       bool   `json:"active"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}


func main() {
	// AWS 세션 생성
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("ap-northeast-2"), // 리전 설정
		},
	}))

	// DynamoDB 클라이언트 생성
	svc := dynamodb.New(sess)

	// 추가할 항목 생성
	newItem := BC_AUTH{
		AccessToken:  "newAccessToken",
		RefreshToken: "newRefreshToken",
		Active:       true,
		CreatedAt:    "2022-04-21 13:00:00",
		UpdatedAt:    "2022-04-21 13:00:00",
	}

	// 새로운 항목을 DynamoDB AttributeValue로 변환
	av, err := dynamodbattribute.MarshalMap(newItem)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// PutItemInput 객체 생성
	input := &dynamodb.PutItemInput{
		TableName: aws.String("BC_AUTH"), // 테이블 이름
		Item:      av,                    // 추가할 항목
	}

	// 항목 추가 요청 전송
	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("PutItem operation successful")
}
```