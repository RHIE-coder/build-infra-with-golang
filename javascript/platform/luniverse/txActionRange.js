const axios = require('axios');
const dotenv = require('dotenv');

const url = "https://api.luniverse.io/tx/v2.0/analytics";

(async()=>{
    dotenv.config()
    const {authToken} = require("./credential")
    const Authorization = `Bearer ${authToken}`

    environmentId = process.env["ConsortiumEnvId"];
    fromDate = "20230503"
    toDate = "20230503"

    const res = await axios({
        baseURL: url,
        method:"GET",
        headers: {
            Authorization,
        },
        data: {
            environmentId,
            fromDate,
            toDate,
        },
    })

    console.log(JSON.stringify(res.data, "", 2))
})()