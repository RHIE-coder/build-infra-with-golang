const Web3 = require('web3');
const web3 = new Web3();

module.exports = async(params) => {
    const infuraClient = require('../utils/infuraio').getClient()
    const path = require('path');
    console.log(params)
    const method = path.basename(__filename, path.extname(__filename));
    const res = await infuraClient.dial({id:1, method, params})
    console.log(res)
    const result = res.result
    const decodedResult = web3.utils.fromWei(result, 'ether');
    console.log(decodedResult) 
}
