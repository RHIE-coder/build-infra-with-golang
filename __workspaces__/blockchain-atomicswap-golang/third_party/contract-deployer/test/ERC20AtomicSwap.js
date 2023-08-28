const {
    loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers } = require("hardhat");
const crypto = require('crypto');

const fixedTransferAmount = 500n;
const fixedSwapAmount = 300n;

describe("ERC20AtomicSwap", function () {

    async function deployFixture() {
        const [owner, addr1, addr2] = await ethers.getSigners();

        const SwapContract = await ethers.deployContract("ERC20AtomicSwap");

        await SwapContract.waitForDeployment();

        const TokenAContract = await ethers.deployContract("TokenA");

        await TokenAContract.waitForDeployment();

        const TokenBContract = await ethers.deployContract("TokenB");

        await TokenBContract.waitForDeployment();

        // Fixtures can return anything you consider useful for your tests
        return { SwapContract, TokenAContract, TokenBContract, owner, addr1, addr2 };
    }

    describe("Deploy Contract", function () {
        it("should set the right owner and right ERC20", async function () {

            const {
                SwapContract,
                TokenAContract,
                TokenBContract,
                owner
            } = await loadFixture(deployFixture);

            expect(await SwapContract.owner()).to.equal(owner.address);
            expect(await TokenAContract.name()).to.equal("atoken");
            expect(await TokenAContract.symbol()).to.equal("aAa");
            expect(await TokenBContract.name()).to.equal("btoken");
            expect(await TokenBContract.symbol()).to.equal("bBb");

        });

        it("should assign the total supply of tokens to the owner", async function () {
            const {
                owner,
                TokenAContract,
                TokenBContract,
            } = await loadFixture(deployFixture);
            const ownerBalanceA = await TokenAContract.balanceOf(owner.address);
            expect(await TokenAContract.totalSupply()).to.equal(ownerBalanceA);
            const ownerBalanceB = await TokenBContract.balanceOf(owner.address);
            expect(await TokenBContract.totalSupply()).to.equal(ownerBalanceB);
        })

        it("validate swaping function", async function () {
            const {
                SwapContract,
                TokenAContract,
                TokenBContract,
                owner,
                addr1,
                addr2,
            } = await loadFixture(deployFixture);

            // send tokens to addr1 and addr2
            await TokenAContract.transfer(addr1.address, fixedTransferAmount);
            await TokenBContract.transfer(addr2.address, fixedTransferAmount);
            const balanceA1 = await TokenAContract.balanceOf(addr1);
            const balanceA2 = await TokenAContract.balanceOf(addr2);
            const balanceB1 = await TokenBContract.balanceOf(addr1);
            const balanceB2 = await TokenBContract.balanceOf(addr2);
            expect(balanceA1).to.be.equal(fixedTransferAmount)
            expect(balanceA2).to.be.equal(0n)
            expect(balanceB1).to.be.equal(0n)
            expect(balanceB2).to.be.equal(fixedTransferAmount)

            // set approval for SwapContract
            await TokenAContract.connect(addr1).approve(SwapContract.target, fixedSwapAmount);
            await TokenBContract.connect(addr2).approve(SwapContract.target, fixedSwapAmount);
            const allowanceOfA = await TokenAContract.allowance(addr1.address, SwapContract.target);
            const allowanceOfB = await TokenBContract.allowance(addr2.address, SwapContract.target);
            expect(allowanceOfA).to.be.equal(300n)
            expect(allowanceOfB).to.be.equal(300n)

            // set secret by addr1
            const secretBytes = crypto.randomBytes(32);
            const secretHexString = secretBytes.toString('hex');
            const secretHash = ethers.keccak256(Buffer.from(secretHexString, 'hex'))
            expect(secretBytes.length).to.be.equal(Buffer.from(secretHexString, 'hex').length)
            expect(secretHash).to.be.equal(ethers.keccak256(secretBytes))

            let result; // tracker

            // addr1 create swap
            result = await SwapContract.createSwap(
                TokenAContract.target,
                addr1.address,
                addr2.address,
                secretHash,
                fixedSwapAmount,
            )

            // addr2 try to know secretHash
            const eventName = "SwapCreated";

            const filter = {
                address: SwapContract.target,
                topics: [ethers.id(eventName)]
            };
            const eventSignature = SwapContract.interface.getEvent(eventName);
            console.log(eventSignature)
            SwapContract.interface.parseLog([
                eventSignature,
                filter.topics,
            ])
            // console.log(filter)

            // result = await SwapContract.createSwap(
            //     TokenBContract.target,
            //     addr2.address,
            //     addr1.address,
            //     secretHash,
            //     fixedSwapAmount,
            // )
            // const ABI = require('../artifacts/contracts/ERC20AtomicSwap.sol/ERC20AtomicSwap.json').abi;

            
        });

    })

});

// const hre = require("hardhat")

// describe("Hardhat Test", function () {
//     async function deployFixture() {
//         console.log(" ------ deployFixture() is invoked ------")
//         // Get the Signers here.
//         const [owner, addr1, addr2] = await ethers.getSigners();

//         // To deploy our contract, we just have to call ethers.deployContract and await
//         // its waitForDeployment() method, which happens once its transaction has been
//         // mined.
//         const hardhatContract = await ethers.deployContract("ERC20AtomicSwap");

//         await hardhatContract.waitForDeployment();

//         // Fixtures can return anything you consider useful for your tests
//         return { hardhatContract, owner, addr1, addr2 };
//     }

//     // You can nest describe calls to create subsections.
//     describe("Deployment", function () {
//         it("Should set the right owner", async function () {
//             const { hardhatContract, owner } = await loadFixture(deployFixture);
//             expect(await hardhatContract.owner()).to.equal(owner.address);
//         });

//         it("accounts handling", async function () {
//             const { owner, addr1, addr2 } = await loadFixture(deployFixture);

//             const baseTx = {
//                 to: addr2.address,
//                 value: ethers.parseEther("1"),
//                 data: "0x",
//             }

//             const chainId = hre.network.config.chainId
//             const nonce = await addr1.getNonce();
//             const feeData = await addr1.provider.getFeeData();
//             const gasLimit = await addr1.provider.estimateGas(baseTx)

//             let blockNumBeforeTx = await addr1.provider.getBlockNumber()
//             let addr1Balance = await addr1.provider.getBalance(addr1)
//             let addr2Balance = await addr2.provider.getBalance(addr2)
//             console.log("block number before send: " + blockNumBeforeTx)
//             console.log("addr1 has (" + addr1.address + ") " + ethers.formatUnits(addr1Balance.toString(), 'ether') + " ETH");
//             console.log("addr2 has (" + addr2.address + ") " + ethers.formatUnits(addr2Balance.toString(), 'ether') + " ETH");

//             const tx = {
//                 type: 2, // EIP1559
//                 nonce,
//                 maxPriorityFeePerGas: feeData["maxPriorityFeePerGas"], // Recommended maxPriorityFeePerGas
//                 maxFeePerGas: feeData["maxFeePerGas"], // Recommended maxFeePerGas
//                 gasLimit,
//                 chainId,
//                 ...baseTx
//             }
//             const mnemonicString = hre.config.networks.hardhat.accounts.mnemonic
//             const ownerWallet = ethers.HDNodeWallet.fromMnemonic(ethers.Mnemonic.fromPhrase(mnemonicString));
//             console.log(ownerWallet.address)
//             console.log(owner.address)
//             const pathArr = ownerWallet.path.split('/');
//             console.log(ownerWallet.path)
//             pathArr[pathArr.length - 1] = '1'
//             addr1Path = pathArr.join("/")
//             console.log(addr1Path)
//             pathArr[pathArr.length - 1] = '2'
//             addr2Path = pathArr.join("/")
//             console.log(addr2Path)

//             const addr1Wallet = ethers.HDNodeWallet.fromMnemonic(ethers.Mnemonic.fromPhrase(mnemonicString), addr1Path)
//             const addr2Wallet = ethers.HDNodeWallet.fromMnemonic(ethers.Mnemonic.fromPhrase(mnemonicString), addr2Path)
//             console.log(addr1Wallet.address)
//             console.log(addr2Wallet.address)
//             console.log(addr1Wallet.address === addr1.address)
//             console.log(addr2Wallet.address === addr2.address)
//             console.log(tx);
//             const signedTx = await addr1Wallet.signTransaction(tx);
//             console.log("Signed Transaction:", signedTx);
//             const result = await addr1.sendTransaction(tx);
//             console.log(result)
//             let blockNumAfterTx = await addr1.provider.getBlockNumber()
//             console.log("block number before send: " + blockNumAfterTx)
//             addr1Balance = await addr1.provider.getBalance(addr1)
//             addr2Balance = await addr2.provider.getBalance(addr2)
//             console.log("addr1 has (" + addr1.address + ") " + ethers.formatUnits(addr1Balance.toString(), 'ether') + " ETH");
//             console.log("addr2 has (" + addr2.address + ") " + ethers.formatUnits(addr2Balance.toString(), 'ether') + " ETH");
//         });

//         it('check context after "it" execute', async function () {
//             console.log(' ### start ### check invoke loadFixture')
//             const [owner, addr1, addr2] = await ethers.getSigners();
//             console.log(' ### end ###check invoke loadFixture')
//             let blockNumAfterTx = await addr1.provider.getBlockNumber()
//             let addr1Balance = await addr1.provider.getBalance(addr1)
//             let addr2Balance = await addr2.provider.getBalance(addr2)
//             console.log("block number after send: " + blockNumAfterTx)
//             console.log("addr1 has (" + addr1.address + ") " + ethers.formatUnits(addr1Balance.toString(), 'ether') + " ETH");
//             console.log("addr2 has (" + addr2.address + ") " + ethers.formatUnits(addr2Balance.toString(), 'ether') + " ETH");
//         })
//     });
// });