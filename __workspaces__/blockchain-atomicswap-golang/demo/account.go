package demo

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func (client *EthereumClient) PendingNonceAt(address string) (string, error) {
	isValid := common.IsHexAddress(address)
	if !isValid {
		return "", fmt.Errorf("address is not valid")
	}
	accountAddress := common.HexToAddress(address)
	nonce, err := client.dialer.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
	return strconv.FormatUint(nonce, 10), nil
}
