package main

import (
	"encoding/json"
	"fmt"
	config "golang/platform/geth"
	"golang/platform/geth/client"
	"golang/platform/geth/client/utils"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func console(topic string, value interface{}) {
	fmt.Println(fmt.Sprintf("[ %s ]: %+v", topic, value))
}

func main() {

	isFileFound := false
	filename := "SimpleStorage.json"
	var metadata []byte
	var err error

	for _, path := range []string{"../../contract", "./contract", "."} {
		metadata, err = os.ReadFile(path + "/" + filename)

		if err != nil {
			continue
		}

		isFileFound = true
		break
	}

	if !isFileFound {
		log.Fatal(err.Error())
	}

	var jsondata map[string]interface{}
	err = json.Unmarshal(metadata, &jsondata)

	if err != nil {
		log.Fatal(err.Error())
	}

	byteToPrint, _ := json.MarshalIndent(jsondata, "", "    ")
	console("FILE DATA", string(byteToPrint))

	abiObject, err := json.Marshal(jsondata["abi"])
	abiString := string(abiObject)
	console("abi", abiString)

	contractAddressObject, ok := jsondata["address"].(map[string]interface{})
	if !ok {
		log.Fatal("invalid type assertion")
	}

	contractAddressString, ok := contractAddressObject["sepolia"].(string)
	if !ok {
		log.Fatal("invalid type assertion")
	}
	console("address", contractAddressString)

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

	// Ignore [start]
	// toAddress := common.HexToAddress("0xa49EAcDaDeF57F4ABC4d52D528945CE4c3834293")

	// 전송할 이더의 양을 설정합니다.   1 ETH = 10^18 wei
	// value := big.NewInt(100000000000000) // 0.0001 ETH
	// console("ether value", value)
	// Ignore [end]

	// 가스 한도 및 가스 가격을 설정합니다.
	// gasLimit := uint64(100000) // 기본 가스 한도
	gasPriceStr := infuraClient.SuggestGasPrice()
	if err != nil {
		log.Fatal(err.Error())
	}
	gasPrice := utils.ConvertStringToBigint(gasPriceStr)
	console("gasPrice", gasPrice)

	contractAddress := common.HexToAddress(contractAddressString)

	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal(err.Error())
	}
	console("abi.JSON", abi)

	inputData, err := abi.Pack("setValue", big.NewInt(41))
	if err != nil {
		log.Fatal(err.Error())
	}
	console("inputData", inputData)

	expectedGasStr := infuraClient.EstimateGas(&ethereum.CallMsg{
		From: fromAddress,
		To:   &contractAddress,
		Data: inputData,
	})
	console("expected gas", expectedGasStr)

	BN_gl := utils.ConvertStringToBigint(expectedGasStr)
	BN_gp := gasPrice

	mulResult := new(big.Int).Mul(BN_gl, BN_gp)
	console("gasPrice * gasLimit", mulResult)

	// 트랜잭션을 생성합니다.
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      BN_gl.Uint64(),
		To:       &contractAddress,
		Value:    big.NewInt(0),
		Data:     inputData,
		// ChainID:  big.NewInt(1), // or any other network ID
	})

	// 트랜잭션에 서명합니다.
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		panic(err)
	}

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

// func sendTransaction() {
// 	// nonce 값을 가져옵니다.
// 	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 스마트 컨트랙트 주소를 설정합니다.
// 	contractAddress := common.HexToAddress("CONTRACT_ADDRESS")

// 	abi, err := abi.JSON(strings.NewReader(CONTRACT_ABI))
// 	if err != nil {
// 		panic(err)
// 	}
// 	input, err := abi.Pack("setValue", big.NewInt(41))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 트랜잭션을 생성합니다.
// 	tx := types.NewTransaction(nonce, contractAddress, nil, gasLimit, gasPrice, input)

// }
