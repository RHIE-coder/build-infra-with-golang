package demo

import (
	"fmt"
)

const (
	POINT_TYPE = iota
	TOKEN_TYPE
	POINT_SWAP_TYPE
	TOKEN_SWAP_TYPE
)

type SwapBridge struct {
	provider  *Provider
	point     *ERC20Contract
	token     *ERC20Contract
	pointSwap *ERC20AtomicSwapContract
	tokenSwap *ERC20AtomicSwapContract
}

func NewSwapBridge(
	provider *Provider,
	pointAddr string,
	tokenAddr string,
	pointSwapAddr string,
	tokenSwapAddr string,
) (*SwapBridge, error) {

	txRequester := NewTransactionRequester(provider)
	bridge := &SwapBridge{}
	bridge.provider = provider

	point := NewERC20Contract()
	token := NewERC20Contract()
	pointSwap := NewERC20AtomicSwapContract()
	tokenSwap := NewERC20AtomicSwapContract()

	point.SetAddressByString(pointAddr)
	token.SetAddressByString(tokenAddr)
	pointSwap.SetAddressByString(pointSwapAddr)
	tokenSwap.SetAddressByString(tokenSwapAddr)

	nameMsg, err := point.Name()
	if err != nil {
		return nil, fmt.Errorf("point: %s", err.Error())
	}
	symbolMsg, err := point.Symbol()
	if err != nil {
		return nil, fmt.Errorf("point: %s", err.Error())
	}
	decimalsMsg, err := point.Decimals()
	if err != nil {
		return nil, fmt.Errorf("point: %s", err.Error())
	}

	nameBytes, err := txRequester.Call(*nameMsg)
	if err != nil {
		return nil, fmt.Errorf("point: %s", err.Error())
	}
	symbolBytes, err := txRequester.Call(*symbolMsg)
	if err != nil {
		return nil, fmt.Errorf("point: %s", err.Error())
	}
	decimalsBytes, err := txRequester.Call(*decimalsMsg)
	if err != nil {
		return nil, fmt.Errorf("point: %s", err.Error())
	}

	bridge.point = point.SetMetaData(
		ContractStringBytesToString(nameBytes),
		ContractStringBytesToString(symbolBytes),
		ContractIntegerBytesToString(decimalsBytes),
	)

	nameMsg, err = token.Name()
	if err != nil {
		return nil, fmt.Errorf("token: %s", err.Error())
	}
	symbolMsg, err = token.Symbol()
	if err != nil {
		return nil, fmt.Errorf("token: %s", err.Error())
	}
	decimalsMsg, err = token.Decimals()
	if err != nil {
		return nil, fmt.Errorf("token: %s", err.Error())
	}

	nameBytes, err = txRequester.Call(*nameMsg)
	if err != nil {
		return nil, fmt.Errorf("token: %s", err.Error())
	}
	symbolBytes, err = txRequester.Call(*symbolMsg)
	if err != nil {
		return nil, fmt.Errorf("token: %s", err.Error())
	}
	decimalsBytes, err = txRequester.Call(*decimalsMsg)
	if err != nil {
		return nil, fmt.Errorf("token: %s", err.Error())
	}

	bridge.token = token.SetMetaData(
		ContractStringBytesToString(nameBytes),
		ContractStringBytesToString(symbolBytes),
		ContractIntegerBytesToString(decimalsBytes),
	)

	nameMsg, err = pointSwap.Name()
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}
	symbolMsg, err = pointSwap.Symbol()
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}
	decimalsMsg, err = pointSwap.Decimals()
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}
	addressOfTargetContractMsg, err := pointSwap.AddressOfTargetContract()
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}

	nameBytes, err = txRequester.Call(*nameMsg)
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}
	symbolBytes, err = txRequester.Call(*symbolMsg)
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}
	decimalsBytes, err = txRequester.Call(*decimalsMsg)
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}
	addressOfTargetContractBytes, err := txRequester.Call(*addressOfTargetContractMsg)
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}

	bridge.pointSwap = pointSwap.SetMetaData(
		ContractStringBytesToString(nameBytes),
		ContractStringBytesToString(symbolBytes),
		ContractIntegerBytesToString(decimalsBytes),
		ContractAddressBytes(addressOfTargetContractBytes).Hex(),
	)

	nameMsg, err = tokenSwap.Name()
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}
	symbolMsg, err = tokenSwap.Symbol()
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}
	decimalsMsg, err = tokenSwap.Decimals()
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}
	addressOfTargetContractMsg, err = tokenSwap.AddressOfTargetContract()
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}

	nameBytes, err = txRequester.Call(*nameMsg)
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}
	symbolBytes, err = txRequester.Call(*symbolMsg)
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}
	decimalsBytes, err = txRequester.Call(*decimalsMsg)
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}
	addressOfTargetContractBytes, err = txRequester.Call(*addressOfTargetContractMsg)
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}

	bridge.tokenSwap = tokenSwap.SetMetaData(
		ContractStringBytesToString(nameBytes),
		ContractStringBytesToString(symbolBytes),
		ContractIntegerBytesToString(decimalsBytes),
		ContractAddressBytes(addressOfTargetContractBytes).Hex(),
	)

	return bridge, nil
}

func (swap *SwapBridge) Name()
