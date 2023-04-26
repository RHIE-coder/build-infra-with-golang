const { ethers } = require('ethers');
require("dotenv").config()
const mnemonic = process.env.MNEMONIC 
const wallet = ethers.Wallet.fromPhrase(mnemonic);
const privateKey = wallet.privateKey;



console.log(ethers)