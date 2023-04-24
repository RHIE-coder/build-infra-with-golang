package main

import (
	"fmt"
	config "golang/platform/geth"
	"golang/platform/geth/client"
	"golang/platform/geth/client/utils"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func console(topic string, value interface{}) {
	fmt.Println(fmt.Sprintf("[ %s ]: %+v", topic, value))
}

func main() {
	config.LoadConfig()
	infuraClient := client.NewClient(
		client.GetInfuraURL(
			config.GetConfig("NETWORK"),
			config.GetConfig("INFURA_API_KEY"),
		),
	)

	privateKey, err := crypto.HexToECDSA(
		config.GetConfig("PRIVATE_KEY"),
	)
	console("privateKey", privateKey)
	if err != nil {
		log.Fatal(err.Error())
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	console("fromAddress", fromAddress) //0xE36AE64156db78dd4797864E9A2f3C1C40625BF3

	nonceStr := infuraClient.PendingNonceAt(
		utils.ConvertAddressToString(fromAddress),
	)

	nonce := utils.ConvertStringToUint64(nonceStr)

	toAddress := common.HexToAddress("0x1Bc8D4d2A7069965CA0436667903aF4cf0f3A144")

	// 전송할 이더의 양을 설정합니다.   1 ETH = 10^18 wei
	value := big.NewInt(1000000000000000) // 0.001 ETH
	console("ether value", value)

	// 가스 한도 및 가스 가격을 설정합니다.
	gasLimit := uint64(21000) // 기본 가스 한도
	gasPriceStr := infuraClient.SuggestGasPrice()
	if err != nil {
		log.Fatal(err.Error())
	}
	gasPrice := utils.ConvertStringToBigint(gasPriceStr)
	console("gasPrice", gasPrice)

	// 트랜잭션을 생성합니다.
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// 트랜잭션에 서명합니다.
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		panic(err)
	}

	// raw 트랜잭션을 생성합니다.
	rawTxBytes, err := signedTx.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Printf("raw transaction: %s\n", string(rawTxBytes))
}
