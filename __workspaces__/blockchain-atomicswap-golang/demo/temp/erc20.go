package demo

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type ERC20Contract struct {
	contractAddress string
	abi             abi.ABI
	rpc             *EthereumClient
	admin           AdminAccount
	name            string
	symbol          string
	decimals        string
}

func NewERC20Controller(client *EthereumClient, address string, admin AdminAccount) ERC20Contract {
	abiJson, _ := abi.JSON(strings.NewReader(""))

	erc20 := ERC20Contract{
		contractAddress: address,
		abi:             abiJson,
		rpc:             client,
		admin:           admin,
	}

	erc20.Name()
	erc20.Symbol()
	erc20.Decimals()

	return erc20
}

func (erc20 ERC20Contract) Name() (string, error) {
	if erc20.name != "" {
		return erc20.name, nil
	}

	inputBytes, err := erc20.abi.Pack("name")

	if err != nil {
		return "", fmt.Errorf("fail to packing 'name' methods")
	}

	contractAddress := common.HexToAddress(erc20.contractAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: inputBytes,
	}
	nameBytes, err := erc20.rpc.GetClient().CallContract(context.Background(), callMsg, nil)

	if err != nil {
		return "", fmt.Errorf("fail to get name")
	}

	erc20.name = string(nameBytes)

	return erc20.name, nil
}

func (erc20 ERC20Contract) Symbol() (string, error) {

	if erc20.symbol != "" {
		return erc20.symbol, nil
	}

	inputBytes, err := erc20.abi.Pack("symbol")

	if err != nil {
		return "", fmt.Errorf("fail to packing 'symbol' methods")
	}

	contractAddress := common.HexToAddress(erc20.contractAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: inputBytes,
	}
	symbolBytes, err := erc20.rpc.GetClient().CallContract(context.Background(), callMsg, nil)

	if err != nil {
		return "", fmt.Errorf("fail to get symbol")
	}

	erc20.symbol = string(symbolBytes)

	return erc20.symbol, nil
}

func (erc20 ERC20Contract) Decimals() (string, error) {
	if erc20.decimals != "" {
		return erc20.decimals, nil
	}
	inputBytes, err := erc20.abi.Pack("decimals")

	if err != nil {
		return "", fmt.Errorf("fail to packing 'decimals' methods")
	}

	contractAddress := common.HexToAddress(erc20.contractAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: inputBytes,
	}
	decimalsBytes, err := erc20.rpc.GetClient().CallContract(context.Background(), callMsg, nil)

	if err != nil {
		return "", fmt.Errorf("fail to get decimals")
	}

	erc20.decimals = new(big.Int).SetBytes(decimalsBytes).String()

	return erc20.decimals, nil
}

func (erc20 *ERC20Contract) EstimateGas(fromAddress string, data []byte) {
	contractAddress := common.HexToAddress(erc20.contractAddress)
	gas, err := erc20.rpc.GetClient().EstimateGas(context.Background(), ethereum.CallMsg{
		From: common.HexToAddress(fromAddress),
		To:   &contractAddress,
		Data: data,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(gas)
}

func (erc20 *ERC20Contract) Approve(senderAddress string, amount string) error {
	parsedAmount, _ := new(big.Int).SetString(amount, 10)
	inputBytes, err := erc20.abi.Pack("approve", common.HexToAddress(senderAddress), parsedAmount)

	if err != nil {
		return fmt.Errorf("fail to packing 'approve' methods : %s", err.Error())
	}

	contractAddress := common.HexToAddress(erc20.contractAddress)
	callMsg := ethereum.CallMsg{
		From: erc20.admin.GetEthAddress(),
		To:   &contractAddress,
		Data: inputBytes,
	}

	erc20.admin.rpc.EstimateGas(callMsg)

	// resultBytes, err := erc20.rpc.GetClient().CallContract(context.Background(), callMsg, nil)

	// if err != nil {
	// 	return fmt.Errorf("fail to invoke approve")
	// }
	// fmt.Println(resultBytes)
	// fmt.Println(string(resultBytes))
	return nil
}

func (erc20 *ERC20Contract) Allowance(ownerAddress string, amount string) (string, error) {
	return "", nil
}

func (erc20 *ERC20Contract) BalanceOf(address string) (string, error) {
	inputBytes, err := erc20.abi.Pack("balanceOf", common.HexToAddress(address))

	if err != nil {
		return "", fmt.Errorf("fail to packing 'balanceOf' methods")
	}

	contractAddress := common.HexToAddress(erc20.contractAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: inputBytes,
	}

	balanceBytes, err := erc20.rpc.GetClient().CallContract(context.Background(), callMsg, nil)

	if err != nil {
		return "", fmt.Errorf("fail to rpc request to call contract")
	}

	balanceBigInt := new(big.Int)
	balanceBigInt.SetBytes(balanceBytes)

	return balanceBigInt.String(), nil
}

func (erc20 *ERC20Contract) Address() string {
	return erc20.contractAddress
}

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
