package demo

import (
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
