package main

import (
	"fmt"
	"math/big"
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
	// client, err := ethclient.Dial(GANACHE_URL)
	// Check(err)
	// abiJson, _ := abi.JSON(strings.NewReader(demo.ERC20_ABI))
	amount := "0.01"
	decimals := "18"

	amountBigFlt := new(big.Float)
	_, _, err := amountBigFlt.Parse(amount, 10)
	if err != nil {
		// return
	}

	decimalsBig := new(big.Int)
	_, isSuccess := decimalsBig.SetString(decimals, 10)
	if !isSuccess {
		// return
	}

	decimalsBig.Exp(big.NewInt(10), decimalsBig, nil)

	// fmt.Println(amountBig)
	fmt.Println(decimalsBig)
}
