require("dotenv").config()
const url = `https://${process.env.NETWORK}.infura.io/v3/${process.env.INFURA_API_KEY}`
const web3 = new (require('web3'))(url);
/* 
The sign method calculates an Ethereum specific signature with: 

    sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))).

*/
module.exports = async(params) => {
    const result = await require("./__common__")(__filename, params)
    if(result === undefined) {
        const address = params[0]
        const message = params[1] 

        console.log(" --- Account --- ")
        const signature = web3.eth.accounts.sign(message, process.env.PRIVATE_KEY)
        console.log(signature)

        const accountAddress = web3.eth.accounts.recover({
            messageHash: signature.messageHash,
            v: signature.v,
            r: signature.r,
            s: signature.s,
            signature: signature.signature
        })
        console.log(accountAddress) 

        const accountAddress2 = web3.eth.accounts.recover(message, signature.signature)
        console.log(accountAddress2)

        const accountAddress3 = web3.eth.accounts.recover(message,  signature.v, signature.r, signature.s)
        console.log(accountAddress3)

        console.log(" --- Eth --- ")
        const signature2 = await web3.eth.sign(message, address)
        console.log(signature2)

        // console.log(" --- Personal --- ")
        // console.log(await web3.eth.personal.ecRecover(message, signature))
        // const signature3 = await web3.eth.personal.sign(message,address,'test1234')
        // console.log(signature3) 
    }
}
