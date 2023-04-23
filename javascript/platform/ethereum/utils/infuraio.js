const JsonRpcRequester = require('./json-rpc-requester')

module.exports.getClient = function() {
    require('dotenv').config()
    const network = process.env.NETWORK
    const apiKey = process.env.INFURA_API_KEY
    const url = `https://${network}.infura.io/v3/${apiKey}`
    return new JsonRpcRequester(url)
}
