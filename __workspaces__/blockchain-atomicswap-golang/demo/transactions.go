package demo

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EIP1559Transaction struct {
	ChainID   *big.Int
	Nonce     uint64
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Gas       uint64
	Value     *big.Int
	To        *common.Address
	Data      []byte
}
