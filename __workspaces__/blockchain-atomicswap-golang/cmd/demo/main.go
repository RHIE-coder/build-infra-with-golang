package main

import (
	"os"
	demo "singaporedemo"
)

const network = "ganache"

// const network = "alchemy-optimism-goerli"

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// // url := ParseEndPoint(network)

	// cmd := os.Args[1]
	// controller := demo.NewAtomicSwapDispacher(
	// 	os.Getenv("LOCAL_TOKEN_ORIGIN"),
	// 	os.Getenv("LOCAL_TOKEN_DESTINATION"),
	// 	os.Getenv("LOCAL_ATOMICSWAP_CONTRACT_ORIGIN"),
	// 	os.Getenv("LOCAL_ATOMICSWAP_CONTRACT_DESTINATION"),
	// )

	// client, err := ethclient.Dial(ParseEndPoint(network))

	// if err != nil {
	// 	fmt.Println("init client error")
	// }

	// switch cmd {
	// case "info":
	// 	controller.Info()
	// case "chainid":
	// 	fmt.Println()
	// }
}

func ParseEndPoint(network string) string {
	switch network {
	case "ganache":
		return os.Getenv("GANACHE_RPC_ENDPOINT")
	case "alchemy-optimism-goerli":
		apiKey := os.Getenv("ALCHEMY_OP_GOERLI_API_KEY")
		return demo.GetAlchemyUrl("optimism-goerli", apiKey)
	}

	return os.Getenv("GANACHE_RPC_ENDPOINT")
}
