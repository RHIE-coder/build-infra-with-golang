package demo

import "fmt"

func GetAlchemyUrl(network string, apiKey string) string {
	switch network {
	case "optimism-goerli":
		return fmt.Sprintf("https://opt-goerli.g.alchemy.com/v2/%s", apiKey)

	case "goerli":
		return fmt.Sprintf("https://eth-goerli.g.alchemy.com/v2/%s", apiKey)
	}
	return ""
}
