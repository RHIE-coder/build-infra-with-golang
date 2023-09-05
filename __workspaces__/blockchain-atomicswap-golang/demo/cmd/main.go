package main

import (
	"fmt"
	"log"
	demo "singaporedemo"
)

const GANACHE_URL = "http://192.168.100.73:10545"
const DEPLOYER = "0x872d3d0d6C5c1C0f5E8f9EEc2c4236cc9b5AB823"
const USER = "0xd5a38dD251Aa8493C03954268FF851290051E634"

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		panic(r)
	// 	}
	// }()

	log.Println("client init")
	client, err := demo.NewClient(GANACHE_URL)
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Println("client Test : Nonce Check")
	fmt.Println(client.PendingNonceAt(DEPLOYER))

	log.Println("call ERC20 : balanceOf")
	point := demo.NewERC20Controller(client, demo.LOCAL_POINT_ADDR)
	balance, err := point.BalanceOf(DEPLOYER)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(balance)
}
