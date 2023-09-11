package demo

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type ERC20 interface {
	Name() string
	GetName() string
	Symbol() (string, error)
	GetSymbol() string
	Decimals() (string, error)
	GetDecimals() string
	Approve(string, string) (bool, error)
	Allowance(string, string) (string, error)
	Transfer(string, string) (bool, error)
	ContractAddress() string
}

type ERC20Contract struct {
	contractAddress common.Address
	abi             abi.ABI
	name            string
	symbol          string
	decimals        string
}

func NewERC20Contract() *ERC20Contract {
	abiJson, _ := abi.JSON(strings.NewReader(ERC20_ABI))
	return &ERC20Contract{
		abi: abiJson,
	}
}

func (erc20 *ERC20Contract) SetAddress(contractAddress string) *ERC20Contract {
	erc20.contractAddress = common.HexToAddress(contractAddress)
	return erc20
}

func (erc20 *ERC20Contract) GetAddress() string {
	return erc20.contractAddress.Hex()
}

func (erc20 *ERC20Contract) SetMetaData(name string, symbol string, decimals string) *ERC20Contract {
	erc20.name = name
	erc20.symbol = symbol
	erc20.decimals = decimals
	return erc20
}

func (erc20 *ERC20Contract) GetMetaData() (string, string, string) {
	return erc20.name, erc20.symbol, erc20.decimals
}

func (erc20 *ERC20Contract) Name() (*ethereum.CallMsg, error) {
	methodName := "name"
	inputBytes, err := erc20.abi.Pack(methodName)
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &erc20.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (erc20 *ERC20Contract) Symbol() (*ethereum.CallMsg, error) {
	methodName := "symbol"
	inputBytes, err := erc20.abi.Pack(methodName)
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &erc20.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (erc20 *ERC20Contract) Decimals() (*ethereum.CallMsg, error) {
	methodName := "decimals"
	inputBytes, err := erc20.abi.Pack(methodName)
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &erc20.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (erc20 *ERC20Contract) BalanceOf(targetAddress string) (*ethereum.CallMsg, error) {
	methodName := "balanceOf"
	inputBytes, err := erc20.abi.Pack(methodName, common.HexToAddress(targetAddress))
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &erc20.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (erc20 *ERC20Contract) Approve(spender string, amount string) (*ethereum.CallMsg, error) {
	methodName := "approve"
	biAmount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, fmt.Errorf("fail to parse the amount")
	}
	inputBytes, err := erc20.abi.Pack(methodName, common.HexToAddress(spender), biAmount)
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &erc20.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (erc20 *ERC20Contract) Allowance(owner string, spender string) (*ethereum.CallMsg, error) {
	methodName := "allowance"
	inputBytes, err := erc20.abi.Pack(methodName, common.HexToAddress(owner), common.HexToAddress(spender))
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &erc20.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (erc20 *ERC20Contract) Transfer(to string, amount string) (*ethereum.CallMsg, error) {
	methodName := "transfer"
	biAmount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, fmt.Errorf("fail to parse the amount")
	}
	inputBytes, err := erc20.abi.Pack(methodName, common.HexToAddress(to), biAmount)
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &erc20.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}
