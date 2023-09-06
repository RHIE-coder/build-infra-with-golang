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
	abiJson, _ := abi.JSON(strings.NewReader(ERC20_ABI))

	erc20 := ERC20Contract{
		contractAddress: address,
		abi:             abiJson,
		rpc:             client,
		admin:           admin,
	}

	return erc20
}

func Name() string {
	return ""
}

func (erc20 *ERC20Contract) Approve(senderAddress string, amount string) error {
	amountBig := new(big.Int)
	amountBig.SetString(amount, 10)
	decimalsBig := new(big.Int)
	decimalsBig.SetString(erc20.decimals, 10)
	decimalsBig.Exp(big.NewInt(10), decimalsBig, nil)

	// inputBytes, err := erc20.abi.Pack("approve", common.HexToAddress(senderAddress))
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
