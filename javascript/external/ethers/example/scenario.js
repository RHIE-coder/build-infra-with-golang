const customLib = require("../libs");
require("dotenv").config(__dirname);
const axios = require("axios")


/* APIs */
const HTTP_METHODS = {
    GET: "get",
    POST: "post",
};
const API_SERVER ="http://localhost:5000";

async function request(reqObj) {

    const client = axios.create({
        baseURL: API_SERVER,
    })

    if(reqObj.data){
        return (await client[reqObj.method](reqObj.url, reqObj.data)).data
    }

    return (await client[reqObj.method](reqObj.url)).data
}

function assertExist(v){
    v = v ?? undefined
    if(v === undefined) {
        throw new Error("not exist")
    }
}

const AccountBalance = function(chains, address) {
    return {
        method: HTTP_METHODS.GET,
        url: `/bc/v1/${chains}/accounts/balance?address=${address}`,
    }
} 

const AccountNonce = function(chains, address) {
    return {
        method: HTTP_METHODS.GET,
        url:`/bc/v1/${chains}/accounts/nonce?address=${address}` ,
    }
} 

const AccountEstimate = function(chains, data) {

    assertExist(data.amount)
    assertExist(data.from)
    assertExist(data.to)

    return {
        method: HTTP_METHODS.POST,
        url:`/bc/v1/${chains}/accounts/transfer/estimate-gas`,
        data: data
    }
} 

const AccountRawTxn = function(chains, data) {

    assertExist(data.amount)
    assertExist(data.from)
    assertExist(data.to)
    assertExist(data.gas)
    assertExist(data.gas_price)

    return {
        method: HTTP_METHODS.POST,
        url:`/bc/v1/${chains}/accounts/transfer/raw-txn`,
        data: data
    }
} 


const TokenBalance = function(chains, address) {
    return {
        method: HTTP_METHODS.GET,
        url:`/bc/v1/${chains}/tokens/balance-of?address=${address}`,
    }
} 

const TokenEstimate = function(chains, data) {
    return {
        method: HTTP_METHODS.POST,
        url:`/bc/v1/${chains}/tokens/transfer/estimate-gas`,
        data: data
    }
}

const TokenRawTxn = function(chains, data) {
    return {
        method: HTTP_METHODS.POST,
        url: `/bc/v1/${chains}/tokens/transfer/raw-txn`,
        data: data,
    }
}


const SendTxn = function(chains, data) {

    assertExist(data.signed_tx)

    return {
        method: HTTP_METHODS.POST,
        url:`/bc/v1/${chains}/transactions/send-txn`,
        data: data
    }
} 



/* Blockchain Network */
const ETH = "ethereum"
const GEO = "goerli"
const SEP = "sepolia"
const LUN = "luniverse"

const ETH_CHAIN_ID = 1
const GEO_CHAIN_ID = 5
const SEP_CHAIN_ID = 11155111
const LUN_CHAIN_ID = 256


/* Account Address */
const OWEN = "0x2894706DEBa1df71735053e8F55F65d34348C051"
const ALICE = "0xa49EAcDaDeF57F4ABC4d52D528945CE4c3834293"
const RHIELU = "0x2D81c2486F2C8a286B067cdEdda2E6815e61DDdA"
const HUBER = "0xe1625a0d89B0fB0BfC3835E91B1FA8475409aE8E"
const ANDY = "0xf44ec05E8d0065252e3a9D2b8334225d3Ee71B4B";

const RHIELU_PK = customLib.getPrivateKeyFromMnemonic(process.env.RHIENY_MNEMONIC).slice(2)
const OWEN_PK =   process.env.OWEN_PRIVATE_KEY
const HUBER_PK =  process.env.HUBER_PRIVATE_KEY;

(async()=>{
    // await GoerliETH();
    // await GoerliToken();
    await LuniverseToken();
    // console.log(require('ethers'))
})();

async function LuniverseToken() {
    let result;

    const network = LUN
    const chainId = LUN_CHAIN_ID

    const tokenSendInfo = {
        amount: "100000000000000",
        from: RHIELU,
        to: ANDY,
    }
    console.log(" --- token balance --- ")
    result = await request(TokenBalance(network, RHIELU))
    console.log(result)

    console.log(" --- account nonce --- ")
    result = await request(AccountNonce(network, RHIELU)) 
    console.log(result)
    const nonce = result.data.nonce

    console.log(" --- token estimate gas --- ")
    result = await request(TokenEstimate(network, {
        ...tokenSendInfo,
    }))
    console.log(result)

    console.log(" --- token raw transation --- ")
    result = await request(TokenRawTxn(network, {
        ...tokenSendInfo,
    }))
    console.log(result)

    const rawTxn = {
        from: result.data.from,
        to: result.data.to,
        gas: result.data.gas_limit,
        gasPrice: result.data.gas_price,
        data: result.data.data,
        nonce: nonce,
    }

    console.log(" >>> sign transaction")
    // const signed_tx = (await customLib.signTransactionLuniverse(rawTxn, RHIELU_PK, chainId)).slice(2)
    const signed_tx = (await customLib.signTransaction(rawTxn, RHIELU_PK, chainId)).slice(2)
    console.log(signed_tx)

    console.log(" --- send transaction --- ")
    result = await request(SendTxn(network, {
        signed_tx,
    }))
    console.log(result)
}

async function GoerliToken() {
    let result;

    const network = GEO
    const chainId = GEO_CHAIN_ID
    const tokenSendInfo = {
        amount: "2000",
        from: HUBER,
        to: OWEN,
    }

    console.log(" --- token balance --- ")
    result = await request(TokenBalance(network, HUBER))
    console.log(result)

    console.log(" --- account nonce --- ")
    result = await request(AccountNonce(network, HUBER)) 
    console.log(result)
    const nonce = result.data.nonce

    console.log(" --- token estimate gas --- ")
    result = await request(TokenEstimate(network, {
        ...tokenSendInfo,
    }))
    console.log(result)
    const estimate = result.data

    console.log(" --- token raw transation --- ")
    result = await request(TokenRawTxn(network, {
        ...tokenSendInfo,
        gas: estimate.gas,
        gas_price: estimate.gas_price,
    }))
    console.log(result)

    const rawTxn = {
        from: result.data.from,
        to: result.data.to,
        gas: result.data.gas,
        gasPrice: result.data.gas_price,
        data: result.data.data,
        nonce: nonce,
    }
    console.log(rawTxn)
    
    console.log(" >>> sign transaction")
    const signed_tx = (await customLib.signTransaction(rawTxn, HUBER_PK, chainId)).slice(2)
    console.log(signed_tx)

    console.log(" --- send transaction --- ")
    result = await request(SendTxn(network, {
        signed_tx,
    }))
    console.log(result)
}

async function GoerliETH() {
    // console.log(require("ethers").toBeHex(49)) // GET NONCE
    let result;

    const network = GEO
    const chainId = GEO_CHAIN_ID
    const coinSendInfo = {
        amount: "1000000000000000",
        from: OWEN,
        to: ALICE,
    }
    console.log(" --- account balance --- ")
    result = await request(AccountBalance(network, OWEN))
    console.log(result)

    console.log(" --- account nonce --- ")
    result = await request(AccountNonce(network, OWEN)) 
    console.log(result)
    const nonce = result.data.nonce

    console.log(" --- account estimate gas --- ")
    result = await request(AccountEstimate(network, {
        ...coinSendInfo,
    }))
    console.log(result)
    const estimate = result.data

    console.log(" --- account raw transation --- ")
    result = await request(AccountRawTxn(network, {
        ...coinSendInfo,
        gas: estimate.gas,
        gas_price: estimate.gas_price,
    }))
    console.log(result)
    const rawTxn = {
        from: result.data.from,
        to: result.data.to,
        value: result.data.value,
        gas: result.data.gas,
        gasPrice: result.data.gas_price,
        nonce: nonce,
    }
    console.log(rawTxn)
    
    console.log(" >>> sign transaction")

    const signed_tx = (await customLib.signTransaction(rawTxn, OWEN_PK, chainId)).slice(2)
    console.log(signed_tx)

    console.log(" --- send transaction --- ")
    result = await request(SendTxn(network, {
        signed_tx,
    }))
    console.log(result)
}

function verboseGas(rawTxn){
    t_gasPrice = BigInt(rawTxn.gasPrice)
    t_gasLimit = BigInt(rawTxn.gas)
    t_parsed_gasPrice = require("ethers").formatUnits(t_gasPrice, "gwei")
    console.log(t_gasPrice)
    console.log("gasPrice ( gwei ): ", t_parsed_gasPrice)
    console.log(t_gasLimit)
}