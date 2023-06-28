//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract A {

    uint num = 100;

    function a() public virtual returns(uint) {
        return 1;
    }

    function b() public virtual returns(uint) {
        return 2;
    }
}

contract B {

    // uint num = 200;

    function a() public virtual returns(uint) {
        return 3;
    }
}

contract Main is B, A{

    function a() public override(B,A) returns(uint) {
        return super.a();
    }

    function b() public override returns(uint) {
        return super.b();
    }

    function getNum() public view returns(uint) {
        return num;
    }

}