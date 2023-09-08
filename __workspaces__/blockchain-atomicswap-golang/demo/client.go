package demo

import "github.com/ethereum/go-ethereum/ethclient"

type Provider struct {
	dialer *ethclient.Client
}

func NewProvider(url string) (*Provider, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Provider{
		dialer: client,
	}, nil
}

func (client *Provider) GetClient() *ethclient.Client {
	return client.dialer
}
