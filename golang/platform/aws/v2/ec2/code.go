package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	LoadEnvFile()
	ACCESS_KEY := os.Getenv("ACCESS_KEY")
	SECRET_KEY := os.Getenv("SECRET_KEY")

	creds := credentials.NewStaticCredentialsProvider(ACCESS_KEY, SECRET_KEY, "")

	// AWS SDK v2의 기본 자격 증명 공급자를 사용하여 구성합니다.
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(creds),
		config.WithRegion("ap-northeast-2"),
	)
	if err != nil {
		panic(err)
	}

	// EC2 클라이언트를 생성합니다.
	svc := ec2.NewFromConfig(cfg)

	// EC2 인스턴스 목록을 가져옵니다.
	resp, err := svc.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		MaxResults: aws.Int32(100), // 최대 100개의 인스턴스를 가져옵니다.
	})
	if err != nil {
		panic(err)
	}

	// 인스턴스 정보를 출력합니다.
	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Instance ID: %s\n", aws.ToString(instance.InstanceId))
			fmt.Printf("Private IP Address: %s\n", aws.ToString(instance.PrivateIpAddress))
			fmt.Printf("Public IP Address: %s\n", aws.ToString(instance.PublicIpAddress))
			fmt.Println()
		}
	}

}
