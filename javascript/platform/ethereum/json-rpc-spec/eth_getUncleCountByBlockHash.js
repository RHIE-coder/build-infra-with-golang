const web3 = new (require('web3'))();

module.exports = async(params) => {
    const result = await require("./__common__")(__filename, params)
    console.log(result)
    console.log(web3.utils.hexToNumberString(result))
}
