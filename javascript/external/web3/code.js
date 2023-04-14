const Web3 = require('web3');

const web3 = new Web3();

const hexString = '0x6848f5974'; 

const utf8String = web3.utils.fromWei(hexString);
console.log(utf8String); 
