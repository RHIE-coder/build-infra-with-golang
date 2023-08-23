// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ERC20AtomicSwap is Ownable{

    struct Swap {
        address tokenAddress;
        address sender;
        address receiver;
        bytes32 secretHash;
        uint256 amount;
    }

    /**
     * @dev 
     */
    enum Stage {
        INVALID,
        PENDING,
        COMPLETED,
        CANCELED
    }

    mapping(bytes32 => Swap) public _swaps;
    mapping(bytes32 => Stage) public _swapStatus;


    error InsufficientAllowance(address tokenAddress, address owner, address spender, uint256 amount,uint256 allowance);

    /**
     * @dev 
     */
    function createSwap(address tokenAddress_,address initiator_, address receiver_, bytes32 secretHash_, uint256 amount_) public onlyOwner {

        require(amount_ != 0, "the amount cannot be zero");

        uint256 balance = IERC20(tokenAddress_).balanceOf(initiator_);
        require(balance >= amount_, "insufficient balance");

        uint256 allowedAmount = IERC20(tokenAddress_).allowance(initiator_, address(this));
        if(allowedAmount == 0 || allowedAmount < amount_) {
            revert InsufficientAllowance(tokenAddress_, initiator_, address(this), amount_, allowedAmount);
        }

        require(_swapStatus[secretHash_] == Stage.INVALID, "hash is already exists");
        
        Swap memory initSwap = Swap({
            tokenAddress: tokenAddress_,
            sender: initiator_,
            receiver: receiver_,
            secretHash: secretHash_,
            amount: amount_
        });

        bool isTransferSuccess = IERC20(tokenAddress_).transferFrom(initiator_, address(this), amount_);
        require(isTransferSuccess, "fail to transfer");

        _swaps[secretHash_]=initSwap;
        _swapStatus[secretHash_]=Stage.PENDING;

    }

    /**
     * @dev Redeem ERC20 token 
     */
    function redeem(bytes memory secret_, bytes32 secretHash_) public onlyOwner {

        require(_swapStatus[secretHash_] != Stage.COMPLETED, "swap is already completed");

        Swap memory pendingSwap = _swaps[secretHash_];

        require(keccak256(abi.encodePacked(secret_)) == pendingSwap.secretHash, "secret is not matched with swap");

        bool isTransferSuccess = IERC20(pendingSwap.tokenAddress).transfer(pendingSwap.receiver, pendingSwap.amount);

        require(isTransferSuccess, "fail to transfer");

        _swapStatus[secretHash_] = Stage.COMPLETED;

    }

    /**
     * @dev Refund ERC20 token
     */
    function refund(bytes32 secretHash_) public onlyOwner {
        require(_swapStatus[secretHash_] != Stage.CANCELED, "swap is already canceled");
        require(_swapStatus[secretHash_] != Stage.COMPLETED, "swap is already completed");

        Swap memory pendingSwap = _swaps[secretHash_];


        bool isTransferSuccess = IERC20(pendingSwap.tokenAddress).transfer(pendingSwap.sender, pendingSwap.amount);

        require(isTransferSuccess, "fail to transfer");

        _swapStatus[secretHash_] = Stage.CANCELED;
    }

    function getSwap(bytes32 secretHash_) public view returns(Swap memory) {
        return _swaps[secretHash_];
    }

    /**
     * @dev Check whether the swap is already redeemed or not
     */
    function isRedeemed(bytes32 secretHash) public view returns(bool) {
        require(_swapStatus[secretHash] != Stage.INVALID, "swap hash is not valid");
        return _swapStatus[secretHash] == Stage.COMPLETED ? true : false;
    }

    
    /**
     * @dev Check whether the swap is already refunded or not
     */
    function isRefunded(bytes32 secretHash) public view returns(bool) {
        require(_swapStatus[secretHash] != Stage.INVALID, "swap hash is not valid");
        return _swapStatus[secretHash] == Stage.CANCELED ? true : false;
    }

}
