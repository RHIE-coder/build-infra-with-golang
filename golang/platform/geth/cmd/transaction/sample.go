package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Ethereum 클라이언트에 연결
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		fmt.Println("Error: cannot connect to the Ethereum network:", err)
		return
	}

	// 보낼 계정의 개인 키
	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY")
	if err != nil {
		fmt.Println("Error: invalid private key:", err)
		return
	}

	// 수신자 주소
	toAddress := common.HexToAddress("RECEIVER_ADDRESS")

	// 가스 가격 및 한도 설정
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Error: cannot get suggested gas price:", err)
		return
	}
	gasLimit := uint64(21000)

	// 전송할 이더 금액 설정
	value := big.NewInt(1000000000000000000) // 1 ETH

	// Nonce 얻기
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("Error: cannot get nonce:", err)
		return
	}

	// Transaction 생성
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// Transaction 서명
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(nil), privateKey)
	if err != nil {
		fmt.Println("Error: cannot sign transaction:", err)
		return
	}

	// Raw Transaction Hex 값 생성
	rawTxBytes := signedTx.Marshal()
	rawTxHex := hex.EncodeToString(rawTxBytes)

	// Raw Transaction 전송
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("Error: cannot send transaction:", err)
		return
	}

	fmt.Println("Transaction sent. Hash:", signedTx.Hash().Hex())
}
