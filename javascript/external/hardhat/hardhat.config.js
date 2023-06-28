require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config()

const NETWORK = process.env.NETWORK

// Go to https://infura.io, sign up, create a new API key
// in its dashboard, and replace "KEY" with it
const INFURA_API_KEY = process.env.INFURA_API_KEY;

// Replace this private key with your Sepolia account private key
// To export your private key from Coinbase Wallet, go to
// Settings > Developer Settings > Show private key
// To export your private key from Metamask, open Metamask and
// go to Account Details > Export Private Key
// Beware: NEVER put real Ether into testing accounts
const PRIVATE_KEY = process.env.PRIVATE_KEY;
console.log(`https://${NETWORK}.infura.io/v3/${INFURA_API_KEY}`)

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.18",
  networks: {
    ganache: {
      url: "http://192.168.100.73:8999",
      chainId:1337,
      accounts: [
        `6ff38a6fcde856869ddba8a1e0058a02cf81742f150607507d5245da607ba48f`,
      ],
    },
    // sepolia: {
    //   url: `https://${NETWORK}.infura.io/v3/${INFURA_API_KEY}`,
    //   accounts: [PRIVATE_KEY],
    // },
  },
};
