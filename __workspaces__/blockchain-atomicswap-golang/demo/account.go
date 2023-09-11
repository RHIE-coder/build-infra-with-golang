package demo

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthereumAccount struct {
	privateKey          *ecdsa.PrivateKey
	publicKey           *ecdsa.PublicKey
	mnemonic            string
	path                string
	index               int
	address             string
	isCreatedByMnemonic bool
}

func NewAccountFromPrivateKey(privateKeyHex string) (*EthereumAccount, error) {
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
	return &EthereumAccount{
		privateKey:          privateKeyECDSA,
		publicKey:           publicKeyECDSA,
		mnemonic:            "",
		path:                "",
		index:               -1,
		address:             address,
		isCreatedByMnemonic: false,
	}, nil
}

func NewAccountFromMnemonic(mnemonic string, index int, path string) (*EthereumAccount, error) {
	return nil, nil
}

func (account *EthereumAccount) GetPrivateKey() string {
	return hex.EncodeToString(crypto.FromECDSA(account.privateKey))
}

func (account *EthereumAccount) GetPublicKey() string {
	return hex.EncodeToString(crypto.FromECDSAPub(account.publicKey))
}

func (account *EthereumAccount) SignTx(tx EIP1559Tx) (*types.Transaction, error) {
	signedTx, err := types.SignTx(
		tx.GetTxForSigning(),
		types.NewCancunSigner(tx.ChainID),
		account.privateKey,
	)

	if err != nil {
		return nil, fmt.Errorf("fail to sign the transaction")
	}

	return signedTx, nil
}
