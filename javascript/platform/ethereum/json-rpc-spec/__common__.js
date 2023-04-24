module.exports = async function(filename, params) {
    const infuraClient = require('../utils/infuraio').getClient()
    const path = require('path');
    console.log(params)
    const method = path.basename(filename, path.extname(filename));
    const res = await infuraClient.dial({id:1, method, params})
    console.log(res)
    return res.result
}