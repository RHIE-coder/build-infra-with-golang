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

	os.Setenv("ACCOUNT_OWEN", "0x2894706debA1DF71735053E8f55f65D34348c051")
	os.Setenv("ACCOUNT_ALICE", "0xa49EAcDaDeF57F4ABC4d52D528945CE4c3834293")
}

func GetConfig(property string) string {
	return os.Getenv(property)
}
