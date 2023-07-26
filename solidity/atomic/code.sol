pragma solidity ^0.8.0;

contract AtomicSwap {
    struct Swap {
        address sender;
        address receiver;
        bytes32 hashedSecret;
        uint256 lockTime;
        uint256 amount;
        bool isWithdrawn;
    }

    mapping(bytes32 => Swap) public swaps;

    event SwapStarted(
        bytes32 indexed swapId,
        address indexed sender,
        address indexed receiver,
        uint256 amount,
        uint256 lockTime
    );

    event SwapCompleted(bytes32 indexed swapId, bytes32 secret);

    modifier onlyNotWithdrawn(bytes32 swapId) {
        require(!swaps[swapId].isWithdrawn, "Swap already withdrawn");
        _;
    }

    function createSwap(bytes32 swapId, address receiver, bytes32 hashedSecret, uint256 lockTime) external payable {
        require(swaps[swapId].sender == address(0), "Swap already exists");
        require(msg.value > 0, "Amount should be greater than 0");
        require(block.timestamp < lockTime, "Lock time should be in the future");

        swaps[swapId] = Swap(msg.sender, receiver, hashedSecret, lockTime, msg.value, false);

        emit SwapStarted(swapId, msg.sender, receiver, msg.value, lockTime);
    }

    function withdraw(bytes32 swapId, bytes32 secret) external onlyNotWithdrawn(swapId) {
        Swap storage swap = swaps[swapId];
        require(msg.sender == swap.receiver, "Only the receiver can withdraw");
        require(block.timestamp >= swap.lockTime, "Lock time has not yet expired");
        require(keccak256(abi.encodePacked(secret)) == swap.hashedSecret, "Invalid secret");

        swap.isWithdrawn = true;
        payable(swap.receiver).transfer(swap.amount);

        emit SwapCompleted(swapId, secret);
    }

    function refund(bytes32 swapId) external onlyNotWithdrawn(swapId) {
        Swap storage swap = swaps[swapId];
        require(msg.sender == swap.sender, "Only the sender can refund");
        require(block.timestamp >= swap.lockTime, "Lock time has not yet expired");

        swap.isWithdrawn = true;
        payable(swap.sender).transfer(swap.amount);
    }
}