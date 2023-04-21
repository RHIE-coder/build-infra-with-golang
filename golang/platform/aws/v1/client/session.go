package client

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession(accessKey string, secretKey string, region string) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		// Region:      aws.String("ap-northeast-2"),
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		HTTPClient:  &http.Client{},
	})

	if err != nil {
		log.Fatal("fail to initialize the aws session(v1) : " + err.Error())
	}

	return sess
}
