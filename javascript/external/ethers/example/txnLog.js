require("dotenv").config(__dirname);
const axios = require("axios");
const url = "https://mainnet.infura.io/v3/" + process.env["INFURA_KEY"];
const { ethers } = require('ethers');

(async()=>{
    const infuraProvider = new ethers.InfuraProvider("homestead", process.env["INFURA_KEY"])
    console.log(infuraProvider)
    const latest_block = await infuraProvider.getBlockNumber('latest');
    console.log(latest_block)
    // Define the ERC20 contract address
    const contractAddress = '0x777fD20c983d6658c1D50b3958B3A1733d1cd1E1'; 

    // Define the Transfer event signature
    const transferEventSig = 'Transfer(address,address,uint256)';
    const topicHash = ethers.utils.keccak256(ethers.utils.toUtf8Bytes(transferEventSig));
    // Define the filter parameters
    const filter = {
    address: contractAddress,
    topics: [topicHash],
    fromBlock: 0,
    toBlock: 'latest',
    };

    // Use the filter to retrieve ERC20 transfer logs
    infuraProvider.getLogs(filter)
    .then(logs => {
        // Process the logs
        console.log(logs);
    })
    .catch(error => {
        // Handle any errors
        console.error(error);
    });
})()