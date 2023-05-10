const axios = require('axios');
const dotenv = require('dotenv');

const url = "https://api.luniverse.io/svc/v2/auth-tokens";

(async()=>{
    dotenv.config()
    const accessKey= process.env["ACCESS_KEY"]
    const secretKey= process.env["SECRET_KEY"]
    const payload = {
        accessKey: accessKey,
        secretKey: secretKey,
        expiresIn: 604800,
    }

    const res = await axios({
        baseURL: url,
        method: "POST",
        data: payload,
    })

    console.log(JSON.stringify(res.data, "", 2))
})()