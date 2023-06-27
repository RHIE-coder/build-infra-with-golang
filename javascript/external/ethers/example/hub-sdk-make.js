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

const ALLOWED_BLOCKCHAIN_NETWORK = [
    "ethereum",
    "goerli",
    "sepolia",
    "luniverse",
]

class BaseAPI {

    #baseURL
    #chainName

    constructor(originURL, chainName) {
        if(originURL === undefined) {
            throw new ReferenceError('"originURL" cannot be undefined')
        }

        if(chainName === undefined) {
            throw new ReferenceError('"chainName" cannot be undefined')
        }

        if (!ALLOWED_BLOCKCHAIN_NETWORK.includes(chainName.toLowerCase())) {
            throw new RangeError(`the blockchain network is not supported`);
        }

        this.#baseURL = originURL
        this.#chainName = chainName
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

    async accountBalance(address) {

        if(address === undefined) {
            throw new ReferenceError('"address" cannot be undefined')
        }

        const reqInfo =  {
            method: "get",
            url: `/bc/v1/${this.#chainName}/accounts/balance?address=${address}`,
        }
        return await request(reqInfo)
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

    async accountEstimate(data) {

        Validator.of(data)
            .assertExist("amount")
            .assertExist("from")
            .assertExist("to")

            const reqInfo =  {
                method: "post",
                url:`/bc/v1/${this.#chainName}/accounts/transfer/estimate-gas`,
                data: data
            }
        return await this.request(reqInfo)
    }

    async accountRawTxn(data) {

        Validator.of(data)
            .assertExist("amount")
            .assertExist("from")
            .assertExist("to")
            .assertExist("gas")
            .assertExist("gas_price")

        const reqInfo =  {
            method: "post",
            url:`/bc/v1/${this.#chainName}/accounts/transfer/raw-txn`,
            data: data
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

    async tokenEstimate(data) {

        Validator.of(data)
            .assertExist("amount")
            .assertExist("from")
            .assertExist("to")

        const reqInfo =  {
            method: "post",
            url:`/bc/v1/${this.#chainName}/tokens/transfer/estimate-gas`,
            data: data
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

class LuniverseAPI extends BaseAPI {
    constructor(originURL) {
        super(originURL, "luniverse");
    }
}

const WalletHandler = {

    ETH_CHAIN_ID : 1,
    GEO_CHAIN_ID : 5,
    SEP_CHAIN_ID : 11155111,
    LUN_CHAIN_ID : 256,
    
    generateMnemonic() {
        return ethers.Wallet.createRandom().mnemonic.phrase;
    },

    createWalletFromMnemonic(mnemonicString, index) {
        index = index ?? 0;
        const childPath = `m/44'/60'/0'/0/${index}`
        const wallet = ethers.Wallet.fromPhrase(mnemonicString)
        return ethers.HDNodeWallet.fromMnemonic(wallet.mnemonic, childPath);
    },

    async sign(wallet, networkId, data) {
        Validator.of(data)
            .assertExist("from")
            .assertExist("to")
            .assertExist("gas")
            .assertExist("gasPrice")
            .assertExist("data")
            .assertExist("nonce")

        console.log(data)
        const tx = new ethers.Transaction()
        
        tx.chainId = networkId;

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

(async()=>{
    await LuniverseTokenScenario()
    // require("dotenv").config(__dirname);
    // const mnemonic = process.env.METAMASK_MNEMONIC
    // const masterWallet = WalletHandler.createMasterWalletFromMnemonic(mnemonic);
    // const childWallet1 = WalletHandler.deriveChildFromWallet(masterWallet, 1);
    // const childWallet2 = WalletHandler.deriveChildFromWallet(masterWallet, 2);
    // console.log(masterWallet.mnemonic.phrase)
    // console.log(childWallet1.mnemonic.phrase)
    // console.log(childWallet2.mnemonic.phrase)

    // const sender = masterWallet.address
    // const receiver = childWallet1.address;

    // const api = new BaseAPI("https://devapi.publish-hub.io", "sepolia")
    // const balance = await api.tokenBalance(sender)
    // console.log(balance)


})()

async function LuniverseTokenScenario() {
    require("dotenv").config(__dirname);

    const mnemonic = process.env.LUNIVERSE_MNEMONIC;
    const wallet0 = WalletHandler.createWalletFromMnemonic(mnemonic);
    const wallet1 = WalletHandler.createWalletFromMnemonic(mnemonic, 1);
    const wallet2 = WalletHandler.createWalletFromMnemonic(mnemonic, 2);

    console.log(wallet0.mnemonic.phrase)
    console.log(wallet1.mnemonic.phrase)
    console.log(wallet2.mnemonic.phrase)

    const sender = wallet0.address
    const receiver = "0xe7517164cBd1943eD5dffe1fbAC14E467Db41a75";

    const api = new LuniverseAPI("http://localhost:5000")

    let resp;
    
    resp = await api.tokenBalance(sender)
    const balance = resp.data
    console.log(resp)
    console.log(balance)


    resp = await api.accountNonce(sender)
    const nonce = resp.data.nonce
    console.log(resp);
    console.log(nonce);




    resp = await api.accountEstimate({
        from: sender,
        to: receiver,
        amount: "0.0001"
    })
    const estimate = resp.data
    console.log(resp);
    console.log(estimate);



    resp = await api.tokenRawTxn({
        amount: "0.0001",
        from:sender,
        to: receiver,
    })
    const raw = resp.data;
    console.log(resp);
    console.log(raw);

    
    const signedTxHex = (await WalletHandler.sign(wallet0, WalletHandler.LUN_CHAIN_ID, {
        from: raw.from,
        to: raw.to,
        gas: raw.gas_limit,
        gasPrice: raw.gas_price,
        nonce: nonce,
        data: raw.data,
    }));
    const signedTx = signedTxHex.slice(2);
    console.log(signedTx);

    resp = await api.sendTxn(signedTx);
    const result = resp.data;
    console.log(resp);
    console.log(result);
}


// const wallet = ethers.Wallet.createRandom()
// const mnem = wallet.mnemonic.phrase;
// console.log(wallet.address)
// const derivedWallet = ethers.Wallet.fromPhrase(mnem)
// console.log(derivedWallet.address)
// console.log("========================")
// console.log(wallet.deriveChild(0))
// console.log(wallet.privateKey)
// console.log(wallet.mnemonic.phrase)
// console.log(wallet.address)

// m/44'/60'/0'/0/0
// m/purpose'/coin_type'/account'/change/address_index
// console.log(wallet.derivePath("1").address)
// console.log(derivedWallet.deriveChild(1).mnemonic)
// console.log(wallet.derivePath("1").address)
// console.log(derivedWallet.deriveChild(1).mnemonic)
// console.log("========================")
// console.log(wallet.derivePath("2").address)
// console.log(derivedWallet.deriveChild(2).mnemonic)
// console.log(wallet.derivePath("2").address)
// console.log(derivedWallet.deriveChild(2).mnemonic)

// const mnemonic = bip39.generateMnemonic();
// console.log(mnemonic)