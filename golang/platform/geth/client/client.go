package client

import (
	"context"
	"golang/platform/geth/client/utils"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
	dialer *ethclient.Client
}

func NewClient(uri string) *EthereumClient {
	client, err := ethclient.Dial(uri)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &EthereumClient{
		dialer: client,
	}
}

func (client *EthereumClient) ChainId() string {

	chainId, err := client.dialer.ChainID(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	return utils.ConvertBigintToString(chainId)
}

func (client *EthereumClient) BlockNumber() string {

	blockNum, err := client.dialer.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	return utils.ConvertUint64ToString(blockNum)
}

func (client *EthereumClient) BalanceAt(address string) string {
	isValid := utils.IsValidAddress(address)
	if !isValid {
		log.Fatal("invalid address")
	}
	accountAddress := utils.ConvertStringToAddress(address)

	balance, err := client.dialer.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	return utils.ConvertBigintToString(balance)
}

func (client *EthereumClient) PendingNonceAt(address string) string {
	isValid := utils.IsValidAddress(address)
	if !isValid {
		log.Fatal("invalid address")
	}
	accountAddress := utils.ConvertStringToAddress(address)
	nonce, err := client.dialer.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
	return utils.ConvertUint64ToString(nonce)
}

func (client *EthereumClient) SuggestGasPrice() string {
	gasPrice, err := client.dialer.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	return utils.ConvertBigintToString(gasPrice)
}
