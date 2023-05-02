package main

import (
	"golang/platform/aws/v1/client"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	LoadEnvFile()
	ACCESS_KEY := os.Getenv("ACCESS_KEY")
	SECRET_KEY := os.Getenv("SECRET_KEY")
	REGION := os.Getenv("REGION")
	session := client.NewSession(ACCESS_KEY, SECRET_KEY, REGION)

}
