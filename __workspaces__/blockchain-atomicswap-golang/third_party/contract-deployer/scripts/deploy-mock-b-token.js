// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {

  const contractName = 'TokenB'

  console.log(`ready to deploy ${contractName}.sol`)
  const token = await hre.ethers.deployContract(contractName);

  console.log(`deploying ${contractName}.sol`);
  const tokenDeployed = await token.waitForDeployment();

  console.log(`the deployed ${contractName} contract address is : ` + tokenDeployed.target);

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
