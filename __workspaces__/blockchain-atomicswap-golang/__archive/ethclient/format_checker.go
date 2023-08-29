package ethclient

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func GetInfuraURL(network string, apiKey string) string {
	return fmt.Sprintf("https://%s.infura.io/v2/%s", network, apiKey)
}

func IsValidAddress(address string) bool {
	return common.IsHexAddress(address)
}
