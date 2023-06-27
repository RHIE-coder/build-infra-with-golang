/* dependencies modules */
const axios = require('axios');
const { ethers } = require('ethers');

class Validator {
   
    #target

    static of(targetObj) {
        return new Validator(targetObj)
    }

    constructor(targetObj) {
        this.#target = targetObj;
    }

    assertExist(prop){
        this.#target[prop] = this.#target[prop] ?? undefined;
        if(this.#target[prop] === undefined) {
            throw new ReferenceError(`the property of "${prop}" is not exist`);
        }
        return this;
    }
}

class Luniverse {

    #baseURL
    #chainName

    constructor(originURL) {
        if(originURL === undefined) {
            throw new ReferenceError('"originURL" cannot be undefined')
        }
        
        this.#baseURL = originURL
        this.#chainName = "luniverse"
    }

    get originURL() {
        return this.#baseURL
    }

    get chainName() {
        return this.#chainName
    }

    async request(reqObj) {

        const client = axios.create({
            baseURL: this.#baseURL,
        })

        if(reqObj.data){
            return (await client[reqObj.method](reqObj.url, reqObj.data)).data
        }

        return (await client[reqObj.method](reqObj.url)).data
    }

    async accountNonce(address) {

        if(address === undefined) {
            throw new ReferenceError('"address" cannot be undefined')
        }

        const reqInfo = {
            method: "get",
            url:`/bc/v1/${this.#chainName}/accounts/nonce?address=${address}` ,
        }
        return await this.request(reqInfo)
    }

    async tokenBalance(address) {

        if(address === undefined) {
            throw new ReferenceError('"address" cannot be undefined')
        }

        const reqInfo =  {
            method: "get",
            url:`/bc/v1/${this.#chainName}/tokens/balance-of?address=${address}`,
        }
        return await this.request(reqInfo)
    }

    async tokenRawTxn(data) {

        const verifying = Validator.of(data)
            .assertExist("from")
            .assertExist("to")
            .assertExist("amount")

        if(this.#chainName !== "luniverse") {
            verifying
                .assertExist("gas")
                .assertExist("gas_price")
        }
        const reqInfo =  {
            method: "post",
            url: `/bc/v1/${this.#chainName}/tokens/transfer/raw-txn`,
            data: data,
        }
        return await this.request(reqInfo)
    }

    async sendTxn(signedTx) {
        if(signedTx === undefined) {
            throw new ReferenceError('"signedTx" cannot be undefined')
        }

        const reqInfo =  {
            method: "post",
            url:`/bc/v1/${this.#chainName}/transactions/send-txn`,
            data: {
                signed_tx: signedTx,
            }
        }
        return await this.request(reqInfo)
    }

}

const WalletHandler = {

    generateMnemonic() {
        return ethers.Wallet.createRandom().mnemonic.phrase;
    },

    createWalletFromMnemonic(mnemonicString, index) {
        index = index ?? 0;
        const childPath = `m/44'/60'/0'/0/${index}`
        const wallet = ethers.Wallet.fromPhrase(mnemonicString)
        return ethers.HDNodeWallet.fromMnemonic(wallet.mnemonic, childPath);
    },

    async sign(wallet, data) {
        Validator.of(data)
            .assertExist("from")
            .assertExist("to")
            .assertExist("gas")
            .assertExist("gasPrice")
            .assertExist("data")
            .assertExist("nonce")

        console.log(data)
        const tx = new ethers.Transaction()
        
        tx.chainId = 256;

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

        return await wallet.signTransaction(tx)
    },
};

module.exports = {
    Luniverse,
    WalletHandler,
}