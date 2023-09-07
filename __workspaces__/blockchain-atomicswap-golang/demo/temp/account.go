package demo

import (
	"github.com/ethereum/go-ethereum/common"
)

type AdminAccount struct {
	address    string
	privateKey string
	rpc        *EthereumClient
}

func NewAccount(client *EthereumClient, address string, privateKey string) AdminAccount {
	return AdminAccount{
		address:    address,
		privateKey: privateKey,
		rpc:        client,
	}
}

func (admin AdminAccount) GetEthAddress() common.Address {
	return common.HexToAddress(admin.address)
}
