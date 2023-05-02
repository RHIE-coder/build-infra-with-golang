package main

import (
	"encoding/json"
	"fmt"
	config "golang/platform/geth"
	"golang/platform/geth/client"
	"golang/platform/geth/client/utils"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

	for _, path := range []string{"../../contract", "../contract", "./contract", "."} {
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

	contractAddress := common.HexToAddress(contractAddressString)
	console("contractAddress", contractAddress)

	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal(err.Error())
	}
	console("abi.JSON", abi)

	inputData, err := abi.Pack("getValue")
	if err != nil {
		log.Fatal(err.Error())
	}
	console("inputData", inputData)
	console("inputData(HEX)", hexutil.Encode(inputData))
	// Set up a new call message to the `getValue` function of the contract
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: inputData,
	}

	// Call the `eth_call` method to execute the call message and get the result
	callData := infuraClient.CallContract(&callMsg)
	console("callData", callData)
	console("callData", hexutil.Encode(callData))

	value, err := utils.HexStringToUint64("0x0000000000000000000000000000000000000000000000000000000000000029")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
