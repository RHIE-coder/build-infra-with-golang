package dynamodb_lib

import (
	"golang/platform/aws/v2/common/session"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	sess *dynamodb.Client
}

func GetClient(accessKey string, secretKey string, region string) (*DynamoDBClient, error) {
	cfg, err := session.GetCredConfig(accessKey, secretKey, region)
	if err != nil {
		return nil, err
	}
	awsCfg := cfg.(aws.Config)
	if err != nil {
		return nil, err
	}

	ddbClient := &DynamoDBClient{
		sess: dynamodb.NewFromConfig(awsCfg),
	}

	return ddbClient, nil
}
