const axios = require('axios');
const dotenv = require('dotenv');
const { v4: uuidv4 } = require('uuid');
const myUUID = uuidv4();
console.log(myUUID);

const actionName = "chosunnftdeprecateDEV"
const url = `https://api.luniverse.io/tx/v2.0/transactions/${actionName}`;

(async()=>{
    dotenv.config()
    const {authToken} = require("./credential")
    const Authorization = `Bearer ${authToken}`

    environmentId = process.env["ConsortiumEnvId"];
    // txHash = "0x71554dce35609bb2d3efade68c52733637c749d9b56721cfba8448c227054a16"; 
    txId = "b8ca9fbd-9c0d-4f07-99ca-1fbdbbee84c9";

    const res = await axios({
        baseURL: url,
        method:"POST",
        headers: {
            Authorization,
            "Content-Type": "application/json"
        },
        data: {
            txId: myUUID,
            from: "0x0d3b0e20b1768f407c88753f0f18653b12e7c9e6",
            inputs: {
                index:1000,
            },
            callbackUrl: "https://d7d4-1-212-236-115.ngrok-free.app/callback/nft/deprecate",
        },
    })

    console.log(JSON.stringify(res.data, "", 2))
})()
/*  
{
  environmentId: '1671600618469641269',
  txId: '8f7a57cb-17c9-4c72-b4c4-d39c67c46819',
  txHash: '0xaa6e946ed1b5548cb24eaad36b098d105ccd83130d56461081aebe95f14d35c7',
  status: 'Fail',
  receipt: {
    blockHash: '0x6d3f372acb6a26cc783ce810483c6c091315d34118b277de6c19ca4ea8d5e2e8',
    blockNumber: '0xaf04b4',
    contractAddress: null,
    cumulativeGasUsed: '0x65de',
    effectiveGasPrice: '0x0',
    from: '0x0d3b0e20b1768f407c88753f0f18653b12e7c9e6',
    gasUsed: '0x65de',
    logs: [],
    logsBloom: '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
    status: '0x0',
    to: '0x71273ba1bd7e9174b932d1eda4cfd63aef4d6d06',
    transactionHash: '0xaa6e946ed1b5548cb24eaad36b098d105ccd83130d56461081aebe95f14d35c7',
    transactionIndex: '0x0',
    type: '0x0'
  }
}
*/
