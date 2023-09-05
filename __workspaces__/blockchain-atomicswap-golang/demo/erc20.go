package demo

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type ERC20Contract struct {
	contractAddress string
	abi             abi.ABI
	rpc             *EthereumClient
}

func NewERC20Controller(client *EthereumClient, address string) *ERC20Contract {
	abiJson, _ := abi.JSON(strings.NewReader(ERC20_ABI))

	return &ERC20Contract{
		contractAddress: address,
		abi:             abiJson,
		rpc:             client,
	}
}

func (erc20 *ERC20Contract) BalanceOf(address string) string {

}

func (erc20 *ERC20Contract) Address() string {
	return erc20.contractAddress
}
