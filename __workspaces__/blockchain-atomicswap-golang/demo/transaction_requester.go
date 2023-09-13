package demo

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionRequester struct {
	provider *Provider
	signer   *EthereumAccount
}

func NewTransactionRequester(provider *Provider) *TransactionRequester {
	return &TransactionRequester{
		provider: provider,
	}
}

func (txRequester *TransactionRequester) SetSigner(account *EthereumAccount) *TransactionRequester {
	txRequester.signer = account
	return txRequester
}

func (txRequester *TransactionRequester) Call(msg ethereum.CallMsg) ([]byte, error) {
	rpc := txRequester.provider.GetClient()
	return rpc.CallContract(context.Background(), msg, nil)
}

func (txRequester *TransactionRequester) SignTransaction(msg ethereum.CallMsg) (*types.Transaction, error) {
	rpc := txRequester.provider.GetClient()
	signer := txRequester.signer

	if signer == nil {
		return nil, fmt.Errorf("the signer is not assigned")
	}
	from := signer.GetAddress()

	msg.From = from

	chainId, err := rpc.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("fail to get chain id: %s", err.Error())
	}

	nonce, err := rpc.PendingNonceAt(context.Background(), signer.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("fail to get nonce: %s", err.Error())
	}

	gasLimit, err := rpc.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, fmt.Errorf("fail to estimate gas usage: %s", err.Error())
	}
	gasPrice, err := rpc.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("fail to get gas price: %s", err.Error())
	}
	tipCap, err := rpc.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, fmt.Errorf("fail to get gas tip cap: %s", err.Error())
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: tipCap,
		GasFeeCap: gasPrice,
		Gas:       gasLimit,
		Value:     msg.Value,
		To:        msg.To,
		Data:      msg.Data,
	})

	signedTx, err := signer.Sign(chainId, tx)
	if err != nil {
		return nil, fmt.Errorf("fail to sign transaction: %s", err.Error())
	}

	return signedTx, nil
}

func (txRequester *TransactionRequester) SendRawTransaction(signedTx *types.Transaction) error {
	rpc := txRequester.provider.GetClient()
	err := rpc.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("fail to call contract: " + err.Error())
	}
	return nil
}

func (txRequester *TransactionRequester) SendTransaction(msg ethereum.CallMsg) error {
	signedTx, err := txRequester.SignTransaction(msg)
	if err != nil {
		return fmt.Errorf("the SignTransaction is failed: %s", err.Error())
	}
	err = txRequester.SendRawTransaction(signedTx)
	if err != nil {
		return fmt.Errorf("the SendRawTransaction is failed: %s", err.Error())
	}
	return nil
}
