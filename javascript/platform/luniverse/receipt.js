const axios = require('axios');
const dotenv = require('dotenv');

const url = "https://api.luniverse.io/tx/v2.0/receipts";

(async()=>{
    dotenv.config()
    const {authToken} = require("./credential")
    const Authorization = `Bearer ${authToken}`

    environmentId = process.env["ConsortiumEnvId"];
    // txHash = "0x71554dce35609bb2d3efade68c52733637c749d9b56721cfba8448c227054a16"; 
    // txId = "b8ca9fbd-9c0d-4f07-99ca-1fbdbbee84c9";
    txId = "8f7a57cb-17c9-4c72-b4c4-d39c67c46819";

    const res = await axios({
        baseURL: url,
        method:"GET",
        headers: {
            Authorization,
        },
        data: {
            environmentId,
            txId,
            // txHash,
        },
    })

    console.log(JSON.stringify(res.data, "", 2))
})()