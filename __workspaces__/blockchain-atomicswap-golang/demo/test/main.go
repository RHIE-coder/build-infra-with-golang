package main

import (
	"context"
	"fmt"
	"math/big"
	demo "singaporedemo"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

const GANACHE_URL = "http://192.168.100.73:10545"
const DEPLOYER = "0x872d3d0d6C5c1C0f5E8f9EEc2c4236cc9b5AB823"
const USER = "0xd5a38dD251Aa8493C03954268FF851290051E634"

func main() {

	client, err := ethclient.Dial(GANACHE_URL)
	Check(err)
	abiJson, _ := abi.JSON(strings.NewReader(demo.ERC20_ABI))
	inputBytes, err := abiJson.Pack("balanceOf", common.HexToAddress(DEPLOYER))
	Check(err)
	contractAddress := common.HexToAddress(demo.LOCAL_POINT_ADDR)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: inputBytes,
	}
	result, err := client.CallContract(context.Background(), callMsg, nil)
	Check(err)
	// 바이트 배열을 *big.Int로 변환
	resultBigInt := new(big.Int)
	resultBigInt.SetBytes(result)

	// resultBigInt를 사용하여 작업 수행
	fmt.Printf("결과: %s\n", resultBigInt.String())

}
