package ethclient

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	return chainId.String()
}

func (client *EthereumClient) BlockNumber() string {

	blockNum, err := client.dialer.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	return ConvertUintToString(blockNum)
}

func (client *EthereumClient) BalanceAt(address string) string {
	isValid := IsValidAddress(address)
	if !isValid {
		log.Fatal("invalid address")
	}
	accountAddress := common.HexToAddress(address)

	balance, err := client.dialer.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	return balance.String()
}

func (client *EthereumClient) PendingNonceAt(address string) string {
	isValid := IsValidAddress(address)
	if !isValid {
		log.Fatal("invalid address")
	}
	accountAddress := common.HexToAddress(address)
	nonce, err := client.dialer.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
	return ConvertUintToString(nonce)
}

func (client *EthereumClient) SuggestGasPrice() string {
	gasPrice, err := client.dialer.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	return gasPrice.String()
}

func (client *EthereumClient) SendTransaction(tx *types.Transaction) error {
	return client.dialer.SendTransaction(context.Background(), tx)
}

func (client *EthereumClient) EstimateGas(msg *ethereum.CallMsg) string {
	expectedGas, err := client.dialer.EstimateGas(context.Background(), *msg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return ConvertUintToString(expectedGas)
}

func (client *EthereumClient) CallContract(msg *ethereum.CallMsg) []byte {
	callData, err := client.dialer.CallContract(context.Background(), *msg, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	return callData
}
