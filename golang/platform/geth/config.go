package root

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("fail to load the environment variables")
	}

	os.Setenv("ACCOUNT_ADDRESS", "0xE36AE64156db78dd4797864E9A2f3C1C40625BF3")
}

func GetConfig(property string) string {
	return os.Getenv(property)
}
