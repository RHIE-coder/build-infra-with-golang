//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Callee {

    event CalleeSenderAddress(address sender);

    function call() public returns(address){
        emit CalleeSenderAddress(msg.sender);
        return msg.sender;
    }
}

contract Caller {

    Callee otherContract;

    event CallerSenderAddress(address sender);

    constructor() public{
        otherContract = new Callee();
    }

    function call() public returns(address) {
        emit CallerSenderAddress(msg.sender);
        otherContract.call();
        return address(otherContract);
    }
}