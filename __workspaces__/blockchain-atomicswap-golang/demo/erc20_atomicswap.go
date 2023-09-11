package demo

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// type AtomicSwap interface {
// 	Name(SwapTargetERC20) (string, error)
// 	Symbol(SwapTargetERC20) (string, error)
// 	Decimals(SwapTargetERC20) (string, error)
// 	addressOfTargetContract() (string, error)
// 	CreateSwap(swap Swap) error
// 	Redeem(secret []byte, secretHash string) error
// 	Refund(secretHash string) error
// 	GetSwap(secretHash string) Swap
// 	GetSwapStatus(secretHash string) Stage
// 	IsRedeemed(secretHash string)
// 	IsRefunded(secretHash string)
// }

/* Swap */
type Swap struct {
	PoolInitiatorAddress string // sender address
	ReceiverAddress      string // receiver address
	SecretHash           string // keccack256 algorithm
	Amount               string // Amount = amount * decimals
}

/* enum Stage */
type Stage int

const (
	INVALID Stage = iota
	PENDING
	COMPLETED
	CANCELED
)

func (s Stage) String() string {
	switch s {

	case INVALID:
		return "invalid"
	case PENDING:
		return "pending"
	case COMPLETED:
		return "completed"
	case CANCELED:
		return "canceled"
	}

	return "unknown"
}

type ERC20AtomicSwapContract struct {
	contractAddress   common.Address
	swapTargetAddress common.Address
	abi               abi.ABI
	name              string
	symbol            string
	decimals          string
}

func NewERC20AtomicSwapContract() *ERC20AtomicSwapContract {
	abiJson, _ := abi.JSON(strings.NewReader(ERC20ATOMICSWAP_ABI))
	return &ERC20AtomicSwapContract{
		abi: abiJson,
	}
}
func (atomicSwap *ERC20AtomicSwapContract) SetAddress(contractAddress string) *ERC20AtomicSwapContract {
	atomicSwap.contractAddress = common.HexToAddress(contractAddress)
	return atomicSwap
}

func (atomicSwap *ERC20AtomicSwapContract) GetAddress() string {
	return atomicSwap.contractAddress.Hex()
}

func (atomicSwap *ERC20AtomicSwapContract) SetMetaData(name string, symbol string, decimals string, swapTargetAddress string) *ERC20AtomicSwapContract {
	// fmt.Println(swapTargetAddress)
	atomicSwap.name = name
	atomicSwap.symbol = symbol
	atomicSwap.decimals = decimals
	atomicSwap.swapTargetAddress = common.HexToAddress(swapTargetAddress)
	return atomicSwap
}

func (atomicSwap *ERC20AtomicSwapContract) GetMetaData() (string, string, string, string) {
	return atomicSwap.name, atomicSwap.symbol, atomicSwap.decimals, atomicSwap.swapTargetAddress.Hex()
}

func (atomicSwap *ERC20AtomicSwapContract) Name() (*ethereum.CallMsg, error) {
	methodName := "name"
	inputBytes, err := atomicSwap.abi.Pack(methodName)
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &atomicSwap.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (atomicSwap *ERC20AtomicSwapContract) Symbol() (*ethereum.CallMsg, error) {
	methodName := "symbol"
	inputBytes, err := atomicSwap.abi.Pack(methodName)
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &atomicSwap.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (atomicSwap *ERC20AtomicSwapContract) Decimals() (*ethereum.CallMsg, error) {
	methodName := "decimals"
	inputBytes, err := atomicSwap.abi.Pack(methodName)
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &atomicSwap.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (atomicSwap *ERC20AtomicSwapContract) AddressOfTargetContract() (*ethereum.CallMsg, error) {
	methodName := "addressOfTargetContract"
	inputBytes, err := atomicSwap.abi.Pack(methodName)
	if err != nil {
		return nil, fmt.Errorf("fail to pack the '%s' method", methodName)
	}
	callMsg := ethereum.CallMsg{
		To:   &atomicSwap.contractAddress,
		Data: inputBytes,
	}
	return &callMsg, nil
}

func (atomicSwap *ERC20AtomicSwapContract) CreateSwap(initiator string, receiver string, secretHash string, amount string) {

}
