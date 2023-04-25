package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	// Create a new transaction
	tx := types.NewTransaction(
		0,                      // nonce
		common.HexToAddress("0x123..."),  // to address
		big.NewInt(10000000000), // value
		21000,                  // gas limit
		big.NewInt(20000000000), // gas price
		[]byte("hello"),        // data
	)

	// Sign the transaction
	privateKey, _ := crypto.HexToECDSA("123...")
	signedTx, _ := types.SignTx(tx, types.HomesteadSigner{}, privateKey)

	// Convert the signed transaction to hex format
	serializedTx, _ := rlp.EncodeToBytes(signedTx)
	hexEncodedTx := hex.EncodeToString(serializedTx)
	fmt.Println(hexEncodedTx)
}