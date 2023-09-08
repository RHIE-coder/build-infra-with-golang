package demo

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	privateKey          *ecdsa.PrivateKey
	publicKey           *ecdsa.PublicKey
	mnemonic            string
	path                string
	index               int
	address             string
	isCreatedByMnemonic bool
}

func NewAccountFromPrivateKey(privateKeyHex string) (*Account, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("fail to get private key")
	}
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("fail to get public key")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return &Account{
		privateKey:          privateKeyECDSA,
		publicKey:           publicKeyECDSA,
		mnemonic:            "",
		path:                "",
		index:               -1,
		address:             address,
		isCreatedByMnemonic: false,
	}, nil
}

func NewAccountFromMnemonic(mnemonic string, index int, path string) (*Account, error) {
	return nil, nil
}

func (account *Account) GetPrivateKey() string {
	return hex.EncodeToString(crypto.FromECDSA(account.privateKey))
}

func (account *Account) GetPublicKey() string {
	return hex.EncodeToString(crypto.FromECDSAPub(account.publicKey))
}

func (account *Account) SignTx(tx types.TxData) {

}
