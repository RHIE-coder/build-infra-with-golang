package demo

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EIP1559Tx struct {
	ChainID   *big.Int
	From      common.Address // the sender of the 'transaction'
	To        common.Address // the destination contract (nil for contract creation)
	Gas       uint64         // if 0, the call executes with near-infinite gas
	GasPrice  *big.Int       // wei <-> gas exchange ratio
	GasFeeCap *big.Int       // EIP-1559 fee cap per gas.
	GasTipCap *big.Int       // EIP-1559 tip per gas.
	Value     *big.Int       // amount of wei sent along with the call
	Data      []byte         // input data, usually an ABI-encoded contract method invocation
	Nonce     uint64         // the transaction count of account
}

func (tx EIP1559Tx) GetMsgToCall() ethereum.CallMsg {

	return ethereum.CallMsg{
		From:  tx.From,
		To:    &tx.To,
		Value: tx.Value,
		Data:  tx.Data,
	}

}

func (tx EIP1559Tx) GetTxForSigning() *types.Transaction {

	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   tx.ChainID,
		Nonce:     tx.Nonce,
		GasTipCap: tx.GasTipCap,
		GasFeeCap: tx.GasFeeCap,
		Gas:       tx.Gas,
		Value:     tx.Value,
		To:        &tx.To,
		Data:      nil,
	})

}
