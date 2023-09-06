package demo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type AdminAccount struct {
	address    string
	privateKey string
	rpc        *EthereumClient
}

func NewAccount(client *EthereumClient, address string, privateKey string) *AdminAccount {
	return &AdminAccount{
		address:    address,
		privateKey: privateKey,
		rpc:        client,
	}
}

func (admin *AdminAccount) PendingNonceAt() (string, error) {
	isValid := common.IsHexAddress(admin.address)

	if !isValid {
		return "", fmt.Errorf("address is not valid")
	}

	nonce, err := admin.rpc.dialer.PendingNonceAt(context.Background(), common.HexToAddress(admin.address))
	if err != nil {
		return "", fmt.Errorf("fail to jsonrpc request")
	}

	return strconv.FormatUint(nonce, 10), nil
}
