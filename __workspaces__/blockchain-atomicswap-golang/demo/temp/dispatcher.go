package demo

// import (
// 	"context"
// 	"fmt"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/common"
// )

// type ContractType int

// const (
// 	POINT ContractType = iota
// 	TOKEN
// 	POINT_SWAP
// 	TOKEN_SWAP
// )

// type SwapDispatcher struct {
// 	provider   *Provider
// 	signer     *EthereumAccount
// 	pointERC20 *ERC20Contract
// 	tokenERC20 *ERC20Contract
// 	pointSwap  *ERC20AtomicSwapContract
// 	tokenSwap  *ERC20AtomicSwapContract
// }

// func NewSwapDispatcher(
// 	provider *Provider,
// 	signer *EthereumAccount,
// 	pointERC20 *ERC20Contract,
// 	tokenERC20 *ERC20Contract,
// 	pointSwap *ERC20AtomicSwapContract,
// 	tokenSwap *ERC20AtomicSwapContract,
// ) *SwapDispatcher {
// 	return &SwapDispatcher{
// 		provider:   provider,
// 		signer:     signer,
// 		pointERC20: pointERC20,
// 		tokenERC20: tokenERC20,
// 		pointSwap:  pointSwap,
// 		tokenSwap:  tokenSwap,
// 	}
// }

// func (dispatcher *SwapDispatcher) GetProvider(contractType ContractType) *Provider {
// 	return dispatcher.provider
// }

// func (dispatcher *SwapDispatcher) GetERC20(contractType ContractType) *ERC20Contract {
// 	switch contractType {
// 	case POINT:
// 		return dispatcher.pointERC20
// 	case TOKEN:
// 		return dispatcher.tokenERC20
// 	default:
// 		return nil
// 	}
// }

// func (dispatcher *SwapDispatcher) GetERC20Swap(contractType ContractType) *ERC20AtomicSwapContract {
// 	switch contractType {
// 	case POINT_SWAP:
// 		return dispatcher.pointSwap
// 	case TOKEN_SWAP:
// 		return dispatcher.tokenSwap
// 	default:
// 		return nil
// 	}
// }

// func (dispatcher *SwapDispatcher) SetMetadataByCall(contractType ContractType) error {
// 	switch contractType {
// 	case POINT:
// 		erc20Point := dispatcher.GetERC20(POINT)

// 		nameMsg, err := erc20Point.Name()
// 		if err != nil {
// 			return fmt.Errorf("point: %s", err.Error())
// 		}
// 		symbolMsg, err := erc20Point.Symbol()
// 		if err != nil {
// 			return fmt.Errorf("point: %s", err.Error())
// 		}
// 		decimalsMsg, err := erc20Point.Decimals()
// 		if err != nil {
// 			return fmt.Errorf("point: %s", err.Error())
// 		}

// 		nameBytes, err := dispatcher.Call(*nameMsg)
// 		if err != nil {
// 			return fmt.Errorf("point: %s", err.Error())
// 		}
// 		symbolBytes, err := dispatcher.Call(*symbolMsg)
// 		if err != nil {
// 			return fmt.Errorf("point: %s", err.Error())
// 		}
// 		decimalsBytes, err := dispatcher.Call(*decimalsMsg)
// 		if err != nil {
// 			return fmt.Errorf("point: %s", err.Error())
// 		}

// 		erc20Point.SetMetaData(
// 			string(nameBytes),
// 			string(symbolBytes),
// 			new(big.Int).SetBytes(decimalsBytes).String(),
// 		)
// 		return nil
// 	case TOKEN:
// 		erc20Token := dispatcher.GetERC20(TOKEN)

// 		nameMsg, err := erc20Token.Name()
// 		if err != nil {
// 			return fmt.Errorf("token: %s", err.Error())
// 		}
// 		symbolMsg, err := erc20Token.Symbol()
// 		if err != nil {
// 			return fmt.Errorf("token: %s", err.Error())
// 		}
// 		decimalsMsg, err := erc20Token.Decimals()
// 		if err != nil {
// 			return fmt.Errorf("token: %s", err.Error())
// 		}

// 		nameBytes, err := dispatcher.Call(*nameMsg)
// 		if err != nil {
// 			return fmt.Errorf("token: %s", err.Error())
// 		}
// 		symbolBytes, err := dispatcher.Call(*symbolMsg)
// 		if err != nil {
// 			return fmt.Errorf("token: %s", err.Error())
// 		}
// 		decimalsBytes, err := dispatcher.Call(*decimalsMsg)
// 		if err != nil {
// 			return fmt.Errorf("token: %s", err.Error())
// 		}

// 		erc20Token.SetMetaData(
// 			string(nameBytes),
// 			string(symbolBytes),
// 			new(big.Int).SetBytes(decimalsBytes).String(),
// 		)
// 		return nil
// 	case POINT_SWAP:
// 		pointSwap := dispatcher.GetERC20Swap(POINT_SWAP)

// 		nameMsg, err := pointSwap.Name()
// 		if err != nil {
// 			return fmt.Errorf("pointSwap: %s", err.Error())
// 		}
// 		symbolMsg, err := pointSwap.Symbol()
// 		if err != nil {
// 			return fmt.Errorf("pointSwap: %s", err.Error())
// 		}
// 		decimalsMsg, err := pointSwap.Decimals()
// 		if err != nil {
// 			return fmt.Errorf("pointSwap: %s", err.Error())
// 		}
// 		addressOfTargetContractMsg, err := pointSwap.AddressOfTargetContract()
// 		if err != nil {
// 			return fmt.Errorf("pointSwap: %s", err.Error())
// 		}

// 		nameBytes, err := dispatcher.Call(*nameMsg)
// 		if err != nil {
// 			return fmt.Errorf("pointSwap: %s", err.Error())
// 		}
// 		symbolBytes, err := dispatcher.Call(*symbolMsg)
// 		if err != nil {
// 			return fmt.Errorf("pointSwap: %s", err.Error())
// 		}
// 		decimalsBytes, err := dispatcher.Call(*decimalsMsg)
// 		if err != nil {
// 			return fmt.Errorf("pointSwap: %s", err.Error())
// 		}
// 		addressOfTargetContractBytes, err := dispatcher.Call(*addressOfTargetContractMsg)
// 		if err != nil {
// 			return fmt.Errorf("pointSwap: %s", err.Error())
// 		}

// 		pointSwap.SetMetaData(
// 			string(nameBytes),
// 			string(symbolBytes),
// 			new(big.Int).SetBytes(decimalsBytes).String(),
// 			common.BytesToAddress(addressOfTargetContractBytes).Hex(),
// 		)
// 		return nil
// 	case TOKEN_SWAP:
// 		tokenSwap := dispatcher.GetERC20Swap(TOKEN_SWAP)

// 		nameMsg, err := tokenSwap.Name()
// 		if err != nil {
// 			return fmt.Errorf("tokenSwap: %s", err.Error())
// 		}
// 		symbolMsg, err := tokenSwap.Symbol()
// 		if err != nil {
// 			return fmt.Errorf("tokenSwap: %s", err.Error())
// 		}
// 		decimalsMsg, err := tokenSwap.Decimals()
// 		if err != nil {
// 			return fmt.Errorf("tokenSwap: %s", err.Error())
// 		}
// 		addressOfTargetContractMsg, err := tokenSwap.AddressOfTargetContract()
// 		if err != nil {
// 			return fmt.Errorf("tokenSwap: %s", err.Error())
// 		}

// 		nameBytes, err := dispatcher.Call(*nameMsg)
// 		if err != nil {
// 			return fmt.Errorf("tokenSwap: %s", err.Error())
// 		}
// 		symbolBytes, err := dispatcher.Call(*symbolMsg)
// 		if err != nil {
// 			return fmt.Errorf("tokenSwap: %s", err.Error())
// 		}
// 		decimalsBytes, err := dispatcher.Call(*decimalsMsg)
// 		if err != nil {
// 			return fmt.Errorf("tokenSwap: %s", err.Error())
// 		}
// 		addressOfTargetContractBytes, err := dispatcher.Call(*addressOfTargetContractMsg)
// 		if err != nil {
// 			return fmt.Errorf("tokenSwap: %s", err.Error())
// 		}
// 		tokenSwap.SetMetaData(
// 			string(nameBytes),
// 			string(symbolBytes),
// 			new(big.Int).SetBytes(decimalsBytes).String(),
// 			common.BytesToAddress(addressOfTargetContractBytes).Hex(),
// 		)
// 		return nil
// 	}
// 	return fmt.Errorf("not found contract type")
// }

// func (dispatcher *SwapDispatcher) Call(msg ethereum.CallMsg) ([]byte, error) {
// 	rpc := dispatcher.provider.GetClient()
// 	return rpc.CallContract(context.Background(), msg, nil)
// }

// func (dispatcher *SwapDispatcher) SendTx(tx ethereum.CallMsg) error {
// 	rpc := dispatcher.provider.GetClient()

// 	gasLimit, err := rpc.EstimateGas(context.Background(), tx)
// 	if err != nil {
// 		return fmt.Errorf("fail to estimate gas usage: %s", err.Error())
// 	}
// 	gasPrice, err := rpc.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		return fmt.Errorf("fail to get gas price: %s", err.Error())
// 	}
// 	tipCap, err := rpc.SuggestGasTipCap(context.Background())
// 	if err != nil {
// 		return fmt.Errorf("fail to get gas tip cap: %s", err.Error())
// 	}

// 	tx.Gas = gasLimit
// 	tx.GasPrice = gasPrice
// 	tx.GasFeeCap = tipCap

// 	signedTx, err := dispatcher.signer.SignTx(tx)
// 	if err != nil {
// 		return fmt.Errorf("fail to sign transaction: " + err.Error())
// 	}

// 	err = rpc.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		return fmt.Errorf("fail to call contract: " + err.Error())
// 	}
// 	return nil
// }
