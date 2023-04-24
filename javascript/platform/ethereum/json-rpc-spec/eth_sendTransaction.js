require("dotenv").config()
const url = `https://${process.env.NETWORK}.infura.io/v3/${process.env.INFURA_API_KEY}`
const web3 = new (require('web3'))(url);

/* 
from: DATA, 20 Bytes - The address the transaction is sent from.
to: DATA, 20 Bytes - (optional when creating new contract) The address the transaction is directed to.
gas: QUANTITY - (optional, default: 90000) Integer of the gas provided for the transaction execution. It will return unused gas.
gasPrice: QUANTITY - (optional, default: To-Be-Determined) Integer of the gasPrice used for each paid gas.
value: QUANTITY - (optional) Integer of the value sent with this transaction.
data: DATA - The compiled code of a contract OR the hash of the invoked method signature and encoded parameters.
nonce: QUANTITY - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.

params: [
  {
    from: "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
    to: "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
    gas: "0x76c0", // 30400
    gasPrice: "0x9184e72a000", // 10000000000000
    value: "0x9184e72a", // 2441406250
    data: "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675",
  },
]
*/

module.exports = async(params) => {
    const result = await require("./__common__")(__filename, params)
    console.log(result)
    const paramObject = JSON.parse(params[0])
}
