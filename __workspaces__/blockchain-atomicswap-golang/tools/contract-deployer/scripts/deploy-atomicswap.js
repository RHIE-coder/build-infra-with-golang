// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {

  const cashDev = "0xC881D558C5Ba65d15667f96B1c66fdEF8255AcA4";
  const pointDev = "0xCF3a2132bEDFf194f4E40891c481bE38F95C382c";

  // const targetContractAddress = cashDev;
  const targetContractAddress = pointDev;

  console.log("ready to deploy ERC20AtomicSwap.sol")
  const erc20AtomicSwap = await hre.ethers.deployContract("ERC20AtomicSwap", [targetContractAddress]); //how

  console.log("deploying ERC20AtomicSwap.sol");
  const deployed = await erc20AtomicSwap.waitForDeployment();

  console.log('the contract address is : ' + deployed.target);

  const result = await deployed.name();
  console.log(result);

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
