// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface Hello {
    function abc() external returns(uint);
    function def(uint num) external;
    function hij(string memory, bytes memory) external returns(string memory);
}

abstract contract World {
    function klm() public {}
    function nop() public virtual returns(bytes memory);
    function qrs() external virtual;
}
contract Rhie {}


contract InterfaceIdTest {

    function howToGetSelector1() public pure returns(bytes4) {
        return bytes4(keccak256("hij(string,bytes)"));
    }

    function howToGetSelector2() public pure returns(bytes4) {
        return Hello.hij.selector;
    }

    function getAbstract() public pure returns(bytes4) {
        return type(World).interfaceId;
    }

    function getInterface() public pure returns(bytes4) {
        return type(Hello).interfaceId;
    }
 zzzzzzzzz
    function getSelectorStringOf(string memory funcName) public pure returns(bytes4){
        return bytes4(keccak256(abi.encodePacked(funcName)));
    }

    function one111() public pure returns(bytes4) {
        return Hello.abc.selector;
    }

    function two222() public pure returns(bytes4) {
        return Hello.def.selector;
    }

    function three333() public pure returns(bytes4) {
        return Hello.hij.selector;
    }

    function getInterfaceCustom() public pure returns(bytes4) {

        // bytes4 selector1 = bytes4(keccak256("abc()"));
        // bytes4 selector2 = bytes4(keccak256("def(uint256)"));
        // bytes4 selector3 = bytes4(keccak256("hij(string,bytes)"));
         
        // return selector1 ^ selector2 ^ selector3;
        return bytes4(keccak256("abc()")) ^ bytes4(keccak256("def(uint256)")) ^ bytes4(keccak256("hij(string,bytes)"));
    }


    function getAbstractCustom() public pure returns(bytes4) {
        return World.klm.selector ^ World.nop.selector ^ World.qrs.selector;
    }

    function getContract() public pure returns(string memory) {
        return type(Rhie).name;
    }
}

contract AbiControl {

}