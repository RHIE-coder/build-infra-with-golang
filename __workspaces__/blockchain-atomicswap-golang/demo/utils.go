package demo

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func ParseDecimalStringToIntegerString(amount string, decimals string) (string, error) {
	amountBigFlt := new(big.Float)
	_, _, err := amountBigFlt.Parse(amount, 10)
	if err != nil {
		return "", fmt.Errorf("amount should be number format")
	}

	decimalsBig := new(big.Int)
	_, isSuccess := decimalsBig.SetString(decimals, 10)
	if !isSuccess {
		return "", fmt.Errorf("need check decimals value")
	}

	decimalsBig.Exp(big.NewInt(10), decimalsBig, nil)
	result := new(big.Float)
	result.Mul(amountBigFlt, new(big.Float).SetInt(decimalsBig))

	return result.Text('f', 0), nil
}

func ParseIntegerStringToDecimalString(amount string, decimals string) (string, error) {
	amountBigInt, isSuccess := new(big.Int).SetString(amount, 10)
	if !isSuccess {
		return "", fmt.Errorf("amount should be number format")
	}

	decimalsNum, err := strconv.ParseInt(decimals, 10, 64)
	if err != nil {
		return "", fmt.Errorf("need check decimals value")
	}

	scaleFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimalsNum), nil)
	result := new(big.Float).Quo(new(big.Float).SetInt(amountBigInt), new(big.Float).SetInt(scaleFactor))
	str := result.Text('f', int(decimalsNum))
	return strings.TrimRight(str, "0"), nil
}

func EthToWei(amount string) (string, error) {
	return ParseDecimalStringToIntegerString(amount, "18")
}

func EthToGwei(amount string) (string, error) {
	return ParseDecimalStringToIntegerString(amount, "9")
}

func WeiToEth(amount string) (string, error) {
	return ParseIntegerStringToDecimalString(amount, "18")
}

func GweiToEth(amount string) (string, error) {
	return ParseIntegerStringToDecimalString(amount, "9")
}

func ContractIntegerBytesToString(bytes []byte) string {
	return new(big.Int).SetBytes(bytes).String()
}

func ContractStringBytesToString(bytes []byte) string {
	return string(bytes)
}

func ContractAddressBytes(bytes []byte) common.Address {
	return common.BytesToAddress(bytes)
}
