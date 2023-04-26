const customLib = require("../libs")
require("dotenv").config(__dirname);
const API_SERVER = "http://localhost:5000";
const axios = require("axios")
const CHAINS = {
    ETH: "ethereum",
    GEO: "goerli",
    SEP: "sepolia",
    LUN: "luniverse"
}

const owen = "0x2894706debA1DF71735053E8f55f65D34348c051";
const alice = "0xa49EAcDaDeF57F4ABC4d52D528945CE4c3834293";
const rhieny = "0x2D81c2486F2C8a286B067cdEdda2E6815e61DDdA";
const huber = "0xe1625a0d89B0fB0BfC3835E91B1FA8475409aE8E";
const andy = "0xf44ec05E8d0065252e3a9D2b8334225d3Ee71B4B";

function GoerliScenario() {

}

(async()=>{
    const rhienyPK = customLib.getPrivateKeyFromMnemonic(process.env.RHIENY_MNEMONIC).slice(2)
    const owenPK =   process.env.OWEN_PRIVATE_KEY
    const huberPK =  process.env.HUBER_PRIVATE_KEY

    const client = axios.create({
        baseURL: API_SERVER,
    })

    const res = await client.get(`/bc/v1/goerli/accounts/balance?address=${owen}`)
    console.log(await res.data.code)
})()

