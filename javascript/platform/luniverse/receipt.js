const axios = require('axios');
const dotenv = require('dotenv');

const url = "https://api.luniverse.io/tx/v2.0/receipts";

(async()=>{
    dotenv.config()
    const {authToken} = require("./credential")
    const Authorization = `Bearer ${authToken}`

    // environmentId = process.env["ConsortiumEnvId"];
    environmentId = process.env["LegacyEnvId"];
    // txHash = "0x71554dce35609bb2d3efade68c52733637c749d9b56721cfba8448c227054a16"; 
    // txId = "b8ca9fbd-9c0d-4f07-99ca-1fbdbbee84c9";
    successTxHash = "0x5a17234ae08715611bdd03c492b5d7fdcbb13dbe2bc1d05a221eb64fca970b86"
    failTxHash = "0xe09f15b8327010bd8a58bd63910b5b2829f50cdef2fafe5e826297067db5354c"
    txId = "8f7a57cb-17c9-4c72-b4c4-d39c67c46819";

    const res = await axios({
        baseURL: url,
        method:"GET",
        headers: {
            Authorization,
        },
        data: {
            environmentId,
            // txId,
            // txHash,
            txHash: successTxHash,
            // failTxHash,
        },
    })

    console.log(JSON.stringify(res.data, "", 2))
})()