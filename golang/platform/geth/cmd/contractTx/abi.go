/*
위 코드에서는 ioutil.ReadFile 함수를 사용하여 CONTRACT_ABI.json 파일을 읽어온 후, json.Unmarshal 함수를 사용하여 JSON 데이터를 abi.ABI 타입의 객체로 변환합니다. 이후 myABI 변수에 저장된 스마트 컨트랙트 ABI 정보를 사용하여 스마트 컨트랙트 함수 호출 등을 구현할 수 있습니다.
*/
/*
[
    {
        "constant": false,
        "inputs": [
            {
                "name": "newValue",
                "type": "uint256"
            }
        ],
        "name": "setValue",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "name": "initialValue",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "getValue",
        "outputs": [
            {
                "name": "",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    }
]
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	// 스마트 컨트랙트 ABI 정보를 파일에서 읽어옴
	abiJSON, err := ioutil.ReadFile("CONTRACT_ABI.json")
	if err != nil {
		fmt.Println("Error: cannot read ABI file:", err)
		return
	}

	// ABI 정보를 abi.ABI 타입의 객체로 변환
	var myABI abi.ABI
	err = json.Unmarshal(abiJSON, &myABI)
	if err != nil {
		fmt.Println("Error: cannot unmarshal ABI JSON:", err)
		return
	}

	// 스마트 컨트랙트 ABI 정보를 사용하여 스마트 컨트랙트 함수 호출 등을 구현할 수 있음
}
