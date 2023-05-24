const Web3 = require('web3');
require("dotenv").config()

const network = "goerli"
// const web3 = new Web3(`https://${process.env["NETWORK"]}.infura.io/v3/${process.env["INFURA_API_KEY"]}`);
const web3 = new Web3(`https://${network}.infura.io/v3/${process.env["INFURA_API_KEY"]}`);


async function main(){
  const blockInfo = await web3.eth.getBlock(8564341)
  console.log(JSON.stringify(blockInfo, null, 2))
  console.log(blockInfo.timestamp)
}

main()