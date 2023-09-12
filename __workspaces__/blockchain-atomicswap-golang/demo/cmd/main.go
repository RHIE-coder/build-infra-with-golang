package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	demo "singaporedemo"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const GANACHE = "http://192.168.100.73:10545"
const MAINNET = "https://eth-mainnet.g.alchemy.com/v2/Jp75GWh8YtUBSOnK0tU_hh9G4KUOLcQk"
const GOERLI = "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
const ADMIN_DEPLOYER = "0x872d3d0d6C5c1C0f5E8f9EEc2c4236cc9b5AB823"
const USER = "0xd5a38dD251Aa8493C03954268FF851290051E634"
const prikey = "6ff38a6fcde856869ddba8a1e0058a02cf81742f150607507d5245da607ba48f"
const secret = "6d733fa2adc868bcb7f263e2e278125dc88e0f67eea1b106d75e77da91fb730359ca5c362c76c5ab8a0032889acfacc2ed2bc848ad97f43e152af6448d1819a1f3f27627f870658225b401057aa78824ae140aa3e5ba878fca693e6bf110b0f366b5d28de9d44b705783a5cbf51557a90c960d090d24aa33f33650d7c0f81b16"
const secretHash = "a757feec92c62bfd91fcef3b0a16a97650aea0999531a75cbc09192dd2089f1c"

func main() {
	client, err := demo.NewProvider(GANACHE)
	if err != nil {
		panic(err)
	}

	signer, err := demo.NewAccountFromPrivateKey(prikey)
	if err != nil {
		panic(err)
	}

	txRequester := demo.NewTransactionRequester(client)

	prepare(txRequester)
	checkBalanceOf(txRequester)
	// transferOwnerPointToUser(txRequester)
	// allowToPointContract(txRequester)
	// allowToTokenContract(txRequester)
	// checkAllowance(txRequester)
	// createSwapPoint(txRequester)
	// checkTxToKnowSecretHash(txRequester) // Optional Now
	// redeemPoint(txRequester)
	// checkTxToKnowSecret(txRequester) // Optional Now
	// redeemToken(txRequester)
}

func prepare(txRequester *demo.TransactionRequester) {
	log.Println("   ---   prepare   ---")

}

func checkBalanceOf(txRequester *demo.TransactionRequester) {
	log.Println("   ---   checkBalanceOf   ---")

}

func transferOwnerPointToUser(txRequester *demo.TransactionRequester) { // faucet

	log.Println("   ---   transferOwnerPointToUser   ---")

}

func allowToPointContract(txRequester *demo.TransactionRequester) {

	log.Println("   ---   allowToPointContract   ---")
}

func allowToTokenContract(txRequester *demo.TransactionRequester) {

	log.Println("   ---   allowToTokenContract   ---")
}

func checkAllowance(txRequester *demo.TransactionRequester) {

	log.Println("   ---   checkAllowance   ---")
}

func createSwapPoint(txRequester *demo.TransactionRequester) {

	log.Println("   ---   createSwapPoint   ---")
}

func checkTxToKnowSecretHash(txRequester *demo.TransactionRequester) {
	// Optional Now
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

func test2() {
	randomData := make([]byte, 128)

	_, err := rand.Read(randomData)
	if err != nil {
		panic(err)
	}

	hash := crypto.Keccak256Hash(randomData)

	fmt.Printf("Random Data: %x\n", randomData)
	fmt.Printf("Keccak256 Hash: %s\n", hash.Hex())
}

func test() {
	rpc, err := demo.NewProvider(GANACHE)
	if err != nil {
		panic(err)
	}
	signer, err := demo.NewAccountFromPrivateKey(prikey)
	if err != nil {
		panic(err)
	}
	fmt.Println(signer)
	client := rpc.GetClient()

	chainId, _ := client.ChainID(context.Background())
	log.Printf("Chain ID: %s", chainId.String())

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Current Gas Price: %s wei", gasPrice.String())

	capFee, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Current Gas Tip Cap: %s wei", capFee.String())

	sender := common.HexToAddress(ADMIN_DEPLOYER)
	receiver := common.HexToAddress(USER)
	amount, _ := demo.EthToWei("0.01")
	sendAmount, _ := new(big.Int).SetString(amount, 10)

	msg := ethereum.CallMsg{
		// From:  sender,
		To:    &receiver,
		Value: sendAmount,
		Data:  nil,
	}

	gasLimit, _ := client.EstimateGas(context.Background(), msg)
	log.Printf("Current Gas Limit: %d gwei", gasLimit)

	// target := ethereum.CallMsg{
	// 	From:      msg.From,
	// 	To:        msg.To,
	// 	Value:     msg.Value,
	// 	Data:      nil,
	// 	Gas:       gasLimit,
	// 	GasPrice:  gasPrice,
	// 	GasTipCap: capFee,
	// }

	nonce, _ := client.PendingNonceAt(context.Background(), sender)
	log.Printf("Current Nonce: %d", nonce)

	privateKey, _ := crypto.HexToECDSA(strings.TrimPrefix(prikey, "0x"))

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: capFee,
		GasFeeCap: gasPrice,
		Gas:       gasLimit,
		Value:     sendAmount,
		To:        &receiver,
		Data:      nil,
	})

	// signer.SignTx(tx)

	signedTx, err := types.SignTx(tx, types.NewCancunSigner(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Signed Transaction: %s", signedTx.Hash())

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println("-------------------------------------------")
	// amount2, _ := demo.EthToWei("0")
	// sendAmount2, _ := new(big.Int).SetString(amount2, 10)

	// fmt.Println(sendAmount2)
	// log.Printf("X2 price: %s", big.NewInt(0).Mul(gasPrice, big.NewInt(2)).String())
	// tx2 := types.NewTx(&types.DynamicFeeTx{
	// 	ChainID:   chainId,
	// 	Nonce:     nonce,
	// 	GasTipCap: big.NewInt(0).Mul(capFee, big.NewInt(2)),
	// 	GasFeeCap: big.NewInt(0).Mul(gasPrice, big.NewInt(2)),
	// 	Gas:       gasLimit * uint64(2),
	// 	Value:     sendAmount2,
	// 	To:        &receiver,
	// 	Data:      nil,
	// })

	// signedTx2, err := types.SignTx(tx2, types.NewLondonSigner(chainId), privateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("success sign tx")

	// err = client.SendTransaction(context.Background(), signedTx2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Re-Signed Transaction: %s", signedTx2.Hash())
}
