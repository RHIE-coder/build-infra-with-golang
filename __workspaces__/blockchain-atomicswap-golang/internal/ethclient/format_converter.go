package ethclient

import (
	"fmt"
	"testing"
)

// import (
// 	"log"
// 	"math/big"
// 	"strconv"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/common/hexutil"
// )

func ConvertUintToString(num interface{}) (string, error) {
	switch num.(type) {
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	default:
		fmt.Errorf("")
	}
	// return strconv.FormatUint(num, 10), nil
	return "", nil
}

// func ConvertStringToAddress(address string) common.Address {
// 	return common.HexToAddress(address)
// }

// func ConvertAddressToString(accountAddress common.Address) string {
// 	return accountAddress.Hex()
// }

// func ConvertStringToUint64(numStr string) uint64 {
// 	num, err := strconv.ParseUint(numStr, 10, 64)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	return num
// }

// func ConvertBigintToString(bigNum *big.Int) string {
// 	return bigNum.String()
// }

// func ConvertStringToBigint(bigNumStr string) *big.Int {
// 	num := &big.Int{}
// 	num, ok := num.SetString(bigNumStr, 10)
// 	if !ok {
// 		log.Fatal("cannot convert string to big.Int type")
// 	}

// 	return num
// }

// func HexStringToUint64(hexStr string) uint64 {
// 	var bigInt big.Int
// 	// bytes, err := hex.DecodeString(hexStr[2:])
// 	bytes, err := hexutil.Decode(hexStr)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	bigInt.SetBytes(bytes)
// 	return bigInt.Uint64()
// }

// func HexToUintString(input string) string {
// 	value := HexToUint64(input)
// 	return strconv.FormatUint(value, 10)

// }

// func HexToUint64(input string) uint64 {
// 	if has0xPrefix(input) {
// 		result, err := strconv.ParseUint(input[2:], 16, 64)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		return result
// 	} else {
// 		return 0
// 	}
// }

// func has0xPrefix(input string) bool {
// 	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
// }

func Test(t *testing.T) {
	var a interface{}

	a.(type)
}
