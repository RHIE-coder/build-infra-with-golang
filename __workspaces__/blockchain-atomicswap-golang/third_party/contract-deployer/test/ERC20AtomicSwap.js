const {
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const hre = require("hardhat")

describe("ERC20AtomicSwap", function () {
  async function deployFixture() {
    // Get the Signers here.
    const [owner, addr1, addr2] = await ethers.getSigners();

    // To deploy our contract, we just have to call ethers.deployContract and await
    // its waitForDeployment() method, which happens once its transaction has been
    // mined.
    const hardhatContract = await ethers.deployContract("ERC20AtomicSwap");

    await hardhatContract.waitForDeployment();

    // Fixtures can return anything you consider useful for your tests
    return { hardhatContract, owner, addr1, addr2 };
  }

  // You can nest describe calls to create subsections.
  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      const { hardhatContract, owner } = await loadFixture(deployFixture);
      expect(await hardhatContract.owner()).to.equal(owner.address);
    });

    it("test", async function () {
      const { owner, addr1, addr2 } = await loadFixture(deployFixture);
      
      const baseTx = {
        to: addr2.address,
        value: ethers.parseEther("1"),
        data:"0x",
      }

      const chainId = hre.network.config.chainId
      const nonce = await addr1.getNonce();
      const feeData = await addr1.provider.getFeeData();
      const gasLimit = await addr1.provider.estimateGas(baseTx)
      
      let blockNumBeforeTx = await addr1.provider.getBlockNumber()
      let addr1Balance = await addr1.provider.getBalance(addr1)
      let addr2Balance = await addr2.provider.getBalance(addr2)
      console.log("block number before send: " + blockNumBeforeTx)
      console.log("addr1 has (" + addr1.address + ") " + ethers.formatUnits(addr1Balance.toString(), 'ether') + " ETH");
      console.log("addr2 has (" + addr2.address + ") " + ethers.formatUnits(addr2Balance.toString(), 'ether') + " ETH");

      const tx = {
        type: 2, // EIP1559
        nonce, 
        maxPriorityFeePerGas: feeData["maxPriorityFeePerGas"], // Recommended maxPriorityFeePerGas
        maxFeePerGas: feeData["maxFeePerGas"], // Recommended maxFeePerGas
        gasLimit,
        chainId,
        ...baseTx
      }
      const wallet1 = ethers.fromMnemonic(hre.config.networks.hardhat.accounts.mnemonic)
      console.log(wallet1.address)
      console.log(owen.address)
      // console.log(tx);
      // console.log(hre.config.networks.hardhat)
      // const signedTx = await addr1.signTransaction(tx);
      // console.log("Signed Transaction:", signedTx);
    });
  });
});
