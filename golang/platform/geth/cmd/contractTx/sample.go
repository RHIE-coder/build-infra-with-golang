package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func sendTransaction() {
	// Ethereum 네트워크에 연결합니다.
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		panic(err)
	}

	// 송신자의 개인 키를 설정합니다.
	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY")
	if err != nil {
		panic(err)
	}

	// nonce 값을 가져옵니다.
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	// 스마트 컨트랙트 주소를 설정합니다.
	contractAddress := common.HexToAddress("CONTRACT_ADDRESS")

	// 스마트 컨트랙트 메서드를 호출하는 데이터를 생성합니다.
	abi, err := abi.JSON(strings.NewReader(CONTRACT_ABI))
	if err != nil {
		panic(err)
	}
	input, err := abi.Pack("setValue", big.NewInt(42))
	if err != nil {
		panic(err)
	}

	// 가스 한도 및 가스 가격을 설정합니다.
	gasLimit := uint64(300000) // 가스 한도는 호출하는 메서드에 따라 다릅니다.
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	// 트랜잭션을 생성합니다.
	tx := types.NewTransaction(nonce, contractAddress, nil, gasLimit, gasPrice, input)

	// 트랜잭션에 서명합니다.
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		panic(err)
	}

	// raw 트랜잭션을 생성합니다.
	rawTxBytes, err := signedTx.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Printf("raw transaction: %s\n", string(rawTxBytes))
}
