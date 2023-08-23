require("@nomicfoundation/hardhat-toolbox");
require('dotenv').config({
  path: require('path').join(__dirname,"..","..",".env")
})

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.19",
  networks:{
    ganache: {
      url:"http://192.168.100.73:10545",
      accounts:[process.env.LOCAL_ADMIN_PRIVATE_KEY],
    },
  }
};
