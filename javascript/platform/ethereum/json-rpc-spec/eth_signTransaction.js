require("dotenv").config()
const url = `https://${process.env.NETWORK}.infura.io/v3/${process.env.INFURA_API_KEY}`
const web3 = new (require('web3'))(url);

module.exports = async(params) => {
    const result = await require("./__common__")(__filename, params)
    console.log(result)
    const paramObject = JSON.parse(params[0])
    console.log(await web3.eth.signTransaction(paramObject))
}
