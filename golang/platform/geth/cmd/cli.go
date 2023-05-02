package main

import (
	"fmt"
	config "golang/platform/geth"
	"golang/platform/geth/client"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetInfuraURL(network string, apiKey string) string {
	return fmt.Sprintf("https://%s.infura.io/v3/%s", network, apiKey)
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("invalid arguments")
	}
	config.LoadConfig()
	client, err := ethclient.Dial(
		client.GetInfuraURL(
			config.GetConfig("NETWORK"),
			config.GetConfig("INFURA_API_KEY"),
		),
	)

	if err != nil {
		panic(err)
	}

	switch args {
		case ""
	}
	// var command string
	// flag.StringVar(&command, "", "", "execution entry point")

	// flag.Parse()
	// fmt.Println(flag.Lookup("").Value)
	// flag 값이 없으면 에러 처리
	// if flag.Lookup("").Value.String() == "-1" {
	// flag.PrintDefaults()
	// fmt.Println("port is required")
	// return
	// }
	// // Create a new transaction
	// tx := types.NewTransaction(
	// 	0,                      // nonce
	// 	common.HexToAddress("0x123..."),  // to address
	// 	big.NewInt(10000000000), // value
	// 	21000,                  // gas limit
	// 	big.NewInt(20000000000), // gas price
	// 	[]byte("hello"),        // data
	// )

	// // Sign the transaction
	// privateKey, _ := crypto.HexToECDSA("123...")
	// signedTx, _ := types.SignTx(tx, types.HomesteadSigner{}, privateKey)

	// // Convert the signed transaction to hex format
	// serializedTx, _ := rlp.EncodeToBytes(signedTx)
	// hexEncodedTx := hex.EncodeToString(serializedTx)
	// fmt.Println(hexEncodedTx)
}
