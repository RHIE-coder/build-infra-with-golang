package demo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
	dialer *ethclient.Client
}

func NewClient(url string) (*EthereumClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &EthereumClient{
		dialer: client,
	}, nil
}

func (client *EthereumClient) GetClient() *ethclient.Client {
	return client.dialer
}

func (client *EthereumClient) EstimateGas(msg ethereum.CallMsg) {
	gas, err := client.dialer.EstimateGas(context.Background(), msg)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(gas)
}

func (client *EthereumClient) PendingNonceAt(address string) (string, error) {
	isValid := common.IsHexAddress(address)

	if !isValid {
		return "", fmt.Errorf("address is not valid")
	}

	nonce, err := client.dialer.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return "", fmt.Errorf("fail to jsonrpc request")
	}

	return strconv.FormatUint(nonce, 10), nil
}
