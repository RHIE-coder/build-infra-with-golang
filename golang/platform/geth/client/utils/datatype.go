package utils

import (
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func ConvertStringToAddress(address string) common.Address {
	return common.HexToAddress(address)
}

func ConvertAddressToString(accountAddress common.Address) string {
	return accountAddress.Hex()
}

func ConvertUint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func ConvertStringToUint64(numStr string) uint64 {
	num, err := strconv.ParseUint(numStr, 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return num
}

func ConvertBigintToString(bigNum *big.Int) string {
	return bigNum.String()
}

func ConvertStringToBigint(bigNumStr string) *big.Int {
	num := &big.Int{}
	num, ok := num.SetString(bigNumStr, 10)
	if !ok {
		log.Fatal("cannot convert string to big.Int type")
	}

	return num
}
