const ethers = require('ethers');
const InputDataDecoder = require('ethereum-input-data-decoder');
const axios = require('axios')
const endpoint = 'http://3.37.74.37:22000';

class JsonRpcRequester{

    #_baseURL
    #_axios

    constructor(baseURL){
        this.#_baseURL = baseURL;
        this.#_axios = axios.create({baseURL})
    }

    get baseURL(){
        return this.#_baseURL
    }

    async dial(req){
        const body = { jsonrpc: "2.0", ...req};
        return (await this.#_axios.post("", body)).data
    }

}

class Utils {
    static HexToNumber(hexString) {
        return parseInt(hexString, 16);
    }

    static NumberToHex(num) {
        return '0x'+num.toString(16);
    }
}

const ABI =  [{"constant":false,"inputs":[{"name":"key","type":"string"},{"name":"value","type":"string"}],"name":"invoke","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getVersion","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"key","type":"string"}],"name":"get","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"strVersion","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}];

async function main() {
    const rpc = new JsonRpcRequester(endpoint);
    const res = await rpc.dial({id:1,method:"eth_blockNumber",params:[]}) 
    const blockNumber = Utils.HexToNumber(res.result)
    console.log(res.result)
    console.log(blockNumber);
    const block = await rpc.dial({id:2,method:'eth_getBlockByNumber',params:[Utils.NumberToHex(1000), true]})
    console.log(block.result)
    console.log(block.result.transactions)
    const decoder = new InputDataDecoder(ABI);
    const result = decoder.decodeData(block.result.transactions[0].input)
    console.log(result.inputs.length)
    console.log(result.inputs[0])
    console.log(JSON.parse(result.inputs[1]))
}

main()

