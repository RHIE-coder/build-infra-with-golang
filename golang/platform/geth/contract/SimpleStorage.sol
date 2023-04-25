// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract SimpleStorage {
    uint256 private _value;
    event Set(address indexed _sender, uint256 _value, uint256 _when);

    function setValue(uint256 value) public {
        emit Set(msg.sender, value, block.timestamp);
        _value = value;
    }
    
    function getValue() public view returns (uint256) {
        return _value;
    }
}