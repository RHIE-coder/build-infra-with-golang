package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	config "golang/platform/geth"
	"golang/platform/geth/client"
	"golang/platform/geth/client/utils"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func console(topic string, value interface{}) {
	fmt.Println(fmt.Sprintf("[ %s ]: %+v", topic, value))
}

func main() {
	config.LoadConfig()
	infuraClient := client.NewClient(
		client.GetInfuraURL(
			"goerli",
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

	toAddress := common.HexToAddress("0xa49EAcDaDeF57F4ABC4d52D528945CE4c3834293")

	// 전송할 이더의 양을 설정합니다.   1 ETH = 10^18 wei
	value := big.NewInt(100000000000000) // 0.0001 ETH
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
	// tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil) // Deprecated
	// tx := types.NewTx(&types.LegacyTx{
	// 	Nonce:    nonce,
	// 	GasPrice: gasPrice,
	// 	Gas:      gasLimit,
	// 	To:       &toAddress,
	// 	Value:    value,
	// 	Data:     nil,
	// 	// ChainID:  big.NewInt(1), // or any other network ID
	// })

	fmt.Println(nonce)
	fmt.Println(gasLimit)
	fmt.Println(toAddress)
	gp, _ := new(big.Int).SetString(utils.HexToUintString("0x1de14a929a"), 10)
	v, _ := new(big.Int).SetString(utils.HexToUintString("0x9184E72A000"), 10)
	// fmt.Println(utils.ConvertStringToBigint("0x1de14a929a"))
	// fmt.Println(utils.ConvertStringToBigint("0x9184E72A000"))
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    utils.HexStringToUint64("0x30"),
		GasPrice: gp,
		Gas:      utils.HexStringToUint64("0x5208"),
		To:       &toAddress,
		Value:    v,
		Data:     nil,
		// ChainID:  big.NewInt(1), // or any other network ID
	})

	// 트랜잭션에 서명합니다.
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("----")
	// Convert the signed transaction to hex format
	serializedTx, _ := rlp.EncodeToBytes(signedTx)
	hexEncodedTx := hex.EncodeToString(serializedTx)
	fmt.Println(hexEncodedTx)
	return

	// raw 트랜잭션을 생성합니다.
	rawTxBytes, err := signedTx.MarshalJSON()
	if err != nil {
		log.Fatal(err.Error())
	}
	console("rawTxBytes", string(rawTxBytes))

	// Send the transaction
	var txToSend types.Transaction
	err = json.Unmarshal(rawTxBytes, &txToSend)
	if err != nil {
		log.Fatal(err.Error())
	}
	console("txToSend", txToSend)
	err = infuraClient.SendTransaction(&txToSend)
	if err != nil {
		log.Fatal(err)
	}
}
