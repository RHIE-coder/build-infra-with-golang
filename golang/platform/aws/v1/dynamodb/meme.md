
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