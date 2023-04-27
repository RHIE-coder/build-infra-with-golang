const { ethers } = require('ethers');

module.exports = {
    getPrivateKeyFromMnemonic(mnemonic) {
        const wallet = ethers.Wallet.fromPhrase(mnemonic);
        const privateKey = wallet.privateKey;
        return privateKey
    },
    
    async signTransaction(data, pk, chainId) {
        const wallet = new ethers.Wallet(pk)
        const tx = new ethers.Transaction()

        tx.chainId = chainId

        tx.type = 0
        tx.to = data.to;
        tx.from = data.from;
        tx.gasLimit = BigInt(data.gas);
        tx.gasPrice = BigInt(data.gasPrice);
        tx.nonce = BigInt(data.nonce);

        if(data.value) {
            tx.value = BigInt(data.value);
        }

        if(data.data) {
            tx.data = data.data
        }

        console.log(tx.toJSON())

        return await wallet.signTransaction(tx)
    },

        
    async signTransactionLuniverse(data, pk, chainId) {
        const wallet = new ethers.Wallet(pk)
        const tx = new ethers.Transaction()

        tx.chainId = chainId

        tx.type = 0
        tx.to = data.to;
        tx.from = data.from;
        tx.gasLimit = BigInt(data.gas);
        tx.gasPrice = BigInt(data.gasPrice);
        tx.nonce = BigInt(data.nonce);
        tx.value = BigInt('0x0')
        tx.data = data.data

        console.log("-------------------------------------")
        console.log(tx.toJSON())
        const btx = ethers.Transaction.from(tx)
        console.log(btx.unsignedHash)
        ethers.encodeRlp([
            0, //nonce
            20*10**9, //gasprice
            
        ])
        console.log("-------------------------------------")

        // const toLuniverse = tx.toJSON()
        // delete toLuniverse.value
        // const signedTx = await wallet.signTransaction(toLuniverse)
        // const checkTx = ethers.Transaction.from(signedTx)
        // console.log("-------------------------------------")
        // console.log(signedTx)
        // const deserialized = checkTx.toJSON()
        // console.log(deserialized)
        // console.log(ethers.hexlify(deserialized))
        // console.log(ethers.encodeRlp(JSON.stringify(deserialized)))
        // console.log("-------------------------------------")
        return 
    }
}