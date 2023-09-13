package main

import (
	"fmt"
	"log"
	demo "singaporedemo"
)

const GANACHE = "http://192.168.100.73:10545"
const MAINNET = "https://eth-mainnet.g.alchemy.com/v2/Jp75GWh8YtUBSOnK0tU_hh9G4KUOLcQk"
const GOERLI = "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
const ADMIN_DEPLOYER = "0x872d3d0d6C5c1C0f5E8f9EEc2c4236cc9b5AB823"
const USER = "0xd5a38dD251Aa8493C03954268FF851290051E634"
const prikey = "6ff38a6fcde856869ddba8a1e0058a02cf81742f150607507d5245da607ba48f"
const userPrikey = "fc22325bc9636151e15fb14b52f84fde54bfc58e9aa85620828b520cb0a62305"
const secret = "0x6d733fa2adc868bcb7f263e2e278125dc88e0f67eea1b106d75e77da91fb730359ca5c362c76c5ab8a0032889acfacc2ed2bc848ad97f43e152af6448d1819a1f3f27627f870658225b401057aa78824ae140aa3e5ba878fca693e6bf110b0f366b5d28de9d44b705783a5cbf51557a90c960d090d24aa33f33650d7c0f81b16"
const secretHash = "0xa757feec92c62bfd91fcef3b0a16a97650aea0999531a75cbc09192dd2089f1c"

func main() {
	client, err := demo.NewProvider(GANACHE)
	if err != nil {
		panic(err)
	}

	swapBridge, err := demo.NewSwapBridge(
		client,
		demo.LOCAL_POINT_ADDR,
		demo.LOCAL_TOKEN_ADDR,
		demo.LOCAL_ATOMICSWAP_POINT_ADDR,
		demo.LOCAL_ATOMICSWAP_TOKEN_ADDR,
	)

	if err != nil {
		panic(err)
	}

	prepare(swapBridge)
	// checkBalanceOf(swapBridge)
	// transferOwnerPointToUser(swapBridge)
	// allowToPointContract(swapBridge)
	// allowToTokenContract(swapBridge)
	// checkAllowance(swapBridge)
	getSwap(swapBridge)
	getSwapStatus(swapBridge)
	isRedeemed(swapBridge)
	isRefunded(swapBridge)
	// createSwapPoint(swapBridge)
	// checkTxToKnowSecretHash(txRequester) // Optional Now
	// redeemPoint(txRequester)
	// checkTxToKnowSecret(txRequester) // Optional Now
	// redeemToken(txRequester)
}

func getSwap(swapBridge *demo.SwapBridge) {
	hexString, _ := demo.RemoveHexPrefix(secretHash)
	swapData, err := swapBridge.GetSwap(demo.POINT_SWAP_TYPE, hexString)
	if err != nil {
		panic(err)
	}
	fmt.Println(swapData)
}
func getSwapStatus(swapBridge *demo.SwapBridge) {

}
func isRedeemed(swapBridge *demo.SwapBridge) {

}
func isRefunded(swapBridge *demo.SwapBridge) {

}

func prepare(swapBridge *demo.SwapBridge) {
	log.Println("   ---   prepare   ---")
	log.Println(swapBridge.GetMetaData(demo.POINT_SWAP_TYPE))
	log.Println(swapBridge.GetMetaData(demo.TOKEN_SWAP_TYPE))
}

func checkBalanceOf(swapBridge *demo.SwapBridge) {
	log.Println("   ---   checkBalanceOf   ---")
	bal1, err := swapBridge.ERC20BalanceOf(demo.POINT_TYPE, ADMIN_DEPLOYER)
	if err != nil {
		panic(err)
	}
	bal2, err := swapBridge.ERC20BalanceOf(demo.POINT_TYPE, USER)
	if err != nil {
		panic(err)
	}
	bal3, err := swapBridge.ERC20BalanceOf(demo.TOKEN_TYPE, ADMIN_DEPLOYER)
	if err != nil {
		panic(err)
	}
	bal4, err := swapBridge.ERC20BalanceOf(demo.TOKEN_TYPE, USER)
	if err != nil {
		panic(err)
	}
	log.Println(bal1)
	log.Println(bal2)
	log.Println(bal3)
	log.Println(bal4)
}

func transferOwnerPointToUser(swapBridge *demo.SwapBridge) { // faucet
	log.Println("   ---   transferOwnerPointToUser   ---")
	signer, err := demo.NewAccountFromPrivateKey(prikey)
	if err != nil {
		panic(err)
	}
	swapBridge.SetSigner(signer)
	amount, err := demo.EthToWei("10")
	if err != nil {
		panic(err)
	}
	err = swapBridge.ERC20Transfer(demo.POINT_TYPE, USER, amount)
	if err != nil {
		panic(err)
	}

}

func allowToPointContract(swapBridge *demo.SwapBridge) {
	log.Println("   ---   allowToPointContract   ---")
	signer, err := demo.NewAccountFromPrivateKey(userPrikey)
	if err != nil {
		panic(err)
	}
	swapBridge.SetSigner(signer)
	amount, err := demo.EthToWei("10")
	if err != nil {
		panic(err)
	}

	err = swapBridge.ERC20Approve(demo.POINT_TYPE, swapBridge.GetContractAddress(demo.POINT_SWAP_TYPE), amount)
	if err != nil {
		panic(err)
	}
}

func allowToTokenContract(swapBridge *demo.SwapBridge) {
	log.Println("   ---   allowToTokenContract   ---")
	ownerSigner, err := demo.NewAccountFromPrivateKey(prikey)
	if err != nil {
		panic(err)
	}
	swapBridge.SetSigner(ownerSigner)
	amount := "10000000000000000000000"
	err = swapBridge.ERC20Approve(demo.TOKEN_TYPE, swapBridge.GetContractAddress(demo.TOKEN_SWAP_TYPE), amount)
	if err != nil {
		panic(err)
	}
}

func checkAllowance(swapBridge *demo.SwapBridge) {
	log.Println("   ---   checkAllowance   ---")
	allowance1, err := swapBridge.ERC20Allowance(demo.POINT_TYPE, USER, swapBridge.GetContractAddress(demo.POINT_SWAP_TYPE))
	if err != nil {
		panic(err)
	}
	allowance2, err := swapBridge.ERC20Allowance(demo.TOKEN_TYPE, ADMIN_DEPLOYER, swapBridge.GetContractAddress(demo.TOKEN_SWAP_TYPE))
	if err != nil {
		panic(err)
	}

	log.Println(allowance1)
	log.Println(allowance2)
}

func createSwapPoint(swapBridge *demo.SwapBridge) {
	log.Println("   ---   createSwapPoint   ---")
	ownerSigner, err := demo.NewAccountFromPrivateKey(prikey)
	if err != nil {
		panic(err)
	}
	swapBridge.SetSigner(ownerSigner)
	log.Printf("the secret key is: %s", secret)
	hexString, err := demo.RemoveHexPrefix(secret)
	if err != nil {
		panic(err)
	}
	secretBytes, err := demo.HexStringToBytes(hexString)
	if err != nil {
		panic(err)
	}
	secretHash := demo.GetSecretHashFrom(secretBytes)
	log.Printf("the secret hash is: %s", secretHash)
	amount, err := demo.EthToWei("1")
	if err != nil {
		panic(err)
	}

	secrethashHex, _ := demo.RemoveHexPrefix(secretHash)
	err = swapBridge.CreateSwap(
		demo.POINT_SWAP_TYPE,
		USER,
		ADMIN_DEPLOYER,
		secrethashHex,
		amount,
	)
	if err != nil {
		panic(err)
	}
}

func checkTxToKnowSecretHash(txRequester *demo.TransactionRequester) {
	// Optional Now
}

func createSwapToken(swapBridge *demo.SwapBridge) {
	log.Println("   ---   createSwapToken   ---")
}

func redeemPoint(txRequester *demo.TransactionRequester) {

	log.Println("   ---   redeemPoint   ---")
}

func checkTxToKnowSecret(txRequester *demo.TransactionRequester) {
	// Optional Now
}

func redeemToken(txRequester *demo.TransactionRequester) {

	log.Println("   ---   redeemToken   ---")
}
