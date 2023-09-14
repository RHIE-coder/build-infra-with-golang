package demo

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum"
)

type ContractType int

const (
	POINT_TYPE ContractType = iota
	TOKEN_TYPE
	POINT_SWAP_TYPE
	TOKEN_SWAP_TYPE
)

type SwapBridge struct {
	txRequester *TransactionRequester
	point       *ERC20Contract
	token       *ERC20Contract
	pointSwap   *ERC20AtomicSwapContract
	tokenSwap   *ERC20AtomicSwapContract
}

func NewSwapBridge(
	provider *Provider,
	pointAddr string,
	tokenAddr string,
	pointSwapAddr string,
	tokenSwapAddr string,
) (*SwapBridge, error) {

	bridge := &SwapBridge{}
	bridge.txRequester = NewTransactionRequester(provider)

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

	nameBytes, err := bridge.txRequester.Call(*nameMsg)
	if err != nil {
		return nil, fmt.Errorf("point: %s", err.Error())
	}
	symbolBytes, err := bridge.txRequester.Call(*symbolMsg)
	if err != nil {
		return nil, fmt.Errorf("point: %s", err.Error())
	}
	decimalsBytes, err := bridge.txRequester.Call(*decimalsMsg)
	if err != nil {
		return nil, fmt.Errorf("point: %s", err.Error())
	}

	bridge.point = point.SetMetaData(
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

	nameBytes, err = bridge.txRequester.Call(*nameMsg)
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}
	symbolBytes, err = bridge.txRequester.Call(*symbolMsg)
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}
	decimalsBytes, err = bridge.txRequester.Call(*decimalsMsg)
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}
	addressOfTargetContractBytes, err := bridge.txRequester.Call(*addressOfTargetContractMsg)
	if err != nil {
		return nil, fmt.Errorf("pointSwap: %s", err.Error())
	}

	bridge.pointSwap = pointSwap.SetMetaData(
		ContractStringBytesToString(nameBytes),
		ContractStringBytesToString(symbolBytes),
		ContractIntegerBytesToString(decimalsBytes),
		ContractAddressBytes(addressOfTargetContractBytes).Hex(),
	)

	pointName, pointSymbol, pointDecimals := bridge.point.GetMetaData()
	pointSwapName, pointSwapSymbol, pointSwapDecimals, poinSwapTarget := bridge.pointSwap.GetMetaData()

	if pointName != pointSwapName {
		return nil, fmt.Errorf("point name is not same")
	}

	if pointSymbol != pointSwapSymbol {
		return nil, fmt.Errorf("point symbol is not same")
	}

	if pointDecimals != pointSwapDecimals {
		return nil, fmt.Errorf("point decimals is not same")
	}

	if point.GetAddressAsString() != poinSwapTarget {
		return nil, fmt.Errorf("point address is not same")
	}

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

	nameBytes, err = bridge.txRequester.Call(*nameMsg)
	if err != nil {
		return nil, fmt.Errorf("token: %s", err.Error())
	}
	symbolBytes, err = bridge.txRequester.Call(*symbolMsg)
	if err != nil {
		return nil, fmt.Errorf("token: %s", err.Error())
	}
	decimalsBytes, err = bridge.txRequester.Call(*decimalsMsg)
	if err != nil {
		return nil, fmt.Errorf("token: %s", err.Error())
	}

	bridge.token = token.SetMetaData(
		ContractStringBytesToString(nameBytes),
		ContractStringBytesToString(symbolBytes),
		ContractIntegerBytesToString(decimalsBytes),
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

	nameBytes, err = bridge.txRequester.Call(*nameMsg)
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}
	symbolBytes, err = bridge.txRequester.Call(*symbolMsg)
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}
	decimalsBytes, err = bridge.txRequester.Call(*decimalsMsg)
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}
	addressOfTargetContractBytes, err = bridge.txRequester.Call(*addressOfTargetContractMsg)
	if err != nil {
		return nil, fmt.Errorf("tokenSwap: %s", err.Error())
	}

	bridge.tokenSwap = tokenSwap.SetMetaData(
		ContractStringBytesToString(nameBytes),
		ContractStringBytesToString(symbolBytes),
		ContractIntegerBytesToString(decimalsBytes),
		ContractAddressBytes(addressOfTargetContractBytes).Hex(),
	)

	tokenName, tokenSymbol, tokenDecimals := bridge.token.GetMetaData()
	tokenSwapName, tokenSwapSymbol, tokenSwapDecimals, poinSwapTarget := bridge.tokenSwap.GetMetaData()

	if tokenName != tokenSwapName {
		return nil, fmt.Errorf("token name is not same")
	}

	if tokenSymbol != tokenSwapSymbol {
		return nil, fmt.Errorf("token symbol is not same")
	}

	if tokenDecimals != tokenSwapDecimals {
		return nil, fmt.Errorf("token decimals is not same")
	}

	if token.GetAddressAsString() != poinSwapTarget {
		return nil, fmt.Errorf("token address is not same")
	}

	return bridge, nil
}

func (swap *SwapBridge) SetSigner(signer *EthereumAccount) {
	swap.txRequester.SetSigner(signer)
}

func (swap *SwapBridge) GetContractAddress(contractType ContractType) string {
	if contractType == POINT_TYPE {
		return swap.point.GetAddressAsString()
	}
	if contractType == TOKEN_TYPE {
		return swap.token.GetAddressAsString()
	}
	if contractType == POINT_SWAP_TYPE {
		return swap.pointSwap.GetAddressAsString()
	}
	if contractType == TOKEN_SWAP_TYPE {
		return swap.tokenSwap.GetAddressAsString()
	}
	return ""
}

func (swap *SwapBridge) GetMetaData(contractType ContractType) map[string]string {
	metadata := make(map[string]string)

	if contractType == POINT_TYPE {
		metadata["name"], metadata["symbol"], metadata["decimals"] = swap.point.GetMetaData()
		return metadata
	}

	if contractType == TOKEN_TYPE {
		metadata["name"], metadata["symbol"], metadata["decimals"] = swap.token.GetMetaData()
		return metadata
	}

	if contractType == POINT_SWAP_TYPE {
		metadata["name"], metadata["symbol"], metadata["decimals"], metadata["addressOfTargetContract"] = swap.pointSwap.GetMetaData()
		return metadata
	}

	if contractType == TOKEN_SWAP_TYPE {
		metadata["name"], metadata["symbol"], metadata["decimals"], metadata["addressOfTargetContract"] = swap.tokenSwap.GetMetaData()
		return metadata
	}

	return nil
}

func (swap *SwapBridge) ERC20BalanceOf(contractType ContractType, addr string) (string, error) {

	var msg *ethereum.CallMsg
	var err error

	if contractType == POINT_TYPE {
		msg, err = swap.point.BalanceOf(addr)
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
	} else if contractType == TOKEN_TYPE {
		msg, err = swap.token.BalanceOf(addr)
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
	} else {
		return "", fmt.Errorf("unsupported contract type(should be erc20)")
	}
	balBytes, err := swap.txRequester.Call(*msg)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return ContractIntegerBytesToString(balBytes), nil
}

func (swap *SwapBridge) ERC20Approve(contractType ContractType, spender string, amount string) error {

	var msg *ethereum.CallMsg
	var err error

	if contractType == POINT_TYPE {
		msg, err = swap.point.Approve(spender, amount)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	} else if contractType == TOKEN_TYPE {
		msg, err = swap.token.Approve(spender, amount)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	} else {
		return fmt.Errorf("unsupported contract type(should be erc20)")
	}
	err = swap.txRequester.SendTransaction(*msg)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (swap *SwapBridge) ERC20Allowance(contractType ContractType, owner string, spender string) (string, error) {

	var msg *ethereum.CallMsg
	var err error

	if contractType == POINT_TYPE {
		msg, err = swap.point.Allowance(owner, spender)
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
	} else if contractType == TOKEN_TYPE {
		msg, err = swap.token.Allowance(owner, spender)
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
	} else {
		return "", fmt.Errorf("unsupported contract type(should be erc20)")
	}
	balBytes, err := swap.txRequester.Call(*msg)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return ContractIntegerBytesToString(balBytes), nil
}

func (swap *SwapBridge) ERC20Transfer(contractType ContractType, to string, amount string) error {

	var msg *ethereum.CallMsg
	var err error

	if contractType == POINT_TYPE {
		msg, err = swap.point.Transfer(to, amount)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	} else if contractType == TOKEN_TYPE {
		msg, err = swap.token.Transfer(to, amount)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	} else {
		return fmt.Errorf("unsupported contract type(should be erc20)")
	}
	err = swap.txRequester.SendTransaction(*msg)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (swap *SwapBridge) CreateSwap(contractType ContractType, initiator string, receiver string, secretHash string, amount string) error {
	var msg *ethereum.CallMsg
	var err error

	if contractType == POINT_SWAP_TYPE {
		msg, err = swap.pointSwap.CreateSwap(initiator, receiver, secretHash, amount)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	} else if contractType == TOKEN_SWAP_TYPE {
		msg, err = swap.tokenSwap.CreateSwap(initiator, receiver, secretHash, amount)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	} else {
		return fmt.Errorf("unsupported contract type(should be erc20AtomicSwap)")
	}
	err = swap.txRequester.SendTransaction(*msg)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (swap *SwapBridge) GetSwap(contractType ContractType, secretHash string) (string, error) {

	var msg *ethereum.CallMsg
	var err error

	if contractType == POINT_SWAP_TYPE {
		msg, err = swap.pointSwap.GetSwap(secretHash)
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
	} else if contractType == TOKEN_SWAP_TYPE {
		msg, err = swap.tokenSwap.GetSwap(secretHash)
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
	} else {
		return "", fmt.Errorf("unsupported contract type(should be erc20AtomicSwap)")
	}
	swapBytesByContract, err := swap.txRequester.Call(*msg)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	fmt.Println("--------")
	unpackedSwapData, err := swap.pointSwap.abi.Unpack("getSwap", swapBytesByContract)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	fmt.Println(unpackedSwapData)

	swapBytes, err := json.Marshal(unpackedSwapData[0])
	if err != nil {
		panic(err)
	}

	var jsonData interface{}

	if err := json.Unmarshal(swapBytes, &jsonData); err != nil {
		panic(err)
	}

	fmt.Println(jsonData)

	fmt.Println("--------")
	return "", nil
}
