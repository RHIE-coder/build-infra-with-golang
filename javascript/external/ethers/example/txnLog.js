require("dotenv").config(__dirname);
const axios = require("axios");
const url = "https://mainnet.infura.io/v3/" + process.env["INFURA_KEY"];
const { ethers } = require('ethers');

(async()=>{
    const infuraProvider = new ethers.InfuraProvider("homestead", process.env["INFURA_KEY"])
    console.log(infuraProvider)
    const latest_block = await infuraProvider.getBlockNumber('latest');
    console.log(latest_block)
})()