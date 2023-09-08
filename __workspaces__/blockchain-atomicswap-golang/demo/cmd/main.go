package main

import (
	"context"
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

func main() {
	test()
}
func WeiToEth(wei *big.Int) *big.Float {
	weiInEth := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(weiInEth, big.NewFloat(1e18))
	return ethValue
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

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainId), privateKey)
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
