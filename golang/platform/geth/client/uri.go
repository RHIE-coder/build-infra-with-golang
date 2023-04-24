package client

import "fmt"

func GetInfuraURL(network string, apiKey string) string {
	return fmt.Sprintf("https://%s.infura.io/v3/%s", network, apiKey)
}
